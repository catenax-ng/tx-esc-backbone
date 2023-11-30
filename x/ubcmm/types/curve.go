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
	S0 = iota
	S1
	S2
	S3
	S4
)

func (c *Curve) p0Y() sdk.Dec {
	return c.S1.startY()
}

func (c *Curve) p1Y() sdk.Dec {
	return c.S2.startY()
}

func (c *Curve) p2Y() sdk.Dec {
	return c.S3.startY()
}

func (c *Curve) p3Y() sdk.Dec {
	return c.S4.startY()
}

func (c *Curve) p0x() sdk.Dec {
	return c.S1.startX()
}
func (c *Curve) p1x() sdk.Dec {
	return c.S2.startX()
}
func (c *Curve) p2x() sdk.Dec {
	return c.S3.startX()
}

func (c *Curve) p3x() sdk.Dec {
	return c.S4.startX()
}

func (c *Curve) setP0Y(p0Y sdk.Dec) {
	c.S0.Y = p0Y
	c.S1.P0Y = p0Y
}

func (c *Curve) setP1Y(p1Y sdk.Dec) {
	c.S1.P1Y = p1Y
	c.S2.P0Y = p1Y
}

func (c *Curve) setP2Y(p2Y sdk.Dec) {
	c.S2.P1Y = p2Y
	c.S3.setP0Y(p2Y)
}

func (c *Curve) setP3Y(p3Y sdk.Dec) {
	c.S3.P1Y = p3Y
}

func (c *Curve) setP0X(p0X sdk.Dec) {
	c.S0.P1X = p0X
	c.S1.setP0X(p0X)
}

func (c *Curve) setP1X(p1X sdk.Dec) {
	c.S1.setP1X(p1X)
	c.S2.setP0X(p1X)
}

func (c *Curve) setP2X(p2X sdk.Dec) {
	c.S2.setP1X(p2X)
	c.S3.setP0X(p2X)
}

func (c *Curve) setP3X(p3X sdk.Dec) {
	c.S3.setP1X(p3X)
	c.S4.setP0X(p3X)
}

// segmentNum returns the segment number for the given point x.
func (c *Curve) segmentNum(x sdk.Dec) int {
	upperBoundsX := []sdk.Dec{c.p0x(), c.p1x(), c.p2x(), c.p3x()}
	segments := []int{S0, S1, S2, S3}

	for i, upperBound := range upperBoundsX {
		if x.LT(upperBound) {
			return segments[i]
		}
	}
	return S4
}

func (c *Curve) a(segNum int) sdk.Dec {
	switch segNum {
	case S1:
		return c.S1.A
	case S2:
		return c.S2.A
	case S3:
		return c.S3.A
	default:
		return sdk.ZeroDec()
	}
}

func (c *Curve) b(segNum int) sdk.Dec {
	switch segNum {
	case S1:
		return c.S1.B
	case S2:
		return c.S2.B
	case S3:
		return c.S3.B
	default:
		return sdk.ZeroDec()
	}
}

func (c *Curve) upperBoundX(segNum int) sdk.Dec {
	switch segNum {
	case S0:
		return c.p0x()
	case S1:
		return c.p1x()
	case S2:
		return c.p2x()
	case S3:
		return c.p3x()
	default:
		return sdk.ZeroDec()
	}
}

func (c *Curve) lowerBoundX(segNum int) sdk.Dec {
	switch segNum {
	case S1:
		return c.p0x()
	case S2:
		return c.p1x()
	case S3:
		return c.p2x()
	case S4:
		return c.p3x()
	default:
		return sdk.ZeroDec()
	}
}

func (c *Curve) lowerBound(segNum int) sdk.Dec {
	switch segNum {
	case S1:
		return c.p0Y()
	case S2:
		return c.p1Y()
	case S3:
		return c.p2Y()
	default:
		return sdk.ZeroDec()
	}
}

func (c *Curve) upperBound(segNum int) sdk.Dec {
	switch segNum {
	case S1:
		return c.p1Y()
	case S2:
		return c.p2Y()
	case S3:
		return c.p3Y()
	default:
		return sdk.ZeroDec()
	}
}

func (c *Curve) t(segNum int) func(x sdk.Dec) (t sdk.Dec) {
	switch segNum {
	case S1:
		return c.S1.t
	case S2:
		return c.S2.t
	case S3:
		return c.S3.t
	default:
		return func(x sdk.Dec) sdk.Dec { return sdk.ZeroDec() }
	}
}

func (c *Curve) deltaX(segNum int) sdk.Dec {
	switch segNum {
	case S1:
		return c.S1.DeltaX
	case S2:
		return c.S2.DeltaX
	case S3:
		return c.S3.DeltaX
	default:
		return sdk.ZeroDec()
	}
}

// integralX12 returns the function for computing the integral for the given
// segment.
func (c *Curve) integralXFn(segNum int) func(x1, x2 sdk.Dec) sdk.Dec {
	switch segNum {
	case S0:
		return c.S0.integralX12
	case S1:
		return c.S1.integralX12
	case S2:
		return c.S2.integralX12
	case S3:
		return c.S3.integralX12
	case S4:
		return c.S4.integralX12
	default:
		return func(sdk.Dec, sdk.Dec) sdk.Dec { return sdk.ZeroDec() }
	}
}

// IsIntegralEqualToBPool checks if BPool is equal to the integral
// under the curve from zero to current supply.
func (c *Curve) IsIntegralEqualToBPool() bool {
	integral := c.integralX12(sdk.ZeroDec(), c.CurrentSupply)
	integral = bankersRoundOff(integral, VoucherMultiplier)

	bPool := bankersRoundOff(c.BPool, VoucherMultiplier)

	return integral.Equal(bPool)
}

func (c *Curve) integralX12(lowerBoundX, upperBoundX sdk.Dec) (vouchers sdk.Dec) {
	segLowerBoundX := c.segmentNum(lowerBoundX)
	segUpperBoundX := c.segmentNum(upperBoundX)

	vouchers = sdk.NewDec(0)
	for ; segLowerBoundX <= segUpperBoundX; segLowerBoundX = segLowerBoundX + 1 {
		x1 := lowerBoundX
		x2 := c.upperBoundX(segLowerBoundX)
		if segLowerBoundX == segUpperBoundX {
			x2 = upperBoundX
		}
		additionalVouchers := c.integralXFn(segLowerBoundX)(x1, x2)
		vouchers = vouchers.Add(additionalVouchers)

		lowerBoundX = c.upperBoundX(segLowerBoundX)
	}
	return vouchers
}

func (c *Curve) slopeX1(x1 sdk.Dec) sdk.Dec {
	switch c.segmentNum(x1) {
	case S0:
		return sdk.ZeroDec()
	case S1:
		return c.S1.firstDerivativeX1(x1)
	case S2:
		return c.S2.firstDerivativeX1(x1)
	case S3:
		return c.S3.firstDerivativeX1(x1)
	case S4:
		return c.S4.firstDerivativeX1(x1)
	default:
		return sdk.ZeroDec()
	}
}

func (c *Curve) y(x sdk.Dec) sdk.Dec {
	switch c.segmentNum(x) {
	case S0:
		return c.S0.y(x)
	case S1:
		return c.S1.y(x)
	case S2:
		return c.S2.y(x)
	case S3:
		return c.S3.y(x)
	case S4:
		return c.S4.y(x)
	default:
		return sdk.ZeroDec()
	}
}

// integralT1 computes the integral of the curve segment with respect to the
// bezier curve parameter "t", from the beginning of the curve until point t1.
func (c *Curve) integralT1(t1 sdk.Dec, seg int) sdk.Dec {
	switch seg {
	case S1:
		return c.S1.integralT1(t1)
	case S2:
		return c.S2.integralT1(t1)
	case S3:
		return c.S3.integralT1(t1)
	default:
		return sdk.ZeroDec()
	}
}
