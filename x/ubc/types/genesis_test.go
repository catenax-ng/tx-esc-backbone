// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package types_test

import (
	"testing"

	"github.com/catenax/esc-backbone/x/ubc/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				Curve: &types.Curve{
					FS0:             new(types.FlatSegment),
					S0:              new(types.BezierSegment),
					S1:              new(types.BezierSegment),
					S2:              new(types.FixedBezierSegment),
					QS3:             new(types.QuadraticSegment),
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
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
