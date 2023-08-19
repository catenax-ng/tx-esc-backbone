// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Sell sells the given amount of tokens against the curve.
// It returns the amount of vouchers released.
//
// It assumes the value of tokens is greater than zero. This condition is
// implemented in the ValidBasic function for the buy message.
func (c *Curve) Sell(tokens sdk.Dec) sdk.Dec {
	xCurrent := c.CurrentSupply
	xNew := c.CurrentSupply.Sub(tokens)

	vouchersOut := c.integralX12(xNew, xCurrent)
	vouchersOut = roundOff(vouchersOut, VoucherMultiplier)

	c.CurrentSupply = xNew
	c.BPool = c.BPool.Sub(vouchersOut)
	// CLARIFY: Should we change BPoolUnder
	return vouchersOut
}
