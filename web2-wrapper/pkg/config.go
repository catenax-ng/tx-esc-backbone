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
package web2wrapper

import (
	"encoding/json"
	"github.com/ignite/cli/ignite/pkg/cosmosaccount"
	"io"
	"os"
)

type BrokerCfg struct {
	Url        string `json:"url"`
	Clientname string `json:"clientname"`
	Topic      string `json:"topic"`
	Queue      string `json:"queue"`
}

type Config struct {
	AddressPrefix  string                       `json:"address_prefix"`
	ChainId        string                       `json:"chain_id"`
	From           string                       `json:"from"`
	HostAddress    string                       `json:"host_address"`
	NodeAddress    string                       `json:"node_address"`
	Fees           string                       `json:"fees"`
	Gas            string                       `json:"gas"`
	Home           string                       `json:"home"`
	LogLevel       string                       `json:"log_level"`
	KeyRingBackend cosmosaccount.KeyringBackend `json:"key_ring_backend"`
	StartBlock     int64                        `json:"start_block"`
	Broker         BrokerCfg                    `json:"broker"`
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
