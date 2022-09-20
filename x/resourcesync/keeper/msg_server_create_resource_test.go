package keeper_test

import (
	"github.com/catenax/esc-backbone/x/resourcesync/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"reflect"
	"testing"
)

func Test_msgServer_CreateResource(t *testing.T) {
	type want struct {
		want    *types.MsgCreateResourceResponse
		wantErr bool
		stored  []types.ResourceMap
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
			"One new resource",
			args{msgs: []*types.MsgCreateResource{
				{
					"creator's address",
					&types.Resource{
						Originator:   Alice,
						OrigResId:    "an Id",
						TargetSystem: "some url",
						ResourceKey:  "target system's key",
						DataHash:     []byte("not empty hash"),
					},
				},
			}},
			[]want{
				{
					&types.MsgCreateResourceResponse{},
					false,
					[]types.ResourceMap{
						{
							types.Resource{
								Originator:   Alice,
								OrigResId:    "an Id",
								TargetSystem: "some url",
								ResourceKey:  "target system's key",
								DataHash:     []byte("not empty hash"),
							},
							nil,
						},
					},
				},
			},
		},
		{
			"Two new resources same originator",
			args{msgs: []*types.MsgCreateResource{
				{
					"creator's address",
					&types.Resource{
						Originator:   Alice,
						OrigResId:    "an Id",
						TargetSystem: "some url",
						ResourceKey:  "target system's key",
						DataHash:     []byte("not empty hash"),
					},
				},
				{
					"creator's address",
					&types.Resource{
						Originator:   Alice,
						OrigResId:    "another Id",
						TargetSystem: "some url",
						ResourceKey:  "target system's key",
						DataHash:     []byte("not empty hash"),
					},
				},
			}},
			[]want{
				{
					&types.MsgCreateResourceResponse{},
					false,
					[]types.ResourceMap{
						{
							types.Resource{
								Originator:   Alice,
								OrigResId:    "an Id",
								TargetSystem: "some url",
								ResourceKey:  "target system's key",
								DataHash:     []byte("not empty hash"),
							},
							nil,
						},
					},
				},
				{
					&types.MsgCreateResourceResponse{},
					false,
					[]types.ResourceMap{
						{
							types.Resource{
								Originator:   Alice,
								OrigResId:    "an Id",
								TargetSystem: "some url",
								ResourceKey:  "target system's key",
								DataHash:     []byte("not empty hash"),
							},
							nil,
						},
						{
							types.Resource{
								Originator:   Alice,
								OrigResId:    "another Id",
								TargetSystem: "some url",
								ResourceKey:  "target system's key",
								DataHash:     []byte("not empty hash"),
							},
							nil,
						},
					},
				},
			},
		},
		{
			"Two new resources same id different originators",
			args{msgs: []*types.MsgCreateResource{
				{
					"creator's address",
					&types.Resource{
						Originator:   Alice,
						OrigResId:    "same Id",
						TargetSystem: "some url",
						ResourceKey:  "target system's key",
						DataHash:     []byte("not empty hash"),
					},
				},
				{
					"creator's address",
					&types.Resource{
						Originator:   Bob,
						OrigResId:    "same Id",
						TargetSystem: "some url",
						ResourceKey:  "target system's key",
						DataHash:     []byte("not empty hash"),
					},
				},
			}},
			[]want{
				{
					&types.MsgCreateResourceResponse{},
					false,
					[]types.ResourceMap{
						{
							types.Resource{
								Originator:   Alice,
								OrigResId:    "same Id",
								TargetSystem: "some url",
								ResourceKey:  "target system's key",
								DataHash:     []byte("not empty hash"),
							},
							nil,
						},
					},
				},
				{
					&types.MsgCreateResourceResponse{},
					false,
					[]types.ResourceMap{
						{
							types.Resource{
								Originator:   Alice,
								OrigResId:    "same Id",
								TargetSystem: "some url",
								ResourceKey:  "target system's key",
								DataHash:     []byte("not empty hash"),
							},
							nil,
						},
						{
							types.Resource{
								Originator:   Bob,
								OrigResId:    "same Id",
								TargetSystem: "some url",
								ResourceKey:  "target system's key",
								DataHash:     []byte("not empty hash"),
							},
							nil,
						},
					},
				},
			},
		},
		{
			"Two new resources same originator same Id",
			args{msgs: []*types.MsgCreateResource{
				{
					"creator's address",
					&types.Resource{
						Originator:   Bob,
						OrigResId:    "same Id",
						TargetSystem: "some url",
						ResourceKey:  "target system's key",
						DataHash:     []byte("not empty hash"),
					},
				},
				{
					"creator's address",
					&types.Resource{
						Originator:   Bob,
						OrigResId:    "same Id",
						TargetSystem: "some url",
						ResourceKey:  "target system's key",
						DataHash:     []byte("not empty hash"),
					},
				},
			}},
			[]want{
				{
					&types.MsgCreateResourceResponse{},
					false,
					[]types.ResourceMap{
						{
							types.Resource{
								Originator:   Bob,
								OrigResId:    "same Id",
								TargetSystem: "some url",
								ResourceKey:  "target system's key",
								DataHash:     []byte("not empty hash"),
							},
							nil,
						},
					},
				},
				{
					nil,
					true,
					[]types.ResourceMap{
						{
							types.Resource{
								Originator:   Bob,
								OrigResId:    "same Id",
								TargetSystem: "some url",
								ResourceKey:  "target system's key",
								DataHash:     []byte("not empty hash"),
							},
							nil,
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
				gotStored := keeper.GetAllResourceMap(ctx)

				if !reflect.DeepEqual(gotStored, tt.wants[i].stored) {
					t.Errorf("After CreateResource() keeper stored = %v, want %v", gotStored, tt.wants[i].stored)
				}

			}

		})
	}
}
