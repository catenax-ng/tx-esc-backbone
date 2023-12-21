// Copyright (c) 2022-2023 Contributors to the Eclipse Foundation
//
// See the NOTICE file(s) distributed with this work for additional
// information regarding copyright ownership.
//
// This program and the accompanying materials are made available under the
// terms of the Apache License, Version 2.0 which is available at
// https://www.apache.org/licenses/LICENSE-2.0.
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.
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
	cdc.RegisterConcrete(&MsgBuy{}, "ubc/Buy", nil)
	cdc.RegisterConcrete(&MsgSell{}, "ubc/Sell", nil)
	cdc.RegisterConcrete(&MsgUndergird{}, "ubc/Undergird", nil)
	cdc.RegisterConcrete(&MsgShiftup{}, "ubc/Shiftup", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgInit{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBuy{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSell{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUndergird{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgShiftup{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
