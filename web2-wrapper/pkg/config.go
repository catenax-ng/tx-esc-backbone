// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package web2wrapper

import (
	"encoding/json"
	"github.com/ignite/cli/ignite/pkg/cosmosaccount"
	"io"
	"os"
)

type Config struct {
	AddressPrefix  string                       `json:"address_prefix"`
	ChainId        string                       `json:"chain_id"`
	From           string                       `json:"from"`
	HostAddress    string                       `json:"host_address"`
	NodeAddress    string                       `json:"node_address"`
	Fees           string                       `json:"fees"`
	Gas            string                       `json:"gas"`
	Home           string                       `json:"home"`
	KeyRingBackend cosmosaccount.KeyringBackend `json:"key_ring_backend"`
	StartBlock     int64                        `json:"start_block"`
	filePath       string
}

func ReadConfig(path string) (*Config, error) {
	config := &Config{
		HostAddress: ":8080",
	}
	configFile, err := os.Open(path)

	if err != nil {
		return nil, err
	}
	defer configFile.Close()
	bz, err := io.ReadAll(configFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bz, config)
	if err != nil {
		return nil, err
	}
	config.filePath = path
	return config, nil
}
func (c *chainClient) SafeConfig() error {
	return nil
}
