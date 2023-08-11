// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (ubc *Curve) UndergirdS01(BPoolAdd sdk.Dec) error {
	if ubc.CurrentSupply.LT(ubc.p2x()) {
		errMsg := "could not undergird, since the currentSupply is not beyond P2"
		return sdkerrors.ErrInvalidRequest.Wrap(errMsg)
	}

	ubc.BPoolUnder = ubc.BPoolUnder.Add(BPoolAdd)
	ubc.BPool = ubc.BPool.Add(BPoolAdd)

	// CLARIFY: If it is fine to use the formula in calcP1X (used
	// internally by fitS1S0, instead of calcP1XMethod2 ?
	ubc.fitS0S1Repeatedly(10)

	// CLARIFY: Fit S1 along with the condition "isA1WithinLimits"
	// is full parameter sweep. And this condition takes several
	// thousand steps to evaluate to true. In tests, this never
	// converged in 10 cycles. So, is it okay that this is ignored ?

	err := ubc.validateCurvature()
	if err != nil {
		return ErrCurveFitting.Wrap("curvature conditions failed after undergirding: " + err.Error())
	}
	return nil
}
