// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (ubc *Ubcobject) p0() sdk.Dec {
	return ubc.S0.P0
}

func (ubc *Ubcobject) p1() sdk.Dec {
	return ubc.S1.P0
}

func (ubc *Ubcobject) p2() sdk.Dec {
	return ubc.S2.P0
}

func (ubc *Ubcobject) p3() sdk.Dec {
	return ubc.S2.P1
}

func (ubc *Ubcobject) p0x() sdk.Dec {
	return ubc.S0.P0X
}
func (ubc *Ubcobject) p1x() sdk.Dec {
	return ubc.S1.P0X
}
func (ubc *Ubcobject) p2x() sdk.Dec {
	return ubc.S2.IntervalP0X
}

func (ubc *Ubcobject) p3x() sdk.Dec {
	return ubc.S2.P1X
}

func (ubc *Ubcobject) setP0(p0 sdk.Dec) {
	ubc.FS0.Y = p0
	ubc.S0.P0 = p0
}

func (ubc *Ubcobject) setP1(p1 sdk.Dec) {
	ubc.S0.P1 = p1
	ubc.S1.P0 = p1
}

func (ubc *Ubcobject) setP2(p2 sdk.Dec) {
	ubc.S1.P1 = p2
	ubc.S2.setP0(p2)
}

func (ubc *Ubcobject) setP3(p3 sdk.Dec) {
	ubc.S2.P1 = p3
}

func (ubc *Ubcobject) setP0X(p0X sdk.Dec) {
	ubc.FS0.X0 = p0X
	ubc.S0.setP0X(p0X)
}

func (ubc *Ubcobject) setP1X(p1X sdk.Dec) {
	ubc.S0.setP1X(p1X)
	ubc.S1.setP0X(p1X)
}

func (ubc *Ubcobject) setP2X(p2X sdk.Dec) {
	ubc.S1.setP1X(p2X)
	ubc.S2.setP0X(p2X)
}

func (ubc *Ubcobject) setP3X(p3X sdk.Dec) {
	ubc.S2.setP1X(p3X)
}
