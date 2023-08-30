// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
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
	// Trading starts at 50% of reference supply.
	c.CurrentSupply = c.RefTokenSupply.Quo(sdk.NewDec(2))

	c.initSegmentsToZero()
	c.NumericalErrorAccumulator = sdk.ZeroDec()

	c.fitS2()
	c.fitS3()
	c.fitS0S1Repeatedly(1)

	// Self-consistency to fit P0 better.
	c.fitS0S1Repeatedly(2)

	return c.validateCurvature()
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

func (c *Curve) initSegmentsToZero() {
	c.FS0 = &FlatSegment{
		X0: sdk.ZeroDec(),
		Y:  sdk.ZeroDec(),
	}
	c.S0 = &BezierSegment{
		P0Y: sdk.ZeroDec(),
		P1Y: sdk.ZeroDec(),
		P0X: sdk.ZeroDec(),
		P1X: sdk.ZeroDec(),
	}
	c.S1 = &BezierSegment{
		P0Y: sdk.ZeroDec(),
		P1Y: sdk.ZeroDec(),
		P0X: sdk.ZeroDec(),
		P1X: sdk.ZeroDec(),
	}
	c.S2 = &FixedBezierSegment{
		BezierSegment: &BezierSegment{
			P0X: sdk.ZeroDec(),
			P0Y: sdk.ZeroDec(),
		},
	}
	c.QS3 = &FixedQuadraticSegment{
		A:             sdk.ZeroDec(),
		B:             sdk.ZeroDec(),
		C:             sdk.ZeroDec(),
		ScalingFactor: sdk.ZeroDec(),
		InitialX0:     sdk.ZeroDec(),
		CurrentX0:     sdk.ZeroDec(),
	}
}

func (c *Curve) fitS2() {
	c.setP3X(c.RefTokenSupply)
	c.setP3Y(c.RefTokenPrice)
	c.setP2X(c.RefTokenSupply.Quo(sdk.NewDec(2)))
	c.setP2Y(c.RefTokenPrice.Quo(c.RefProfitFactor))

	c.calcS2AB()
}

func (c *Curve) calcS2AB() {
	factor := sdk.NewDec(1).Quo(sdk.NewDec(3))
	c.S2.B = c.S2.P1Y.Sub(factor.Mul(c.SlopeP3.Mul(c.S2.DeltaX)))
	c.S2.A = c.S2.P0Y.Add(factor.Mul(c.SlopeP2.Mul(c.S2.DeltaX)))
}

func (c *Curve) fitS3() {
	c.QS3.ScalingFactor = sdk.NewDec(1e9)

	curvatureP3 := c.S2.curvatureAtEnd()
	c.calcS3ABC(curvatureP3, c.SlopeP3, c.S2.endY(), c.S2.endX())
}

func (c *Curve) calcS3ABC(curvatureP3, slopeP3, p3, p3X sdk.Dec) {
	x3Scaled := p3X.Quo(c.QS3.ScalingFactor)

	c.QS3.A = curvatureP3.Mul(c.QS3.ScalingFactor).Quo(sdk.NewDec(2))
	c.QS3.B = c.SlopeP3.Mul(c.QS3.ScalingFactor).
		Sub(sdk.NewDec(2).Mul(c.QS3.A).Mul(x3Scaled))
	c.QS3.C = p3.Sub(c.QS3.A.Mul(x3Scaled.Power(2))).Sub(c.QS3.B.Mul(x3Scaled))
}

// fitS0S1Repeatedly fits S0 and S1 repeatedly for given count of repititions,
// for the sake of self consistency.
func (c *Curve) fitS0S1Repeatedly(repititions uint) {
	for i := uint(0); i < repititions; i++ {
		c.calcP1X()
		g0 := c.calcG0()
		g1 := c.calcG1(g0)

		c.S1.B = c.calcS1B()
		c.setP0Y(c.calcP0(g0, g1))
		c.setP1Y(c.calcP1())

		c.S0.A = c.p0Y()
		c.S0.B = c.calcS0B()

		c.S1.A = c.calcS1A()
	}
}

func (c *Curve) calcP1X() {
	factor := sdk.NewDec(1).Sub(c.FactorFy).Mul(c.FactorFxy)
	deltaX1 := factor.Mul(c.p2Y().Sub(c.p0Y()))
	c.setP1X(c.p2x().Sub(deltaX1))
}

func (c *Curve) calcP1XMethod2() {
	x1 := c.p2x().Sub(c.FactorFxy.Mul(c.p2Y().Sub(c.p1Y())))
	c.setP1X(x1)
}

func (c *Curve) calcG0() sdk.Dec {
	part1 := sdk.NewDec(-3).Quo(sdk.NewDec(2)).Mul(c.S0.DeltaX)
	part2 := sdk.NewDecWithPrec(5, 1).Mul(c.S1.DeltaX.Power(2).Quo(c.S0.DeltaX))
	return (part1.Sub(part2)).Mul(c.FactorFy)
}

func (c *Curve) calcG1(g0 sdk.Dec) sdk.Dec {
	part1 := sdk.NewDec(4).Mul(c.S0.DeltaX)
	factorPart2 := sdk.NewDec(2).Mul(c.S1.DeltaX)
	part2 := factorPart2.Mul(sdk.NewDec(1).Sub(c.FactorFy))
	part3 := sdk.NewDec(4).Mul(c.FS0.X0)
	return g0.Add(part1).Add(part2).Add(part3)
	//TODO: Implement the condition in GetFS0X0 from the prototype.
}

func (c *Curve) calcS1B() sdk.Dec {
	factorPart1 := c.S1.DeltaX.Quo(c.S2.DeltaX)
	part1 := factorPart1.Mul(c.p2Y().Sub(c.S2.A))
	return part1.Add(c.p2Y())
}

func (c *Curve) calcP0(g0, g1 sdk.Dec) sdk.Dec {
	factor := sdk.NewDec(1).Quo(g1)
	part1 := sdk.NewDec(4).Mul(c.BPoolUnder)
	part2 := c.p2Y().Mul(g0.
		Sub(sdk.NewDec(2).Mul(c.S1.DeltaX.Mul(c.FactorFy))).
		Sub(c.S1.DeltaX))
	part3 := c.S1.DeltaX.Mul(c.S1.B)
	return factor.Mul(part1.Add(part2).Sub(part3))
}

func (c *Curve) calcP1() sdk.Dec {
	part1 := c.FactorFy.Mul(c.p2Y())
	part2 := (sdk.NewDec(1).Sub(c.FactorFy)).Mul(c.p0Y())
	return part1.Add(part2)
}

func (c *Curve) calcS0B() sdk.Dec {
	return sdk.NewDecWithPrec(5, 1).Mul(c.p0Y().Add(c.p1Y()))
}

func (c *Curve) calcS1A() sdk.Dec {
	factorPart1 := sdk.NewDecWithPrec(5, 1).Mul(c.S1.DeltaX.Quo(c.S0.DeltaX))
	part1 := factorPart1.Mul(c.p1Y().Sub(c.p0Y()))
	return part1.Add(c.p1Y())
}

func (c *Curve) validateCurvature() error {
	factor := (c.S1.DeltaX.Quo(c.S2.DeltaX)).Mul(
		c.S2.A.Sub(c.p2Y()))

	c1 := sdk.NewDecWithPrec(5, 1).Mul(c.p1Y().Sub(factor).Add(c.p2Y()))
	if !c1.GT(c.S1.A) {
		return errors.Errorf("curvature condition 1 failed")
	}

	c2 := sdk.NewDec(-2).Mul(factor).Add(c.p2Y())
	if !c2.LT(c.S1.A) {
		return errors.Errorf("curvature condition 2 failed")
	}

	c3 := c.p2Y().Sub(sdk.NewDec(3).Mul(factor))
	if !c.p1Y().GT(c3) {
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
	c4 := sdk.NewDecWithPrec(5, 1).Mul(c.p0Y().Add(c.p1Y()))
	if !c.S0.B.GT(c.p0Y()) && c.S0.B.Equal(c4) {
		return errors.Errorf("curvature condition 4 failed")
	}

	if !(c.p1Y().Sub(c.p0Y())).GTE(c.S0.firstDerivativeT1(sdk.ZeroDec())) {
		return errors.Errorf("curvature condition 5 failed")
	}

	if c.p0Y().LT(sdk.ZeroDec()) ||
		c.p1Y().LT(sdk.ZeroDec()) ||
		c.S0.A.LT(sdk.ZeroDec()) ||
		c.S0.B.LT(sdk.ZeroDec()) ||
		c.S1.A.LT(sdk.ZeroDec()) ||
		c.S1.B.LT(sdk.ZeroDec()) {
		return errors.Errorf("curvature condition 6 failed")
	}
	return nil
}
