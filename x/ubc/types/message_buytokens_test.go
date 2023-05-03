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

func TestMsgBuytokens_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgBuytokens
		err  error
	}{
		{
			name: "valid",
			msg: MsgBuytokens{
				Buyer: sample.AccAddress(),
				Value: "10" + CaxDenom,
			},
			err: nil,
		}, {
			name: "zero value",
			msg: MsgBuytokens{
				Buyer: sample.AccAddress(),
				Value: "0" + CaxDenom,
			},
			err: ErrInvalidArg,
		}, {
			name: "negative value",
			msg: MsgBuytokens{
				Buyer: sample.AccAddress(),
				Value: "-5" + CaxDenom,
			},
			err: sdkerrors.ErrInvalidCoins,
		}, {
			name: "invalid address",
			msg: MsgBuytokens{
				Buyer: "invalid_address",
				Value: "10" + CaxDenom,
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid value",
			msg: MsgBuytokens{
				Buyer: sample.AccAddress(),
				Value: "abcd",
			},
			err: sdkerrors.ErrInvalidCoins,
		}, {
			name: "invalid denom",
			msg: MsgBuytokens{
				Buyer: sample.AccAddress(),
				Value: "2" + CaxDenom + "x",
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
