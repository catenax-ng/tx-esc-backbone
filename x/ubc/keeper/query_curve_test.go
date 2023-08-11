// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/catenax/esc-backbone/testutil/keeper"
	"github.com/catenax/esc-backbone/testutil/nullify"
	"github.com/catenax/esc-backbone/x/ubc/types"
)

func TestUbcobjectQuery(t *testing.T) {
	keeper, ctx := keepertest.UbcKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestUbcobject(keeper, ctx)
	tests := []struct {
		desc     string
		request  *types.QueryGetUbcobjectRequest
		response *types.QueryGetUbcobjectResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetUbcobjectRequest{},
			response: &types.QueryGetUbcobjectResponse{Ubcobject: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Ubcobject(wctx, tc.request)
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
