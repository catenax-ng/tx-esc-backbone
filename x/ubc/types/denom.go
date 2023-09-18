// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	// SystemTokenMultiplier is the multiplier for cax tokens
	SystemTokenMultiplier = 1e6

	// VoucherMultiplier is the multiplier for voucher
	VoucherMultiplier = 1e6

	// SystemTokenDenom is the denominator for cax tokens
	SystemTokenDenom = "ucax"

	// VoucherDenom is the denominator for voucher
	VoucherDenom = "uvoucher"
)

func isValidDenom(denom string) bool {
	return denom == VoucherDenom || denom == SystemTokenDenom
}

func validateVoucherCoin(value string) error {
	coin, err := sdk.ParseCoinNormalized(value)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "(%s)", err)
	}
	if coin.Denom != VoucherDenom {
		return sdkerrors.Wrapf(ErrInvalidArg, "invalid denom")
	}
	if coin.Amount.IsZero() || coin.Amount.IsNegative() {
		return sdkerrors.Wrapf(ErrInvalidArg, "amount should be a positive integer")
	}
	return nil
}
