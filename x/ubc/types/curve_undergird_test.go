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

func Test_Curve_Undergird(t *testing.T) {
	ubc := validUbcParams()
	require.NoError(t, ubc.Fit())

	// Check the paramters before undergirding
	IsEqualDecimal(t, "0.000000000000000000", ubc.FS0.X0)
	IsEqualDecimal(t, "0.012826277713841738", ubc.FS0.Y)

	IsEqualDecimal(t, "0.012826277713841738", ubc.S0.P0Y)
	IsEqualDecimal(t, "0.012826277713841738", ubc.S0.A)
	IsEqualDecimal(t, "0.021543649942457564", ubc.S0.B)
	IsEqualDecimal(t, "0.030261022171073390", ubc.S0.P1Y)
	IsEqualDecimal(t, "0.000000000000000000", ubc.S0.P0X)
	IsEqualDecimal(t, "1891524601.156705919616317034", ubc.S0.P1X)
	IsEqualDecimal(t, "1891524601.156705919616317034", ubc.S0.DeltaX)

	IsEqualDecimal(t, "0.030261022171073390", ubc.S1.P0Y)
	IsEqualDecimal(t, "0.035369595779934340", ubc.S1.A)
	IsEqualDecimal(t, "0.063050820038556864", ubc.S1.B)
	IsEqualDecimal(t, "0.100000000000000000", ubc.S1.P1Y)
	IsEqualDecimal(t, "1891524601.156705919616317034", ubc.S1.P0X)
	IsEqualDecimal(t, "3000000000.000000000000000000", ubc.S1.P1X)
	IsEqualDecimal(t, "1108475398.843294080383682966", ubc.S1.DeltaX)

	IsEqualDecimal(t, "0.100000000000000000", ubc.S2.P0Y)
	IsEqualDecimal(t, "0.200000000000000000", ubc.S2.A)
	IsEqualDecimal(t, "0.333333333000000001", ubc.S2.B)
	IsEqualDecimal(t, "1.000000000000000000", ubc.S2.P1Y)
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

	require.NoError(t, ubc.UndergirdS01(sdk.NewDec(100e5)))

	// Check the paramters after undergirding
	// Parameters that don't change
	IsEqualDecimal(t, "110000000", ubc.BPool)
	IsEqualDecimal(t, "110000000", ubc.BPoolUnder)
	IsEqualDecimal(t, "3000000000.000000000000000000", ubc.S1.P1X)
	IsEqualDecimal(t, "0.100000000000000000", ubc.S2.P0Y)
	IsEqualDecimal(t, "0.200000000000000000", ubc.S2.A)
	IsEqualDecimal(t, "0.333333333000000001", ubc.S2.B)
	IsEqualDecimal(t, "1.000000000000000000", ubc.S2.P1Y)
	IsEqualDecimal(t, "3000000000.000000000000000000", ubc.S2.P0X)
	IsEqualDecimal(t, "6000000000.000000000000000000", ubc.S2.P1X)
	IsEqualDecimal(t, "3000000000.000000000000000000", ubc.S2.DeltaX)
	IsEqualDecimal(t, "3000000000.000000000000000000", ubc.S2.IntervalP0X)
	IsEqualDecimal(t, "0.000000000000000000", ubc.FS0.X0)
	IsEqualDecimal(t, "0.000000000000000000", ubc.S0.P0X)
	IsEqualDecimal(t, "0.100000000000000000", ubc.S1.P1Y)

	// Parameters that change
	IsEqualDecimal(t, "0.018376535626862294", ubc.FS0.Y)
	IsEqualDecimal(t, "0.018376535626862294", ubc.S0.P0Y)
	IsEqualDecimal(t, "0.018376535626862294", ubc.S0.A)
	IsEqualDecimal(t, "0.026538882064176064", ubc.S0.B)
	IsEqualDecimal(t, "0.034701228501489835", ubc.S0.P1Y)
	// IsEqualDecimal(t, "1966150664.271770977111675239", ubc.S0.P1X)
	IsEqualDecimal(t, "1966150664.271770980278195239", ubc.S0.P1X) // Alternate

	// IsEqualDecimal(t, "1966150664.271770977111675239", ubc.S0.DeltaX)
	IsEqualDecimal(t, "1966150664.271770980278195239", ubc.S0.DeltaX) // Alternate
	IsEqualDecimal(t, "0.034701228501489835", ubc.S1.P0Y)
	IsEqualDecimal(t, "0.038993186689407263", ubc.S1.A)
	IsEqualDecimal(t, "0.065538355475725699", ubc.S1.B)
	// IsEqualDecimal(t, "1966150664.271770977111675239", ubc.S1.P0X)
	IsEqualDecimal(t, "1966150664.271770980278195239", ubc.S1.P0X) // Alternate
	// IsEqualDecimal(t, "1033849335.728229022888324761", ubc.S1.DeltaX)
	IsEqualDecimal(t, "1033849335.728229019721804761", ubc.S1.DeltaX) // Alternate

}
