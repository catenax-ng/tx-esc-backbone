// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgInit{}, "ubc/Init", nil)
	cdc.RegisterConcrete(&MsgBuytokens{}, "ubc/Buytokens", nil)
	cdc.RegisterConcrete(&MsgSelltokens{}, "ubc/Selltokens", nil)
	cdc.RegisterConcrete(&MsgUndergird{}, "ubc/Undergird", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgInit{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBuytokens{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSelltokens{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUndergird{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
