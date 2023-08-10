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

const TypeMsgSelltokens = "selltokens"

var _ sdk.Msg = &MsgSelltokens{}

func NewMsgSelltokens(seller string, value string) *MsgSelltokens {
	return &MsgSelltokens{
		Seller: seller,
		Value:  value,
	}
}

func (msg *MsgSelltokens) Route() string {
	return RouterKey
}

func (msg *MsgSelltokens) Type() string {
	return TypeMsgSelltokens
}

func (msg *MsgSelltokens) GetSigners() []sdk.AccAddress {
	seller, err := sdk.AccAddressFromBech32(msg.Seller)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{seller}
}

func (msg *MsgSelltokens) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSelltokens) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Seller)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid seller address (%s)", err)
	}
	return validateTokenCoin(msg.Value)
}
