// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Buy buys the given amount of tokens against the curve. It
// returns the amount of vouchers used.
//
// It assumes the value of tokens is greater than zero. This condition is
// implemented in the ValidBasic function for the buy message.
func (c *Curve) Buy(tokens sdk.Dec) sdk.Dec {
	xCurrent := c.CurrentSupply
	xNew := c.CurrentSupply.Add(tokens)

	vouchersUsed := c.integralX12(xCurrent, xNew)
	vouchersUsed = roundOff(vouchersUsed, VoucherMultiplier)

	c.CurrentSupply = xNew
	c.BPool = c.BPool.Add(vouchersUsed)
	// CLARIFY: Should we change BPoolUnder
	return vouchersUsed
}

func roundOff(t sdk.Dec, multiplier int64) sdk.Dec {
	// CLARIFY: Is the rounding off strategy correct ?
	return t.MulInt64(VoucherMultiplier).
		TruncateDec().
		QuoInt64(VoucherMultiplier)
}
