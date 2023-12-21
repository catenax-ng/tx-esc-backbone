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
package main

import (
	"context"
	"cosmossdk.io/log"
	"fmt"
	web2wrapper "github.com/catenax/web2-wrapper/pkg"
	"os"
)

func main() {
	rootCmd := web2wrapper.NewRootCmd(func(ctx context.Context, cfg *web2wrapper.Config, logger log.Logger) {
		client, err := web2wrapper.NewChainClient(ctx, logger, cfg)
		if err != nil {
			logger.Error(fmt.Sprintf("Cannot connect to chain %v", err))
			os.Exit(1)
		}
		broker, err := web2wrapper.NewNatsBrokerFor(cfg)
		if err != nil {
			logger.Error(fmt.Sprintf("Cannot connect to nats %v", err))
			os.Exit(1)
		}
		web2wrapper.NewDucttape(broker, client, logger).Forward(ctx, cfg.StartBlock)
	})
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
