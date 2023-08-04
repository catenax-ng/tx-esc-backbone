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
		broker, err := web2wrapper.NewNatsBrokerFor(cfg)
		if err != nil {
			panic(err)
		}
		for msg := range broker.Receive(ctx) {
			logger.Info(fmt.Sprintf("Wrapper subscribe: %s", msg))
		}
	})
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
