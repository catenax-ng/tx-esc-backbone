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
	"github.com/pkg/errors"
)

func (k msgServer) Init(goCtx context.Context, msg *types.MsgInit) (*types.MsgInitResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}
	ubc, err := msg.ParseUbcobject()
	if err != nil {
		return nil, err
	}

	if err = ubc.Fit(); err != nil {
		return nil, errors.Wrap(types.ErrCurveFitting, err.Error())
	}
	k.SetUbcobject(ctx, *ubc)

	creatorAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, errors.Wrap(types.ErrFundHandling, "creator address: "+err.Error())
	}

	initcvoucher := sdk.NewCoin(
		types.VoucherDenom,
		ubc.BPool.RoundInt().MulRaw(types.VoucherMultiplier))
	err = k.bankKeeper.SendCoinsFromAccountToModule(
		ctx,
		creatorAddress,
		types.ModuleName,
		sdk.NewCoins(initcvoucher))
	if err != nil {
		return nil, errors.Wrap(types.ErrFundHandling, "insufficient cvoucher: "+err.Error())
	}

	acaxToMint := sdk.NewCoin(
		types.CaxDenom,
		ubc.BPool.RoundInt().MulRaw(types.CaxMultiplier))
	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(acaxToMint))
	if err != nil {
		return nil, errors.Wrap(types.ErrFundHandling, "minting acax: "+err.Error())
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(
		ctx,
		types.ModuleName,
		creatorAddress,
		sdk.NewCoins(acaxToMint))
	if err != nil {
		return nil, errors.Wrap(types.ErrFundHandling, "transfering minted acax: "+err.Error())
	}

	return &types.MsgInitResponse{}, nil
}
