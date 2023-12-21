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
package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/catenax/esc-backbone/testutil/keeper"
	"github.com/catenax/esc-backbone/testutil/nullify"
	"github.com/catenax/esc-backbone/x/ubcmm/keeper"
	"github.com/catenax/esc-backbone/x/ubcmm/types"
)

func createTestCurve(keeper *keeper.Keeper, ctx sdk.Context) types.Curve {
	item := types.Curve{
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
	keeper.SetCurve(ctx, item)
	return item
}

func TestCurveGet(t *testing.T) {
	keeper, ctx := keepertest.UbcmmKeeper(t)
	item := createTestCurve(keeper, ctx)
	rst, found := keeper.GetCurve(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestCurveRemove(t *testing.T) {
	keeper, ctx := keepertest.UbcmmKeeper(t)
	createTestCurve(keeper, ctx)
	keeper.RemoveCurve(ctx)
	_, found := keeper.GetCurve(ctx)
	require.False(t, found)
}
