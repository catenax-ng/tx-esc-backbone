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
	"encoding/json"
	"errors"
	"math/big"
	"net/http"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Write test.
// Check availability of REF and TEST nodes.
// Check if from & to account exist.
// Check if from account has sufficient balance for a token transfer.
// From TEST node, submit a transaction (token transfer). Record transaction hash (txhash) returned.
// Get transaction by hash from TEST node. Extract the block height in the response.
// Get the block from REF node (or nodes) of the same block height.
// Check if the transaction is included in the block.
func TestWrite(t *testing.T) {

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

	//  Check Test node
	t.Run("check_test_node", func(t *testing.T) {
		nodeID, err = CheckTestNode(testHost)
		require.NoError(t, err)
		t.Logf("Test node ID: %s", nodeID)
		t.Logf("Test node host name: %s", testHost)
	})

	//  Check Reference node
	t.Run("check_reference_node", func(t *testing.T) {
		nodeID, refHost, err = CheckRefNode(refHostsList)
		require.NoError(t, err)
		t.Logf("Reference node ID: %s", nodeID)
		t.Logf("Reference node host name: %s", refHost)
	})

	//  Get fund from faucet
	t.Run("get_fund_from_faucet", func(t *testing.T) {
		err = getFundFromFaucet(t, httpProtocol, testHost, faucetHost, cfg)
		assert.NoError(t, err)
	})

	//  Create transaction (Test node)
	t.Run("create_transaction", func(t *testing.T) {
		txHash, err := createTxn(testHost, cfg)
		require.NoError(t, err)
		t.Logf("Transaction hash: %s", txHash)
		time.Sleep(10 * time.Second)

		//  Get transaction by hash (Test node)
		txResponse, err := ApiGetTxnByHash(testHost, txHash)
		require.NoError(t, err)
		txnBlockHeight := txResponse.Height
		t.Logf("Transaction block height: %d", txnBlockHeight)

		//  Get transaction block by height (Test node)
		testBlockRes, err := ApiGetBlockByHeight(testHost, txnBlockHeight)
		require.NoError(t, err)

		//  Get transaction block by height (Reference node)
		refBlockRes, err := ApiGetBlockByHeight(refHost, txnBlockHeight)
		require.NoError(t, err)

		//  Compare transactions data hash
		testDataHash := GetBlockDataHash(testBlockRes.GetBlock())
		t.Logf("Test node transaction data hash: %s",
			base64.StdEncoding.EncodeToString(testDataHash))

		refDataHash := GetBlockDataHash(refBlockRes.GetBlock())
		t.Logf("Reference node transaction data hash: %s",
			base64.StdEncoding.EncodeToString(refDataHash))

		isMatching := bytes.Equal(testDataHash, refDataHash)
		if assert.True(t, isMatching, "Transactions are not matching") {
			t.Log("Transactions are matching")
		}
	})
}

// Create a transaction sending tokens from account A to account B.
func createTxn(testHost string,
	cfg map[string]string) (string, error) {

	txBytes, err := CreateSignedTxn(testHost, cfg)
	if err != nil {
		return "", err
	}

	txHash, _, err := ApiBroadcastSignedTxn(testHost, txBytes)
	if err != nil {
		return "", err
	}

	return txHash, nil
}

// Create a transaction sending token from account A to account B.
func createTxnCmd(cfg map[string]string) (string, error) {

	cmdLine := "printf" + " " + "'" + cfg["PassPhrase"] + "\n" + "'" + " | " +
		cfg["App"] + " " +
		"tx" + " " + "bank" + " " + "send" + " " +
		cfg["FromAccount"] + " " + cfg["ToAccount"] + " " + cfg["TxfAmount"] + cfg["TxfDenom"] + " " +
		"--yes" + " " +
		"--chain-id" + " " + cfg["ChainID"] + " " +
		"--home" + " " + cfg["HomeDir"]
	cmd := exec.Command("bash", "-c", cmdLine)

	stdout, err := cmd.Output()
	if err != nil {
		return "", err
	}

	res := string(stdout)

	idxTxHash := strings.Index(res, "txhash:")
	txHash := res[idxTxHash+8:]
	//  Remove \n character at the end.
	txHash = txHash[:len(txHash)-1]
	return txHash, nil
}

// Check if the account balance is sufficient for a given amount.
func checkSufficientBalance(balanceAmount *big.Int, queryAmount string) error {

	queryAmt, _ := new(big.Int).SetString(queryAmount, 10)

	//  -1: if balanceAmount < queryAmt
	//   0: if balanceAmount = queryAmt
	//  +1: if balanceAmount > queryAmt
	comparison := balanceAmount.Cmp(queryAmt)
	if comparison == -1 {
		err := errors.New("From account does not have sufficient balance.")
		return err
	}

	return nil
}

// Get fund from faucet.
func getFundFromFaucet(t *testing.T,
	httpProtocol string,
	testHost string,
	faucetHost string,
	cfg map[string]string) error {

	payloadData := map[string]string{
		"address": cfg["FromAccount"],
		"denom":   cfg["TxfDenom"],
	}

	payloadByte, err := json.Marshal(payloadData)
	require.NoError(t, err)

	url := httpProtocol + faucetHost
	httpClient := &http.Client{}
	res, err := httpClient.Post(url, "application/json", bytes.NewBuffer(payloadByte))
	require.NoError(t, err)
	res.Body.Close()

	//  Check from account has sufficient balances.
	checkBalanceFunc := func() bool {
		balance, err := ApiGetBalances(testHost, cfg["FromAccount"], cfg["TxfDenom"])
		require.NoError(t, err)
		balAmount := balance.Amount
		err = checkSufficientBalance(balAmount.BigInt(), cfg["TxfAmount"])
		return err == nil
	}
	checkBalanceMsg := "Checking for sufficient balance"

	switch res.StatusCode {
	case http.StatusMethodNotAllowed:
		t.Log("Request to faucet declined due to multiple retries in single time window")
		require.True(t, checkBalanceFunc(), checkBalanceMsg)
	case http.StatusOK:
		t.Log("Request to faucet succeeded")
		require.Eventually(t,
			checkBalanceFunc,
			10*time.Second,
			500*time.Millisecond,
			checkBalanceMsg)
	default:
		t.Fatal("Request to faucet failed", res)
	}

	return nil
}
