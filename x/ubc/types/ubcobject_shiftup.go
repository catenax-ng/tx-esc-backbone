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

func (ubc *Ubcobject) ShiftUp(BPoolAdd, DegirdingFactor sdk.Dec) error {
	if ubc.CurrentSupply.LT(ubc.p2x()) {
		errMsg := "could not undergird, since the currentSupply is not beyond P2"
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
	P2Xnew, P2new, err := ubc.computeP2New(p2XOld, dxsh)
	if err != nil {
		return err
	}
	dP2 := P2new.Sub(ubc.p2())
	dBPool := ubc.integralX12(p2XOld, P2Xnew)

	ubc.shiftP0P1P2(dxsh, dP2)

	// CLARIFY: The value of BPoolUnder is not rounded off to voucherMultiplier. How to do it ?
	ubc.BPoolUnder = ubc.BPoolUnder.Add(BPoolAdd).Add(dBPool)
	ubc.BPool = ubc.BPool.Add(BPoolAdd)

	ubc.fitS0S1Repeatedly(4)

	return nil
}

func (ubc *Ubcobject) computeP2New(P2XOld, dx sdk.Dec) (sdk.Dec, sdk.Dec, error) {
	P2XNew := P2XOld.Add(dx)
	if P2XNew.GTE(ubc.CurrentSupply) {
		return sdk.ZeroDec(), sdk.ZeroDec(), errors.Errorf("P2X was shifted beyond the current supply")
	}

	return P2XNew, ubc.y(P2XNew), nil
}

func (ubc *Ubcobject) computeDx(p2XOld, BPoolAddFactored sdk.Dec) (sdk.Dec, error) {
	slopeAtP2XOld := ubc.slopeX1(p2XOld)

	part1 := slopeAtP2XOld.Mul(p2XOld).Add(ubc.p0()).Sub(ubc.p2()).Quo(slopeAtP2XOld)

	part2a := part1.Power(2)
	part2b := sdk.NewDec(2).Mul(BPoolAddFactored).Quo(slopeAtP2XOld)
	part2, err := part2a.Add(part2b).ApproxSqrt()
	if err != nil {
		return sdk.ZeroDec(), errors.Wrap(err, "evaluating approx square root for computing dxsh")
	}

	dx := part2.Sub(part1)
	return dx, nil
}

func (ubc *Ubcobject) shiftP0P1P2(dxsh, dP2 sdk.Dec) {
	ubc.setP0X(ubc.p0x().Add(dxsh))
	ubc.setP1X(ubc.p1x().Add(dxsh))
	ubc.setP2X(ubc.p2x().Add(dxsh))

	ubc.setP0(ubc.p0().Add(dP2))
	ubc.setP1(ubc.p1().Add(dP2))
	ubc.setP2(ubc.p2().Add(dP2))
}
