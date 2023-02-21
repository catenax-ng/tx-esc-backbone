package keeper_test

import (
	"context"
	"github.com/catenax/esc-backbone/testutil"
	"testing"

	keepertest "github.com/catenax/esc-backbone/testutil/keeper"
	"github.com/catenax/esc-backbone/x/resourcesync/keeper"
	"github.com/catenax/esc-backbone/x/resourcesync/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

const (
	Alice = testutil.Alice
	Bob   = testutil.Bob
	Carol = testutil.Carol
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

func createValidResouceKey(originator string, origResId string) types.ResourceKey {
	resourceKey, err := types.NewResourceKey(originator, origResId)
	if err != nil {
		panic(err)
	}
	return resourceKey
}
