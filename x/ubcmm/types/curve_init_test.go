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
package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Curve_Fit_Happy(t *testing.T) {
	t.Run("primary set of valid params", func(t *testing.T) {
		c := validCurve()
		require.NoError(t, c.Fit())
	})

	t.Run("alternate set of valid params", func(t *testing.T) {
		type test struct {
			name     string
			modifier func(c *Curve)
		}

		tests := []test{
			{"BPoolUnder_alternate1",
				func(c *Curve) {
					c.BPoolUnder = sdk.NewDec(150e6)
					c.BPool = sdk.NewDec(150e6)
				}},
			{"BPoolUnder_alternate2",
				func(c *Curve) {
					c.BPoolUnder = sdk.NewDec(90e6)
					c.BPool = sdk.NewDec(90e6)
				}},
		}

		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				c := validCurve()
				tc.modifier(&c)
				assert.NoError(t, c.Fit())
			})
		}
	})
}

func Test_Curve_Fit_Error(t *testing.T) {
	type test struct {
		name     string
		modifier func(c *Curve)
	}

	tests := []test{
		{"RefTokenSupply_isNotSet",
			func(c *Curve) { c.RefTokenSupply = sdk.Dec{} }},
		{"RefTokenPrice_isNotSet",
			func(c *Curve) { c.RefTokenPrice = sdk.Dec{} }},
		{"RefProfitFactor_isNotSet",
			func(c *Curve) { c.RefProfitFactor = sdk.Dec{} }},
		{"BPoolUnder_isNotSet",
			func(c *Curve) { c.BPoolUnder = sdk.Dec{} }},
		{"SlopeP2_isNotSet",
			func(c *Curve) { c.SlopeP2 = sdk.Dec{} }},
		{"SlopeP3_isNotSet",
			func(c *Curve) { c.SlopeP3 = sdk.Dec{} }},
		{"FactorFy_isNotSet",
			func(c *Curve) { c.FactorFy = sdk.Dec{} }},
		{"FactorFxy_isNotSet",
			func(c *Curve) { c.FactorFxy = sdk.Dec{} }},
		{"RefTokenSupply_isTooLess",
			func(c *Curve) { c.RefTokenSupply = sdk.NewDec(6e6) }},
		{"RefTokenSupply_isTooHigh",
			func(c *Curve) { c.RefTokenSupply = sdk.NewDec(6e10) }},
		{"RefTokenTokenPrice_isTooLess",
			func(c *Curve) { c.RefTokenPrice = sdk.NewDecWithPrec(1, 1) }},
		{"RefTokenTokenPrice_isTooHigh",
			func(c *Curve) { c.RefTokenPrice = sdk.NewDec(10) }},
		{"BPoolUnder_isTooLess",
			func(c *Curve) { c.BPoolUnder = sdk.NewDec(100e3) }},
		{"BPoolUnder_isTooHigh",
			func(c *Curve) { c.BPoolUnder = sdk.NewDec(100e10) }},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			c := validCurve()
			tc.modifier(&c)
			assert.Error(t, c.Fit())
		})
	}
}
func BenchmarkUbcFit(b *testing.B) {
	c := validCurve()
	for i := 0; i < b.N; i++ {
		c.Fit()
	}
}

func IsEqualDecimal(t *testing.T, expected string, actual sdk.Dec) {
	t.Helper()
	expectedDec, err := sdk.NewDecFromStr(expected)
	require.NoError(t, err)
	assert.Equal(t, expectedDec, actual)
}

func validCurve() Curve {
	return Curve{
		RefTokenSupply:  sdk.NewDec(6e9),
		RefTokenPrice:   sdk.NewDec(1),
		RefProfitFactor: sdk.NewDec(10),
		BPool:           sdk.NewDec(100e6),
		BPoolUnder:      sdk.NewDec(100e6),
		SlopeP2:         sdk.NewDecWithPrec(3, 1).Quo(sdk.NewDec(3e9)),
		SlopeP3:         sdk.NewDec(2).Quo(sdk.NewDec(3e9)),
		FactorFy:        sdk.NewDecWithPrec(2, 1),
		FactorFxy:       sdk.NewDec(15832600001),
	}
}
