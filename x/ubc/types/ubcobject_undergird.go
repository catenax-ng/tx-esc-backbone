// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (ubc *Ubcobject) UndergirdS01(BPoolAdd sdk.Dec) {
	ubc.BPoolUnder = ubc.BPoolUnder.Add(BPoolAdd)
	ubc.BPool = ubc.BPool.Add(BPoolAdd)

	cycles := 0
	for i := 0; i < 10; i++ {
		cycles = cycles + 1

		ubc.calcP1XMethod2()
		ubc.fitS1S0GivenP1X()

		// CLARIFY: Fit S1 along with this condition is full parameter sweep.
		// And this condition takes several thousand steps to evaluate to true.
		// In tests, this never converged in 10 cycles.
		if ubc.isA1WithinLimits() {
			break
		}
	}
}

func (ubc *Ubcobject) isA1WithinLimits() bool {
	// Allowed difference for a1 (convergence loop)
	allowedA1Difference := sdk.NewDecWithPrec(1, 4) // 0.0001 <=> 0.01%

	factor1 := sdk.NewDec(-2).Quo(sdk.NewDec(3)).Mul(ubc.S1.DeltaX)
	factor2 := ubc.S2.firstDerivativeT1(sdk.NewDec(0)).Quo(ubc.S2.DeltaX)
	lowerBoundA1 := factor1.Mul(factor2).Add(ubc.p2())

	part1 := sdk.NewDecWithPrec(5, 1).Mul(ubc.p1())
	part2 := sdk.NewDec(1).Quo(sdk.NewDec(6)).Mul(ubc.S1.DeltaX).Mul(factor2)
	part3 := sdk.NewDecWithPrec(5, 1).Mul(ubc.p2())
	upperBoundA1 := part1.Sub(part2).Add(part3)

	midA1 := lowerBoundA1.Add(upperBoundA1).Quo(sdk.NewDec(2))

	lower := midA1.Mul(sdk.NewDec(1).Sub(allowedA1Difference))
	upper := midA1.Mul(sdk.NewDec(1).Add(allowedA1Difference))
	return lower.LT(ubc.S1.A) && upper.GT(ubc.S1.A)
}
