// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
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
