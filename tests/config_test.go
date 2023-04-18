// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
//
//
// Follow the instructions from the documentation
// "Validator Setup Instructions (Testnet)".
// Setup a new validator node and join the Testnet.
// https://confluence.catena-x.net/display/CORE/Validator+Setup+Instructions+(Testnet)
//
// Enter the specific host names and information in this config.
// For CI Pipeline automated test, set cfg["ValidatorAccount"] = "".
// To check if a particular validator is one of the block-proposers,
// set cfg["ValidatorAccount"] to the specific account.

package txn_test

import (
	"os"
	"path"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdktypes "github.com/cosmos/cosmos-sdk/types"

	catxapp "github.com/catenax/esc-backbone/app"
)

var (
	refHostsList []string
	cfg          map[string]string
	escKeyring   keyring.Keyring
)

const (
	Bech32AccountPrefix       = catxapp.AccountAddressPrefix
	Bech32ValidatorAddrPrefix = Bech32AccountPrefix +
		sdktypes.PrefixValidator +
		sdktypes.PrefixOperator
	Bech32ConsensusAddrPrefix = Bech32AccountPrefix +
		sdktypes.PrefixValidator +
		sdktypes.PrefixConsensus

	testHost     = "validator1-csms-grpc.dev.demo.catena-x.net:443"
	faucetHost   = "faucet-faucet.dev.demo.catena-x.net/"
	httpProtocol = "https://"
)

func init() {

	refHostsList = []string{
		"validator2-csms-grpc.dev.demo.catena-x.net:443",
		"validator3-csms-grpc.dev.demo.catena-x.net:443",
		"validator4-csms-grpc.dev.demo.catena-x.net:443",
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		os.Exit(-1)
	}

	cfg = make(map[string]string)
	cfg["App"] = path.Join(homeDir, "go", "bin", "esc-backboned")
	cfg["ValidatorAccount"] = ""
	cfg["TxfAmount"] = "5"
	cfg["TxfDenom"] = "ncaxdemo"
	cfg["ChainID"] = "catenax-testnet-1"
	cfg["HomeDir"] = path.Join(homeDir, ".esc-backbone")
	cfg["PassPhrase"] = "password"
	cfg["Fee"] = "2000000"
	cfg["GasLimit"] = "2000000"
	cfg["KeyringBackend"] = keyring.BackendTest

	config := sdktypes.GetConfig()
	config.SetBech32PrefixForAccount(Bech32AccountPrefix, sdktypes.PrefixPublic)
	config.SetBech32PrefixForValidator(Bech32ValidatorAddrPrefix, sdktypes.PrefixPublic)
	config.SetBech32PrefixForConsensusNode(Bech32ConsensusAddrPrefix, sdktypes.PrefixPublic)

	escKeyring, err = NewKeyring(cfg)
	if err != nil {
		os.Exit(-1)
	}
}
