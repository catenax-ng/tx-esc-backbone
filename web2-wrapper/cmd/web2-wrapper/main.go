// Copyright (c) 2022-2023 - for information on the respective copyright owner
// see the NOTICE file and/or the repository at
// https://github.com/catenax-ng/product-esc-backbone-code
//
// SPDX-License-Identifier: Apache-2.0
package main

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func main() {
	var configFile string
	rootCmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			ctx := context.Background()
			config, err := ReadConfig(configFile)
			if err != nil {
				log.Fatal(err)
			}
			client, err := NewChainClient(ctx, *config)
			if err != nil {
				log.Fatal(err)
			}
			client.Poll(ctx)
			newRouter(config, client).handleRequests()
		},
	}
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "wrapper-config.json", "File to read configuration from")
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
