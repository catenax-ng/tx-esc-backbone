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

const TypeMsgDeleteResource = "delete_resource"

var _ sdk.Msg = &MsgDeleteResource{}

func NewMsgDeleteResource(creator string, originator string, origResId string) *MsgDeleteResource {
	return &MsgDeleteResource{
		Creator:    creator,
		Originator: originator,
		OrigResId:  origResId,
	}
}

func (msg *MsgDeleteResource) Route() string {
	return RouterKey
}

func (msg *MsgDeleteResource) Type() string {
	return TypeMsgDeleteResource
}

func (msg *MsgDeleteResource) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteResource) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteResource) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	_, err = NewResourceKeyForDelete(msg)
	if err != nil {
		return sdkerrors.Wrapf(ErrInvalidResource, "resource to delete invalid: %s", err)
	}
	return nil
}
