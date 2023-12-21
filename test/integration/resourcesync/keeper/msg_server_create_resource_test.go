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
)

const (
	TARGET_SYS1 = "http://some-registry"
	TARGET_SYS2 = "ftp://a-registry"
)

func (suite *KeeperIntegrationTestSuite) TestCreateResourceIsSaved() {
	suite.T().Skipf("Skip until fixed with https://jira.catena-x.net/browse/CGE-269")
	// FIXME https://jira.catena-x.net/browse/CGE-269
	goCtx := sdk.WrapSDKContext(suite.ctx)
	originator := alice
	origResId := "1"
	suite.msgServer.CreateResource(goCtx, &types.MsgCreateResource{
		Creator: bob,
		Entry: &types.Resource{
			Originator:   originator,
			OrigResId:    origResId,
			TargetSystem: TARGET_SYS1,
			ResourceKey:  "target_sys1_1",
			DataHash:     []byte("some hash"),
		},
	})
	keeper := suite.app.ResourcesyncKeeper
	resKey, err := types.NewResourceKey(originator, origResId)
	suite.Require().Nilf(err, "Failed resource key creation for (originator=%s , origResId=%s)", originator, origResId)
	resourceMap, found := keeper.GetResourceMap(suite.ctx, resKey)
	suite.Require().True(found)
	suite.Require().EqualValues(types.ResourceMap{
		Resource: types.Resource{
			Originator:   originator,
			OrigResId:    origResId,
			TargetSystem: TARGET_SYS1,
			ResourceKey:  "target_sys1_1",
			DataHash:     []byte("some hash"),
		},
		AuditLogs: nil,
	}, resourceMap)

	expectedTypedEvents := []proto.Message{
		&types.EventCreateResource{
			Creator: bob,
			Resource: types.Resource{
				Originator:   originator,
				OrigResId:    origResId,
				TargetSystem: TARGET_SYS1,
				ResourceKey:  "target_sys1_1",
				DataHash:     []byte("some hash"),
			},
		},
	}

	var actualTypedEvents []proto.Message
	abciEvents := suite.ctx.EventManager().ABCIEvents()
	for _, event := range abciEvents {
		typedEvent, err := sdk.ParseTypedEvent(event)
		suite.Require().Nilf(err, "Failed parsing event", event)
		actualTypedEvents = append(actualTypedEvents, typedEvent)
	}

	suite.Require().EqualValues(expectedTypedEvents, actualTypedEvents)
}
