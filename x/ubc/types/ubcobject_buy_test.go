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

func Test_UbcObject_BuyWithVouchers(t *testing.T) {
	ubc := validUbcParams()
	require.NoError(t, ubc.Fit())

	initialBPoolUnder := ubc.BPoolUnder
	initialBPool := ubc.BPool
	initialCurrentSupply := ubc.CurrentSupply

	tokensToBuy := sdk.NewDecWithPrec(10000, 6)
	vouchersUsed := ubc.BuyExactTokens(tokensToBuy)

	assert.Equal(t, ubc.BPoolUnder, initialBPoolUnder)
	assert.Equal(t, ubc.BPool, initialBPool.Add(vouchersUsed))
	assert.Equal(t, ubc.CurrentSupply, initialCurrentSupply.Add(tokensToBuy))
}

func Test_UbcObject_BuyTokensFor(t *testing.T) {
	ubc := validUbcParams()
	require.NoError(t, ubc.Fit())

	initialBPoolUnder := ubc.BPoolUnder
	initialBPool := ubc.BPool
	initialCurrentSupply := ubc.CurrentSupply

	tokens, vouchersUsed, err := ubc.BuyTokensFor(sdk.NewDecWithPrec(9970, 6))
	require.NoError(t, err)

	assert.Equal(t, ubc.BPoolUnder, initialBPoolUnder)
	assert.Equal(t, ubc.BPool, initialBPool.Add(vouchersUsed))
	assert.Equal(t, ubc.CurrentSupply, initialCurrentSupply.Add(tokens))
}

func BenchmarkUbcBuyNTokensForVouchers(b *testing.B) {
	ubc := validUbcParams()
	ubc.Fit()

	tokens := sdk.NewDecWithPrec(10000, 6)
	for i := 0; i < b.N; i++ {
		ubc.BuyExactTokens(tokens)
	}
}

func BenchmarkUbcBuyTokensForNVouchers(b *testing.B) {
	ubc := validUbcParams()
	ubc.Fit()

	vouchersIn := sdk.NewDecWithPrec(9970, 6)
	for i := 0; i < b.N; i++ {
		ubc.BuyTokensFor(vouchersIn)
	}
}
