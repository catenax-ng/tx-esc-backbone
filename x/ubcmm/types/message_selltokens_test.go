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

func TestMsgSell_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgSell
		err  error
	}{
		{
			name: "valid",
			msg: MsgSell{
				Seller: sample.AccAddress(),
				Value:  "10" + SystemTokenDenom,
			},
			err: nil,
		}, {
			name: "zero value",
			msg: MsgSell{
				Seller: sample.AccAddress(),
				Value:  "0" + SystemTokenDenom,
			},
			err: ErrInvalidArg,
		}, {
			name: "negative value",
			msg: MsgSell{
				Seller: sample.AccAddress(),
				Value:  "-5" + SystemTokenDenom,
			},
			err: sdkerrors.ErrInvalidCoins,
		}, {
			name: "invalid address",
			msg: MsgSell{
				Seller: "invalid_address",
				Value:  "10" + SystemTokenDenom,
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid value",
			msg: MsgSell{
				Seller: sample.AccAddress(),
				Value:  "abcd",
			},
			err: sdkerrors.ErrInvalidCoins,
		}, {
			name: "invalid denom",
			msg: MsgSell{
				Seller: sample.AccAddress(),
				Value:  "2" + SystemTokenDenom + "x",
			},
			err: ErrInvalidArg,
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
