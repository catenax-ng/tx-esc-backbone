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
			S0: &types.FlatSegment{
				Y:   sdk.ZeroDec(),
				P1X: sdk.ZeroDec(),
			},
			S1: zeroBezierSegment(),
			S2: zeroBezierSegment(),
			S3: &types.FixedBezierSegment{
				BezierSegment: zeroBezierSegment(),
				IntervalP0X:   sdk.ZeroDec(),
			},
			S4: &types.FixedQuadraticSegment{
				A:             sdk.ZeroDec(),
				B:             sdk.ZeroDec(),
				C:             sdk.ZeroDec(),
				ScalingFactor: sdk.ZeroDec(),
				InitialX0:     sdk.ZeroDec(),
				CurrentX0:     sdk.ZeroDec(),
			},
			RefProfitFactor:           sdk.ZeroDec(),
			RefTokenSupply:            sdk.ZeroDec(),
			RefTokenPrice:             sdk.ZeroDec(),
			BPool:                     sdk.ZeroDec(),
			BPoolUnder:                sdk.ZeroDec(),
			FactorFy:                  sdk.ZeroDec(),
			FactorFxy:                 sdk.ZeroDec(),
			TradingPoint:              sdk.ZeroDec(),
			CurrentSupply:             sdk.ZeroDec(),
			SlopeP2:                   sdk.ZeroDec(),
			SlopeP3:                   sdk.ZeroDec(),
			NumericalErrorAccumulator: sdk.ZeroDec(),
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
