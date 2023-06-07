// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package types_test

import (
	"testing"

	"github.com/catenax/esc-backbone/x/ubc/types"
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

				Ubcobject: &types.Ubcobject{
					FS0:             "75",
					S0:              "71",
					S1:              "32",
					S2:              "77",
					QS3:             "52",
					RefProfitFactor: "100",
					RefTokenSupply:  "36",
					RefTokenPrice:   "94",
					BPool:           "64",
					BPoolUnder:      "17",
					FactorFy:        "52",
					FactorFxy:       "63",
					TradingPoint:    "17",
					CurrentSupply:   "91",
					Slopep2:         "64",
					Slopep3:         "57",
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
