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
)

var (
	// endPoint maps the segment to the end point of a segment.
	endPoint = map[segN]pointN{
		s0: p0,
		s1: p1,
		s2: p2,
		s3: p3,
	}
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

	firstSegment = s0
	lastSegment  = s4
)

const (
	p0 pointN = iota
	p1
	p2
	p3
)

// PopulateSegments populates the segment array, which can be used to access common properties of a segment using its index.
//
// Since, this is not stored on the blockchain, to prevent unnecessary storage.
func (c *Curve) PopulateSegments() {
	c.Segments = Segments([]Segment{c.S0, c.S1, c.S2, c.S3, c.S4})
}

// pX returns the x co-ordinate for the point of the curve.
func (c *Curve) pX(pN pointN) sdk.Dec {
	return c.Segments[startPointOf[pN]].startX()
}

// pY returns the x co-ordinate for the point of the curve.
func (c *Curve) pY(pN pointN) sdk.Dec {
	return c.Segments[startPointOf[pN]].startY()
}

func (c *Curve) setPX(point pointN, value sdk.Dec) {
	c.Segments[endPointOf[point]].setP1X(value)
	c.Segments[startPointOf[point]].setP0X(value)
}

func (c *Curve) setPY(point pointN, value sdk.Dec) {
	c.Segments[endPointOf[point]].setP1Y(value)
	c.Segments[startPointOf[point]].setP0Y(value)
}

func (c *Curve) segN(x sdk.Dec) segN {
	for i := firstSegment; i < lastSegment; i++ {
		if x.LT(c.pX(endPoint[i])) {
			return i
		}
	}
	return lastSegment
}

func (c *Curve) upperBoundX(segNum segN) sdk.Dec {
	if segNum >= 4 {
		return sdk.NewDec(-1)
	}
	return c.pX(endPoint[segNum])
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
	segLowerBoundX := c.segN(lowerBoundX)
	segUpperBoundX := c.segN(upperBoundX)

	vouchers = sdk.NewDec(0)
	for ; segLowerBoundX <= segUpperBoundX; segLowerBoundX = segLowerBoundX + 1 {
		x1 := lowerBoundX
		x2 := c.upperBoundX(segLowerBoundX)
		if segLowerBoundX == segUpperBoundX {
			x2 = upperBoundX
		}
		additionalVouchers := c.Segments[segLowerBoundX].integralX12(x1, x2)
		vouchers = vouchers.Add(additionalVouchers)

		lowerBoundX = c.upperBoundX(segLowerBoundX)
	}
	return vouchers
}

func (c *Curve) slopeX1(x1 sdk.Dec) sdk.Dec {
	return c.Segments[c.segN(x1)].firstDerivativeX1(x1)
}

func (c *Curve) y(x sdk.Dec) sdk.Dec {
	return c.Segments[c.segN(x)].y(x)
}
