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

// BuyExactTokens buys the given amount of tokens against the curve. It
// returns the amount of vouchers used.
//
// It assumes the value of tokens is greater than zero. This condition is
// implemented in the ValidBasic function for the buy message.
func (ubc *Ubcobject) BuyExactTokens(tokens sdk.Dec) sdk.Dec {
	xCurrent := ubc.CurrentSupply
	xNew := ubc.CurrentSupply.Add(tokens)

	segXCurrent := ubc.segmentNum(xCurrent)
	segXNew := ubc.segmentNum(xNew)

	var vouchersUsed = sdk.NewDec(0)
	for ; segXCurrent <= segXNew; segXCurrent = segXCurrent + 1 {
		x1 := xCurrent
		x2 := ubc.upperBoundX(segXCurrent)
		if segXCurrent == segXNew {
			x2 = xNew
		}
		additionalVouchers := ubc.integralXFn(segLowerBoundX)(x1, x2)
		vouchersUsed = vouchersUsed.Add(additionalVouchers)

		xCurrent = ubc.upperBoundX(segXCurrent)
	}
	vouchersUsed = roundOff(vouchersUsed, VoucherMultiplier)

	ubc.CurrentSupply = xNew
	ubc.BPool = ubc.BPool.Add(vouchersUsed)
	// CLARIFY: Should we change BPoolUnder
	return vouchersUsed
}

// BuyTokensFor buys the tokens for given amount of vouchers against the curve,
// also deducting the fees. It returns the amount of tokens and amount of
// vouchers used.
//
// It assumes the value of tokens is greater than zero. This condition is
// implemented in the ValidBasic function for the buy message.
func (ubc *Ubcobject) BuyTokensFor(vouchersIn sdk.Dec) (sdk.Dec, sdk.Dec, error) {
	tokens, err := ubc.calcApproxTokens(vouchersIn)
	if err != nil {
		return sdk.ZeroDec(), sdk.ZeroDec(), err
	}

	vouchersUsed := ubc.BuyExactTokens(tokens)
	vouchersRemaining := vouchersIn.Sub(vouchersUsed)
	// What if the vouchers are valid amount, but insufficent to buy tokens,
	// we shold not err in thsi case, but rather return 0 in BuyTokensFor
	if !vouchersRemaining.LTE(sdk.NewDec(VoucherMultiplier)) {
		vouchersCorrection, tokensCorrection, err := ubc.BuyTokensFor(vouchersRemaining)
		if err != nil {
			return sdk.ZeroDec(), sdk.ZeroDec(), err
		}
		vouchersUsed = vouchersUsed.Add(vouchersCorrection)
		tokens = tokens.Add(tokensCorrection)
	}

	return tokens, vouchersUsed, nil
}

// calcApproxTokens calculates the approx amount of tokens equivalent to the
// given amount of vouchers.
//
// It uses taylor expansion. Hence the results are approximate.
func (ubc *Ubcobject) calcApproxTokens(vouchersIn sdk.Dec) (sdk.Dec, error) {
	xCurrent := ubc.CurrentSupply
	segment := ubc.segmentNum(xCurrent)

	tCurrent := ubc.t(segment)(xCurrent)
	approxDeltaT, err := ubc.approxDeltaT(tCurrent, vouchersIn, segment)
	if err != nil {
		return sdk.ZeroDec(), err
	}
	correctedDeltaT := ubc.correctCalcErrInDeltaT(tCurrent, approxDeltaT, segment)
	approxTokens := ubc.deltaX(segment).Mul(correctedDeltaT)
	return roundOff(approxTokens, SystemTokenMultiplier), nil

}

// correctCalcErrInDeltaT corrects the deltaT calculated using approximation
// methods to prevent loss of funds due to calculation errors.
//
// We use taylor series appromiation of the bezier segments to calculate the
// value of deltaT. This approximate value of deltaT could at times be slightly
// larger than the actual value, which if used as such could lead to loss of
// funds for the curve. Hence, we use a mechanism to correct the approximate
// value to ensure it is always less than or equal to the actual value of
// deltaT.
//
// The degree of error is described using the constant factorOfSafety, used in
// this correction.
func (ubc *Ubcobject) correctCalcErrInDeltaT(tCurrent, deltaT sdk.Dec, segment int) sdk.Dec {
	factorOfSafety := sdk.NewDecWithPrec(1, 4)

	tNew := tCurrent.Add(deltaT)
	approxIntegralT := ubc.approxIntegralT12(tNew, tCurrent, segment)
	exactIntegralT := ubc.integralT1(tCurrent.Add(deltaT), segment)
	return approxIntegralT.Quo(exactIntegralT).Sub(factorOfSafety).Mul(deltaT)
}

// approxIntegralT12 computes the approximate integral of the bezier curve
// between the points tNew and tCurrent using taylor series expansion.
func (ubc *Ubcobject) approxIntegralT12(tNew sdk.Dec, tCurrent sdk.Dec, segment int) sdk.Dec {
	partA := ubc.taylorCoeffA(tCurrent, segment)
	partB := ubc.taylorCoeffB(tCurrent, segment).Mul(tNew.Sub(tCurrent))
	partC := ubc.taylorCoeffC(tCurrent, segment).Mul(tNew.Sub(tCurrent).Power(2))
	return partA.Add(partB).Add(partC)
}

func (ubc *Ubcobject) taylorCoeffA(t sdk.Dec, seg int) (a sdk.Dec) {
	return ubc.integralT1(t, seg)
}

func (ubc *Ubcobject) taylorCoeffB(t sdk.Dec, seg int) (b sdk.Dec) {
	Pi := computePolyFor(t, []term{{-4, 3}, {12, 2}, {-12, 1}, {4, 0}}).Mul(ubc.lowerBound(seg))
	ai := computePolyFor(t, []term{{12, 3}, {-24, 2}, {12, 1}}).Mul(ubc.a(seg))
	bi := computePolyFor(t, []term{{12, 3}, {-12, 2}}).Mul(ubc.b(seg))
	Pi1 := computePolyFor(t, []term{{4, 3}}).Mul(ubc.upperBound(seg))
	return Pi.Add(ai).Sub(bi).Add(Pi1)
}

func (ubc *Ubcobject) taylorCoeffC(t sdk.Dec, seg int) (c sdk.Dec) {
	Pi := computePolyFor(t, []term{{-12, 2}, {24, 1}, {-12, 0}}).Mul(ubc.lowerBound(seg))
	ai := computePolyFor(t, []term{{36, 2}, {-48, 1}, {12, 0}}).Mul(ubc.a(seg))
	bi := computePolyFor(t, []term{{36, 2}, {-24, 1}}).Mul(ubc.b(seg))
	Pi1 := computePolyFor(t, []term{{12, 2}}).Mul(ubc.upperBound(seg))
	return sdk.NewDecWithPrec(5, 1).Mul(Pi.Add(ai).Sub(bi).Add(Pi1))
}

// approxDeltaT computes the value of approxDeltaT to mint tokens worth
// "vouchersIn".
//
// It uses taylor expansion. Hence the results are approximate.
func (ubc *Ubcobject) approxDeltaT(tCurrent, vouchersIn sdk.Dec, segment int) (sdk.Dec, error) {
	b := ubc.taylorCoeffB(tCurrent, segment)
	c := ubc.taylorCoeffC(tCurrent, segment)

	part1 := sdk.NewDec(-1).Mul(b).Quo(sdk.NewDec(2).Mul(c))

	part2a := b.Power(2).Quo(sdk.NewDec(4).Mul(c.Power(2)))
	part2b := sdk.NewDec(4).Mul(vouchersIn).Quo(c.Mul(ubc.deltaX(segment)))
	part2, err := part2a.Add(part2b).ApproxSqrt()
	if err != nil {
		return sdk.ZeroDec(), errors.Wrap(ErrComputation, "calculating approx square root")
	}
	return part1.Add(part2), nil
}

// term is a term in a polynomial equation.
type term struct {
	coefficient int64
	exponent    uint64
}

// computePolyFor returns the value of polynomial p(x) constructed using the given
// terms for the point x1.
//
// Eg: "computePolyFor(2, []term{{36, 2}, {-48, 1}, {12, 0}})" returns the value of
// "36(x^2) - 48x + 12" at x=2.
func computePolyFor(x1 sdk.Dec, terms []term) sdk.Dec {
	// We don't use powers > 4. If we do in future, this will err and we can fix it.
	const maxPow = 4
	x1Pows := [maxPow + 1]sdk.Dec{}
	x1Pows[0] = sdk.OneDec()
	for i := 1; i <= maxPow; i++ {
		x1Pows[i] = x1Pows[i-1].Mul(x1)
	}

	sum := sdk.ZeroDec()
	for _, term := range terms {
		sum = sum.Add(sdk.NewDec(term.coefficient).Mul(x1Pows[term.exponent]))
	}
	return sum
}

func roundOff(t sdk.Dec, multiplier int64) sdk.Dec {
	// CLARIFY: Is the rounding off strategy correct ?
	return t.MulInt64(VoucherMultiplier).
		TruncateDec().
		QuoInt64(VoucherMultiplier)
}
