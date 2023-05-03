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

func (k msgServer) Buytokens(goCtx context.Context, msg *types.MsgBuytokens) (*types.MsgBuytokensResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// These will not err, as error has been checked in ValidateBasic.
	buyer, _ := sdk.AccAddressFromBech32(msg.Buyer)
	coin, _ := sdk.ParseCoinNormalized(msg.Value)

	ubc, found := k.GetUbcobject(ctx)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "curve is not initialized")
	}

	// TODO: Consume gas from gas meter and make it a param that can be
	// changed later using a param.

	ubc, vouchersSpentCoin, tokensCoin, err := buyTokens(ctx, coin, ubc)
	if err != nil {
		return nil, err
	}

	err = k.takeVoucherAndGiveTokens(ctx, buyer, vouchersSpentCoin, tokensCoin)
	if err != nil {
		return nil, types.ErrFundHandling.Wrap(err.Error())
	}

	k.SetUbcobject(ctx, ubc)

	return &types.MsgBuytokensResponse{
		Tokensbought:  tokensCoin.String(),
		Vouchersspent: vouchersSpentCoin.String(),
	}, nil

}

func buyTokens(ctx sdk.Context, coin sdk.Coin, ubc types.Ubcobject) (types.Ubcobject, sdk.Coin, sdk.Coin, error) {
	var vouchersSpentCoin, tokensCoin sdk.Coin
	var err error

	if coin.Denom == types.SystemTokenDenom {
		tokensCoin = coin
		ubc, vouchersSpentCoin = buyExactTokens(tokensCoin, ubc)
	} else {
		vouchersSpentCoin, tokensCoin, err = buyTokensFor(coin, ubc)
		if err != nil {
			return ubc, sdk.Coin{}, sdk.Coin{}, err
		}
	}

	return ubc, vouchersSpentCoin, tokensCoin, nil
}

func buyExactTokens(tokensCoin sdk.Coin, ubc types.Ubcobject) (types.Ubcobject, sdk.Coin) {
	tokens := sdk.NewDecFromInt(tokensCoin.Amount).QuoInt64(types.SystemTokenMultiplier)
	vouchersSpent := ubc.BuyExactTokens(tokens)

	fee := vouchersSpent.Mul(feePercentage)

	vouchersSpent = vouchersSpent.Add(fee).MulInt64(types.VoucherMultiplier)

	return ubc, sdk.NewCoin(types.VoucherDenom, vouchersSpent.Ceil().TruncateInt())
}

func buyTokensFor(vouchersInCoin sdk.Coin, ubc types.Ubcobject) (sdk.Coin, sdk.Coin, error) {
	vouchersToSpendInt := subFees(vouchersInCoin.Amount)
	feeInt := vouchersInCoin.Amount.Sub(vouchersToSpendInt)
	vouchersToSpend := sdk.NewDecFromInt(vouchersToSpendInt).QuoInt64(types.VoucherMultiplier)

	tokens, vouchersSpent, err := ubc.BuyTokensFor(vouchersToSpend)
	if err != nil {
		return sdk.Coin{}, sdk.Coin{}, err
	}

	tokensInt := tokens.MulInt64(types.SystemTokenMultiplier).Ceil().TruncateInt()
	tokensCoin := sdk.NewCoin(types.SystemTokenDenom, tokensInt)
	vouchersSpentInt := vouchersSpent.MulInt64(types.VoucherMultiplier).Ceil().TruncateInt()
	vouchersSpentInt = vouchersSpentInt.Add(feeInt)
	vouchersSpentCoin := sdk.NewCoin(types.VoucherDenom, vouchersSpentInt)

	return vouchersSpentCoin, tokensCoin, nil
}
func (k msgServer) takeVoucherAndGiveTokens(ctx sdk.Context, buyer sdk.AccAddress, vouchers, tokens sdk.Coin) error {
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

var feePercentage = sdk.NewDecWithPrec(3, 3)

func subFees(v sdk.Int) sdk.Int {
	return v.MulRaw(997).QuoRaw(1000)
}
