// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package types

import (
	"testing"

	"github.com/catenax/esc-backbone/testutil/sample"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgShiftup_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgShiftup
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgShiftup{
				Operator:        "invalid_address",
				Voucherstoadd:   "10" + VoucherDenom,
				Degirdingfactor: sdk.ZeroDec(),
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid vouchers to add (zero)",
			msg: MsgShiftup{
				Operator:        sample.AccAddress(),
				Voucherstoadd:   "0" + VoucherDenom,
				Degirdingfactor: sdk.ZeroDec(),
			},
			err: ErrInvalidArg,
		}, {
			name: "invalid vouchers to add (negative)",
			msg: MsgShiftup{
				Operator:        sample.AccAddress(),
				Voucherstoadd:   "-10" + VoucherDenom,
				Degirdingfactor: sdk.ZeroDec(),
			},
			err: sdkerrors.ErrInvalidCoins,
		}, {
			name: "invalid vouchers to add (decimal)",
			msg: MsgShiftup{
				Operator:        sample.AccAddress(),
				Voucherstoadd:   "10.0" + VoucherDenom,
				Degirdingfactor: sdk.ZeroDec(),
			},
			err: nil,
		}, {
			name: "invalid vouchers to add (non number)",
			msg: MsgShiftup{
				Operator:        sample.AccAddress(),
				Voucherstoadd:   "abcd" + VoucherDenom,
				Degirdingfactor: sdk.ZeroDec(),
			},
			err: sdkerrors.ErrInvalidCoins,
		}, {
			name: "invalid vouchers to add (empty)",
			msg: MsgShiftup{
				Operator:        sample.AccAddress(),
				Voucherstoadd:   "",
				Degirdingfactor: sdk.ZeroDec(),
			},
			err: sdkerrors.ErrInvalidCoins,
		}, {
			name: "valid address",
			msg: MsgShiftup{
				Operator:        sample.AccAddress(),
				Voucherstoadd:   "10" + VoucherDenom,
				Degirdingfactor: sdk.ZeroDec(),
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
