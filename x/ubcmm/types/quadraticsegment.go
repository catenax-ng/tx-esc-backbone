// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// integralX12 computes the integral of the curve segment with respect to "x",
// between limits x1 and x2, in the scaled domain.
func (fqseg *FixedQuadraticSegment) integralX12(x1, x2 sdk.Dec) sdk.Dec {
	x1 = fqseg.scaleX(x1)

	part1 := x1.Power(3).Mul(fqseg.A.Quo(sdk.NewDec(3)))
	part2 := x1.Power(2).Mul(fqseg.B.Quo(sdk.NewDec(2)))
	part3 := x1.Mul(fqseg.C)
	integralX12 := part1.Add(part2).Add(part3)

	return fqseg.deScaleX(integralX12)

}

// firstDerivativeX1 computes the first derivate of the curve segment with respect to
// the "x", at the point x1.
func (fqseg *FixedQuadraticSegment) firstDerivativeX1(x1 sdk.Dec) (y sdk.Dec) {
	x1 = fqseg.scaleX(x1)
	firstDerivativeX1 := sdk.NewDec(2).Mul(fqseg.A).Mul(x1).Add(fqseg.B)
	return fqseg.deScaleX(firstDerivativeX1)
}

// y returns the y value for the given x.
func (fqseg *FixedQuadraticSegment) y(x1 sdk.Dec) (y sdk.Dec) {
	x1 = fqseg.scaleX(x1)
	return fqseg.A.Mul(x1.Power(2)).Add(fqseg.B.Mul(x1)).Add(fqseg.C)
}

func (fqseg *FixedQuadraticSegment) scaleX(x1 sdk.Dec) sdk.Dec {
	return x1.Quo(fqseg.ScalingFactor)
}

func (fqseg *FixedQuadraticSegment) deScaleX(x1 sdk.Dec) sdk.Dec {
	return x1.Mul(fqseg.ScalingFactor)
}

var _ view = (*FixedQuadraticSegment)(nil)

// startX returns the x-value of the start of the visible part of the curve.
func (fqseg *FixedQuadraticSegment) startX() sdk.Dec {
	return fqseg.CurrentX0
}

// endX returns the x-value of the end of the visible part of the curve.
func (fqseg *FixedQuadraticSegment) endX() sdk.Dec {
	return sdk.NewDec(1e17)
}

// startY returns the y-value of the start of the visible part of the line
func (fqseg *FixedQuadraticSegment) startY() sdk.Dec {
	return fqseg.y(fqseg.CurrentX0)
}

// endY returns the y-value of the end of the visible part of the line.
func (fqseg *FixedQuadraticSegment) endY() sdk.Dec {
	return sdk.NewDec(1e17)
}
