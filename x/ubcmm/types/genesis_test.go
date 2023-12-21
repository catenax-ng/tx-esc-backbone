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
package types_test

import (
	"testing"

	"github.com/catenax/esc-backbone/x/ubcmm/types"
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
					S0:                        new(types.FlatSegment),
					S1:                        new(types.BezierSegment),
					S2:                        new(types.BezierSegment),
					S3:                        new(types.FixedBezierSegment),
					S4:                        new(types.FixedQuadraticSegment),
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
