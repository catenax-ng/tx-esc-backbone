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

func Test_Curve_Sell(t *testing.T) {
	c := validCurve()
	require.NoError(t, c.Fit())

	initialBPoolUnder := c.BPoolUnder
	initialBPool := c.BPool
	initialCurrentSupply := c.CurrentSupply

	tokensToSell := sdk.NewDecWithPrec(10000, 6)
	vouchersOut := c.Sell(tokensToSell)

	assert.Equal(t, c.BPoolUnder, initialBPoolUnder)
	assert.Equal(t, c.BPool, initialBPool.Sub(vouchersOut))
	assert.Equal(t, c.CurrentSupply, initialCurrentSupply.Sub(tokensToSell))
}
