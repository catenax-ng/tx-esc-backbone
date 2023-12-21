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
package keeper_test

import (
	"github.com/catenax/esc-backbone/x/resourcesync/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/require"

	"reflect"
	"testing"
)

func Test_msgServer_DeleteResource(t *testing.T) {

	type want struct {
		want    *types.MsgDeleteResourceResponse
		wantErr bool
		stored  []types.ResourceMap
		events  []proto.Message
	}
	type args struct {
		resources []*types.Resource
		msgs      []*types.MsgDeleteResource
	}
	tests := []struct {
		name  string
		args  args
		wants []want
	}{
		{
			name: "Delete a resource",
			args: args{
				resources: []*types.Resource{
					{
						Originator:   Alice,
						OrigResId:    "an Id",
						TargetSystem: "some url",
						ResourceKey:  "target system's key",
						DataHash:     []byte("not empty hash"),
					},
				},

				msgs: []*types.MsgDeleteResource{
					{
						Creator:    "creator's address",
						Originator: Alice,
						OrigResId:  "an Id",
					},
				},
			},
			wants: []want{
				{
					want:    &types.MsgDeleteResourceResponse{},
					wantErr: false,
					stored:  nil,
					events: []proto.Message{
						&types.EventDeleteResource{
							Creator: "creator's address",
							Resource: types.Resource{
								Originator:   Alice,
								OrigResId:    "an Id",
								TargetSystem: "some url",
								ResourceKey:  "target system's key",
								DataHash:     []byte("not empty hash"),
							},
						},
					},
				},
			},
		},
		{
			name: "Delete first of two resources",
			args: args{
				resources: []*types.Resource{
					{
						Originator:   Alice,
						OrigResId:    "an Id",
						TargetSystem: "some url",
						ResourceKey:  "target system's key",
						DataHash:     []byte("not empty hash"),
					},
					{
						Originator:   Alice,
						OrigResId:    "another Id",
						TargetSystem: "some url",
						ResourceKey:  "target system's key",
						DataHash:     []byte("not empty hash"),
					},
				},

				msgs: []*types.MsgDeleteResource{
					{
						Creator:    "creator's address",
						Originator: Alice,
						OrigResId:  "an Id",
					},
				},
			},
			wants: []want{
				{
					want:    &types.MsgDeleteResourceResponse{},
					wantErr: false,
					stored: []types.ResourceMap{
						{
							Resource: types.Resource{
								Originator:   Alice,
								OrigResId:    "another Id",
								TargetSystem: "some url",
								ResourceKey:  "target system's key",
								DataHash:     []byte("not empty hash"),
							},
						},
					},
					events: []proto.Message{
						&types.EventDeleteResource{
							Creator: "creator's address",
							Resource: types.Resource{
								Originator:   Alice,
								OrigResId:    "an Id",
								TargetSystem: "some url",
								ResourceKey:  "target system's key",
								DataHash:     []byte("not empty hash"),
							},
						},
					},
				},
			},
		},
		{
			name: "Delete second of two resources",
			args: args{
				resources: []*types.Resource{
					{
						Originator:   Alice,
						OrigResId:    "an Id",
						TargetSystem: "some url",
						ResourceKey:  "target system's key",
						DataHash:     []byte("not empty hash"),
					},
					{
						Originator:   Alice,
						OrigResId:    "another Id",
						TargetSystem: "some url",
						ResourceKey:  "target system's key",
						DataHash:     []byte("not empty hash"),
					},
				},

				msgs: []*types.MsgDeleteResource{
					{
						Creator:    "creator's address",
						Originator: Alice,
						OrigResId:  "another Id",
					},
				},
			},
			wants: []want{
				{
					want:    &types.MsgDeleteResourceResponse{},
					wantErr: false,
					stored: []types.ResourceMap{
						{
							Resource: types.Resource{
								Originator:   Alice,
								OrigResId:    "an Id",
								TargetSystem: "some url",
								ResourceKey:  "target system's key",
								DataHash:     []byte("not empty hash"),
							},
						},
					},
					events: []proto.Message{
						&types.EventDeleteResource{
							Creator: "creator's address",
							Resource: types.Resource{
								Originator:   Alice,
								OrigResId:    "another Id",
								TargetSystem: "some url",
								ResourceKey:  "target system's key",
								DataHash:     []byte("not empty hash"),
							},
						},
					},
				},
			},
		},
		{
			name: "Delete two of two resources",
			args: args{
				resources: []*types.Resource{
					{
						Originator:   Alice,
						OrigResId:    "an Id",
						TargetSystem: "some url",
						ResourceKey:  "target system's key",
						DataHash:     []byte("not empty hash"),
					},
					{
						Originator:   Bob,
						OrigResId:    "another Id",
						TargetSystem: "some url",
						ResourceKey:  "target system's key",
						DataHash:     []byte("not empty hash"),
					},
				},

				msgs: []*types.MsgDeleteResource{
					{
						Creator:    "creator's address",
						Originator: Bob,
						OrigResId:  "another Id",
					},
					{
						Creator:    "creator's address",
						Originator: Alice,
						OrigResId:  "an Id",
					},
				},
			},
			wants: []want{
				{
					want:    &types.MsgDeleteResourceResponse{},
					wantErr: false,
					stored: []types.ResourceMap{
						{
							Resource: types.Resource{
								Originator:   Alice,
								OrigResId:    "an Id",
								TargetSystem: "some url",
								ResourceKey:  "target system's key",
								DataHash:     []byte("not empty hash"),
							},
						},
					},
					events: []proto.Message{
						&types.EventDeleteResource{
							Creator: "creator's address",
							Resource: types.Resource{
								Originator:   Bob,
								OrigResId:    "another Id",
								TargetSystem: "some url",
								ResourceKey:  "target system's key",
								DataHash:     []byte("not empty hash"),
							},
						},
					},
				},
				{
					want:    &types.MsgDeleteResourceResponse{},
					wantErr: false,
					stored:  nil,
					events: []proto.Message{
						&types.EventDeleteResource{
							Creator: "creator's address",
							Resource: types.Resource{
								Originator:   Bob,
								OrigResId:    "another Id",
								TargetSystem: "some url",
								ResourceKey:  "target system's key",
								DataHash:     []byte("not empty hash"),
							},
						},
						&types.EventDeleteResource{
							Creator: "creator's address",
							Resource: types.Resource{
								Originator:   Alice,
								OrigResId:    "an Id",
								TargetSystem: "some url",
								ResourceKey:  "target system's key",
								DataHash:     []byte("not empty hash"),
							},
						},
					},
				},
			},
		},
		{
			name: "Try delete a not existing  resource",
			args: args{
				resources: []*types.Resource{
					{
						Originator:   Alice,
						OrigResId:    "an Id",
						TargetSystem: "some url",
						ResourceKey:  "target system's key",
						DataHash:     []byte("not empty hash"),
					},
				},

				msgs: []*types.MsgDeleteResource{
					{
						Creator:    "creator's address",
						Originator: Alice,
						OrigResId:  "nonexistent Id",
					},
				},
			},
			wants: []want{
				{
					want:    nil,
					wantErr: true,
					stored: []types.ResourceMap{
						{
							Resource: types.Resource{
								Originator:   Alice,
								OrigResId:    "an Id",
								TargetSystem: "some url",
								ResourceKey:  "target system's key",
								DataHash:     []byte("not empty hash"),
							},
						},
					},
					events: nil,
				},
			},
		},
		{
			name: "Try delete same resource twice",
			args: args{
				resources: []*types.Resource{
					{
						Originator:   Alice,
						OrigResId:    "an Id",
						TargetSystem: "some url",
						ResourceKey:  "target system's key",
						DataHash:     []byte("not empty hash"),
					},
				},

				msgs: []*types.MsgDeleteResource{
					{
						Creator:    "creator's address",
						Originator: Alice,
						OrigResId:  "an Id",
					},
					{
						Creator:    "creator's address",
						Originator: Alice,
						OrigResId:  "an Id",
					},
				},
			},
			wants: []want{
				{
					want:    &types.MsgDeleteResourceResponse{},
					wantErr: false,
					stored:  nil,
					events: []proto.Message{
						&types.EventDeleteResource{
							Creator: "creator's address",
							Resource: types.Resource{
								Originator:   Alice,
								OrigResId:    "an Id",
								TargetSystem: "some url",
								ResourceKey:  "target system's key",
								DataHash:     []byte("not empty hash"),
							},
						},
					},
				},
				{
					want:    nil,
					wantErr: true,
					stored:  nil,
					events: []proto.Message{
						&types.EventDeleteResource{
							Creator: "creator's address",
							Resource: types.Resource{
								Originator:   Alice,
								OrigResId:    "an Id",
								TargetSystem: "some url",
								ResourceKey:  "target system's key",
								DataHash:     []byte("not empty hash"),
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msgServer, keeper, goCtx := setupMsgServer(t)
			ctx := sdk.UnwrapSDKContext(goCtx)
			require.NotNil(t, ctx)
			for _, resource := range tt.args.resources {
				resourceMap := types.NewResourceMap(*resource)
				keeper.SetResourceMap(ctx, resourceMap)
			}
			for i, msg := range tt.args.msgs {
				got, err := msgServer.DeleteResource(goCtx, msg)
				if (err != nil) != tt.wants[i].wantErr {
					t.Errorf("DeleteResource() error = %v, wantErr %v", err, tt.wants[i].wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.wants[i].want) {
					t.Errorf("DeleteResource() got = %v, want %v", got, tt.wants[i].want)
				}

				gotStored := keeper.GetAllResourceMap(ctx)
				if !reflect.DeepEqual(gotStored, tt.wants[i].stored) {
					t.Errorf("After DeleteResource() keeper stored = %v, want %v", gotStored, tt.wants[i].stored)
				}
				abciEvents := ctx.EventManager().ABCIEvents()
				var gotTypedEvents []proto.Message
				for _, event := range abciEvents {
					typedEvent, err := sdk.ParseTypedEvent(event)
					require.Nil(t, err)
					gotTypedEvents = append(gotTypedEvents, typedEvent)
				}
				if !reflect.DeepEqual(gotTypedEvents, tt.wants[i].events) {
					t.Errorf("DeleteResource() emitted events = %v, want %v", gotTypedEvents, tt.wants[i].events)
				}
			}

		})
	}
}
