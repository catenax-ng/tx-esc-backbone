// Copyright (c) 2022-2023 Contributors to the Eclipse Foundation
//
// See the NOTICE file(s) distributed with this work for additional
// information regarding copyright ownership.
//
// This program and the accompanying materials are made available under the
// terms of the Apache License, Version 2.0 which is available at
// https://www.apache.org/licenses/LICENSE-2.0.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.
//
// SPDX-License-Identifier: Apache-2.0
package cli_test

import (
	"fmt"
	"testing"

	tmcli "github.com/cometbft/cometbft/libs/cli"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/status"

	"github.com/catenax/esc-backbone/testutil/network"
	"github.com/catenax/esc-backbone/testutil/nullify"
	"github.com/catenax/esc-backbone/x/ubcmm/client/cli"
	"github.com/catenax/esc-backbone/x/ubcmm/types"
)

func networkWithCurveObjects(t *testing.T) (*network.Network, types.Curve) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	curve := &types.Curve{
		RefTokenSupply:            sdk.ZeroDec(),
		RefTokenPrice:             sdk.ZeroDec(),
		RefProfitFactor:           sdk.ZeroDec(),
		BPool:                     sdk.ZeroDec(),
		BPoolUnder:                sdk.ZeroDec(),
		SlopeP2:                   sdk.ZeroDec(),
		SlopeP3:                   sdk.ZeroDec(),
		FactorFy:                  sdk.ZeroDec(),
		FactorFxy:                 sdk.ZeroDec(),
		TradingPoint:              sdk.ZeroDec(),
		CurrentSupply:             sdk.ZeroDec(),
		NumericalErrorAccumulator: sdk.ZeroDec(),
	}
	nullify.Fill(&curve)
	state.Curve = curve
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), *state.Curve
}

func TestShowCurve(t *testing.T) {
	net, obj := networkWithCurveObjects(t)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	tests := []struct {
		desc string
		args []string
		err  error
		obj  types.Curve
	}{
		{
			desc: "get",
			args: common,
			obj:  obj,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			var args []string
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowCurve(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetCurveResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.Curve)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.Curve),
				)
			}
		})
	}
}
