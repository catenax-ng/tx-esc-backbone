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
	"encoding/json"
	abci "github.com/cometbft/cometbft/abci/types"
	tmtypes "github.com/cometbft/cometbft/types"
	"io"
	"testing"
	"time"

	escapp "github.com/catenax/esc-backbone/app"
	"github.com/catenax/esc-backbone/testutil"
	"github.com/catenax/esc-backbone/x/resourcesync/keeper"
	"github.com/catenax/esc-backbone/x/resourcesync/types"
	dbm "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/suite"
)

const (
	alice = testutil.Alice
	bob   = testutil.Bob
	carol = testutil.Carol
)

type KeeperIntegrationTestSuite struct {
	suite.Suite

	app         *escapp.App
	msgServer   types.MsgServer
	ctx         sdktypes.Context
	queryClient types.QueryClient
}

func TestResourceSyncKeeperTestSuite(t *testing.T) {
	t.Skipf("Skip until fixed with https://jira.catena-x.net/browse/CGE-269")
	// FIXME https://jira.catena-x.net/browse/CGE-269
	suite.Run(t, new(KeeperIntegrationTestSuite))
}

var DefaultConsensusParams = &tmproto.ConsensusParams{
	Block: &tmproto.BlockParams{
		MaxBytes: 200000,
		MaxGas:   2000000,
	},
	Evidence: &tmproto.EvidenceParams{
		MaxAgeNumBlocks: 302400,
		MaxAgeDuration:  504 * time.Hour, // 3 weeks is the max duration
		MaxBytes:        10000,
	},
	Validator: &tmproto.ValidatorParams{
		PubKeyTypes: []string{
			tmtypes.ABCIPubKeyTypeEd25519,
		},
	},
}

type KeeperInit = func(app *escapp.App, ctx sdktypes.Context)

func InitBankKeeper(app *escapp.App, ctx sdktypes.Context) {
	app.BankKeeper.SetParams(ctx, banktypes.DefaultParams())
}
func InitAccountKeeper(app *escapp.App, ctx sdktypes.Context) {
	app.AccountKeeper.SetParams(ctx, authtypes.DefaultParams())
}

func (suite *KeeperIntegrationTestSuite) SetupTest() {
	suite.setupTest(InitAccountKeeper, InitBankKeeper)
}
func (suite *KeeperIntegrationTestSuite) setupTest(keeperInits ...KeeperInit) {
	invariantCheckEveryNBlocks := uint(0)
	var traceStore io.Writer
	var skipUpgradeHeights map[int64]bool
	loadLatest := true
	encCdc := escapp.MakeEncodingConfig()
	app := escapp.New(
		log.NewNopLogger(),
		dbm.NewMemDB(),
		traceStore,
		loadLatest,
		skipUpgradeHeights,
		escapp.DefaultNodeHome,
		invariantCheckEveryNBlocks,
		encCdc,
		simtestutil.EmptyAppOptions{},
	)
	genesisState := escapp.NewDefaultGenesisState(encCdc.Marshaler)
	encodedState, err := json.Marshal(genesisState)
	if err != nil {
		panic(err)
	}
	app.InitChain(
		abci.RequestInitChain{
			Validators:      abci.ValidatorUpdates{},
			ConsensusParams: DefaultConsensusParams, // Maybe its not working
			AppStateBytes:   encodedState,
		},
	)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{Time: time.Now()})
	for _, keeperInit := range keeperInits {
		keeperInit(app, ctx)
	}

	queryHelper := baseapp.NewQueryServerTestHelper(ctx, app.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, app.ResourcesyncKeeper)

	suite.app = app
	suite.ctx = ctx
	suite.msgServer = keeper.NewMsgServerImpl(app.ResourcesyncKeeper)
	suite.queryClient = types.NewQueryClient(queryHelper)
}
