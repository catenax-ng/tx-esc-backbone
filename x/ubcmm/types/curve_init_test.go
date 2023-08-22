// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
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

		IsEqualDecimal(t, "0.000000000000000000", c.FS0.X0)
		IsEqualDecimal(t, "0.012826277713841738", c.FS0.Y)

		IsEqualDecimal(t, "0.012826277713841738", c.S0.P0Y)
		IsEqualDecimal(t, "0.012826277713841738", c.S0.A)
		IsEqualDecimal(t, "0.021543649942457564", c.S0.B)
		IsEqualDecimal(t, "0.030261022171073390", c.S0.P1Y)
		IsEqualDecimal(t, "0.000000000000000000", c.S0.P0X)
		IsEqualDecimal(t, "1891524601.156705919616317034", c.S0.P1X)
		IsEqualDecimal(t, "1891524601.156705919616317034", c.S0.DeltaX)

		IsEqualDecimal(t, "0.030261022171073390", c.S1.P0Y)
		IsEqualDecimal(t, "0.035369595779934340", c.S1.A)
		IsEqualDecimal(t, "0.063050820038556864", c.S1.B)
		IsEqualDecimal(t, "0.100000000000000000", c.S1.P1Y)
		IsEqualDecimal(t, "1891524601.156705919616317034", c.S1.P0X)
		IsEqualDecimal(t, "3000000000.000000000000000000", c.S1.P1X)
		IsEqualDecimal(t, "1108475398.843294080383682966", c.S1.DeltaX)

		IsEqualDecimal(t, "0.100000000000000000", c.S2.P0Y)
		IsEqualDecimal(t, "0.200000000000000000", c.S2.A)
		IsEqualDecimal(t, "0.333333333000000001", c.S2.B)
		IsEqualDecimal(t, "1.000000000000000000", c.S2.P1Y)
		IsEqualDecimal(t, "3000000000.000000000000000000", c.S2.P0X)
		IsEqualDecimal(t, "6000000000.000000000000000000", c.S2.P1X)
		IsEqualDecimal(t, "3000000000.000000000000000000", c.S2.DeltaX)
		IsEqualDecimal(t, "3000000000.000000000000000000", c.S2.IntervalP0X)

		IsEqualDecimal(t, "0.533333334000000000", c.QS3.A)
		IsEqualDecimal(t, "-5.733333341000000000", c.QS3.B)
		IsEqualDecimal(t, "16.200000022000000000", c.QS3.C)
		IsEqualDecimal(t, "1000000000.000000000000000000", c.QS3.ScalingFactor)
		IsEqualDecimal(t, "6000000000.000000000000000000", c.QS3.InitialX0)
		IsEqualDecimal(t, "6000000000.000000000000000000", c.QS3.CurrentX0)
	})

	t.Run("alternate set of valid params", func(t *testing.T) {
		type test struct {
			name     string
			modifier func(c *Curve)
		}

		tests := []test{
			{"BPoolUnder_alternate1",
				func(c *Curve) { c.BPoolUnder = sdk.NewDec(150e6) }},
			{"BPoolUnder_alternate2",
				func(c *Curve) { c.BPoolUnder = sdk.NewDec(90e6) }},
			{"BPool_notSet", // Because BPool is not used in Fit.
				func(c *Curve) { c.BPool = sdk.Dec{} }},
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
