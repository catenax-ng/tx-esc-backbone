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
	segN   int

	segment interface {
		setP0X(sdk.Dec)
		setP1X(sdk.Dec)
		setP0Y(sdk.Dec)
		setP1Y(sdk.Dec)

		y(x sdk.Dec) sdk.Dec
	}
)

var (
	// endPointOf maps the end point of a segment to the segment.
	endPointOf = map[pointN]segN{
		p0: s0,
		p1: s1,
		p2: s2,
		p3: s3,
	}
	// startPointOf maps the start point of a segment to the segment.
	startPointOf = map[pointN]segN{
		p0: s1,
		p1: s2,
		p2: s3,
		p3: s4,
	}
)

const (
	s0 segN = iota
	s1
	s2
	s3
	s4
)

const (
	p0 pointN = iota
	p1
	p2
	p3
)

func (c *Curve) segments(segNum segN) segment {
	s := []segment{c.S0, c.S1, c.S2, c.S3, c.S4}
	return s[segNum]
}

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

func (c *Curve) setPX(point pointN, value sdk.Dec) {
	c.segments(endPointOf[point]).setP1X(value)
	c.segments(startPointOf[point]).setP0X(value)
}

func (c *Curve) setPY(point pointN, value sdk.Dec) {
	c.segments(endPointOf[point]).setP1Y(value)
	c.segments(startPointOf[point]).setP0Y(value)
}

// segmentNum returns the segment number for the given point x.
func (c *Curve) segmentNum(x sdk.Dec) segN {
	upperBoundsX := []sdk.Dec{c.pX(p0), c.pX(p1), c.pX(p2), c.pX(p3)}
	segments := []segN{s0, s1, s2, s3}

	for i, upperBound := range upperBoundsX {
		if x.LT(upperBound) {
			return segments[i]
		}
	}
	return s4
}

func (c *Curve) upperBoundX(segN segN) sdk.Dec {
	switch segN {
	case s0:
		return c.pX(p0)
	case s1:
		return c.pX(p1)
	case s2:
		return c.pX(p2)
	case s3:
		return c.pX(p3)
	default:
		return sdk.ZeroDec()
	}
}

// integralX12 returns the function for computing the integral for the given
// segment.
func (c *Curve) integralXFn(segN segN) func(x1, x2 sdk.Dec) sdk.Dec {
	switch segN {
	case s0:
		return c.S0.integralX12
	case s1:
		return c.S1.integralX12
	case s2:
		return c.S2.integralX12
	case s3:
		return c.S3.integralX12
	case s4:
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
	case s0:
		return sdk.ZeroDec()
	case s1:
		return c.S1.firstDerivativeX1(x1)
	case s2:
		return c.S2.firstDerivativeX1(x1)
	case s3:
		return c.S3.firstDerivativeX1(x1)
	case s4:
		return c.S4.firstDerivativeX1(x1)
	default:
		return sdk.ZeroDec()
	}
}

func (c *Curve) y(x sdk.Dec) sdk.Dec {
	return c.segments(c.segmentNum(x)).y(x)
}
