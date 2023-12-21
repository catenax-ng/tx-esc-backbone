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
package ubcmm

import (
	"math/rand"

	"github.com/catenax/esc-backbone/testutil/sample"
	ubcsimulation "github.com/catenax/esc-backbone/x/ubcmm/simulation"
	"github.com/catenax/esc-backbone/x/ubcmm/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = ubcsimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgInit = "op_weight_msg_init"
	// TODO: Determine the simulation weight value
	defaultWeightMsgInit int = 100

	opWeightMsgBuy = "op_weight_msg_buy"
	// TODO: Determine the simulation weight value
	defaultWeightMsgBuy int = 100

	opWeightMsgSell = "op_weight_msg_sell"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSell int = 100

	opWeightMsgUndergird = "op_weight_msg_undergird"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUndergird int = 100

	opWeightMsgShiftup = "op_weight_msg_shiftup"
	// TODO: Determine the simulation weight value
	defaultWeightMsgShiftup int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	ubcGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&ubcGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgInit int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgInit, &weightMsgInit, nil,
		func(_ *rand.Rand) {
			weightMsgInit = defaultWeightMsgInit
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgInit,
		ubcsimulation.SimulateMsgInit(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgBuy int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgBuy, &weightMsgBuy, nil,
		func(_ *rand.Rand) {
			weightMsgBuy = defaultWeightMsgBuy
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgBuy,
		ubcsimulation.SimulateMsgBuy(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSell int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSell, &weightMsgSell, nil,
		func(_ *rand.Rand) {
			weightMsgSell = defaultWeightMsgSell
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSell,
		ubcsimulation.SimulateMsgSell(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUndergird int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUndergird, &weightMsgUndergird, nil,
		func(_ *rand.Rand) {
			weightMsgUndergird = defaultWeightMsgUndergird
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUndergird,
		ubcsimulation.SimulateMsgUndergird(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgShiftup int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgShiftup, &weightMsgShiftup, nil,
		func(_ *rand.Rand) {
			weightMsgShiftup = defaultWeightMsgShiftup
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgShiftup,
		ubcsimulation.SimulateMsgShiftup(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgInit,
			defaultWeightMsgInit,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				ubcsimulation.SimulateMsgInit(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgBuy,
			defaultWeightMsgBuy,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				ubcsimulation.SimulateMsgBuy(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgSell,
			defaultWeightMsgSell,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				ubcsimulation.SimulateMsgSell(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUndergird,
			defaultWeightMsgUndergird,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				ubcsimulation.SimulateMsgUndergird(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgShiftup,
			defaultWeightMsgShiftup,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				ubcsimulation.SimulateMsgShiftup(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
