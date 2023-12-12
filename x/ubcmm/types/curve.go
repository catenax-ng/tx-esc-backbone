// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type (
	pointN int
)

const (
	S0 = iota
	S1
	S2
	S3
	S4
)

const (
	p0 pointN = iota
	p1
	p2
	p3
)

func (c *Curve) view() []view {
	return []view{c.S1, c.S2, c.S3, c.S4}
}

// pX returns the x co-ordinate for the point of the curve.
func (c *Curve) pX(pN pointN) sdk.Dec {
	return c.view()[pN].startX()
}

// pY returns the x co-ordinate for the point of the curve.
func (c *Curve) pY(pN pointN) sdk.Dec {
	return c.view()[pN].startY()
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
	upperBoundsX := []sdk.Dec{c.pX(p0), c.pX(p1), c.pX(p2), c.pX(p3)}
	segments := []int{S0, S1, S2, S3}

	for i, upperBound := range upperBoundsX {
		if x.LT(upperBound) {
			return segments[i]
		}
	}
	return S4
}

func (c *Curve) upperBoundX(segNum int) sdk.Dec {
	switch segNum {
	case S0:
		return c.pX(p0)
	case S1:
		return c.pX(p1)
	case S2:
		return c.pX(p2)
	case S3:
		return c.pX(p3)
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
