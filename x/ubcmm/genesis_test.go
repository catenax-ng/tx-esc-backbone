// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package ubcmm_test

import (
	"testing"

	keepertest "github.com/catenax/esc-backbone/testutil/keeper"
	"github.com/catenax/esc-backbone/testutil/nullify"
	"github.com/catenax/esc-backbone/x/ubcmm"
	"github.com/catenax/esc-backbone/x/ubcmm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {

	zeroBezierSegment := func() *types.BezierSegment {
		return &types.BezierSegment{
			P0Y:    sdk.ZeroDec(),
			A:      sdk.ZeroDec(),
			B:      sdk.ZeroDec(),
			P1Y:    sdk.ZeroDec(),
			P0X:    sdk.ZeroDec(),
			P1X:    sdk.ZeroDec(),
			DeltaX: sdk.ZeroDec(),
		}
	}

	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		Curve: &types.Curve{
			FS0: &types.FlatSegment{
				Y:  sdk.ZeroDec(),
				X0: sdk.ZeroDec(),
			},
			S0: zeroBezierSegment(),
			S1: zeroBezierSegment(),
			S2: &types.FixedBezierSegment{
				BezierSegment: zeroBezierSegment(),
				IntervalP0X:   sdk.ZeroDec(),
			},
			QS3: &types.FixedQuadraticSegment{
				A:             sdk.ZeroDec(),
				B:             sdk.ZeroDec(),
				C:             sdk.ZeroDec(),
				ScalingFactor: sdk.ZeroDec(),
				InitialX0:     sdk.ZeroDec(),
				CurrentX0:     sdk.ZeroDec(),
			},
			RefProfitFactor: sdk.ZeroDec(),
			RefTokenSupply:  sdk.ZeroDec(),
			RefTokenPrice:   sdk.ZeroDec(),
			BPool:           sdk.ZeroDec(),
			BPoolUnder:      sdk.ZeroDec(),
			FactorFy:        sdk.ZeroDec(),
			FactorFxy:       sdk.ZeroDec(),
			TradingPoint:    sdk.ZeroDec(),
			CurrentSupply:   sdk.ZeroDec(),
			SlopeP2:         sdk.ZeroDec(),
			SlopeP3:         sdk.ZeroDec(),
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.UbcmmKeeper(t)
	ubcmm.InitGenesis(ctx, *k, genesisState)
	got := ubcmm.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.Curve, got.Curve)
	// this line is used by starport scaffolding # genesis/test/assert
}
