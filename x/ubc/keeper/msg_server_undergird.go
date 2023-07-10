// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package keeper

import (
	"context"

	"github.com/catenax/esc-backbone/x/ubc/types"
	"github.com/cometbft/cometbft/crypto"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/pkg/errors"
)

func (k msgServer) Undergird(goCtx context.Context, msg *types.MsgUndergird) (*types.MsgUndergirdResponse, error) {
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

	// CLARIFY: How should the number 1e8 be arrived at based on input parameters ???
	vouchersFromOperator := sdk.NewCoin(types.VoucherDenom, sdk.NewInt(1e8))
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
		return &types.MsgUndergirdResponse{}, err
	}

	err = ubc.UndergirdS01(sdk.NewDecFromInt(vouchersToAdd.Amount.QuoRaw(types.VoucherMultiplier)))
	if err != nil {
		return &types.MsgUndergirdResponse{}, err
	}

	k.SetUbcobject(ctx, ubc)

	return &types.MsgUndergirdResponse{}, nil
}

func (k msgServer) isModuleBalanceSufficient(ubc types.Ubcobject, ctx sdk.Context, vouchersToAdd sdk.Int) bool {
	moduleAddress := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))
	moduleVoucherBalance := k.bankKeeper.GetBalance(ctx, moduleAddress, types.VoucherDenom).Amount
	bPoolUnder := ubc.BPoolUnder.MulInt64(types.VoucherMultiplier).TruncateInt()
	balanceDifference := moduleVoucherBalance.Sub(bPoolUnder)

	return !balanceDifference.LT(vouchersToAdd)
}
