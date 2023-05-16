// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SellExactTokens sells the given amount of tokens against the curve.
// It returns the amount of vouchers released.
//
// It assumes the value of tokens is greater than zero. This condition is
// implemented in the ValidBasic function for the buy message.
func (ubc *Ubcobject) SellExactTokens(tokens sdk.Dec) sdk.Dec {
	xCurrent := ubc.CurrentSupply
	xNew := ubc.CurrentSupply.Sub(tokens)

	segXCurrent := ubc.segmentNum(xCurrent)
	segXNew := ubc.segmentNum(xNew)

	var vouchersOut = sdk.NewDec(0)
	for ; segXCurrent >= segXNew; segXCurrent = segXCurrent - 1 {
		x1 := ubc.lowerBoundX(segXCurrent)
		x2 := xCurrent
		if segXCurrent == segXNew || segXCurrent == FS0 {
			x1 = xNew
		}
		additionalVouchers := ubc.integralX12(segXCurrent)(x1, x2)
		vouchersOut = vouchersOut.Add(additionalVouchers)

		xCurrent = ubc.lowerBoundX(segXCurrent)
	}
	vouchersOut = roundOff(vouchersOut, VoucherMultiplier)

	ubc.CurrentSupply = xNew
	ubc.BPool = ubc.BPool.Sub(vouchersOut)
	// CLARIFY: Should we change BPoolUnder
	return vouchersOut
}
