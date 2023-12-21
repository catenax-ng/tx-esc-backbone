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

package txn_test

import (
	"bytes"
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	tmtypes "github.com/cometbft/cometbft/proto/tendermint/types"
	grpcsvc "github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
)

// Number of latest blocks to get for comparison.
const NumOfLatestBlocks int64 = 5

// Read test.
// Check availability of REF and TEST nodes.
// Get latest block from TEST node.
// Get latest block height from REF node.
// Get N number of latest blocks from REF node.
// Check if the TEST block exists in the N number of latest blocks from REF node.
// Check if the contents of these blocks (of same height) are identical.
func TestRead(t *testing.T) {

	var err error
	var nodeID string
	var refHost string

	//  Setup
	err = createTestAccounts(cfg)
	require.NoError(t, err)
	t.Logf("From account: %s", cfg["FromAccount"])
	t.Logf("To account: %s", cfg["ToAccount"])

	//  Tear down
	t.Cleanup(func() {
		err = deleteTestAccounts(cfg)
		require.NoError(t, err)
	})

	//  Check Test Node
	t.Run("check_test_node", func(t *testing.T) {
		nodeID, err = CheckTestNode(testHost)
		require.NoError(t, err)
		t.Logf("Test node ID: %s", nodeID)
		t.Logf("Test node host name: %s", testHost)
	})

	//  Check Reference Node
	t.Run("check_reference_node", func(t *testing.T) {
		nodeID, refHost, err = CheckRefNode(refHostsList)
		require.NoError(t, err)
		t.Logf("Reference node ID: %s", nodeID)
		t.Logf("Reference node host name: %s", refHost)
	})

	//  Check Latest Block Exists
	t.Run("check_latest_block_exists", func(t *testing.T) {

		//  Test Node
		testLatestBlockRes, err := ApiGetLatestBlock(testHost)
		require.NoError(t, err)
		testBlock := testLatestBlockRes.GetBlock()
		testHeader := testBlock.GetHeader()
		testLatestBlockHeight := testHeader.GetHeight()
		t.Logf("Test node latest block height: %d", testLatestBlockHeight)

		//  Reference Node
		refLatestBlockRes, err := ApiGetLatestBlock(refHost)
		require.NoError(t, err)
		refBlock := refLatestBlockRes.GetBlock()
		refHeader := refBlock.GetHeader()
		refLatestBlockHeight := refHeader.GetHeight()
		t.Logf("Reference node latest block height: %d", refLatestBlockHeight)

		//  Reference Node (Latest N Blocks)
		refLatestBlocksList, err := getLatestBlocksList(refHost, refLatestBlockHeight)
		require.NoError(t, err)

		//  Compare transactions
		refBlockRes, found := refBlockExists(testLatestBlockRes, refLatestBlocksList)
		if assert.True(t, found, "Test block does not exist in the Reference node.") {
			t.Log("Test block is found in the Reference node")
		}

		isMatching := compareTxs(testLatestBlockRes.GetBlock(), refBlockRes.GetBlock())
		if assert.True(t, isMatching, "Blocks are not matching") {
			t.Log("Blocks are matching")
		}
	})
}

// Get N number of latest blocks.
func getLatestBlocksList(host string, height int64) ([]*grpcsvc.GetBlockByHeightResponse, error) {

	blocksList := []*grpcsvc.GetBlockByHeightResponse{}
	blocksToGet := NumOfLatestBlocks

	if height <= blocksToGet {
		blocksToGet = height
	}

	for h := height; h > (height - blocksToGet); h-- {
		blockRes, err := ApiGetBlockByHeight(host, h)
		if err != nil {
			return blocksList, err
		}
		blocksList = append(blocksList, blockRes)
	}

	return blocksList, nil
}

// Check if the block from Test node exists in the list
// of latest N blocks from Reference node.
func refBlockExists(testLatestBlockRes *grpcsvc.GetLatestBlockResponse,
	refLatestBlocksList []*grpcsvc.GetBlockByHeightResponse) (*grpcsvc.GetBlockByHeightResponse, bool) {

	var refLastestBlock *grpcsvc.GetBlockByHeightResponse

	testBlockId := testLatestBlockRes.GetBlockId()
	testBlockIdHash := base64.StdEncoding.EncodeToString(testBlockId.GetHash())
	testBlock := testLatestBlockRes.GetBlock()
	testHeader := testBlock.GetHeader()
	testBlockHeight := testHeader.GetHeight()

	for _, refLastestBlock := range refLatestBlocksList {
		refBlockId := refLastestBlock.GetBlockId()
		refBlockIdHash := base64.StdEncoding.EncodeToString(refBlockId.GetHash())
		refBlock := refLastestBlock.GetBlock()
		refHeader := refBlock.GetHeader()
		refBlockHeight := refHeader.GetHeight()

		if (testBlockIdHash == refBlockIdHash) && (testBlockHeight == refBlockHeight) {
			return refLastestBlock, true
		}
	}

	return refLastestBlock, false
}

// Compare the transactions using the block data hashes.
func compareTxs(testBlock, refBlock *tmtypes.Block) bool {

	testDataHash := GetBlockDataHash(testBlock)
	refDataHash := GetBlockDataHash(refBlock)
	isMatching := bytes.Equal(testDataHash, refDataHash)
	return isMatching
}
