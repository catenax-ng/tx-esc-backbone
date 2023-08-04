package main

import (
	"context"
	"cosmossdk.io/log"
	"fmt"
	resourcesyncmoduletypes "github.com/catenax/esc-backbone/x/resourcesync/types"
	web2wrapper "github.com/catenax/web2-wrapper/pkg"
	"os"
)

func main() {
	rootCmd := web2wrapper.NewRootCmd(func(ctx context.Context, cfg *web2wrapper.Config, logger log.Logger) {
		broker, err := web2wrapper.NewNatsBrokerFor(cfg)
		defer broker.Close()
		if err != nil {
			panic(err)
		}
		broker.Submit(
			&web2wrapper.Msg{
				Res: resourcesyncmoduletypes.Resource{
					Originator:   "orig",
					OrigResId:    "oriresid6",
					TargetSystem: "testsys",
					ResourceKey:  "abcde",
					DataHash:     nil,
				},
				Mod: web2wrapper.CREATE,
				Src: web2wrapper.CLIENT,
			},
		)
		select {
		case <-ctx.Done():
		}
	})
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
