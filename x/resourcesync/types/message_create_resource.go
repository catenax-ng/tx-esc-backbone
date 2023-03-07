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

const TypeMsgCreateResource = "create_resource"

var _ sdk.Msg = &MsgCreateResource{}

func NewMsgCreateResource(creator string, entry *Resource) *MsgCreateResource {
	return &MsgCreateResource{
		Creator: creator,
		Entry:   entry,
	}
}

func (msg *MsgCreateResource) Route() string {
	return RouterKey
}

func (msg *MsgCreateResource) Type() string {
	return TypeMsgCreateResource
}

func (msg *MsgCreateResource) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateResource) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateResource) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	err = msg.Entry.Validate()
	if err != nil {
		return sdkerrors.Wrapf(ErrInvalidResource, "resource to create invalid: %s", err)
	}
	return nil
}
