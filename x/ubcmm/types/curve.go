// Copyright (c) 2022-2023 Contributors to the Eclipse Foundation
//
// See the NOTICE file(s) distributed with this work for additional
// information regarding copyright ownership.
//
// This program and the accompanying materials are made available under the
// terms of the Apache License, Version 2.0 which is available at
// https://www.apache.org/licenses/LICENSE-2.0.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.
//
// SPDX-License-Identifier: Apache-2.0
package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// PopulateSegments populates the segment array, which can be used to access common properties of a segment using its index.
//
// Since, this is not stored on the blockchain, to prevent unnecessary storage.
func (c *Curve) PopulateSegments() {
	c.Segments = Segments([]Segment{c.S0, c.S1, c.S2, c.S3, c.S4})
}

// pX returns the x ordinate for the point of the curve.
func (c *Curve) pX(point int) sdk.Dec {
	if !c.isValidPoint(point) {
		return sdk.NewDec(-1)
	}
	return c.Segments[point+1].startX()
}

// pY returns the y ordinate for the point of the curve.
func (c *Curve) pY(point int) sdk.Dec {
	if !c.isValidPoint(point) {
		return sdk.NewDec(-1)
	}
	return c.Segments[point+1].startY()
}

func (c *Curve) setPX(point int, value sdk.Dec) {
	if !c.isValidPoint(point) {
		return
	}
	c.Segments[point].setP1X(value)
	c.Segments[point+1].setP0X(value)
}

func (c *Curve) setPY(point int, value sdk.Dec) {
	if !c.isValidPoint(point) {
		return
	}
	c.Segments[point].setP1Y(value)
	c.Segments[point+1].setP0Y(value)
}

func (c *Curve) isValidPoint(point int) bool {
	return point >= 0 && point < len(c.Segments)-1
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
		x2 := c.Segments[segLowerBoundX].endX()
		if segLowerBoundX == segUpperBoundX {
			x2 = upperBoundX
		}
		additionalVouchers := c.Segments[segLowerBoundX].integralX12(x1, x2)
		vouchers = vouchers.Add(additionalVouchers)

		lowerBoundX = c.Segments[segLowerBoundX].endX()
	}
	return vouchers
}

func (c *Curve) slopeX1(x1 sdk.Dec) sdk.Dec {
	return c.Segments[c.segN(x1)].firstDerivativeX1(x1)
}

func (c *Curve) y(x sdk.Dec) sdk.Dec {
	return c.Segments[c.segN(x)].y(x)
}

func (c *Curve) segN(x sdk.Dec) int {
	for i, segment := range c.Segments {
		if x.LTE(segment.endX()) {
			return i
		}
	}
	// Since last segment is unbounded, it is possible that the condition
	// inside for loop never succeeds for very very large x.
	return len(c.Segments) - 1
}
