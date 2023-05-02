// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package ubc_test

import (
	"testing"

	keepertest "github.com/catenax/esc-backbone/testutil/keeper"
	"github.com/catenax/esc-backbone/testutil/nullify"
	"github.com/catenax/esc-backbone/x/ubc"
	"github.com/catenax/esc-backbone/x/ubc/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {

	zeroSegment := func() *types.Segment {
		return &types.Segment{
			P0:     sdk.ZeroDec(),
			A:      sdk.ZeroDec(),
			B:      sdk.ZeroDec(),
			P1:     sdk.ZeroDec(),
			P0X:    sdk.ZeroDec(),
			P1X:    sdk.ZeroDec(),
			DeltaX: sdk.ZeroDec(),
		}
	}

	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		Ubcobject: &types.Ubcobject{
			FS0: &types.Flatsegment{
				Y:  sdk.ZeroDec(),
				X0: sdk.ZeroDec(),
			},
			S0: zeroSegment(),
			S1: zeroSegment(),
			S2: &types.Fixedsegment{
				Segment:     zeroSegment(),
				IntervalP0X: sdk.ZeroDec(),
			},
			QS3: &types.Quadraticsegment{
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

	k, ctx := keepertest.UbcKeeper(t)
	ubc.InitGenesis(ctx, *k, genesisState)
	got := ubc.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.Ubcobject, got.Ubcobject)
	// this line is used by starport scaffolding # genesis/test/assert
}
