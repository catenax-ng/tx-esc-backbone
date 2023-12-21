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
	"context"
	"github.com/catenax/esc-backbone/testutil"
	"testing"

	keepertest "github.com/catenax/esc-backbone/testutil/keeper"
	"github.com/catenax/esc-backbone/x/resourcesync/keeper"
	"github.com/catenax/esc-backbone/x/resourcesync/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

const (
	Alice = testutil.Alice
	Bob   = testutil.Bob
	Carol = testutil.Carol
)

func setupMsgServer(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
	k, ctx := keepertest.ResourcesyncKeeper(t)
	return keeper.NewMsgServerImpl(*k), *k, sdk.WrapSDKContext(ctx)
}

func TestMsgServer(t *testing.T) {
	ms, keeper, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, keeper)
	require.NotNil(t, ctx)
}

func createValidResouceKey(originator string, origResId string) types.ResourceKey {
	resourceKey, err := types.NewResourceKey(originator, origResId)
	if err != nil {
		panic(err)
	}
	return resourceKey
}
