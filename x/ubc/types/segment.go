// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// setP0X sets the value of parameter P0X.
func (sg *Segment) setP0X(P0X sdk.Dec) {
	sg.P0X = P0X
	sg.updateDeltaX()
}

// setP1X sets the value of parameter P1X.
func (sg *Segment) setP1X(P1X sdk.Dec) {
	sg.P1X = P1X
	sg.updateDeltaX()
}

// updateDeltaX updates the value of parameter "DeltaX", based on the newer
// values of P0X and P1X.
//
// This function should be called each time P0X or P1X is modified.
func (sg *Segment) updateDeltaX() {
	sg.DeltaX = sg.P1X.Sub(sg.P0X)
}

// firstDerivativeT1 computes the first derivate of the curve segment with respect to
// the bezier curve parameter "t", at the point t1.
//
// The caller should ensure t1 lies within the range of t values for which the
// curve segment is defined.
func (sg *Segment) firstDerivativeT1(t1 sdk.Dec) sdk.Dec {
	Pi := sdk.NewDec(-1).Mul(sdk.NewDec(1).Sub(t1).Power(2)).Mul(sg.P0)
	ai := sdk.NewDec(1).Sub(sdk.NewDec(3).Mul(t1)).Mul(sdk.NewDec(1).Sub(t1)).Mul(sg.A)
	bi := t1.Mul(sdk.NewDec(2).Sub(sdk.NewDec(3).Mul(t1))).Mul(sg.B)
	Pi1 := t1.Power(2).Mul(sg.P1)
	return sdk.NewDec(3).Mul(Pi.Add(ai).Add(bi).Add(Pi1))
}

// integralX12 computes the integral of the curve segment with respect to "x",
// between limits x1 and x2.
func (s *Segment) integralX12(x1, x2 sdk.Dec) sdk.Dec {
	// CLARIFY: Moving out 0.25 creates a computation error in the 10th decimal place.
	integralX1 := sdk.NewDecWithPrec(25, 2).Mul(s.integralT1(s.t(x1)))
	integralX2 := sdk.NewDecWithPrec(25, 2).Mul(s.integralT1(s.t(x2)))
	return s.DeltaX.Mul(integralX2.Sub(integralX1))
}

// t computes the value of t (bezier curve parameter) for x1.
func (s *Segment) t(x1 sdk.Dec) sdk.Dec {
	return x1.Sub(s.P0X).Quo(s.DeltaX)
}

// integralT1 computes the integral of the curve segment with respect to the
// bezier curve parameter "t", from the beginning of the curve until point t1.
func (s *Segment) integralT1(t1 sdk.Dec) sdk.Dec {
	Pi := computePolyFor(t1, []term{{-1, 4}, {4, 3}, {-6, 2}, {4, 1}}).Mul(s.P0)
	ai := computePolyFor(t1, []term{{3, 4}, {-8, 3}, {6, 2}}).Mul(s.A)
	bi := computePolyFor(t1, []term{{3, 4}, {-4, 3}}).Mul(s.B)
	Pi1 := t1.Power(4).Mul(s.P1)

	return (Pi.Add(ai).Sub(bi).Add(Pi1))
}
