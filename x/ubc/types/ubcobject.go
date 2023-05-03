// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	FS0 = iota
	S0
	S1
	S2
	QS3
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

// segmentNum returns the segment number for the given point x.
func (ubc *Ubcobject) segmentNum(x sdk.Dec) int {
	upperBounds := []sdk.Dec{ubc.p0x(), ubc.p1x(), ubc.p2x(), ubc.p3x()}
	segments := []int{FS0, S0, S1, S2}

	for i, upperBound := range upperBounds {
		if x.LT(upperBound) {
			return segments[i]
		}
	}
	return QS3
}

func (ubc *Ubcobject) a(segNum int) sdk.Dec {
	switch segNum {
	case S0:
		return ubc.S0.A
	case S1:
		return ubc.S1.A
	case S2:
		return ubc.S2.A
	default:
		return sdk.ZeroDec()
	}
}

func (ubc *Ubcobject) b(segNum int) sdk.Dec {
	switch segNum {
	case S0:
		return ubc.S0.B
	case S1:
		return ubc.S1.B
	case S2:
		return ubc.S2.B
	default:
		return sdk.ZeroDec()
	}
}

func (ubc *Ubcobject) upperBoundX(segNum int) sdk.Dec {
	switch segNum {
	case FS0:
		return ubc.p0x()
	case S0:
		return ubc.p1x()
	case S1:
		return ubc.p2x()
	case S2:
		return ubc.p3x()
	default:
		return sdk.ZeroDec()
	}
}

func (ubc *Ubcobject) lowerBound(segNum int) sdk.Dec {
	switch segNum {
	case S0:
		return ubc.p0()
	case S1:
		return ubc.p1()
	case S2:
		return ubc.p2()
	default:
		return sdk.ZeroDec()
	}
}

func (ubc *Ubcobject) upperBound(segNum int) sdk.Dec {
	switch segNum {
	case S0:
		return ubc.p1()
	case S1:
		return ubc.p2()
	case S2:
		return ubc.p3()
	default:
		return sdk.ZeroDec()
	}
}

func (ubc *Ubcobject) t(segNum int) func(x sdk.Dec) (t sdk.Dec) {
	switch segNum {
	case S0:
		return ubc.S0.t
	case S1:
		return ubc.S1.t
	case S2:
		return ubc.S2.t
	default:
		return func(x sdk.Dec) sdk.Dec { return sdk.ZeroDec() }
	}
}

func (ubc *Ubcobject) deltaX(segNum int) sdk.Dec {
	switch segNum {
	case S0:
		return ubc.S0.DeltaX
	case S1:
		return ubc.S1.DeltaX
	case S2:
		return ubc.S2.DeltaX
	default:
		return sdk.ZeroDec()
	}
}

// integralX12 computes the function for computing the integral for the given
// bezier segment.
func (ubc *Ubcobject) integralX12(segNum int) func(x1, x2 sdk.Dec) sdk.Dec {
	switch segNum {
	case FS0:
		return ubc.FS0.integralX12
	case S0:
		return ubc.S0.integralX12
	case S1:
		return ubc.S1.integralX12
	case S2:
		return ubc.S2.integralX12
	case QS3:
		return ubc.QS3.integralX12
	default:
		return func(sdk.Dec, sdk.Dec) sdk.Dec { return sdk.ZeroDec() }
	}
}

// integralT1 computes the integral of the curve segment with respect to the
// bezier curve parameter "t", from the beginning of the curve until point t1.
func (ubc *Ubcobject) integralT1(t1 sdk.Dec, seg int) sdk.Dec {
	switch seg {
	case S0:
		return ubc.S0.integralT1(t1)
	case S1:
		return ubc.S1.integralT1(t1)
	case S2:
		return ubc.S2.integralT1(t1)
	default:
		return sdk.ZeroDec()
	}
}
