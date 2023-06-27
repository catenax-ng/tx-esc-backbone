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
func (fx *Fixedsegment) setP0X(P0X sdk.Dec) {
	fx.IntervalP0X = P0X
	// If P0X is zero, then it was not set before, so set it.
	if fx.P0X.IsZero() {
		fx.P0X = P0X
	}
	fx.updateDeltaX()
}

// setP0 sets the value of parameter P0, only if it was not already set (it is
// zero).
func (fx *Fixedsegment) setP0(P0 sdk.Dec) {
	// CLARIFY: Reasoning for this if condition.
	if fx.P0.IsZero() {
		fx.P0 = P0
	}
}

// curvatureAtEnd computes curvature at the end of the curve segment.
func (fx *Fixedsegment) curvatureAtEnd() sdk.Dec {

	secondDerivativeT1 := func(fx *Fixedsegment, t sdk.Dec) (y sdk.Dec) {
		Pi := sdk.NewDec(1).Sub(t).Mul(fx.P0)
		ai := sdk.NewDec(3).Mul(t).Sub(sdk.NewDec(2)).Mul(fx.A)
		bi := sdk.NewDec(1).Sub(sdk.NewDec(3).Mul(t)).Mul(fx.B)
		Pi1 := t.Mul(fx.P1)
		return sdk.NewDec(6).Mul(Pi.Add(ai).Add(bi).Add(Pi1))
	}

	t1 := fx.t(fx.P1X)
	// CLARIFY: Reference for this calculation.
	return secondDerivativeT1(fx, t1).Quo(fx.DeltaX)
}
