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

func (k msgServer) Sell(goCtx context.Context, msg *types.MsgSell) (*types.MsgSellResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// These will not err, as error has been checked in ValidateBasic.
	seller, _ := sdk.AccAddressFromBech32(msg.Seller)
	tokensCoin, _ := sdk.ParseCoinNormalized(msg.Value)

	if !(tokensCoin.Denom == types.SystemTokenDenom) {
		errMsg := "amount should be specified in system token denom"
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, errMsg)
	}

	ubc, found := k.GetCurve(ctx)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "curve is not initialized")
	}

	// TODO: Consume gas from gas meter and make it a param that can be
	// changed later using a param.

	ubc, vouchersEarnedCoin := sellExactTokens(tokensCoin, ubc)

	err := k.takeTokensAndGiveVouchers(ctx, seller, vouchersEarnedCoin, tokensCoin)
	if err != nil {
		return nil, err
	}

	k.SetCurve(ctx, ubc)

	return &types.MsgSellResponse{
		Tokenssold:     tokensCoin.String(),
		Vouchersearned: vouchersEarnedCoin.String(),
	}, nil

}

func sellExactTokens(tokensCoin sdk.Coin, ubc types.Curve) (types.Curve, sdk.Coin) {
	tokens := sdk.NewDecFromInt(tokensCoin.Amount).QuoInt64(types.SystemTokenMultiplier)
	vouchersOut := ubc.Sell(tokens)
	vouchersEarned := subFeesDec(vouchersOut)
	return ubc, sdk.NewCoin(types.VoucherDenom, vouchersEarned.MulInt64(types.VoucherMultiplier).TruncateInt())
}

func (k msgServer) takeTokensAndGiveVouchers(ctx sdk.Context, seller sdk.AccAddress, vouchers, tokens sdk.Coin) error {
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, seller, types.ModuleName, sdk.NewCoins(tokens))
	if err != nil {
		return nil
	}
	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(tokens))
	if err != nil {
		return nil
	}
	return k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, seller, sdk.NewCoins(vouchers))
}

func subFeesDec(v sdk.Dec) sdk.Dec {
	return v.Mul(sdk.NewDecWithPrec(997, 3))
}
