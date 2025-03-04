package flags

import "github.com/urfave/cli/v2"

const evnVarPrefix = "MARKET"

func prefixEnvVars(name string) []string {
	return []string{evnVarPrefix + "_" + name}
}

var (

	// RpcHostFlag RPC Service
	RpcHostFlag = &cli.StringFlag{
		Name:     "rpc-host",
		Usage:    "The host of the rpc",
		EnvVars:  prefixEnvVars("RPC_HOST"),
		Required: true,
	}
	RpcPortFlag = &cli.IntFlag{
		Name:     "rpc-port",
		Usage:    "The port of the rpc",
		EnvVars:  prefixEnvVars("RPC_PORT"),
		Value:    8987,
		Required: true,
	}
)

var requireFlags = []cli.Flag{
	RpcHostFlag,
	RpcPortFlag,
}

var optionalFlags = []cli.Flag{}

func init() {
	Flags = append(requireFlags, optionalFlags...)
}

var Flags []cli.Flag
