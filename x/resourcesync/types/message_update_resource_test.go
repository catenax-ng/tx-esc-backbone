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

func TestMsgUpdateResource_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateResource
		err  error
	}{
		{
			name: "invalid creator address, valid resource",
			msg: MsgUpdateResource{
				Creator: "invalid_address",
				Entry: &Resource{
					Originator:   sample.AccAddress(),
					OrigResId:    "some id",
					TargetSystem: "some target system",
					ResourceKey:  "some res key",
					DataHash:     []byte("some hase"),
				},
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "valid address, valid resource",
			msg: MsgUpdateResource{
				Creator: sample.AccAddress(),
				Entry: &Resource{
					Originator:   sample.AccAddress(),
					OrigResId:    "some id",
					TargetSystem: "some target system",
					ResourceKey:  "some res key",
					DataHash:     []byte("some hase"),
				},
			},
		},
		{
			name: "invalid address, invalid resource",
			msg: MsgUpdateResource{
				Creator: "invalid_address",
				Entry: &Resource{
					Originator:   "invalid_address",
					OrigResId:    "some id",
					TargetSystem: "some target system",
					ResourceKey:  "some res key",
					DataHash:     []byte("some hase"),
				},
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "valid address, invalid resource",
			msg: MsgUpdateResource{
				Creator: sample.AccAddress(),
				Entry: &Resource{
					Originator:   "invalid_address",
					OrigResId:    "some id",
					TargetSystem: "some target system",
					ResourceKey:  "some res key",
					DataHash:     []byte("some hase"),
				},
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
