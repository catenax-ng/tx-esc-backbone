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
package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	// SystemTokenMultiplier is the multiplier for cax tokens
	SystemTokenMultiplier = 1e9

	// VoucherMultiplier is the multiplier for voucher
	VoucherMultiplier = 1e2

	// SystemTokenDenom is the denominator for cax tokens
	SystemTokenDenom = "ncax"

	// VoucherDenom is the denominator for voucher
	VoucherDenom = "cvoucher"
)

func validateTokenCoin(value string) error {
	coin, err := sdk.ParseCoinNormalized(value)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "(%s)", err)
	}
	if coin.Denom != SystemTokenDenom {
		return sdkerrors.Wrapf(ErrInvalidArg, "invalid denom")
	}
	if coin.Amount.IsZero() {
		return sdkerrors.Wrapf(ErrInvalidArg, "amount is zero")
	}
	return nil
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
