package services

import (
	"context"
	"fmt"
	"github.com/the-web3/market-services/database"
	"net"
	"sync/atomic"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/ethereum/go-ethereum/log"

	"github.com/the-web3/market-services/proto/market"
)

const MaxRecvMessageSize = 1024 * 1024 * 30000

type MarketRpcConfig struct {
	Host string
	Port int
}

type MarketRpcService struct {
	*MarketRpcConfig

	db *database.DB

	market.UnimplementedMarketServicesServer
	stopped atomic.Bool
}

func NewMarketRpcService(conf *MarketRpcConfig, db *database.DB) (*MarketRpcService, error) {
	return &MarketRpcService{
		MarketRpcConfig: conf,
		db:              db,
	}, nil
}

func (ms *MarketRpcService) Start(ctx context.Context) error {
	go func(ms *MarketRpcService) {
		rpcAddr := fmt.Sprintf("%s:%d", ms.MarketRpcConfig.Host, ms.MarketRpcConfig.Port)
		listener, err := net.Listen("tcp", rpcAddr)
		if err != nil {
			log.Error("Could not start tcp listener. ")
		}

		opt := grpc.MaxRecvMsgSize(MaxRecvMessageSize)

		gs := grpc.NewServer(
			opt,
			grpc.ChainUnaryInterceptor(
				nil,
			),
		)

		reflection.Register(gs)
		market.RegisterMarketServicesServer(gs, ms)

		log.Info("grpc info", "addr", listener.Addr())

		if err := gs.Serve(listener); err != nil {
			log.Error("start rpc server fail", "err", err)
		}
	}(ms)
	return nil
}

func (ms *MarketRpcService) Stop(ctx context.Context) error {
	ms.stopped.Store(true)
	return nil
}

func (ms *MarketRpcService) Stopped() bool {
	return ms.stopped.Load()
}
