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
	vouchersToAdd, _ := sdk.ParseCoinNormalized(msg.Voucherstoadd)

	err := ubc.ShiftUp(sdk.NewDecFromInt(vouchersToAdd.Amount.QuoRaw(types.VoucherMultiplier)), msg.Degirdingfactor)
	if err != nil {
		return &types.MsgShiftupResponse{}, err
	}

	k.SetUbcobject(ctx, ubc)

	return &types.MsgShiftupResponse{}, nil
}
