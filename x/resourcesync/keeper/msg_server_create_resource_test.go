package keeper_test

import (
	"github.com/catenax/esc-backbone/x/resourcesync/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func Test_msgServer_CreateResource(t *testing.T) {
	type want struct {
		want    *types.MsgCreateResourceResponse
		wantErr bool
		stored  []types.ResourceMap
		events  []proto.Message
	}
	type args struct {
		msgs []*types.MsgCreateResource
	}
	tests := []struct {
		name  string
		args  args
		wants []want
	}{
		{
			name: "One new resource",
			args: args{msgs: []*types.MsgCreateResource{
				{
					Creator: "creator's address",
					Entry: &types.Resource{
						Originator:   Alice,
						OrigResId:    "an Id",
						TargetSystem: "some url",
						ResourceKey:  "target system's key",
						DataHash:     []byte("not empty hash"),
					},
				},
			}},
			wants: []want{
				{
					want:    &types.MsgCreateResourceResponse{},
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
						&types.EventCreateResource{
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
			name: "Two new resources same originator",
			args: args{msgs: []*types.MsgCreateResource{
				{
					Creator: "creator's address",
					Entry: &types.Resource{
						Originator:   Alice,
						OrigResId:    "an Id",
						TargetSystem: "some url",
						ResourceKey:  "target system's key",
						DataHash:     []byte("not empty hash"),
					},
				},
				{
					Creator: "creator's address",
					Entry: &types.Resource{
						Originator:   Alice,
						OrigResId:    "another Id",
						TargetSystem: "some url",
						ResourceKey:  "target system's key",
						DataHash:     []byte("not empty hash"),
					},
				},
			}},
			wants: []want{
				{
					want:    &types.MsgCreateResourceResponse{},
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
						&types.EventCreateResource{
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
					want:    &types.MsgCreateResourceResponse{},
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
						&types.EventCreateResource{
							Creator: "creator's address",
							Resource: types.Resource{
								Originator:   Alice,
								OrigResId:    "an Id",
								TargetSystem: "some url",
								ResourceKey:  "target system's key",
								DataHash:     []byte("not empty hash"),
							},
						},
						&types.EventCreateResource{
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
			name: "Two new resources same id different originators",
			args: args{msgs: []*types.MsgCreateResource{
				{
					Creator: "creator's address",
					Entry: &types.Resource{
						Originator:   Alice,
						OrigResId:    "same Id",
						TargetSystem: "some url",
						ResourceKey:  "target system's key",
						DataHash:     []byte("not empty hash"),
					},
				},
				{
					Creator: "creator's address",
					Entry: &types.Resource{
						Originator:   Bob,
						OrigResId:    "same Id",
						TargetSystem: "some url",
						ResourceKey:  "target system's key",
						DataHash:     []byte("not empty hash"),
					},
				},
			}},
			wants: []want{
				{
					want:    &types.MsgCreateResourceResponse{},
					wantErr: false,
					stored: []types.ResourceMap{
						{
							Resource: types.Resource{
								Originator:   Alice,
								OrigResId:    "same Id",
								TargetSystem: "some url",
								ResourceKey:  "target system's key",
								DataHash:     []byte("not empty hash"),
							},
						},
					},
					events: []proto.Message{
						&types.EventCreateResource{
							Creator: "creator's address",
							Resource: types.Resource{
								Originator:   Alice,
								OrigResId:    "same Id",
								TargetSystem: "some url",
								ResourceKey:  "target system's key",
								DataHash:     []byte("not empty hash"),
							},
						},
					},
				},
				{
					want:    &types.MsgCreateResourceResponse{},
					wantErr: false,
					stored: []types.ResourceMap{
						{
							Resource: types.Resource{
								Originator:   Alice,
								OrigResId:    "same Id",
								TargetSystem: "some url",
								ResourceKey:  "target system's key",
								DataHash:     []byte("not empty hash"),
							},
						},
						{
							Resource: types.Resource{
								Originator:   Bob,
								OrigResId:    "same Id",
								TargetSystem: "some url",
								ResourceKey:  "target system's key",
								DataHash:     []byte("not empty hash"),
							},
						},
					},
					events: []proto.Message{
						&types.EventCreateResource{
							Creator: "creator's address",
							Resource: types.Resource{
								Originator:   Alice,
								OrigResId:    "same Id",
								TargetSystem: "some url",
								ResourceKey:  "target system's key",
								DataHash:     []byte("not empty hash"),
							},
						},
						&types.EventCreateResource{
							Creator: "creator's address",
							Resource: types.Resource{
								Originator:   Bob,
								OrigResId:    "same Id",
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
			name: "Two new resources same originator same Id",
			args: args{msgs: []*types.MsgCreateResource{
				{
					Creator: "creator's address",
					Entry: &types.Resource{
						Originator:   Bob,
						OrigResId:    "same Id",
						TargetSystem: "some url",
						ResourceKey:  "target system's key",
						DataHash:     []byte("not empty hash"),
					},
				},
				{
					Creator: "creator's address",
					Entry: &types.Resource{
						Originator:   Bob,
						OrigResId:    "same Id",
						TargetSystem: "some url",
						ResourceKey:  "target system's key",
						DataHash:     []byte("not empty hash"),
					},
				},
			}},
			wants: []want{
				{
					want:    &types.MsgCreateResourceResponse{},
					wantErr: false,
					stored: []types.ResourceMap{
						{
							Resource: types.Resource{
								Originator:   Bob,
								OrigResId:    "same Id",
								TargetSystem: "some url",
								ResourceKey:  "target system's key",
								DataHash:     []byte("not empty hash"),
							},
						},
					},
					events: []proto.Message{
						&types.EventCreateResource{
							Creator: "creator's address",
							Resource: types.Resource{
								Originator:   Bob,
								OrigResId:    "same Id",
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
					stored: []types.ResourceMap{
						{
							Resource: types.Resource{
								Originator:   Bob,
								OrigResId:    "same Id",
								TargetSystem: "some url",
								ResourceKey:  "target system's key",
								DataHash:     []byte("not empty hash"),
							},
						},
					},
					events: []proto.Message{
						&types.EventCreateResource{
							Creator: "creator's address",
							Resource: types.Resource{
								Originator:   Bob,
								OrigResId:    "same Id",
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
			for i, msg := range tt.args.msgs {
				got, err := msgServer.CreateResource(goCtx, msg)
				if (err != nil) != tt.wants[i].wantErr {
					t.Errorf("CreateResource() error = %v, wantErr %v", err, tt.wants[i].wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.wants[i].want) {
					t.Errorf("CreateResource() got = %v, want %v", got, tt.wants[i].want)
				}
				ctx := sdk.UnwrapSDKContext(goCtx)
				require.NotNil(t, ctx)
				gotStored := keeper.GetAllResourceMap(ctx)
				if !reflect.DeepEqual(gotStored, tt.wants[i].stored) {
					t.Errorf("After CreateResource() keeper stored = %v, want %v", gotStored, tt.wants[i].stored)
				}
				abciEvents := ctx.EventManager().ABCIEvents()
				var gotTypedEvents []proto.Message
				for _, event := range abciEvents {
					typedEvent, err := sdk.ParseTypedEvent(event)
					require.Nil(t, err)
					gotTypedEvents = append(gotTypedEvents, typedEvent)
				}
				if !reflect.DeepEqual(gotTypedEvents, tt.wants[i].events) {
					t.Errorf("CreateResource() emitted events = %v, want %v", gotTypedEvents, tt.wants[i].events)
				}
			}

		})
	}
}
