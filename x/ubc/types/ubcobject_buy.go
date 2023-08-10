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
func (ubc *Ubcobject) Buy(tokens sdk.Dec) sdk.Dec {
	xCurrent := ubc.CurrentSupply
	xNew := ubc.CurrentSupply.Add(tokens)

	vouchersUsed := ubc.integralX12(xCurrent, xNew)
	vouchersUsed = roundOff(vouchersUsed, VoucherMultiplier)

	ubc.CurrentSupply = xNew
	ubc.BPool = ubc.BPool.Add(vouchersUsed)
	// CLARIFY: Should we change BPoolUnder
	return vouchersUsed
}

// term is a term in a polynomial equation.
type term struct {
	coefficient int64
	exponent    uint64
}

// computePolyFor returns the value of polynomial p(x) constructed using the given
// terms for the point x1.
//
// Eg: "computePolyFor(2, []term{{36, 2}, {-48, 1}, {12, 0}})" returns the value of
// "36(x^2) - 48x + 12" at x=2.
func computePolyFor(x1 sdk.Dec, terms []term) sdk.Dec {
	// We don't use powers > 4. If we do in future, this will err and we can fix it.
	const maxPow = 4
	x1Pows := [maxPow + 1]sdk.Dec{}
	x1Pows[0] = sdk.OneDec()
	for i := 1; i <= maxPow; i++ {
		x1Pows[i] = x1Pows[i-1].Mul(x1)
	}

	sum := sdk.ZeroDec()
	for _, term := range terms {
		sum = sum.Add(sdk.NewDec(term.coefficient).Mul(x1Pows[term.exponent]))
	}
	return sum
}

func roundOff(t sdk.Dec, multiplier int64) sdk.Dec {
	// CLARIFY: Is the rounding off strategy correct ?
	return t.MulInt64(VoucherMultiplier).
		TruncateDec().
		QuoInt64(VoucherMultiplier)
}
