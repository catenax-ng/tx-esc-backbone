package resourcesync

import (
	"context"
	"testing"
	"time"

	"github.com/catenax/esc-backbone/x/resourcesync/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/suite"

	"github.com/catenax/esc-backbone/testutil/network"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	_ "github.com/cosmos/cosmos-sdk/x/distribution"
	_ "github.com/cosmos/cosmos-sdk/x/gov"
)

func TestE2ETestSuite(t *testing.T) {
	suite.Run(t, new(E2ETestSuite))
}

func (s *E2ETestSuite) SetupSuite() {
	s.T().Log("setting up e2e test suite")

	s.cfg = network.DefaultConfig()

	s.network = network.New(s.T(), s.cfg)

	s.Require().NoError(s.network.WaitForNextBlock())

	s.queryClient = tmservice.NewServiceClient(s.network.Validators[0].ClientCtx)
}

func (s *E2ETestSuite) TearDownSuite() {
	s.T().Log("tearing down e2e test suite")
}

type ResourceOption = func(r *types.Resource) *types.Resource

func WithTargetSystem(targetSystem string) ResourceOption {
	return func(r *types.Resource) *types.Resource {
		r.TargetSystem = targetSystem
		return r
	}
}

func WithResourceKey(resourceKey string) ResourceOption {
	return func(r *types.Resource) *types.Resource {
		r.ResourceKey = resourceKey
		return r
	}
}
func expectCreateResource(originator string, origResid string, options ...ResourceOption) AssertTypedEvent {
	return func(a *require.Assertions, e proto.Message) {
		event, ok := e.(*types.EventCreateResource)
		a.True(ok, "Event is not a EventCreateResource: %s", e)
		expect := types.Resource{
			Originator:   originator,
			OrigResId:    origResid,
			TargetSystem: "",
			ResourceKey:  "",
			DataHash:     nil,
		}
		for _, applyOption := range options {
			applyOption(&expect)
		}
		a.EqualValues(expect, event.Resource)
	}
}

func expectDeleteResource(originator string, origResid string, options ...ResourceOption) AssertTypedEvent {
	return func(a *require.Assertions, e proto.Message) {
		event, ok := e.(*types.EventDeleteResource)
		a.True(ok, "Event is not a EventDeleteResource: %s", e)
		expect := types.Resource{Originator: originator, OrigResId: origResid}
		for _, applyOption := range options {
			applyOption(&expect)
		}
		a.EqualValues(expect, event.Resource)
	}
}

func expectUpdateResource(originator string, origResid string, options ...ResourceOption) AssertTypedEvent {
	return func(a *require.Assertions, e proto.Message) {
		event, ok := e.(*types.EventUpdateResource)
		a.True(ok, "Event is not a EventUpdateResource: %s", e)
		expect := types.Resource{Originator: originator, OrigResId: origResid}
		for _, applyOption := range options {
			applyOption(&expect)
		}
		a.EqualValues(expect, event.Resource)
	}
}

func (s *E2ETestSuite) TestCreateResourceEventSubscription() {
	val := s.network.Validators[0]
	res, err := s.queryClient.GetNodeInfo(context.Background(), &tmservice.GetNodeInfoRequest{})
	s.Require().NoError(err)
	s.Require().Equal(res.ApplicationVersion.AppName, version.NewInfo().AppName)
	socketQuery := "tm.event='Tx' AND message.action='/escbackbone.resourcesync.MsgCreateResource'"
	testContext := s.waitForTypedEvent(val, socketQuery, time.Second*10,
		expectCreateResource(val.Address.String(), "1"),
	)
	go s.execCreateResourceCmd(val, types.Resource{Originator: val.Address.String(), OrigResId: "1"})
	select {
	case <-testContext.Done():
		if context.Canceled != testContext.Err() {
			s.Require().NoError(testContext.Err())
		}
		break
	}
}

func (s *E2ETestSuite) TestDeleteResourceEventSubscription() {
	val := s.network.Validators[0]
	res, err := s.queryClient.GetNodeInfo(context.Background(), &tmservice.GetNodeInfoRequest{})
	s.Require().NoError(err)
	s.Require().Equal(res.ApplicationVersion.AppName, version.NewInfo().AppName)
	socketQuery := "tm.event='Tx' AND message.action='/escbackbone.resourcesync.MsgDeleteResource'"
	testContext := s.waitForTypedEvent(val, socketQuery, time.Second*20,
		expectDeleteResource(val.Address.String(), "2"),
	)
	go func() {
		s.execCreateResourceCmd(val, types.Resource{Originator: val.Address.String(), OrigResId: "2"})
		// Waiting for the next block, so that, the account sequence number used for transaction create matches the chains state.
		// If more than one transaction of the same sender shall be in one block. The transactions have to be modified in more detail.
		s.network.WaitForNextBlock()
		s.execDeleteResourceCmd(val, val.Address.String(), "2")
	}()
	select {
	case <-testContext.Done():
		if context.Canceled != testContext.Err() {
			s.Require().NoError(testContext.Err())
		}
	}
}

func (s *E2ETestSuite) TestUpdateResourceEventSubscription() {
	val := s.network.Validators[0]
	res, err := s.queryClient.GetNodeInfo(context.Background(), &tmservice.GetNodeInfoRequest{})
	s.Require().NoError(err)
	s.Require().Equal(res.ApplicationVersion.AppName, version.NewInfo().AppName)
	socketQuery := "tm.event='Tx' AND message.action='/escbackbone.resourcesync.MsgUpdateResource'"
	testContext := s.waitForTypedEvent(val, socketQuery, time.Second*20,
		expectUpdateResource(val.Address.String(), "3", WithResourceKey("updated key"), WithTargetSystem("updated target system.")),
	)
	go func() {
		s.execCreateResourceCmd(val, types.Resource{Originator: val.Address.String(), OrigResId: "3"})
		// Waiting for the next block, so that, the account sequence number used for transaction create matches the chains state.
		// If more than one transaction of the same sender shall be in one block. The transactions have to be modified in more detail.
		s.network.WaitForNextBlock()
		s.execUpdateResourceCmd(val, types.Resource{Originator: val.Address.String(), OrigResId: "3", ResourceKey: "updated key", TargetSystem: "updated target system."})
	}()
	select {
	case <-testContext.Done():
		if context.Canceled != testContext.Err() {
			s.Require().NoError(testContext.Err())
		}
	}
}

func (s *E2ETestSuite) TestCreateDuplicateResourceEventSubscription() {
	val := s.network.Validators[0]
	res, err := s.queryClient.GetNodeInfo(context.Background(), &tmservice.GetNodeInfoRequest{})
	s.Require().NoError(err)
	s.Require().Equal(res.ApplicationVersion.AppName, version.NewInfo().AppName)
	socketQuery := "tm.event='Tx' AND message.action='/escbackbone.resourcesync.MsgCreateResource'"
	testContext := s.waitForTypedEvent(val, socketQuery, time.Second*10,
		expectCreateResource(val.Address.String(), "4"),
	)
	go func() {
		s.execCreateResourceCmd(val, types.Resource{Originator: val.Address.String(), OrigResId: "4"})
		// Waiting for the next block, so that, the account sequence number used for transaction create matches the chains state.
		// If more than one transaction of the same sender shall be in one block. The transactions have to be modified in more detail.
		s.network.WaitForNextBlock()
		s.execCreateResourceCmd(val, types.Resource{Originator: val.Address.String(), OrigResId: "4"})
	}()
	select {
	case <-testContext.Done():
		if context.Canceled != testContext.Err() {
			s.Require().NoError(testContext.Err())
		}
		break
	}
}
