package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/catenax/esc-backbone/testutil/keeper"
	"github.com/catenax/esc-backbone/testutil/nullify"
	"github.com/catenax/esc-backbone/x/resourcesync/keeper"
	"github.com/catenax/esc-backbone/x/resourcesync/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNResourceMap(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.ResourceMap {
	items := make([]types.ResourceMap, n)
	for i := range items {
		items[i].Originator = strconv.Itoa(i)
		items[i].OrigResId = strconv.Itoa(i)

		keeper.SetResourceMap(ctx, items[i])
	}
	return items
}

func TestResourceMapGet(t *testing.T) {
	keeper, ctx := keepertest.ResourcesyncKeeper(t)
	items := createNResourceMap(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetResourceMap(ctx,
			item.Originator,
			item.OrigResId,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestResourceMapRemove(t *testing.T) {
	keeper, ctx := keepertest.ResourcesyncKeeper(t)
	items := createNResourceMap(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveResourceMap(ctx,
			item.Originator,
			item.OrigResId,
		)
		_, found := keeper.GetResourceMap(ctx,
			item.Originator,
			item.OrigResId,
		)
		require.False(t, found)
	}
}

func TestResourceMapGetAll(t *testing.T) {
	keeper, ctx := keepertest.ResourcesyncKeeper(t)
	items := createNResourceMap(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllResourceMap(ctx)),
	)
}
