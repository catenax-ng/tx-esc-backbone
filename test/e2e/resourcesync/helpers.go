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
package resourcesync

import (
	"context"
	"encoding/json"
	"time"

	"github.com/catenax/esc-backbone/x/resourcesync/client/cli"
	"github.com/catenax/esc-backbone/x/resourcesync/types"
	coretypes "github.com/cometbft/cometbft/rpc/core/types"
	tmtypes "github.com/cometbft/cometbft/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/testutil"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type E2ETestSuite struct {
	suite.Suite

	cfg         network.Config
	network     *network.Network
	queryClient tmservice.ServiceClient
}

func (s *E2ETestSuite) execCommand(val *network.Validator, cmd *cobra.Command, args []string) {
	clientCtx := val.ClientCtx
	cmd.SetArgs(args)
	cmd.Flags().Set(flags.FlagFrom, val.Address.String())
	cmd.Flags().Set(flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String())
	cmd.Flags().Set(flags.FlagSkipConfirmation, "true")
	_, out := testutil.ApplyMockIO(cmd)
	clientCtx = clientCtx.WithOutput(out)

	ctx := context.Background()
	ctx = context.WithValue(ctx, client.ClientContextKey, &clientCtx)
	if err := cmd.ExecuteContext(ctx); err != nil {
		s.Require().Failf("Command failed", "Error: {}\nOutput:\n{}", err, out)
	} else {
		s.T().Log(cmd)
		s.T().Log(out)
	}
}
func (s *E2ETestSuite) execCreateResourceCmd(val *network.Validator, resource types.Resource) {
	cmd := cli.CmdCreateResource()
	r, err := json.Marshal(resource)
	s.Require().NoError(err)
	args := []string{string(r)}
	s.execCommand(val, cmd, args)
}

func (s *E2ETestSuite) execUpdateResourceCmd(val *network.Validator, resource types.Resource) {
	cmd := cli.CmdUpdateResource()
	r, err := json.Marshal(resource)
	s.Require().NoError(err)
	args := []string{string(r)}
	s.execCommand(val, cmd, args)
}

func (s *E2ETestSuite) execDeleteResourceCmd(val *network.Validator, originator string, origResId string) {
	cmd := cli.CmdDeleteResource()
	args := []string{originator, origResId}
	s.execCommand(val, cmd, args)
}

type AssertTypedEvent = func(a *require.Assertions, event proto.Message)

func (s *E2ETestSuite) waitForTypedEvent(val *network.Validator, socketQuery string, duration time.Duration, assert ...AssertTypedEvent) context.Context {
	subscriptionContext, cancelFunc := context.WithTimeout(context.Background(), duration)
	events, err := val.RPCClient.Subscribe(subscriptionContext, "", socketQuery)
	s.Require().NoError(err)
	go func() {
		defer func() {
			_ = val.RPCClient.Unsubscribe(context.Background(), "", socketQuery)
		}()
		assertIndex := 0
		for {
			select {
			case <-subscriptionContext.Done():
				s.Require().False(assertIndex < len(assert), "Not all expected events were received.")
				return
			case event := <-events:
				for _, typedEvent := range convertToTypedEvents(event) {
					assert[assertIndex](s.Require(), typedEvent)
					s.T().Logf("Event %d : %T - %+v", assertIndex, typedEvent, typedEvent)
					assertIndex++
					if assertIndex == len(assert) {
						cancelFunc()
					}
				}
			}
		}
	}()
	return subscriptionContext
}

func convertToTypedEvents(event coretypes.ResultEvent) []proto.Message {
	var typedEvents []proto.Message
	if marshalled, ok := event.Data.(tmtypes.EventDataTx); ok {
		for _, abciEvent := range marshalled.Result.Events {
			if typedEvent, err := sdk.ParseTypedEvent(abciEvent); err == nil && typedEvent != nil {
				typedEvents = append(typedEvents, typedEvent)
			}
		}
	}
	return typedEvents
}
