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
	"context"
	"crypto/rand"
	"crypto/tls"
	"encoding/hex"
	"errors"
	"math/big"
	"strconv"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/catenax/esc-backbone/app"
	"github.com/cometbft/cometbft/proto/tendermint/p2p"
	tmtypes "github.com/cometbft/cometbft/proto/tendermint/types"
	grpcsvc "github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	clienttx "github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	xauthsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	xstaketypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// Get the node info.
func ApiGetNodeInfo(grpcHost string) (*p2p.DefaultNodeInfo, error) {

	var nodeInfo *p2p.DefaultNodeInfo

	grpcConn, err := CreateGrpcConn(grpcHost)
	if err != nil {
		return nodeInfo, err
	}
	defer grpcConn.Close()

	serviceClient := grpcsvc.NewServiceClient(grpcConn)
	nodeInfoRes, err := serviceClient.GetNodeInfo(
		context.Background(),
		&grpcsvc.GetNodeInfoRequest{},
	)
	if err != nil {
		return nodeInfo, err
	}

	nodeInfo = nodeInfoRes.GetDefaultNodeInfo()

	return nodeInfo, nil
}

// Get the transaction by hash.
func ApiGetTxnByHash(grpcHost, txHash string) (*sdktypes.TxResponse, error) {

	var txResponse *sdktypes.TxResponse

	grpcConn, err := CreateGrpcConn(grpcHost)
	if err != nil {
		return txResponse, err
	}
	defer grpcConn.Close()

	serviceClient := txtypes.NewServiceClient(grpcConn)
	txRes, err := serviceClient.GetTx(
		context.Background(),
		&txtypes.GetTxRequest{Hash: txHash},
	)
	if err != nil {
		return txResponse, err
	}

	txResponse = txRes.GetTxResponse()

	return txResponse, nil
}

// Get the latest block.
func ApiGetLatestBlock(grpcHost string) (*grpcsvc.GetLatestBlockResponse, error) {

	var latestBlockRes *grpcsvc.GetLatestBlockResponse

	grpcConn, err := CreateGrpcConn(grpcHost)
	if err != nil {
		return latestBlockRes, err
	}
	defer grpcConn.Close()

	serviceClient := grpcsvc.NewServiceClient(grpcConn)
	latestBlockRes, err = serviceClient.GetLatestBlock(
		context.Background(),
		&grpcsvc.GetLatestBlockRequest{},
	)
	if err != nil {
		return latestBlockRes, err
	}

	return latestBlockRes, nil
}

// Get the transaction block by height.
func ApiGetBlockByHeight(grpcHost string, height int64) (*grpcsvc.GetBlockByHeightResponse, error) {

	var blockRes *grpcsvc.GetBlockByHeightResponse

	grpcConn, err := CreateGrpcConn(grpcHost)
	if err != nil {
		return blockRes, err
	}
	defer grpcConn.Close()

	serviceClient := grpcsvc.NewServiceClient(grpcConn)
	blockRes, err = serviceClient.GetBlockByHeight(
		context.Background(),
		&grpcsvc.GetBlockByHeightRequest{Height: height},
	)
	if err != nil {
		return blockRes, err
	}

	return blockRes, nil
}

// Get the account information.
func ApiGetAccount(grpcHost, acctAddr string) (authtypes.BaseAccount, error) {

	var baseAccount authtypes.BaseAccount

	addr, err := sdktypes.AccAddressFromBech32(acctAddr)
	if err != nil {
		return baseAccount, err
	}

	grpcConn, err := CreateGrpcConn(grpcHost)
	if err != nil {
		return baseAccount, err
	}
	defer grpcConn.Close()

	accountClient := authtypes.NewQueryClient(grpcConn)
	accountRes, err := accountClient.Account(
		context.Background(),
		&authtypes.QueryAccountRequest{Address: addr.String()},
	)
	if err != nil {
		return baseAccount, err
	}

	accountData := accountRes.GetAccount().Value
	if err := baseAccount.XXX_Unmarshal(accountData); err != nil {
		return baseAccount, err
	}

	return baseAccount, nil
}

// Get the balances of a given account address.
func ApiGetBalances(grpcHost, acctAddr, denom string) (*sdktypes.Coin, error) {

	var balance *sdktypes.Coin

	addr, err := sdktypes.AccAddressFromBech32(acctAddr)
	if err != nil {
		return balance, err
	}

	grpcConn, err := CreateGrpcConn(grpcHost)
	if err != nil {
		return balance, err
	}
	defer grpcConn.Close()

	bankClient := banktypes.NewQueryClient(grpcConn)
	bankRes, err := bankClient.Balance(
		context.Background(),
		&banktypes.QueryBalanceRequest{Address: addr.String(), Denom: denom},
	)
	if err != nil {
		return balance, err
	}

	balance = bankRes.GetBalance()

	return balance, nil
}

// Get the list of validators in the validator set.
func ApiGetValidatorSet(grpcHost string) ([]xstaketypes.Validator, error) {

	var validatorList []xstaketypes.Validator

	grpcConn, err := CreateGrpcConn(grpcHost)
	if err != nil {
		return validatorList, err
	}
	defer grpcConn.Close()

	queryClient := xstaketypes.NewQueryClient(grpcConn)

	validatorsRes, err := queryClient.Validators(
		context.Background(),
		&xstaketypes.QueryValidatorsRequest{},
	)
	if err != nil {
		return validatorList, err
	}

	validatorList = validatorsRes.GetValidators()

	return validatorList, nil
}

// Broadcast the signed transations.
func ApiBroadcastSignedTxn(testHost string, txBytes []byte) (string, int64, error) {

	grpcConn, err := CreateGrpcConn(testHost)
	if err != nil {
		return "", int64(0), err
	}
	defer grpcConn.Close()

	serviceClient := txtypes.NewServiceClient(grpcConn)

	broadcastTxRes, err := serviceClient.BroadcastTx(
		context.Background(),
		&txtypes.BroadcastTxRequest{
			Mode:    txtypes.BroadcastMode_BROADCAST_MODE_BLOCK,
			TxBytes: txBytes,
		},
	)
	if err != nil {
		return "", int64(0), err
	}

	txResponse := broadcastTxRes.GetTxResponse()

	if txResponse.Code != 0 {
		return "", int64(0), errors.New("TxResponse.Code: " + strconv.FormatUint(uint64(txResponse.Code), 10))
	}

	return txResponse.TxHash, txResponse.Height, nil
}

// Create a GRPC client connnection.
func CreateGrpcConn(grpcHost string) (*grpc.ClientConn, error) {

	return CreateGrpcConnTls(grpcHost)
}

// Create a GRPC client connnection with TLS.
func CreateGrpcConnTls(grpcHost string) (*grpc.ClientConn, error) {

	var grpcConn *grpc.ClientConn
	var err error

	tlsCfg := &tls.Config{
		InsecureSkipVerify: true,
	}
	grpcConn, err = grpc.Dial(
		grpcHost,
		grpc.WithTransportCredentials(credentials.NewTLS(tlsCfg)),
	)

	return grpcConn, err
}

// Create a GRPC client connnection without TLS.
func CreateGrpcConnInsecure(grpcHost string) (*grpc.ClientConn, error) {

	var grpcConn *grpc.ClientConn
	var err error

	grpcConn, err = grpc.Dial(
		grpcHost,
		grpc.WithInsecure(),
	)

	return grpcConn, err
}

// Get the ID of the Test node.
func CheckTestNode(grpcHost string) (string, error) {

	nodeInfo, err := ApiGetNodeInfo(grpcHost)
	if err != nil {
		return "", err
	}

	nodeId := nodeInfo.GetDefaultNodeID()

	return nodeId, nil
}

// Get the 1st available Reference node from a list of Reference nodes host name.
func CheckRefNode(grpcHostsList []string) (string, string, error) {

	for idx := 0; idx < len(grpcHostsList); idx++ {
		grpcHost := grpcHostsList[idx]
		nodeInfo, err := ApiGetNodeInfo(grpcHost)
		if err == nil {
			nodeId := nodeInfo.GetDefaultNodeID()
			return nodeId, grpcHost, nil
		}
	}

	return "", "", errors.New("No Reference node is available.")
}

// Check if the specified host is a validator exists in the validator set.
func ExistInValidatorSet(grpcHost, accountAddr string) (xstaketypes.Validator, string, int, error) {

	var validator xstaketypes.Validator

	validatorList, err := ApiGetValidatorSet(grpcHost)
	if err != nil {
		return validator, "", 0, err
	}

	nodeInfo, err := ApiGetNodeInfo(grpcHost)
	if err != nil {
		return validator, "", 0, err
	}
	moniker := nodeInfo.Moniker

	numofValidators := len(validatorList)
	for _, validator = range validatorList {
		accAddr, err := GetAccountAddress(validator)
		if err != nil {
			return validator, "", 0, err
		}
		if accountAddr == "" && validator.Description.Moniker == moniker {
			return validator, accAddr.String(), numofValidators, nil
		}
		if accAddr.String() == accountAddr {
			return validator, accAddr.String(), numofValidators, nil
		}
	}

	return validator, "", 0, errors.New("Validator not found in the validator set.")
}

// Get DataHash of the block.
func GetBlockDataHash(block *tmtypes.Block) []byte {

	header := block.GetHeader()
	dataHash := header.GetDataHash()

	return dataHash
}

// Get the account address of the validator.
func GetAccountAddress(validator xstaketypes.Validator) (sdktypes.AccAddress, error) {

	var accAddr sdktypes.AccAddress

	valAddr, err := GetValidatorAddress(validator)
	if err != nil {
		return accAddr, err
	}

	accAddr, err = sdktypes.AccAddressFromHexUnsafe(hex.EncodeToString(valAddr.Bytes()))
	if err != nil {
		return accAddr, err
	}

	return accAddr, nil
}

// Get the validator address of the validator.
func GetValidatorAddress(validator xstaketypes.Validator) (sdktypes.ValAddress, error) {

	var valAddr sdktypes.ValAddress

	valAddr, err := sdktypes.ValAddressFromBech32(validator.OperatorAddress)
	if err != nil {
		return valAddr, err
	}

	return valAddr, nil
}

// Get the ed25519 public key from the validator.
func GetConsensusPublicKey(validator xstaketypes.Validator) (ed25519.PubKey, error) {

	var pubKey ed25519.PubKey

	cssPubKey := validator.ConsensusPubkey
	if err := pubKey.XXX_Unmarshal(cssPubKey.Value); err != nil {
		return pubKey, err
	}

	return pubKey, nil
}

// Instantiate new keyring.
func NewKeyring(cfg map[string]string) (keyring.Keyring, error) {

	var kr keyring.Keyring

	reader := strings.NewReader("")
	reader.Reset(cfg["PassPhrase"] + "\n")
	kr, err := keyring.New(
		sdktypes.KeyringServiceName(),
		cfg["KeyringBackend"],
		cfg["HomeDir"],
		reader,
		app.MakeEncodingConfig().Marshaler,
	)

	if err != nil {
		return kr, err
	}

	return kr, nil
}

// Create a transaction sending tokens from account A to account B.
func CreateSignedTxn(testHost string,
	cfg map[string]string) ([]byte, error) {

	encodingConfig := app.MakeEncodingConfig()
	txConfig := encodingConfig.TxConfig
	txBuilder := txConfig.NewTxBuilder()

	fromAddr, err := sdktypes.AccAddressFromBech32(cfg["FromAccount"])
	if err != nil {
		return nil, err
	}

	toAddr, err := sdktypes.AccAddressFromBech32(cfg["ToAccount"])
	if err != nil {
		return nil, err
	}

	amount, err := strconv.ParseInt(cfg["TxfAmount"], 10, 64)
	if err != nil {
		return nil, err
	}
	coin := sdktypes.NewInt64Coin(cfg["TxfDenom"], amount)
	coins := sdktypes.NewCoins(coin)

	msg := banktypes.NewMsgSend(fromAddr, toAddr, coins)

	if err := txBuilder.SetMsgs(msg); err != nil {
		return nil, err
	}

	fee, err := strconv.ParseInt(cfg["Fee"], 10, 64)
	if err != nil {
		return nil, err
	}
	feeAmt := sdktypes.Coin{Denom: cfg["TxfDenom"], Amount: sdktypes.NewInt(fee)}
	feeAmount := sdktypes.NewCoins(feeAmt)
	txBuilder.SetFeeAmount(feeAmount)

	gasLimit, err := strconv.ParseUint(cfg["GasLimit"], 10, 64)
	txBuilder.SetGasLimit(gasLimit)

	keyInfo, err := escKeyring.KeyByAddress(fromAddr)
	if err != nil {
		return nil, err
	}

	exportedPrivateKey, err := escKeyring.ExportPrivKeyArmor(keyInfo.Name, cfg["PassPhrase"])
	if err != nil {
		return nil, err
	}

	privateKey, _, err := crypto.UnarmorDecryptPrivKey(exportedPrivateKey, cfg["PassPhrase"])
	if err != nil {
		return nil, err
	}

	account, err := ApiGetAccount(testHost, cfg["FromAccount"])
	if err != nil {
		return nil, err
	}
	accNumber := account.AccountNumber
	accSequence := account.Sequence

	privKeyList := []cryptotypes.PrivKey{privateKey}
	accNumberList := []uint64{accNumber}
	accSeqList := []uint64{accSequence}

	//  First round: Gather all signer infos.
	var signatureList []signing.SignatureV2
	for idx, privKey := range privKeyList {
		sig := signing.SignatureV2{
			PubKey: privKey.PubKey(),
			Data: &signing.SingleSignatureData{
				SignMode:  encodingConfig.TxConfig.SignModeHandler().DefaultMode(),
				Signature: nil,
			},
			Sequence: accSeqList[idx],
		}
		signatureList = append(signatureList, sig)
	}

	if err := txBuilder.SetSignatures(signatureList...); err != nil {
		return nil, err
	}

	//  Second round: All signer info are set, each signer can sign.
	var signingList []signing.SignatureV2
	for idx, privKey := range privKeyList {
		signerData := xauthsigning.SignerData{
			ChainID:       cfg["ChainID"],
			AccountNumber: accNumberList[idx],
			Sequence:      accSeqList[idx],
		}
		signMode := encodingConfig.TxConfig.SignModeHandler().DefaultMode()
		sig, err := clienttx.SignWithPrivKey(signMode, signerData, txBuilder, privKey,
			encodingConfig.TxConfig, accSeqList[idx])
		if err != nil {
			return nil, err
		}
		signingList = append(signingList, sig)
	}

	if err := txBuilder.SetSignatures(signingList...); err != nil {
		return nil, err
	}

	txBytes, err := encodingConfig.TxConfig.TxEncoder()(txBuilder.GetTx())
	if err != nil {
		return nil, err
	}

	return txBytes, nil
}

// Create account.
func createAccount(cfg map[string]string) (string, error) {

	keyringAlgo, _ := escKeyring.SupportedAlgorithms()
	signatureAlgo, err := keyring.NewSigningAlgoFromString("secp256k1", keyringAlgo)
	if err != nil {
		return "", err
	}
	uid := genRandomString(40)
	record, _, err := escKeyring.NewMnemonic(uid, keyring.English, "", cfg["PassPhrase"], signatureAlgo)
	if err != nil {
		return "", err
	}

	pubKey, err := record.GetPubKey()
	if err != nil {
		return "", err
	}

	acc, err := sdktypes.AccAddressFromHexUnsafe(pubKey.Address().String())
	if err != nil {
		return "", err
	}

	return acc.String(), nil
}

// Generate a random number up to the specific max.
func genRandomNumber(max int64) int64 {
	randomNum, _ := rand.Int(rand.Reader, big.NewInt(max))
	return randomNum.Int64()
}

// Generate a random string of the specific length.
func genRandomString(strLen int) string {

	const alphabets = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	max := int64(strLen)
	buf := make([]byte, strLen)
	for idx := range buf {
		randomNum, _ := rand.Int(rand.Reader, big.NewInt(max))
		buf[idx] = alphabets[randomNum.Int64()]
	}
	return string(buf)
}

// Create test accounts.
func createTestAccounts(cfg map[string]string) error {
	fromAccount, err := createAccount(cfg)
	if err != nil {
		return err
	}
	toAccount, err := createAccount(cfg)
	if err != nil {
		return err
	}
	cfg["FromAccount"] = fromAccount
	cfg["ToAccount"] = toAccount

	return nil
}

// Delete test accounts.
func deleteTestAccounts(cfg map[string]string) error {

	fromAddr, err := sdktypes.AccAddressFromBech32(cfg["FromAccount"])
	if err != nil {
		return err
	}
	toAddr, err := sdktypes.AccAddressFromBech32(cfg["ToAccount"])
	if err != nil {
		return err
	}
	err = escKeyring.DeleteByAddress(fromAddr)
	if err != nil {
		return err
	}
	err = escKeyring.DeleteByAddress(toAddr)
	if err != nil {
		return err
	}

	return nil
}
