// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (sg *Segment) setP0X(P0X sdk.Dec) {
	sg.P0X = P0X
	sg.calcDeltaX()
}

func (sg *Segment) setP1X(P1X sdk.Dec) {
	sg.P1X = P1X
	sg.calcDeltaX()
}

func (sg *Segment) calcDeltaX() {
	sg.DeltaX = sg.P1X.Sub(sg.P0X)
}

func (sg *Segment) firstDerivativeT(t sdk.Dec) (y sdk.Dec) {
	return sg.gamma1T(t)

}

func (sg *Segment) gamma1T(t sdk.Dec) (y sdk.Dec) {
	Pi := sdk.NewDec(-1).Mul(sdk.NewDec(1).Sub(t).Power(2)).Mul(sg.P0)
	ai := sdk.NewDec(1).Sub(sdk.NewDec(3).Mul(t)).Mul(sdk.NewDec(1).Sub(t)).Mul(sg.A)
	bi := t.Mul(sdk.NewDec(2).Sub(sdk.NewDec(3).Mul(t))).Mul(sg.B)
	Pi1 := t.Power(2).Mul(sg.P1)
	return sdk.NewDec(3).Mul(Pi.Add(ai).Add(bi).Add(Pi1))
}
