package cli_test

import (
	"fmt"
	"github.com/catenax/esc-backbone/testutil"
	"strconv"
	"testing"

	tmcli "github.com/cometbft/cometbft/libs/cli"
	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/catenax/esc-backbone/testutil/network"
	"github.com/catenax/esc-backbone/testutil/nullify"
	"github.com/catenax/esc-backbone/x/resourcesync/client/cli"
	"github.com/catenax/esc-backbone/x/resourcesync/types"
)

const (
	Alice = testutil.Alice
	Bob   = testutil.Bob
	Carol = testutil.Carol
)

// Prevent strconv unused error
var _ = strconv.IntSize

func networkWithResourceMapObjects(t *testing.T, n int) (*network.Network, []types.ResourceMap) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	addresses := []string{Alice, Bob}
	for i := 0; i < n; i++ {
		resourceMap := types.ResourceMap{
			Resource: types.Resource{
				Originator: addresses[i%len(addresses)],
				OrigResId:  strconv.Itoa(i),
			},
		}
		nullify.Fill(&resourceMap)
		state.ResourceMapList = append(state.ResourceMapList, resourceMap)
	}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), state.ResourceMapList
}

func TestShowResourceMap(t *testing.T) {
	net, objs := networkWithResourceMapObjects(t, 2)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	tests := []struct {
		desc         string
		idOriginator string
		idOrigResId  string

		args []string
		err  error
		obj  types.ResourceMap
	}{
		{
			desc:         "found",
			idOriginator: objs[0].Originator,
			idOrigResId:  objs[0].OrigResId,

			args: common,
			obj:  objs[0],
		},
		{
			desc:         "not found",
			idOriginator: strconv.Itoa(100000),
			idOrigResId:  strconv.Itoa(100000),

			args: common,
			err:  status.Error(codes.NotFound, "not found"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				tc.idOriginator,
				tc.idOrigResId,
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowResourceMap(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetResourceMapResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.ResourceMap)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.ResourceMap),
				)
			}
		})
	}
}

func TestListResourceMap(t *testing.T) {
	net, objs := networkWithResourceMapObjects(t, 5)

	ctx := net.Validators[0].ClientCtx
	request := func(next []byte, offset, limit uint64, total bool) []string {
		args := []string{
			fmt.Sprintf("--%s=json", tmcli.OutputFlag),
		}
		if next == nil {
			args = append(args, fmt.Sprintf("--%s=%d", flags.FlagOffset, offset))
		} else {
			args = append(args, fmt.Sprintf("--%s=%s", flags.FlagPageKey, next))
		}
		args = append(args, fmt.Sprintf("--%s=%d", flags.FlagLimit, limit))
		if total {
			args = append(args, fmt.Sprintf("--%s", flags.FlagCountTotal))
		}
		return args
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(objs); i += step {
			args := request(nil, uint64(i), uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListResourceMap(), args)
			require.NoError(t, err)
			var resp types.QueryAllResourceMapResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.ResourceMap), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.ResourceMap),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(objs); i += step {
			args := request(next, 0, uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListResourceMap(), args)
			require.NoError(t, err)
			var resp types.QueryAllResourceMapResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.ResourceMap), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.ResourceMap),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		args := request(nil, 0, uint64(len(objs)), true)
		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListResourceMap(), args)
		require.NoError(t, err)
		var resp types.QueryAllResourceMapResponse
		require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
		require.NoError(t, err)
		require.Equal(t, len(objs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(objs),
			nullify.Fill(resp.ResourceMap),
		)
	})
}
