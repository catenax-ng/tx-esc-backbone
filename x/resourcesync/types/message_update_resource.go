package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateResource = "update_resource"

var _ sdk.Msg = &MsgUpdateResource{}

func NewMsgUpdateResource(creator string, entry *Resource) *MsgUpdateResource {
	return &MsgUpdateResource{
		Creator: creator,
		Entry:   entry,
	}
}

func (msg *MsgUpdateResource) Route() string {
	return RouterKey
}

func (msg *MsgUpdateResource) Type() string {
	return TypeMsgUpdateResource
}

func (msg *MsgUpdateResource) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateResource) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateResource) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
