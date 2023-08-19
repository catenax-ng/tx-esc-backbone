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

func (ubc *Curve) ShiftUp(BPoolAdd, DegirdingFactor sdk.Dec) error {
	if ubc.CurrentSupply.LT(ubc.p2x()) {
		errMsg := "could not shiftup, since the currentSupply is not beyond P2"
		return sdkerrors.ErrInvalidRequest.Wrap(errMsg)
	}

	p2XOld := ubc.p2x()

	// Calculate the factored BPool that should be added.
	// This value is increased by the DegirdingFactor. In the end only the original amount is added.
	// DegiridingFactor [1...inf]
	BPoolAddFactored := BPoolAdd.Mul(DegirdingFactor)

	dxsh, err := ubc.computeDx(p2XOld, BPoolAddFactored)
	if err != nil {
		return err
	}
	P2Xnew, p2YNew, err := ubc.computeP2New(p2XOld, dxsh)
	if err != nil {
		return err
	}
	dY := p2YNew.Sub(ubc.p2Y())
	dBPool := ubc.integralX12(p2XOld, P2Xnew)

	ubc.shiftP0P1P2(dxsh, dY)

	// CLARIFY: The value of BPoolUnder is not rounded off to voucherMultiplier. How to do it ?
	ubc.BPoolUnder = ubc.BPoolUnder.Add(BPoolAdd).Add(dBPool)
	ubc.BPool = ubc.BPool.Add(BPoolAdd)

	ubc.fitS0S1Repeatedly(4)

	return nil
}

func (ubc *Curve) computeP2New(P2XOld, dx sdk.Dec) (sdk.Dec, sdk.Dec, error) {
	P2XNew := P2XOld.Add(dx)
	if P2XNew.GTE(ubc.CurrentSupply) {
		return sdk.ZeroDec(), sdk.ZeroDec(), errors.Errorf("P2X was shifted beyond the current supply")
	}

	return P2XNew, ubc.y(P2XNew), nil
}

func (ubc *Curve) computeDx(p2XOld, BPoolAddFactored sdk.Dec) (sdk.Dec, error) {
	slopeAtP2XOld := ubc.slopeX1(p2XOld)

	part1 := slopeAtP2XOld.Mul(p2XOld).Add(ubc.p0Y()).Sub(ubc.p2Y()).Quo(slopeAtP2XOld)

	part2a := part1.Power(2)
	part2b := sdk.NewDec(2).Mul(BPoolAddFactored).Quo(slopeAtP2XOld)
	part2, err := part2a.Add(part2b).ApproxSqrt()
	if err != nil {
		return sdk.ZeroDec(), errors.Wrap(err, "evaluating approx square root for computing dxsh")
	}

	dx := part2.Sub(part1)
	return dx, nil
}

func (ubc *Curve) shiftP0P1P2(dxsh, dY sdk.Dec) {
	ubc.setP0X(ubc.p0x().Add(dxsh))
	ubc.setP1X(ubc.p1x().Add(dxsh))
	ubc.setP2X(ubc.p2x().Add(dxsh))

	ubc.setP0Y(ubc.p0Y().Add(dY))
	ubc.setP1Y(ubc.p1Y().Add(dY))
	ubc.setP2Y(ubc.p2Y().Add(dY))
}
