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

// firstDerivativeX1 computes the first derivate of the curve segment with respect to
// the "x", at the point x1.
func (sg *Segment) firstDerivativeX1(x1 sdk.Dec) (y sdk.Dec) {
	return sg.firstDerivativeT1(sg.t(x1)).Quo(sg.DeltaX)
}

// firstDerivativeT1 computes the first derivate of the curve segment with respect to
// the bezier curve parameter "t", at the point t1.
func (sg *Segment) firstDerivativeT1(t1 sdk.Dec) sdk.Dec {
	Pi := sdk.NewDec(-1).Mul(sdk.NewDec(1).Sub(t1).Power(2)).Mul(sg.P0)
	ai := sdk.NewDec(1).Sub(sdk.NewDec(3).Mul(t1)).Mul(sdk.NewDec(1).Sub(t1)).Mul(sg.A)
	bi := t1.Mul(sdk.NewDec(2).Sub(sdk.NewDec(3).Mul(t1))).Mul(sg.B)
	Pi1 := t1.Power(2).Mul(sg.P1)
	return sdk.NewDec(3).Mul(Pi.Add(ai).Add(bi).Add(Pi1))
}

// y returns the y value for the given x.
func (sg *Segment) y(x sdk.Dec) sdk.Dec {
	t := sg.t(x)

	// math.Pow((1-t), 3) * sg.P0
	Pi := sdk.NewDec(1).Sub(t).Power(3).Mul(sg.P0)
	// 3 * t * math.Pow((1-t), 2) * sg.A
	ai := sdk.NewDec(3).Mul(t).Mul(sdk.NewDec(1).Sub(t).Power(2)).Mul(sg.A)
	// 3 * math.Pow(t, 2) * (1 - t) * sg.B
	bi := sdk.NewDec(3).Mul(t.Power(2)).Mul(sdk.NewDec(1).Sub(t)).Mul(sg.B)
	// math.Pow(t, 3) * sg.P1
	Pi1 := t.Power(3).Mul(sg.P1)

	return Pi.Add(ai).Add(bi).Add(Pi1)
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
