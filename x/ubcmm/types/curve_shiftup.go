// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/pkg/errors"
)

func (c *Curve) ShiftUp(BPoolAdd, DegirdingFactor sdk.Dec) error {
	if c.CurrentSupply.LT(c.pX(p2)) {
		errMsg := "could not shiftup, since the currentSupply is not beyond P2"
		return sdkerrors.ErrInvalidRequest.Wrap(errMsg)
	}

	p2XOld := c.pX(p2)

	// Calculate the factored BPool that should be added.
	// This value is increased by the DegirdingFactor. In the end only the original amount is added.
	// DegiridingFactor [1...inf]
	BPoolAddFactored := BPoolAdd.Mul(DegirdingFactor)

	dX, err := c.computeDx(p2XOld, BPoolAddFactored)
	if err != nil {
		return err
	}
	P2XNew, p2YNew, err := c.computeP2New(p2XOld, dX)
	if err != nil {
		return err
	}
	dY := p2YNew.Sub(c.pY(p2))
	dBPool := c.integralX12(p2XOld, P2XNew)

	c.shiftP0P1P2(dX, dY)

	// CLARIFY: The value of BPoolUnder is not rounded off to voucherMultiplier. How to do it ?
	c.BPoolUnder = c.BPoolUnder.Add(BPoolAdd).Add(dBPool)
	c.BPool = c.BPool.Add(BPoolAdd)

	c.fitS1S2Repeatedly(4)

	return nil
}

func (c *Curve) computeP2New(P2XOld, dx sdk.Dec) (sdk.Dec, sdk.Dec, error) {
	P2XNew := P2XOld.Add(dx)
	if P2XNew.GTE(c.CurrentSupply) {
		return sdk.ZeroDec(), sdk.ZeroDec(), errors.Errorf("P2X was shifted beyond the current supply")
	}

	return P2XNew, c.y(P2XNew), nil
}

func (c *Curve) computeDx(p2XOld, BPoolAddFactored sdk.Dec) (sdk.Dec, error) {
	slopeAtP2XOld := c.slopeX1(p2XOld)

	part1 := slopeAtP2XOld.Mul(p2XOld).Add(c.pY(p0)).Sub(c.pY(p2)).Quo(slopeAtP2XOld)

	part2a := part1.Power(2)
	part2b := sdk.NewDec(2).Mul(BPoolAddFactored).Quo(slopeAtP2XOld)
	part2, err := part2a.Add(part2b).ApproxSqrt()
	if err != nil {
		return sdk.ZeroDec(), errors.Wrap(err, "evaluating approx square root for computing dxsh")
	}

	dx := part2.Sub(part1)
	return dx, nil
}

func (c *Curve) shiftP0P1P2(dX, dY sdk.Dec) {
	c.setP0X(c.pX(p0).Add(dX))
	c.setP1X(c.pX(p1).Add(dX))
	c.setP2X(c.pX(p2).Add(dX))

	c.setP0Y(c.pY(p0).Add(dY))
	c.setP1Y(c.pY(p1).Add(dY))
	c.setP2Y(c.pY(p2).Add(dY))
}
