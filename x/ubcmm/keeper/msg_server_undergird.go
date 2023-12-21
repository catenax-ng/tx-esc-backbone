// Copyright (c) 2022-2023 Contributors to the Eclipse Foundation
//
// See the NOTICE file(s) distributed with this work for additional
// information regarding copyright ownership.
//
// This program and the accompanying materials are made available under the
// terms of the Apache License, Version 2.0 which is available at
// https://www.apache.org/licenses/LICENSE-2.0.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.
//
// SPDX-License-Identifier: Apache-2.0
package keeper

import (
	"context"

	"github.com/catenax/esc-backbone/x/ubcmm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Undergird(goCtx context.Context, msg *types.MsgUndergird) (*types.MsgUndergirdResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	ubc, found := k.GetCurve(ctx)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "curve is not initialized")
	}

	// These will not err, as error has been checked in ValidateBasic.
	vouchersToAdd, _ := sdk.ParseCoinNormalized(msg.Voucherstoadd)

	err := ubc.UndergirdS02(sdk.NewDecFromInt(vouchersToAdd.Amount.QuoRaw(types.VoucherMultiplier)))
	if err != nil {
		return &types.MsgUndergirdResponse{}, err
	}

	k.SetCurve(ctx, ubc)

	return &types.MsgUndergirdResponse{}, nil
}
