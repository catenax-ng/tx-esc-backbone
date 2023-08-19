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

// setP0 sets the value of parameter P0, only if it was not already set (it is
// zero).
func (fbseg *FixedBezierSegment) setP0(P0 sdk.Dec) {
	// CLARIFY: Reasoning for this if condition.
	if fbseg.P0.IsZero() {
		fbseg.P0 = P0
	}
}

// curvatureAtEnd computes curvature at the end of the curve segment.
func (fbseg *FixedBezierSegment) curvatureAtEnd() sdk.Dec {

	secondDerivativeT1 := func(fbseg *FixedBezierSegment, t sdk.Dec) (y sdk.Dec) {
		Pi := sdk.NewDec(1).Sub(t).Mul(fbseg.P0)
		ai := sdk.NewDec(3).Mul(t).Sub(sdk.NewDec(2)).Mul(fbseg.A)
		bi := sdk.NewDec(1).Sub(sdk.NewDec(3).Mul(t)).Mul(fbseg.B)
		Pi1 := t.Mul(fbseg.P1)
		return sdk.NewDec(6).Mul(Pi.Add(ai).Add(bi).Add(Pi1))
	}

	t1 := fbseg.t(fbseg.P1X)
	// CLARIFY: Reference for this calculation.
	return secondDerivativeT1(fbseg, t1).Quo(fbseg.DeltaX)
}
