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
