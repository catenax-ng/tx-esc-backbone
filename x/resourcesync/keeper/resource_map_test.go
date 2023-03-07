// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
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
	addresses := []string{Alice, Bob}
	for i := range items {
		items[i].Originator = addresses[i%len(addresses)]
		items[i].OrigResId = strconv.Itoa(i)
		keeper.SetResourceMap(ctx, items[i])
	}
	return items
}

func TestResourceMapGet(t *testing.T) {
	keeper, ctx := keepertest.ResourcesyncKeeper(t)
	items := createNResourceMap(keeper, ctx, 10)
	for _, item := range items {
		resKey := createValidResouceKey(item.Originator, item.OrigResId)
		rst, found := keeper.GetResourceMap(ctx, resKey)
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
		resKey := createValidResouceKey(item.Originator, item.OrigResId)
		keeper.RemoveResourceMap(ctx, resKey)
		_, found := keeper.GetResourceMap(ctx, resKey)
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

func TestKeeper_HasResourceMapFor(t *testing.T) {
	keeper, ctx := keepertest.ResourcesyncKeeper(t)
	_ = createNResourceMap(keeper, ctx, 2)

	require.True(t, keeper.HasResourceMapFor(ctx, createValidResouceKey(Alice, "0")))
	require.True(t, keeper.HasResourceMapFor(ctx, createValidResouceKey(Bob, "1")))
	require.False(t, keeper.HasResourceMapFor(ctx, createValidResouceKey(Alice, "1")))
	require.False(t, keeper.HasResourceMapFor(ctx, createValidResouceKey(Bob, "0")))
	require.False(t, keeper.HasResourceMapFor(ctx, createValidResouceKey(Carol, "0")))
	require.False(t, keeper.HasResourceMapFor(ctx, createValidResouceKey(Carol, "1")))
	require.False(t, keeper.HasResourceMapFor(ctx, createValidResouceKey(Carol, "2")))

}

func TestKeeper_RemoveAndGetResourceMap(t *testing.T) {
	keeper, ctx := keepertest.ResourcesyncKeeper(t)
	_ = createNResourceMap(keeper, ctx, 2)
	var result *types.ResourceMap
	var found bool
	result, found = keeper.RemoveAndGetResourceMap(ctx, createValidResouceKey(Alice, "0"))
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&types.ResourceMap{Resource: types.Resource{Originator: Alice, OrigResId: "0"}}),
		nullify.Fill(result),
	)
	result, found = keeper.RemoveAndGetResourceMap(ctx, createValidResouceKey(Bob, "1"))
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&types.ResourceMap{Resource: types.Resource{Originator: Bob, OrigResId: "1"}}),
		nullify.Fill(result),
	)
	result, found = keeper.RemoveAndGetResourceMap(ctx, createValidResouceKey(Alice, "1"))
	require.False(t, found)
	require.Nil(t, result)
	result, found = keeper.RemoveAndGetResourceMap(ctx, createValidResouceKey(Bob, "0"))
	require.False(t, found)
	require.Nil(t, result)
	result, found = keeper.RemoveAndGetResourceMap(ctx, createValidResouceKey(Carol, "0"))
	require.False(t, found)
	require.Nil(t, result)
	result, found = keeper.RemoveAndGetResourceMap(ctx, createValidResouceKey(Carol, "1"))
	require.False(t, found)
	require.Nil(t, result)
	result, found = keeper.RemoveAndGetResourceMap(ctx, createValidResouceKey(Carol, "2"))
	require.False(t, found)
	require.Nil(t, result)
}
