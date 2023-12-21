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
	"github.com/pkg/errors"
)

// Fit fits each of the segments of the undergirding bonding curve using the
// given set of parameters.
//
// It returns an error if any of the parameters was not set or if the resulting
// curve fitted is invalid.
func (c *Curve) Fit() error {
	if err := c.validateParameters(); err != nil {
		return err
	}
	c.initParameters()
	c.initSegmentsToZero()
	c.fitS3()
	c.fitS4()

	return c.FitUntilConvergence()
}

func (c *Curve) FitUntilConvergence() error {
	const maxIterations = 10
	i := 0
	for ; i < maxIterations; i++ {
		c.fitS1S2()
		if c.IsIntegralEqualToBPool() {
			break
		}
	}
	if i == maxIterations {
		return ErrCurveFitting.Wrapf("no convergence in %d iterations", i)
	}

	err := c.validateCurvature()
	if err != nil {
		return ErrCurveFitting.Wrap(err.Error())
	}
	return nil

}

func (c *Curve) validateParameters() error {
	if c.RefTokenSupply.IsNil() {
		return errors.Errorf("RefTokenSupply is not set")
	}
	if c.RefTokenPrice.IsNil() {
		return errors.Errorf("RefTokenPrice is not set")
	}
	if c.RefProfitFactor.IsNil() {
		return errors.Errorf("RefProfitFactor is not set")
	}
	if c.BPoolUnder.IsNil() {
		return errors.Errorf("BPoolUnder is not set")
	}
	if c.SlopeP2.IsNil() {
		return errors.Errorf("SlopeP2 is not set")
	}
	if c.SlopeP3.IsNil() {
		return errors.Errorf("SlopeP3 is not set")
	}
	if c.FactorFy.IsNil() {
		return errors.Errorf("FactorFy is not set")
	}
	if c.FactorFxy.IsNil() {
		return errors.Errorf("FactorFxy is not set")
	}
	return nil
}

func (c *Curve) initParameters() {
	c.NumericalErrorAccumulator = sdk.ZeroDec()

	// Trading starts at 50% of reference supply.
	c.CurrentSupply = c.RefTokenSupply.Quo(sdk.NewDec(2))
}

func (c *Curve) initSegmentsToZero() {
	c.S0 = &FlatSegment{
		P1X: sdk.ZeroDec(),
		Y:   sdk.ZeroDec(),
	}
	c.S1 = &BezierSegment{
		P0Y: sdk.ZeroDec(),
		P1Y: sdk.ZeroDec(),
		P0X: sdk.ZeroDec(),
		P1X: sdk.ZeroDec(),
	}
	c.S2 = &BezierSegment{
		P0Y: sdk.ZeroDec(),
		P1Y: sdk.ZeroDec(),
		P0X: sdk.ZeroDec(),
		P1X: sdk.ZeroDec(),
	}
	c.S3 = &FixedBezierSegment{
		BezierSegment: &BezierSegment{
			P0X: sdk.ZeroDec(),
			P0Y: sdk.ZeroDec(),
		},
	}
	c.S4 = &FixedQuadraticSegment{
		A:             sdk.ZeroDec(),
		B:             sdk.ZeroDec(),
		C:             sdk.ZeroDec(),
		ScalingFactor: sdk.ZeroDec(),
		InitialX0:     sdk.ZeroDec(),
		CurrentX0:     sdk.ZeroDec(),
	}
	c.PopulateSegments()
}

func (c *Curve) fitS3() {
	c.setPX(3, c.RefTokenSupply)
	c.setPY(3, c.RefTokenPrice)
	c.setPX(2, c.RefTokenSupply.Quo(sdk.NewDec(2)))
	c.setPY(2, c.RefTokenPrice.Quo(c.RefProfitFactor))

	c.calcS3AB()
}

func (c *Curve) calcS3AB() {
	factor := sdk.NewDec(1).Quo(sdk.NewDec(3))
	c.S3.B = c.S3.P1Y.Sub(factor.Mul(c.SlopeP3.Mul(c.S3.DeltaX)))
	c.S3.A = c.S3.P0Y.Add(factor.Mul(c.SlopeP2.Mul(c.S3.DeltaX)))
}

func (c *Curve) fitS4() {
	c.S4.ScalingFactor = sdk.NewDec(1e9)

	curvatureP3 := c.S3.curvatureAtEnd()
	c.calcS4ABC(curvatureP3, c.SlopeP3, c.S3.endY(), c.S3.endX())
}

func (c *Curve) calcS4ABC(curvatureP3, slopeP3, p3, p3X sdk.Dec) {
	x3Scaled := p3X.Quo(c.S4.ScalingFactor)

	c.S4.A = curvatureP3.Mul(c.S4.ScalingFactor).Quo(sdk.NewDec(2))
	c.S4.B = c.SlopeP3.Mul(c.S4.ScalingFactor).
		Sub(sdk.NewDec(2).Mul(c.S4.A).Mul(x3Scaled))
	c.S4.C = p3.Sub(c.S4.A.Mul(x3Scaled.Power(2))).Sub(c.S4.B.Mul(x3Scaled))
}

func (c *Curve) fitS1S2() {
	c.fitS1S2Repeatedly(1)
}

// fitS1S2Repeatedly fits S1 and S2 repeatedly for given count of repititions,
// for the sake of self consistency.
func (c *Curve) fitS1S2Repeatedly(repititions uint) {
	for i := uint(0); i < repititions; i++ {
		c.calcP1X()
		g0 := c.calcG0()
		g1 := c.calcG1(g0)

		c.S2.B = c.calcS2B()
		c.setPY(0, c.calcP0(g0, g1))
		c.setPY(1, c.calcP1())

		c.S1.A = c.pY(0)
		c.S1.B = c.calcS1B()

		c.S2.A = c.calcS1A()
	}
}

func (c *Curve) calcP1X() {
	factor := sdk.NewDec(1).Sub(c.FactorFy).Mul(c.FactorFxy)
	deltaX1 := factor.Mul(c.pY(2).Sub(c.pY(0)))
	c.setPX(1, c.pX(2).Sub(deltaX1))
}

func (c *Curve) calcP1XMethod2() {
	x1 := c.pX(2).Sub(c.FactorFxy.Mul(c.pY(2).Sub(c.pY(1))))
	c.setPX(1, x1)
}

func (c *Curve) calcG0() sdk.Dec {
	part1 := sdk.NewDec(-3).Quo(sdk.NewDec(2)).Mul(c.S1.DeltaX)
	part2 := sdk.NewDecWithPrec(5, 1).Mul(c.S2.DeltaX.Power(2).Quo(c.S1.DeltaX))
	return (part1.Sub(part2)).Mul(c.FactorFy)
}

func (c *Curve) calcG1(g0 sdk.Dec) sdk.Dec {
	part1 := sdk.NewDec(4).Mul(c.S1.DeltaX)
	factorPart2 := sdk.NewDec(2).Mul(c.S2.DeltaX)
	part2 := factorPart2.Mul(sdk.NewDec(1).Sub(c.FactorFy))
	part3 := sdk.NewDec(4).Mul(c.S0.P1X)
	return g0.Add(part1).Add(part2).Add(part3)
	//TODO: Implement the condition in GetFS0X0 from the prototype.
}

func (c *Curve) calcS2B() sdk.Dec {
	factorPart1 := c.S2.DeltaX.Quo(c.S3.DeltaX)
	part1 := factorPart1.Mul(c.pY(2).Sub(c.S3.A))
	return part1.Add(c.pY(2))
}

func (c *Curve) calcP0(g0, g1 sdk.Dec) sdk.Dec {
	factor := sdk.NewDec(1).Quo(g1)
	part1 := sdk.NewDec(4).Mul(c.BPoolUnder)
	part2 := c.pY(2).Mul(g0.
		Sub(sdk.NewDec(2).Mul(c.S2.DeltaX.Mul(c.FactorFy))).
		Sub(c.S2.DeltaX))
	part3 := c.S2.DeltaX.Mul(c.S2.B)
	return factor.Mul(part1.Add(part2).Sub(part3))
}

func (c *Curve) calcP1() sdk.Dec {
	part1 := c.FactorFy.Mul(c.pY(2))
	part2 := (sdk.NewDec(1).Sub(c.FactorFy)).Mul(c.pY(0))
	return part1.Add(part2)
}

func (c *Curve) calcS1B() sdk.Dec {
	return sdk.NewDecWithPrec(5, 1).Mul(c.pY(0).Add(c.pY(1)))
}

func (c *Curve) calcS1A() sdk.Dec {
	factorPart1 := sdk.NewDecWithPrec(5, 1).Mul(c.S2.DeltaX.Quo(c.S1.DeltaX))
	part1 := factorPart1.Mul(c.pY(1).Sub(c.pY(0)))
	return part1.Add(c.pY(1))
}

func (c *Curve) validateCurvature() error {
	factor := (c.S2.DeltaX.Quo(c.S3.DeltaX)).Mul(
		c.S3.A.Sub(c.pY(2)))

	c1 := sdk.NewDecWithPrec(5, 1).Mul(c.pY(1).Sub(factor).Add(c.pY(2)))
	if !c1.GT(c.S2.A) {
		return errors.Errorf("curvature condition 1 failed")
	}

	c2 := sdk.NewDec(-2).Mul(factor).Add(c.pY(2))
	if !c2.LT(c.S2.A) {
		return errors.Errorf("curvature condition 2 failed")
	}

	c3 := c.pY(2).Sub(sdk.NewDec(3).Mul(factor))
	if !c.pY(1).GT(c3) {
		return errors.Errorf("curvature condition 3 failed")
	}

	// Note: In actual, P0 < b0 <= c4. But in current implementation, b0 is
	// set to its maximum value c4 in order to have more even slope and
	// token distribution in P1. For more details on this, ref Section
	// 2.4.6 in the master thesis.
	//
	// In future, if b0 is computed using alternate methods and allowed to
	// take values lesser than c4, then this Equals in this condition
	// should be modified to less than or equals.
	c4 := sdk.NewDecWithPrec(5, 1).Mul(c.pY(0).Add(c.pY(1)))
	if !c.S1.B.GT(c.pY(0)) && c.S1.B.Equal(c4) {
		return errors.Errorf("curvature condition 4 failed")
	}

	if !(c.pY(1).Sub(c.pY(0))).GTE(c.S1.firstDerivativeT1(sdk.ZeroDec())) {
		return errors.Errorf("curvature condition 5 failed")
	}

	if c.pY(0).LT(sdk.ZeroDec()) ||
		c.pY(1).LT(sdk.ZeroDec()) ||
		c.S1.A.LT(sdk.ZeroDec()) ||
		c.S1.B.LT(sdk.ZeroDec()) ||
		c.S2.A.LT(sdk.ZeroDec()) ||
		c.S2.B.LT(sdk.ZeroDec()) {
		return errors.Errorf("curvature condition 6 failed")
	}
	return nil
}
