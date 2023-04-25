// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package types

import sdk "github.com/cosmos/cosmos-sdk/types"

func (fx *Fixedsegment) setP0X(P0X sdk.Dec) {
	fx.IntervalP0X = P0X
	// If P0X is zero, then it was not set before, so set it.
	if fx.P0X.IsZero() {
		fx.P0X = P0X
	}
	fx.calcDeltaX()
}

func (fx *Fixedsegment) setP0(P0 sdk.Dec) {
	if fx.P0.IsZero() {
		fx.P0 = P0
	}
}

func (fixsg *Fixedsegment) gamma2T(t sdk.Dec) (y sdk.Dec) {
	Pi := sdk.NewDec(1).Sub(t).Mul(fixsg.P0)
	ai := sdk.NewDec(3).Mul(t).Sub(sdk.NewDec(2)).Mul(fixsg.A)
	bi := sdk.NewDec(1).Sub(sdk.NewDec(3).Mul(t)).Mul(fixsg.B)
	Pi1 := t.Mul(fixsg.P1)

	return sdk.NewDec(6).Mul(Pi.Add(ai).Add(bi).Add(Pi1))
}

func (fx *Fixedsegment) secondDerivativeX(x sdk.Dec) (y sdk.Dec) {
	t := x.Sub(fx.P0X).Quo(fx.DeltaX)
	return fx.gamma2T(t).Quo(fx.DeltaX)
}
