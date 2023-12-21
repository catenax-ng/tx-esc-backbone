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
package resourcesync_test

import (
	"testing"

	keepertest "github.com/catenax/esc-backbone/testutil/keeper"
	"github.com/catenax/esc-backbone/testutil/nullify"
	"github.com/catenax/esc-backbone/x/resourcesync"
	"github.com/catenax/esc-backbone/x/resourcesync/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		ResourceMapList: []types.ResourceMap{
			{
				Resource: types.Resource{
					Originator:   "0",
					OrigResId:    "0",
					TargetSystem: "0",
					ResourceKey:  "0",
					DataHash:     nil,
				},
				AuditLogs: nil,
			},
			{
				Resource: types.Resource{
					Originator:   "1",
					OrigResId:    "1",
					TargetSystem: "1",
					ResourceKey:  "1",
					DataHash:     nil,
				},
				AuditLogs: nil,
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ResourcesyncKeeper(t)
	resourcesync.InitGenesis(ctx, *k, genesisState)
	got := resourcesync.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.ResourceMapList, got.ResourceMapList)
	// this line is used by starport scaffolding # genesis/test/assert
}
