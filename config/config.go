package config

import (
	"github.com/urfave/cli/v2"

	"github.com/the-web3/market-services/flags"
)

type Config struct {
	RpcServer ServerConfig
}

type ServerConfig struct {
	Host string
	Port int
}

func NewConfig(ctx *cli.Context) Config {
	return Config{
		RpcServer: ServerConfig{
			Host: ctx.String(flags.RpcHostFlag.Name),
			Port: ctx.Int(flags.RpcPortFlag.Name),
		},
	}
}
