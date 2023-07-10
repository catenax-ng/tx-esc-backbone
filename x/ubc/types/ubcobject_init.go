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
func (ubc *Ubcobject) Fit() error {
	if err := ubc.validateParameters(); err != nil {
		return err
	}
	// Trading starts at 50% of reference supply.
	ubc.CurrentSupply = ubc.RefTokenSupply.Quo(sdk.NewDec(2))

	ubc.initSegmentsToZero()
	ubc.fitS2()
	ubc.fitS3()
	ubc.fitS1S0()

	// Self-consistency to fit P0 better.
	ubc.fitS1S0()
	ubc.fitS1S0()

	return ubc.validateCurvature()
}

func (ubc *Ubcobject) validateParameters() error {
	if ubc.RefTokenSupply.IsNil() {
		return errors.Errorf("RefTokenSupply is not set")
	}
	if ubc.RefTokenPrice.IsNil() {
		return errors.Errorf("RefTokenPrice is not set")
	}
	if ubc.RefProfitFactor.IsNil() {
		return errors.Errorf("RefProfitFactor is not set")
	}
	if ubc.BPoolUnder.IsNil() {
		return errors.Errorf("BPoolUnder is not set")
	}
	if ubc.SlopeP2.IsNil() {
		return errors.Errorf("SlopeP2 is not set")
	}
	if ubc.SlopeP3.IsNil() {
		return errors.Errorf("SlopeP3 is not set")
	}
	if ubc.FactorFy.IsNil() {
		return errors.Errorf("FactorFy is not set")
	}
	if ubc.FactorFxy.IsNil() {
		return errors.Errorf("FactorFxy is not set")
	}
	return nil
}

func (ubc *Ubcobject) initSegmentsToZero() {
	ubc.FS0 = &Flatsegment{
		X0: sdk.ZeroDec(),
		Y:  sdk.ZeroDec(),
	}
	ubc.S0 = &Segment{
		P0:  sdk.ZeroDec(),
		P1:  sdk.ZeroDec(),
		P0X: sdk.ZeroDec(),
		P1X: sdk.ZeroDec(),
	}
	ubc.S1 = &Segment{
		P0:  sdk.ZeroDec(),
		P1:  sdk.ZeroDec(),
		P0X: sdk.ZeroDec(),
		P1X: sdk.ZeroDec(),
	}
	ubc.S2 = &Fixedsegment{
		Segment: &Segment{
			P0X: sdk.ZeroDec(),
			P0:  sdk.ZeroDec(),
		},
	}
	ubc.QS3 = &Quadraticsegment{}
}

func (ubc *Ubcobject) fitS2() {
	ubc.setP3X(ubc.RefTokenSupply)
	ubc.setP3(ubc.RefTokenPrice)
	ubc.setP2X(ubc.RefTokenSupply.Quo(sdk.NewDec(2)))
	ubc.setP2(ubc.RefTokenPrice.Quo(ubc.RefProfitFactor))

	ubc.calcS2AB()
}

func (ubc *Ubcobject) calcS2AB() {
	factor := sdk.NewDec(1).Quo(sdk.NewDec(3))
	ubc.S2.B = ubc.S2.P1.Sub(factor.Mul(ubc.SlopeP3.Mul(ubc.S2.DeltaX)))
	ubc.S2.A = ubc.S2.P0.Add(factor.Mul(ubc.SlopeP2.Mul(ubc.S2.DeltaX)))
}

func (ubc *Ubcobject) fitS3() {
	ubc.QS3.ScalingFactor = sdk.NewDec(1e9)

	curvatureP3 := ubc.S2.curvatureAtEnd()
	ubc.calcS3ABC(curvatureP3, ubc.SlopeP3, ubc.p3(), ubc.p3x())
}

func (ubc *Ubcobject) calcS3ABC(curvatureP3, slopeP3, p3, p3X sdk.Dec) {
	x3Scaled := p3X.Quo(ubc.QS3.ScalingFactor)

	ubc.QS3.A = curvatureP3.Mul(ubc.QS3.ScalingFactor).Quo(sdk.NewDec(2))
	ubc.QS3.B = ubc.SlopeP3.Mul(ubc.QS3.ScalingFactor).
		Sub(sdk.NewDec(2).Mul(ubc.QS3.A).Mul(x3Scaled))
	ubc.QS3.C = p3.Sub(ubc.QS3.A.Mul(x3Scaled.Power(2))).Sub(ubc.QS3.B.Mul(x3Scaled))
}

func (ubc *Ubcobject) fitS1S0() {
	ubc.calcP1X()

	g0 := ubc.calcG0()
	g1 := ubc.calcG1(g0)

	ubc.S1.B = ubc.calcS1B()
	ubc.setP0(ubc.calcP0(g0, g1))
	ubc.setP1(ubc.calcP1())

	ubc.S0.A = ubc.p0()
	ubc.S0.B = ubc.calcS0B()

	ubc.S1.A = ubc.calcS1A()
}

func (ubc *Ubcobject) calcP1X() {
	factor := sdk.NewDec(1).Sub(ubc.FactorFy).Mul(ubc.FactorFxy)
	deltaX1 := factor.Mul(ubc.p2().Sub(ubc.p0()))
	ubc.setP1X(ubc.p2x().Sub(deltaX1))
}

func (ubc *Ubcobject) calcG0() sdk.Dec {
	part1 := sdk.NewDec(-3).Quo(sdk.NewDec(2)).Mul(ubc.S0.DeltaX)
	part2 := sdk.NewDecWithPrec(5, 1).Mul(ubc.S1.DeltaX.Power(2).Quo(ubc.S0.DeltaX))
	return (part1.Sub(part2)).Mul(ubc.FactorFy)
}

func (ubc *Ubcobject) calcG1(g0 sdk.Dec) sdk.Dec {
	part1 := sdk.NewDec(4).Mul(ubc.S0.DeltaX)
	factorPart2 := sdk.NewDec(2).Mul(ubc.S1.DeltaX)
	part2 := factorPart2.Mul(sdk.NewDec(1).Sub(ubc.FactorFy))
	part3 := sdk.NewDec(4).Mul(ubc.FS0.X0)
	return g0.Add(part1).Add(part2).Add(part3)
	//TODO: Implement the condition in GetFS0X0 from the prototype.
}

func (ubc *Ubcobject) calcS1B() sdk.Dec {
	factorPart1 := ubc.S1.DeltaX.Quo(ubc.S2.DeltaX)
	part1 := factorPart1.Mul(ubc.p2().Sub(ubc.S2.A))
	return part1.Add(ubc.p2())
}

func (ubc *Ubcobject) calcP0(g0, g1 sdk.Dec) sdk.Dec {
	factor := sdk.NewDec(1).Quo(g1)
	part1 := sdk.NewDec(4).Mul(ubc.BPoolUnder)
	part2 := ubc.p2().Mul(g0.
		Sub(sdk.NewDec(2).Mul(ubc.S1.DeltaX.Mul(ubc.FactorFy))).
		Sub(ubc.S1.DeltaX))
	part3 := ubc.S1.DeltaX.Mul(ubc.S1.B)
	return factor.Mul(part1.Add(part2).Sub(part3))
}

func (ubc *Ubcobject) calcP1() sdk.Dec {
	part1 := ubc.FactorFy.Mul(ubc.p2())
	part2 := (sdk.NewDec(1).Sub(ubc.FactorFy)).Mul(ubc.p0())
	return part1.Add(part2)
}

func (ubc *Ubcobject) calcS0B() sdk.Dec {
	return sdk.NewDecWithPrec(5, 1).Mul(ubc.p0().Add(ubc.p1()))
}

func (ubc *Ubcobject) calcS1A() sdk.Dec {
	factorPart1 := sdk.NewDecWithPrec(5, 1).Mul(ubc.S1.DeltaX.Quo(ubc.S0.DeltaX))
	part1 := factorPart1.Mul(ubc.p1().Sub(ubc.p0()))
	return part1.Add(ubc.p1())
}

func (ubc *Ubcobject) validateCurvature() error {
	factor := (ubc.S1.DeltaX.Quo(ubc.S2.DeltaX)).Mul(
		ubc.S2.A.Sub(ubc.p2()))

	c1 := sdk.NewDecWithPrec(5, 1).Mul(ubc.p1().Sub(factor).Add(ubc.p2()))
	if !c1.GT(ubc.S1.A) {
		return errors.Errorf("curvature condition 1 failed")
	}

	c2 := sdk.NewDec(-2).Mul(factor).Add(ubc.p2())
	if !c2.LT(ubc.S1.A) {
		return errors.Errorf("curvature condition 2 failed")
	}

	c3 := ubc.p2().Sub(sdk.NewDec(3).Mul(factor))
	if !ubc.p1().GT(c3) {
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
	c4 := sdk.NewDecWithPrec(5, 1).Mul(ubc.p0().Add(ubc.p1()))
	if !ubc.S0.B.GT(ubc.p0()) && ubc.S0.B.Equal(c4) {
		return errors.Errorf("curvature condition 4 failed")
	}

	if !(ubc.p1().Sub(ubc.p0())).GTE(ubc.S0.firstDerivativeTAtZero()) {
		return errors.Errorf("curvature condition 5 failed")
	}

	if ubc.p0().LT(sdk.ZeroDec()) ||
		ubc.p1().LT(sdk.ZeroDec()) ||
		ubc.S0.A.LT(sdk.ZeroDec()) ||
		ubc.S0.B.LT(sdk.ZeroDec()) ||
		ubc.S1.A.LT(sdk.ZeroDec()) ||
		ubc.S1.B.LT(sdk.ZeroDec()) {
		return errors.Errorf("curvature condition 6 failed")
	}
	return nil
}
