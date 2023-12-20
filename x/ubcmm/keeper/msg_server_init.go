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
	"github.com/pkg/errors"
)

func (k msgServer) Init(goCtx context.Context, msg *types.MsgInit) (*types.MsgInitResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	// This will not err, as error has been checked in ValidateBasic.
	ubc, _ := msg.ParseCurve()
	_ = ubc.Fit() // TODO: Consumer gas from gas meter and make it a param that can be changed later using a param.

	k.SetCurve(ctx, *ubc)

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

	tokensToMint := sdk.NewCoin(
		types.SystemTokenDenom,
		ubc.BPool.RoundInt().MulRaw(types.SystemTokenMultiplier))
	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(tokensToMint))
	if err != nil {
		return nil, errors.Wrap(types.ErrFundHandling, "minting tokens: "+err.Error())
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(
		ctx,
		types.ModuleName,
		creatorAddress,
		sdk.NewCoins(tokensToMint))
	if err != nil {
		return nil, errors.Wrap(types.ErrFundHandling, "transfering minted tokens: "+err.Error())
	}

	k.bankKeeper.SetSendEnabled(ctx, types.VoucherDenom, false)

	return &types.MsgInitResponse{}, nil
}
