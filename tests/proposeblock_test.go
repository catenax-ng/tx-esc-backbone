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
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	sdktypes "github.com/cosmos/cosmos-sdk/types"
	xstaketypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// Propose block test.
// Check if TEST node is in the validator set.
// Get the Proposer Address from the block obtained by height.
// Check if the Proposer Address matches the address of the TEST node.
// May need to repeat the steps by creating another transaction depending on the number
// of validators in the validator set as the proposer is elected in round-robin.
func TestProposeBlock(t *testing.T) {

	var err error
	var validator xstaketypes.Validator
	var proposerAddr string
	var accountAddr string
	var isMatching bool
	var numofValidators int
	var txBytes []byte
	var txHash string
	var txHeight int64

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

	//  Check Test node is a validator
	t.Run("check_test_node_validator", func(t *testing.T) {
		validator, accountAddr, numofValidators, err = ExistInValidatorSet(testHost, cfg["ValidatorAccount"])
		require.NoError(t, err)
		t.Logf("Operator address: %s", validator.OperatorAddress)
		t.Logf("Account address: %s", accountAddr)
		t.Logf("Moniker: %s", validator.Description.Moniker)
		t.Logf("Number of validators: %d", numofValidators)
	})

	//  Get fund from faucet
	t.Run("get_fund_from_faucet", func(t *testing.T) {
		err = getFundFromFaucet(t, httpProtocol, testHost, faucetHost, cfg)
		assert.NoError(t, err)
	})

	//  Check proposer (Test node)
	t.Run("check_proposer", func(t *testing.T) {
		for counter := 0; counter < numofValidators; counter++ {
			t.Logf("Round: %d", counter+1)
			//  Create transaction (Test node)
			txBytes, err = CreateSignedTxn(testHost, cfg)
			require.NoError(t, err)
			t.Log("  Signed transaction created")

			//  Broadcast transaction (Test node)
			txHash, txHeight, err = ApiBroadcastSignedTxn(testHost, txBytes)
			require.NoError(t, err)
			t.Logf("  Transaction hash: %s", txHash)
			t.Logf("  Transaction height: %d", txHeight)

			//  Check proposer (Test node)
			proposerAddr, accountAddr, isMatching, err = checkProposer(testHost, txHeight, validator)
			require.NoError(t, err)
			t.Logf("  Proposer address: %s", proposerAddr)
			t.Logf("  Account address: %s", accountAddr)

			if isMatching {
				break
			}
		}

		if assert.True(t, isMatching, "Test node is not a proposer") {
			t.Log("Test node is a proposer")
		}
	})
}

// Check if the validator is the proposer.
func checkProposer(testHost string,
	blockHeight int64,
	validator xstaketypes.Validator) (string, string, bool, error) {

	blockRes, err := ApiGetBlockByHeight(testHost, blockHeight)
	if err != nil {
		return "", "", false, err
	}

	block := blockRes.GetBlock()
	header := block.GetHeader()
	proposerAddress := header.GetProposerAddress()
	proposerAddr := base64.StdEncoding.EncodeToString(proposerAddress)

	pubKey, err := GetConsensusPublicKey(validator)
	if err != nil {
		return "", "", false, err
	}

	accountAddress := sdktypes.AccAddress(pubKey.Address().Bytes())
	accountAddr := base64.StdEncoding.EncodeToString(accountAddress)

	isMatching := proposerAddr == accountAddr

	return proposerAddr, accountAddr, isMatching, nil
}
