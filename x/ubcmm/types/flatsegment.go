// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// integralX12 computes the integral of the curve segment with respect to "x",
// between limits x1 and x2.
func (fseg *FlatSegment) integralX12(x1, x2 sdk.Dec) sdk.Dec {
	return fseg.integralX1(x2).Sub(fseg.integralX1(x1))
}

// integralX1 computes the integral of the curve segment with respect to "x",
// from the beginning of the curve until point x1.
func (fseg *FlatSegment) integralX1(x1 sdk.Dec) sdk.Dec {
	return fseg.Y.Mul(x1)
}

// y returns the y value for the given x.
func (fseg *FlatSegment) y(x sdk.Dec) sdk.Dec {
	return fseg.Y
}

var _ view = (*FlatSegment)(nil)

// startX returns the x-value of the start of the visible part of the line
func (fseg *FlatSegment) startX() sdk.Dec {
	return sdk.ZeroDec() // TODO: Check if this is correct.
}

// endX returns the x-value of the end of the visible part of the line.
func (fseg *FlatSegment) endX() sdk.Dec {
	return fseg.P1X
}

// startY returns the y-value of the start of the visible part of the line
func (fseg *FlatSegment) startY() sdk.Dec {
	return fseg.Y
}

// endY returns the y-value of the end of the visible part of the line.
func (fseg *FlatSegment) endY() sdk.Dec {
	return fseg.Y
}
