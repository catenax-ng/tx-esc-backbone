// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// setP0X sets the value of parameter P0X. It always sets the value of
// IntervalP0X and sets the value of P0X, only if it was not already set (it is
// zero). It also updates the value of DeltaX.
func (fbseg *FixedBezierSegment) setP0X(P0X sdk.Dec) {
	fbseg.IntervalP0X = P0X
	// If P0X is zero, then it was not set before, so set it.
	if fbseg.P0X.IsZero() {
		fbseg.P0X = P0X
	}
	fbseg.updateDeltaX()
}

// setP0Y sets the value of parameter P0, only if it was not already set (it is
// zero).
func (fbseg *FixedBezierSegment) setP0Y(P0 sdk.Dec) {
	// CLARIFY: Reasoning for this if condition.
	if fbseg.P0Y.IsZero() {
		fbseg.P0Y = P0
	}
}

// curvatureAtEnd computes curvature at the end of the curve segment.
func (fbseg *FixedBezierSegment) curvatureAtEnd() sdk.Dec {

	secondDerivativeT1 := func(fbseg *FixedBezierSegment, t sdk.Dec) (y sdk.Dec) {
		Pi := sdk.NewDec(1).Sub(t).Mul(fbseg.P0Y)
		ai := sdk.NewDec(3).Mul(t).Sub(sdk.NewDec(2)).Mul(fbseg.A)
		bi := sdk.NewDec(1).Sub(sdk.NewDec(3).Mul(t)).Mul(fbseg.B)
		Pi1 := t.Mul(fbseg.P1Y)
		return sdk.NewDec(6).Mul(Pi.Add(ai).Add(bi).Add(Pi1))
	}

	t1 := fbseg.t(fbseg.P1X)
	// CLARIFY: Reference for this calculation.
	return secondDerivativeT1(fbseg, t1).Quo(fbseg.DeltaX)
}

var _ view = (*FixedBezierSegment)(nil)

// startX returns the x-value of the start of the visible part of the curve.
func (fbseg *FixedBezierSegment) startX() sdk.Dec {
	return fbseg.IntervalP0X
}

// endX returns the x-value of the end of the visible part of the curve.
func (fbseg *FixedBezierSegment) endX() sdk.Dec {
	return fbseg.P1X
}

// startY returns the y-value of the start of the visible part of the line
func (fbseg *FixedBezierSegment) startY() sdk.Dec {
	return fbseg.y(fbseg.IntervalP0X)
}

// endY returns the y-value of the end of the visible part of the line.
func (fbseg *FixedBezierSegment) endY() sdk.Dec {
	return fbseg.P1Y
}
