// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/catenax/esc-backbone/testutil/keeper"
	"github.com/catenax/esc-backbone/testutil/nullify"
	"github.com/catenax/esc-backbone/x/ubc/keeper"
	"github.com/catenax/esc-backbone/x/ubc/types"
)

func createTestUbcobject(keeper *keeper.Keeper, ctx sdk.Context) types.Ubcobject {
	item := types.Ubcobject{
		RefTokenSupply:  sdk.ZeroDec(),
		RefTokenPrice:   sdk.ZeroDec(),
		RefProfitFactor: sdk.ZeroDec(),
		BPool:           sdk.ZeroDec(),
		BPoolUnder:      sdk.ZeroDec(),
		SlopeP2:         sdk.ZeroDec(),
		SlopeP3:         sdk.ZeroDec(),
		FactorFy:        sdk.ZeroDec(),
		FactorFxy:       sdk.ZeroDec(),
		TradingPoint:    sdk.ZeroDec(),
		CurrentSupply:   sdk.ZeroDec(),
	}
	keeper.SetUbcobject(ctx, item)
	return item
}

func TestUbcobjectGet(t *testing.T) {
	keeper, ctx := keepertest.UbcKeeper(t)
	item := createTestUbcobject(keeper, ctx)
	rst, found := keeper.GetUbcobject(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestUbcobjectRemove(t *testing.T) {
	keeper, ctx := keepertest.UbcKeeper(t)
	createTestUbcobject(keeper, ctx)
	keeper.RemoveUbcobject(ctx)
	_, found := keeper.GetUbcobject(ctx)
	require.False(t, found)
}
