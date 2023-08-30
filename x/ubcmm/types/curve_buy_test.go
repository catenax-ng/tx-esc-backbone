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

func Test_Curve_Buy(t *testing.T) {
	c := validCurve()
	require.NoError(t, c.Fit())

	initialBPoolUnder := c.BPoolUnder
	initialBPool := c.BPool
	initialCurrentSupply := c.CurrentSupply

	tokensToBuy := sdk.NewDecWithPrec(10000, 6)
	vouchersUsed := c.Buy(tokensToBuy)

	assert.Equal(t, c.BPoolUnder, initialBPoolUnder)
	assert.Equal(t, c.BPool, initialBPool.Add(vouchersUsed))
	assert.Equal(t, c.CurrentSupply, initialCurrentSupply.Add(tokensToBuy))
}

func BenchmarkCurveBuy(b *testing.B) {
	c := validCurve()
	c.Fit()

	tokens := sdk.NewDecWithPrec(10000, 6)
	for i := 0; i < b.N; i++ {
		c.Buy(tokens)
	}
}

func TestRoundOff(t *testing.T) {
	testCases := []struct {
		name       string
		input      string
		multiplier int64
		expected   string
		difference string
	}{
		{"Positive: Small Fraction, Round Down", "1.2345", 100, "1.2300", "0.0045"},
		{"Positive: Whole Number, No Rounding", "5.0000", 100, "5.0000", "0.0000"},
		{"Positive: Half, No Rounding", "1.5000", 100, "1.5000", "0.0000"},
		{"Positive: Larger Fraction, No Rounding", "2.5000", 100, "2.5000", "0.0000"},
		{"Positive: Large Whole Number, No Rounding", "3.5000", 100, "3.5000", "0.0000"},

		// Negative input cases
		{"Negative: Small Fraction, Round Down", "-6.789", 100, "-6.790", "0.0010"},
		{"Negative: Small Fraction, No Rounding", "-6.789", 1000, "-6.789", "0.0000"},

		// Cases with banker's rounding ties
		{"Tie: Small Fraction, Round Down", "1.2345", 1000, "1.2340", "0.0005"},
		{"Tie: Small Fraction, Round Down", "1.225", 100, "1.220", "0.005"},
		{"Tie: Small Fraction, Round Up", "1.275", 100, "1.280", "-0.005"},
		{"Tie: Small Fraction, Round Down", "1.725", 100, "1.720", "0.005"},
		{"Tie: Small Fraction, Round Up", "1.775", 100, "1.780", "-0.005"},

		// Edge cases
		{"Zero Value", "0", 100, "0", "0"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			inputDec := sdk.MustNewDecFromStr(tc.input)
			expectedDec := sdk.MustNewDecFromStr(tc.expected)
			differenceDec := sdk.MustNewDecFromStr(tc.difference)

			c := Curve{
				NumericalErrorAccumulator: sdk.ZeroDec(),
			}
			rounded := c.ubcRoundOff(inputDec, tc.multiplier)
			assert.True(t, rounded.Equal(expectedDec))
			assert.True(t, c.NumericalErrorAccumulator.Equal(differenceDec))
		})

	}
}
