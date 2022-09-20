package types

import (
	"testing"

	"github.com/catenax/esc-backbone/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateResource_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateResource
		err  error
	}{
		{
			name: "invalid creator address, valid resource",
			msg: MsgCreateResource{
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
			msg: MsgCreateResource{
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
			msg: MsgCreateResource{
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
			msg: MsgCreateResource{
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
