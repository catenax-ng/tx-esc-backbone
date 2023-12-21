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
package resourcesync

import (
	"math/rand"

	"github.com/catenax/esc-backbone/testutil/sample"
	resourcesyncsimulation "github.com/catenax/esc-backbone/x/resourcesync/simulation"
	"github.com/catenax/esc-backbone/x/resourcesync/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = resourcesyncsimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgCreateResource = "op_weight_msg_create_resource"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateResource int = 100

	opWeightMsgDeleteResource = "op_weight_msg_delete_resource"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteResource int = 100

	opWeightMsgUpdateResource = "op_weight_msg_update_resource"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateResource int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	resourcesyncGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&resourcesyncGenesis)
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

	var weightMsgCreateResource int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateResource, &weightMsgCreateResource, nil,
		func(_ *rand.Rand) {
			weightMsgCreateResource = defaultWeightMsgCreateResource
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateResource,
		resourcesyncsimulation.SimulateMsgCreateResource(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteResource int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteResource, &weightMsgDeleteResource, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteResource = defaultWeightMsgDeleteResource
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteResource,
		resourcesyncsimulation.SimulateMsgDeleteResource(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateResource int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateResource, &weightMsgUpdateResource, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateResource = defaultWeightMsgUpdateResource
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateResource,
		resourcesyncsimulation.SimulateMsgUpdateResource(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateResource,
			defaultWeightMsgCreateResource,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				resourcesyncsimulation.SimulateMsgCreateResource(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteResource,
			defaultWeightMsgDeleteResource,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				resourcesyncsimulation.SimulateMsgDeleteResource(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateResource,
			defaultWeightMsgUpdateResource,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				resourcesyncsimulation.SimulateMsgUpdateResource(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
