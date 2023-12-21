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
