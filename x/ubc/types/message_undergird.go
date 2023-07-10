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

const TypeMsgUndergird = "undergird"

var _ sdk.Msg = &MsgUndergird{}

func NewMsgUndergird(operator string, voucherstoadd string) *MsgUndergird {
	return &MsgUndergird{
		Operator:      operator,
		Voucherstoadd: voucherstoadd,
	}
}

func (msg *MsgUndergird) Route() string {
	return RouterKey
}

func (msg *MsgUndergird) Type() string {
	return TypeMsgUndergird
}

func (msg *MsgUndergird) GetSigners() []sdk.AccAddress {
	operator, err := sdk.AccAddressFromBech32(msg.Operator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{operator}
}

func (msg *MsgUndergird) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUndergird) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Operator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid operator address (%s)", err)
	}
	return nil
}
