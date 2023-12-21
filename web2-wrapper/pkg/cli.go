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
	"context"
	"cosmossdk.io/log"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
)

type ConfigConsumer func(context.Context, *Config, log.Logger)

func NewRootCmd(entryFn ConfigConsumer) *cobra.Command {
	var configFile string
	rootCmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			cfg, err := ReadConfig(configFile)
			if err != nil {
				fmt.Printf("Cannot read config %v", err)
				os.Exit(1)
			}
			logLevel, err := zerolog.ParseLevel(cfg.LogLevel)
			if err != nil {
				fmt.Printf("Cannot read config %v", err)
				os.Exit(1)
			}
			logger := log.NewLogger(os.Stderr, log.LevelOption(logLevel))
			ctx, cancel := context.WithCancel(context.Background())
			signals := make(chan os.Signal, 1)
			defer func() {
				signal.Stop(signals)
				cancel()
			}()
			go func() {
				select {
				case <-signals:
					cancel()
				case <-ctx.Done():

				}
			}()
			entryFn(ctx, cfg, logger)
		},
	}
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "wrapper-config.json", "File to read configuration from")
	return rootCmd
}
