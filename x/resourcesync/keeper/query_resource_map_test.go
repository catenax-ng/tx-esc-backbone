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
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/catenax/esc-backbone/testutil/keeper"
	"github.com/catenax/esc-backbone/testutil/nullify"
	"github.com/catenax/esc-backbone/x/resourcesync/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestResourceMapQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.ResourcesyncKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNResourceMap(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetResourceMapRequest
		response *types.QueryGetResourceMapResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetResourceMapRequest{
				Originator: msgs[0].Originator,
				OrigResId:  msgs[0].OrigResId,
			},
			response: &types.QueryGetResourceMapResponse{ResourceMap: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetResourceMapRequest{
				Originator: msgs[1].Originator,
				OrigResId:  msgs[1].OrigResId,
			},
			response: &types.QueryGetResourceMapResponse{ResourceMap: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetResourceMapRequest{
				Originator: Carol,
				OrigResId:  strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc:    "InvalidRequest",
			request: nil,
			err:     status.Error(codes.InvalidArgument, "invalid request"),
		},
		{
			desc: "Invalid originator request",
			request: &types.QueryGetResourceMapRequest{
				Originator: "invalid",
				OrigResId:  "5",
			},
			err: status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.ResourceMap(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestResourceMapQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.ResourcesyncKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNResourceMap(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllResourceMapRequest {
		return &types.QueryAllResourceMapRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ResourceMapAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ResourceMap), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ResourceMap),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ResourceMapAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.ResourceMap), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.ResourceMap),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ResourceMapAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.ResourceMap),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ResourceMapAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
