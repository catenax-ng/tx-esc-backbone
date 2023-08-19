// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// integralX12 computes the integral of the curve segment with respect to "x",
// between limits x1 and x2, in the scaled domain.
func (qsg *QuadraticSegment) integralX12(x1, x2 sdk.Dec) sdk.Dec {
	x1 = qsg.scaleX(x1)

	part1 := x1.Power(3).Mul(qsg.A.Quo(sdk.NewDec(3)))
	part2 := x1.Power(2).Mul(qsg.B.Quo(sdk.NewDec(2)))
	part3 := x1.Mul(qsg.C)
	integralX12 := part1.Add(part2).Add(part3)

	return qsg.deScaleX(integralX12)

}

// firstDerivativeX1 computes the first derivate of the curve segment with respect to
// the "x", at the point x1.
func (qsg *QuadraticSegment) firstDerivativeX1(x1 sdk.Dec) (y sdk.Dec) {
	x1 = qsg.scaleX(x1)
	firstDerivativeX1 := sdk.NewDec(2).Mul(qsg.A).Mul(x1).Add(qsg.B)
	return qsg.deScaleX(firstDerivativeX1)
}

// y returns the y value for the given x.
func (qsg *QuadraticSegment) y(x1 sdk.Dec) (y sdk.Dec) {
	x1 = qsg.scaleX(x1)
	return qsg.A.Mul(x1.Power(2)).Add(qsg.B.Mul(x1)).Add(qsg.C)
}

func (qsg *QuadraticSegment) scaleX(x1 sdk.Dec) sdk.Dec {
	return x1.Quo(qsg.ScalingFactor)
}

func (qsg *QuadraticSegment) deScaleX(x1 sdk.Dec) sdk.Dec {
	return x1.Mul(qsg.ScalingFactor)
}
