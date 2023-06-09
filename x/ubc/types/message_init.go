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

const TypeMsgInit = "init"

var _ sdk.Msg = &MsgInit{}

func NewMsgInit(creator string, refTokenSupply string, refTokenPrice string, refProfitFactor string, bPool string, bPoolUnder string, slopeP2 string, slopeP3 string, factorFy string, factorFxy string) *MsgInit {
	return &MsgInit{
		Creator:         creator,
		RefTokenSupply:  refTokenSupply,
		RefTokenPrice:   refTokenPrice,
		RefProfitFactor: refProfitFactor,
		BPool:           bPool,
		BPoolUnder:      bPoolUnder,
		SlopeP2:         slopeP2,
		SlopeP3:         slopeP3,
		FactorFy:        factorFy,
		FactorFxy:       factorFxy,
	}
}

func (msg *MsgInit) Route() string {
	return RouterKey
}

func (msg *MsgInit) Type() string {
	return TypeMsgInit
}

func (msg *MsgInit) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgInit) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgInit) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
