// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package keeper

import (
	"context"

	"github.com/catenax/esc-backbone/x/ubc/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/pkg/errors"
)

func (k msgServer) Shiftup(goCtx context.Context, msg *types.MsgShiftup) (*types.MsgShiftupResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	ubc, found := k.GetUbcobject(ctx)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "curve is not initialized")
	}

	// These will not err, as error has been checked in ValidateBasic.
	operator, _ := sdk.AccAddressFromBech32(msg.Operator)
	vouchersToAdd, _ := sdk.ParseCoinNormalized(msg.Voucherstoadd)
	degirdingFactor, _ := sdk.NewDecFromStr(msg.Degirdingfactor)

	// CLARIFY: How should the number 1e8 be arrived at based on input parameters ???
	// CLARIFY: PoC uses 1e10 * VoucherDenom. But looking at message processing code, this seems like an error ?
	// Also, the module does not have so many vouchers, so exclude the multiplier.
	vouchersFromOperator := sdk.NewCoin(types.VoucherDenom, sdk.NewInt(1e10))
	err := k.bankKeeper.SendCoinsFromAccountToModule(
		ctx,
		operator,
		types.ModuleName,
		sdk.NewCoins(vouchersFromOperator))
	if err != nil {
		return nil, errors.Wrap(types.ErrFundHandling, "insufficient vouchers: "+err.Error())
	}

	if k.isModuleBalanceSufficient(ubc, ctx, vouchersToAdd.Amount) {
		err = types.ErrFundHandling.Wrap("insufficient vouchers")
		return &types.MsgShiftupResponse{}, err
	}

	err = ubc.ShiftUp(sdk.NewDecFromInt(vouchersToAdd.Amount.QuoRaw(types.VoucherMultiplier)), degirdingFactor)
	if err != nil {
		return &types.MsgShiftupResponse{}, err
	}

	k.SetUbcobject(ctx, ubc)

	return &types.MsgShiftupResponse{}, nil
}
