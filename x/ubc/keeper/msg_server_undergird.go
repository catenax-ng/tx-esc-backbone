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
)

func (k msgServer) Undergird(goCtx context.Context, msg *types.MsgUndergird) (*types.MsgUndergirdResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgUndergirdResponse{}, nil
}
