// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package types_test

import (
	"testing"

	"github.com/catenax/esc-backbone/x/resourcesync/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

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
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated resourceMap",
			genState: &types.GenesisState{
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
							Originator:   "0",
							OrigResId:    "0",
							TargetSystem: "1",
							ResourceKey:  "1",
							DataHash:     nil,
						},
						AuditLogs: nil,
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}

func TestDefaultGenesisIsCorrect(t *testing.T) {
	require.EqualValues(t,
		&types.GenesisState{
			ResourceMapList: []types.ResourceMap{},
		},
		types.DefaultGenesis())
}
