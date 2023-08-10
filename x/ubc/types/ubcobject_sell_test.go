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

func Test_UbcObject_Sell(t *testing.T) {
	ubc := validUbcParams()
	require.NoError(t, ubc.Fit())

	initialBPoolUnder := ubc.BPoolUnder
	initialBPool := ubc.BPool
	initialCurrentSupply := ubc.CurrentSupply

	tokensToSell := sdk.NewDecWithPrec(10000, 6)
	vouchersOut := ubc.Sell(tokensToSell)

	assert.Equal(t, ubc.BPoolUnder, initialBPoolUnder)
	assert.Equal(t, ubc.BPool, initialBPool.Sub(vouchersOut))
	assert.Equal(t, ubc.CurrentSupply, initialCurrentSupply.Sub(tokensToSell))
}
