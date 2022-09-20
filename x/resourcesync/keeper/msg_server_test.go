package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/catenax/esc-backbone/testutil/keeper"
	"github.com/catenax/esc-backbone/x/resourcesync/keeper"
	"github.com/catenax/esc-backbone/x/resourcesync/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

const (
	Alice = "cosmos1vvwed6f6uu4hjm6s3dqdfrrfqkwvg2dcag3ecf"
	Bob   = "cosmos1w73d7jg8f46qx354hj62d3pa5kfncc47nw5rx2"
)

func setupMsgServer(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
	k, ctx := keepertest.ResourcesyncKeeper(t)
	return keeper.NewMsgServerImpl(*k), *k, sdk.WrapSDKContext(ctx)
}

func TestMsgServer(t *testing.T) {
	ms, keeper, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, keeper)
	require.NotNil(t, ctx)
}
