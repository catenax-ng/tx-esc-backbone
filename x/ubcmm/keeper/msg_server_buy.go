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

func (k msgServer) Buy(goCtx context.Context, msg *types.MsgBuy) (*types.MsgBuyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// These will not err, as error has been checked in ValidateBasic.
	buyer, _ := sdk.AccAddressFromBech32(msg.Buyer)
	coin, _ := sdk.ParseCoinNormalized(msg.Value)

	ubc, found := k.GetCurve(ctx)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "curve is not initialized")
	}

	// TODO: Consume gas from gas meter and make it a param that can be
	// changed later using a param.

	ubc, vouchersSpentCoin := buy(coin, ubc)

	err := k.takeVouchersAndGiveTokens(ctx, buyer, vouchersSpentCoin, coin)
	if err != nil {
		return nil, types.ErrFundHandling.Wrap(err.Error())
	}

	k.SetCurve(ctx, ubc)

	return &types.MsgBuyResponse{
		Tokensbought:  coin.String(),
		Vouchersspent: vouchersSpentCoin.String(),
	}, nil

}

func buy(tokensCoin sdk.Coin, ubc types.Curve) (types.Curve, sdk.Coin) {
	tokens := sdk.NewDecFromInt(tokensCoin.Amount).QuoInt64(types.SystemTokenMultiplier)
	vouchersSpent := ubc.Buy(tokens)

	fee := vouchersSpent.Mul(feePercentageDec)

	vouchersSpent = vouchersSpent.Add(fee).MulInt64(types.VoucherMultiplier)

	return ubc, sdk.NewCoin(types.VoucherDenom, vouchersSpent.Ceil().TruncateInt())
}

func (k msgServer) takeVouchersAndGiveTokens(ctx sdk.Context, buyer sdk.AccAddress, vouchers, tokens sdk.Coin) error {
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, buyer, types.ModuleName, sdk.NewCoins(vouchers))
	if err != nil {
		return err
	}

	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(tokens))
	if err != nil {
		return err
	}
	return k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, buyer, sdk.NewCoins(tokens))
}

var feePercentageDec = sdk.NewDecWithPrec(3, 3)

func subFeesInt(v sdk.Int) sdk.Int {
	return v.MulRaw(997).QuoRaw(1000)
}
