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

const TypeMsgBuytokens = "buytokens"

var _ sdk.Msg = &MsgBuytokens{}

func NewMsgBuytokens(buyer string, value string) *MsgBuytokens {
	return &MsgBuytokens{
		Buyer: buyer,
		Value: value,
	}
}

func (msg *MsgBuytokens) Route() string {
	return RouterKey
}

func (msg *MsgBuytokens) Type() string {
	return TypeMsgBuytokens
}

func (msg *MsgBuytokens) GetSigners() []sdk.AccAddress {
	buyer, err := sdk.AccAddressFromBech32(msg.Buyer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{buyer}
}

func (msg *MsgBuytokens) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBuytokens) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Buyer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid buyer address (%s)", err)
	}
	return validateTokenCoin(msg.Value)
}
