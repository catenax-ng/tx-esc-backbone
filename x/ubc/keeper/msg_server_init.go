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

	return &types.MsgInitResponse{}, nil
}
