// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package types

import (
	"testing"

	"github.com/catenax/esc-backbone/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgUndergird_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUndergird
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUndergird{
				Operator:      "invalid_address",
				Voucherstoadd: "10" + VoucherDenom,
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid vouchers to add (zero)",
			msg: MsgUndergird{
				Operator:      sample.AccAddress(),
				Voucherstoadd: "0" + VoucherDenom,
			},
			err: ErrInvalidArg,
		}, {
			name: "invalid vouchers to add (negative)",
			msg: MsgUndergird{
				Operator:      sample.AccAddress(),
				Voucherstoadd: "-10" + VoucherDenom,
			},
			err: sdkerrors.ErrInvalidCoins,
		}, {
			name: "invalid vouchers to add (decimal)",
			msg: MsgUndergird{
				Operator:      sample.AccAddress(),
				Voucherstoadd: "10.0" + VoucherDenom,
			},
			err: nil,
		}, {
			name: "invalid vouchers to add (non number)",
			msg: MsgUndergird{
				Operator:      sample.AccAddress(),
				Voucherstoadd: "abcd" + VoucherDenom,
			},
			err: sdkerrors.ErrInvalidCoins,
		}, {
			name: "invalid vouchers to add (empty)",
			msg: MsgUndergird{
				Operator:      sample.AccAddress(),
				Voucherstoadd: "",
			},
			err: sdkerrors.ErrInvalidCoins,
		}, {
			name: "valid address",
			msg: MsgUndergird{
				Operator:      sample.AccAddress(),
				Voucherstoadd: "10" + VoucherDenom,
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
