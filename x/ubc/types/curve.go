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

func (ubc *Curve) p0Y() sdk.Dec {
	return ubc.S0.P0Y
}

func (ubc *Curve) p1Y() sdk.Dec {
	return ubc.S1.P0Y
}

func (ubc *Curve) p2Y() sdk.Dec {
	return ubc.S2.P0Y
}

func (ubc *Curve) p3Y() sdk.Dec {
	return ubc.S2.P1Y
}

func (ubc *Curve) p0x() sdk.Dec {
	return ubc.S0.P0X
}
func (ubc *Curve) p1x() sdk.Dec {
	return ubc.S1.P0X
}
func (ubc *Curve) p2x() sdk.Dec {
	return ubc.S2.IntervalP0X
}

func (ubc *Curve) p3x() sdk.Dec {
	return ubc.S2.P1X
}

func (ubc *Curve) setP0Y(p0Y sdk.Dec) {
	ubc.FS0.Y = p0Y
	ubc.S0.P0Y = p0Y
}

func (ubc *Curve) setP1Y(p1Y sdk.Dec) {
	ubc.S0.P1Y = p1Y
	ubc.S1.P0Y = p1Y
}

func (ubc *Curve) setP2Y(p2Y sdk.Dec) {
	ubc.S1.P1Y = p2Y
	ubc.S2.setP0Y(p2Y)
}

func (ubc *Curve) setP3Y(p3Y sdk.Dec) {
	ubc.S2.P1Y = p3Y
}

func (ubc *Curve) setP0X(p0X sdk.Dec) {
	ubc.FS0.X0 = p0X
	ubc.S0.setP0X(p0X)
}

func (ubc *Curve) setP1X(p1X sdk.Dec) {
	ubc.S0.setP1X(p1X)
	ubc.S1.setP0X(p1X)
}

func (ubc *Curve) setP2X(p2X sdk.Dec) {
	ubc.S1.setP1X(p2X)
	ubc.S2.setP0X(p2X)
}

func (ubc *Curve) setP3X(p3X sdk.Dec) {
	ubc.S2.setP1X(p3X)
}

// segmentNum returns the segment number for the given point x.
func (ubc *Curve) segmentNum(x sdk.Dec) int {
	upperBoundsX := []sdk.Dec{ubc.p0x(), ubc.p1x(), ubc.p2x(), ubc.p3x()}
	segments := []int{FS0, S0, S1, S2}

	for i, upperBound := range upperBoundsX {
		if x.LT(upperBound) {
			return segments[i]
		}
	}
	return QS3
}

func (ubc *Curve) a(segNum int) sdk.Dec {
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

func (ubc *Curve) b(segNum int) sdk.Dec {
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

func (ubc *Curve) upperBoundX(segNum int) sdk.Dec {
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

func (ubc *Curve) lowerBoundX(segNum int) sdk.Dec {
	switch segNum {
	case S0:
		return ubc.p0x()
	case S1:
		return ubc.p1x()
	case S2:
		return ubc.p2x()
	case QS3:
		return ubc.p3x()
	default:
		return sdk.ZeroDec()
	}
}

func (ubc *Curve) lowerBound(segNum int) sdk.Dec {
	switch segNum {
	case S0:
		return ubc.p0Y()
	case S1:
		return ubc.p1Y()
	case S2:
		return ubc.p2Y()
	default:
		return sdk.ZeroDec()
	}
}

func (ubc *Curve) upperBound(segNum int) sdk.Dec {
	switch segNum {
	case S0:
		return ubc.p1Y()
	case S1:
		return ubc.p2Y()
	case S2:
		return ubc.p3Y()
	default:
		return sdk.ZeroDec()
	}
}

func (ubc *Curve) t(segNum int) func(x sdk.Dec) (t sdk.Dec) {
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

func (ubc *Curve) deltaX(segNum int) sdk.Dec {
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

// integralX12 returns the function for computing the integral for the given
// segment.
func (ubc *Curve) integralXFn(segNum int) func(x1, x2 sdk.Dec) sdk.Dec {
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

func (ubc *Curve) integralX12(lowerBoundX, upperBoundX sdk.Dec) (vouchers sdk.Dec) {
	segLowerBoundX := ubc.segmentNum(lowerBoundX)
	segUpperBoundX := ubc.segmentNum(upperBoundX)

	vouchers = sdk.NewDec(0)
	for ; segLowerBoundX <= segUpperBoundX; segLowerBoundX = segLowerBoundX + 1 {
		x1 := lowerBoundX
		x2 := ubc.upperBoundX(segLowerBoundX)
		if segLowerBoundX == segUpperBoundX {
			x2 = upperBoundX
		}
		additionalVouchers := ubc.integralXFn(segLowerBoundX)(x1, x2)
		vouchers = vouchers.Add(additionalVouchers)

		lowerBoundX = ubc.upperBoundX(segLowerBoundX)
	}
	return vouchers
}

func (ubc *Curve) slopeX1(x1 sdk.Dec) sdk.Dec {
	switch ubc.segmentNum(x1) {
	case FS0:
		return sdk.ZeroDec()
	case S0:
		return ubc.S0.firstDerivativeX1(x1)
	case S1:
		return ubc.S1.firstDerivativeX1(x1)
	case S2:
		return ubc.S2.firstDerivativeX1(x1)
	case QS3:
		return ubc.QS3.firstDerivativeX1(x1)
	default:
		return sdk.ZeroDec()
	}
}

func (ubc *Curve) y(x sdk.Dec) sdk.Dec {
	switch ubc.segmentNum(x) {
	case FS0:
		return ubc.FS0.y(x)
	case S0:
		return ubc.S0.y(x)
	case S1:
		return ubc.S1.y(x)
	case S2:
		return ubc.S2.y(x)
	case QS3:
		return ubc.QS3.y(x)
	default:
		return sdk.ZeroDec()
	}
}

// integralT1 computes the integral of the curve segment with respect to the
// bezier curve parameter "t", from the beginning of the curve until point t1.
func (ubc *Curve) integralT1(t1 sdk.Dec, seg int) sdk.Dec {
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
