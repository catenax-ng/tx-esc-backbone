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

func Test_UbcObject_Fit_Happy(t *testing.T) {
	t.Run("primary set of valid params", func(t *testing.T) {
		ubc := validUbcParams()
		require.NoError(t, ubc.Fit())

		IsEqualDecimal(t, "0.000000000000000000", ubc.FS0.X0)
		IsEqualDecimal(t, "0.012826277713841738", ubc.FS0.Y)

		IsEqualDecimal(t, "0.012826277713841738", ubc.S0.P0)
		IsEqualDecimal(t, "0.012826277713841738", ubc.S0.A)
		IsEqualDecimal(t, "0.021543649942457564", ubc.S0.B)
		IsEqualDecimal(t, "0.030261022171073390", ubc.S0.P1)
		IsEqualDecimal(t, "0.000000000000000000", ubc.S0.P0X)
		IsEqualDecimal(t, "1891524601.156705919616317034", ubc.S0.P1X)
		IsEqualDecimal(t, "1891524601.156705919616317034", ubc.S0.DeltaX)

		IsEqualDecimal(t, "0.030261022171073390", ubc.S1.P0)
		IsEqualDecimal(t, "0.035369595779934340", ubc.S1.A)
		IsEqualDecimal(t, "0.063050820038556864", ubc.S1.B)
		IsEqualDecimal(t, "0.100000000000000000", ubc.S1.P1)
		IsEqualDecimal(t, "1891524601.156705919616317034", ubc.S1.P0X)
		IsEqualDecimal(t, "3000000000.000000000000000000", ubc.S1.P1X)
		IsEqualDecimal(t, "1108475398.843294080383682966", ubc.S1.DeltaX)

		IsEqualDecimal(t, "0.100000000000000000", ubc.S2.P0)
		IsEqualDecimal(t, "0.200000000000000000", ubc.S2.A)
		IsEqualDecimal(t, "0.333333333000000001", ubc.S2.B)
		IsEqualDecimal(t, "1.000000000000000000", ubc.S2.P1)
		IsEqualDecimal(t, "3000000000.000000000000000000", ubc.S2.P0X)
		IsEqualDecimal(t, "6000000000.000000000000000000", ubc.S2.P1X)
		IsEqualDecimal(t, "3000000000.000000000000000000", ubc.S2.DeltaX)
		IsEqualDecimal(t, "3000000000.000000000000000000", ubc.S2.IntervalP0X)

		IsEqualDecimal(t, "0.533333334000000000", ubc.QS3.A)
		IsEqualDecimal(t, "-5.733333341000000000", ubc.QS3.B)
		IsEqualDecimal(t, "16.200000022000000000", ubc.QS3.C)
		IsEqualDecimal(t, "1000000000.000000000000000000", ubc.QS3.ScalingFactor)
		assert.Zero(t, ubc.QS3.InitialX0)
		assert.Zero(t, ubc.QS3.CurrentX0)
	})

	t.Run("alternate set of valid params", func(t *testing.T) {
		type test struct {
			name     string
			modifier func(ubc *Curve)
		}

		tests := []test{
			{"BPoolUnder_alternate1",
				func(ubc *Curve) { ubc.BPoolUnder = sdk.NewDec(150e6) }},
			{"BPoolUnder_alternate2",
				func(ubc *Curve) { ubc.BPoolUnder = sdk.NewDec(90e6) }},
			{"BPool_notSet", // Because BPool is not used in Fit.
				func(ubc *Curve) { ubc.BPool = sdk.Dec{} }},
		}

		for _, tc := range tests {
			t.Run(tc.name, func(t *testing.T) {
				ubc := validUbcParams()
				tc.modifier(&ubc)
				assert.NoError(t, ubc.Fit())
			})
		}
	})
}

func Test_UbcObject_Fit_Error(t *testing.T) {
	type test struct {
		name     string
		modifier func(ubc *Curve)
	}

	tests := []test{
		{"RefTokenSupply_isNotSet",
			func(ubc *Curve) { ubc.RefTokenSupply = sdk.Dec{} }},
		{"RefTokenPrice_isNotSet",
			func(ubc *Curve) { ubc.RefTokenPrice = sdk.Dec{} }},
		{"RefProfitFactor_isNotSet",
			func(ubc *Curve) { ubc.RefProfitFactor = sdk.Dec{} }},
		{"BPoolUnder_isNotSet",
			func(ubc *Curve) { ubc.BPoolUnder = sdk.Dec{} }},
		{"SlopeP2_isNotSet",
			func(ubc *Curve) { ubc.SlopeP2 = sdk.Dec{} }},
		{"SlopeP3_isNotSet",
			func(ubc *Curve) { ubc.SlopeP3 = sdk.Dec{} }},
		{"FactorFy_isNotSet",
			func(ubc *Curve) { ubc.FactorFy = sdk.Dec{} }},
		{"FactorFxy_isNotSet",
			func(ubc *Curve) { ubc.FactorFxy = sdk.Dec{} }},
		{"RefTokenSupply_isTooLess",
			func(ubc *Curve) { ubc.RefTokenSupply = sdk.NewDec(6e6) }},
		{"RefTokenSupply_isTooHigh",
			func(ubc *Curve) { ubc.RefTokenSupply = sdk.NewDec(6e10) }},
		{"RefTokenTokenPrice_isTooLess",
			func(ubc *Curve) { ubc.RefTokenPrice = sdk.NewDecWithPrec(1, 1) }},
		{"RefTokenTokenPrice_isTooHigh",
			func(ubc *Curve) { ubc.RefTokenPrice = sdk.NewDec(10) }},
		{"BPoolUnder_isTooLess",
			func(ubc *Curve) { ubc.BPoolUnder = sdk.NewDec(100e3) }},
		{"BPoolUnder_isTooHigh",
			func(ubc *Curve) { ubc.BPoolUnder = sdk.NewDec(100e10) }},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ubc := validUbcParams()
			tc.modifier(&ubc)
			assert.Error(t, ubc.Fit())
		})
	}
}
func BenchmarkUbcFit(b *testing.B) {
	ubc := validUbcParams()
	for i := 0; i < b.N; i++ {
		ubc.Fit()
	}
}

func IsEqualDecimal(t *testing.T, expected string, actual sdk.Dec) {
	t.Helper()
	expectedDec, err := sdk.NewDecFromStr(expected)
	require.NoError(t, err)
	assert.Equal(t, expectedDec, actual)
}

func validUbcParams() Curve {
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
