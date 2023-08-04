// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
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
			logger.Error("Cannot connect to chain %v", err)
			os.Exit(1)
		}
		broker, err := web2wrapper.NewNatsBrokerFor(cfg)
		if err != nil {
			logger.Error("Cannot connect to nats %v", err)
			os.Exit(1)
		}
		web2wrapper.NewDucttape(broker, client, logger).Forward(ctx, cfg.StartBlock)
	})
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
