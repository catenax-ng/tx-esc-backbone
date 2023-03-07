// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
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
