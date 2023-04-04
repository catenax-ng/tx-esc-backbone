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

package txn_test

import (
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
)

// Get list of Reference node (GRPC server) host names.
func getRefHostsList() []string {

	return []string{
		"validator2-csms-grpc.dev.demo.catena-x.net:443",
		"validator3-csms-grpc.dev.demo.catena-x.net:443",
		"validator4-csms-grpc.dev.demo.catena-x.net:443",
	}
}

// Get the Test node (GRPC server) host name.
func getTestHost() string {

	return "validator1-csms-grpc.dev.demo.catena-x.net:443"
}

// Get the faucet host name.
func getFaucetHost() string {

	return "faucet-faucet.dev.demo.catena-x.net/"
}

// Get the http protocol.
func httpProtocol() string {

	return "https://"
}

// Get the Test node configurations.
func getTestNodeConfig() map[string]string {

	cfg := make(map[string]string)

	cfg["App"] = "/home/<user>/<golib>/bin/esc-backboned"
	cfg["ValidatorAccount"] = "catenax105gtxtvscdywtzwcn46n60sfmkqwjy53078vum"
	cfg["FromAccount"] = "catenax14r7fw8vl6tk9gf6a4km9ef9j5xycu6mzg4n0av"
	cfg["ToAccount"] = "catenax192s9m0tjua7f9enwlklgwk5zu2t956zn89cvqv"
	cfg["TxfAmount"] = "5"
	cfg["TxfDenom"] = "ncaxdemo"
	cfg["ChainID"] = "catenax-testnet-1"
	cfg["HomeDir"] = "/home/<user>/.esc-backbone"
	cfg["PassPhrase"] = "password"
	cfg["Fee"] = "2000000"
	cfg["GasLimit"] = "2000000"
	cfg["KeyringBackend"] = keyring.BackendTest

	return cfg
}
