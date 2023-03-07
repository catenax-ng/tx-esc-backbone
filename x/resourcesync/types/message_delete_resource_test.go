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

func TestMsgDeleteResource_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteResource
		err  error
	}{
		{
			name: "invalid creator address, valid resource",
			msg: MsgDeleteResource{
				Creator:    "invalid_address",
				Originator: sample.AccAddress(),
				OrigResId:  "some id",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "valid address, valid resource",
			msg: MsgDeleteResource{
				Creator:    sample.AccAddress(),
				Originator: sample.AccAddress(),
				OrigResId:  "some id",
			},
		},
		{
			name: "invalid address, invalid resource",
			msg: MsgDeleteResource{
				Creator:    "invalid_address",
				Originator: "invalid_address",
				OrigResId:  "some id",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "valid address, invalid resource",
			msg: MsgDeleteResource{
				Creator:    sample.AccAddress(),
				Originator: "invalid_address",
				OrigResId:  "some id",
			},
			err: ErrInvalidResource,
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
