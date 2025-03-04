package main

import (
	"context"
	"fmt"
	"github.com/the-web3/market-services/services"

	"github.com/urfave/cli/v2"

	"github.com/the-web3/market-services/common/cliapp"
	"github.com/the-web3/market-services/config"
	flags2 "github.com/the-web3/market-services/flags"
)

func runRpc(ctx *cli.Context, shutdown context.CancelCauseFunc) (cliapp.Lifecycle, error) {
	fmt.Println("running grpc services...")
	cfg := config.NewConfig(ctx)

	grpcServerCfg := &services.MarketRpcConfig{
		Host: cfg.RpcServer.Host,
		Port: cfg.RpcServer.Port,
	}

	return services.NewMarketRpcService(grpcServerCfg)
}

func NewCli(GitCommit string, GitData string) *cli.App {
	flags := flags2.Flags
	return &cli.App{
		Version:              "0.0.1",
		Description:          "An  market services with rpc",
		EnableBashCompletion: true,
		Commands: []*cli.Command{
			{
				Name:        "rpc",
				Flags:       flags,
				Description: "Run rpc services",
				Action:      cliapp.LifecycleCmd(runRpc),
			},
			{
				Name:        "version",
				Description: "Show project version",
				Action: func(ctx *cli.Context) error {
					cli.ShowVersion(ctx)
					return nil
				},
			},
		},
	}
}
