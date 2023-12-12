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

func (c *Curve) UndergirdS02(BPoolAdd sdk.Dec) error {
	if c.CurrentSupply.LT(c.pX(p2)) {
		errMsg := "could not undergird, since the currentSupply is not beyond P2"
		return sdkerrors.ErrInvalidRequest.Wrap(errMsg)
	}

	c.BPoolUnder = c.BPoolUnder.Add(BPoolAdd)
	c.BPool = c.BPool.Add(BPoolAdd)

	return c.FitUntilConvergence()
}
