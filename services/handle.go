package services

import (
	"context"

	"github.com/the-web3/market-services/proto/market"
)

func (ms *MarketRpcService) GetSupportAsset(ctx context.Context, in *market.SupportAssetRequest) (*market.SupportAssetResponse, error) {
	return &market.SupportAssetResponse{
		ReturnCode: 100,
		Message:    "support this asset",
		IsSupport:  true,
	}, nil
}

func (ms *MarketRpcService) GetMarketPrice(ctx context.Context, in *market.MarketPriceRequest) (*market.MarketPriceResponse, error) {
	var marketPrice []*market.MarketPrice
	var coinRate []*market.OfficialCoinRate
	return &market.MarketPriceResponse{
		ReturnCode:       100,
		Message:          "get asset market success",
		MarketPrice:      marketPrice,
		OfficialCoinRate: coinRate,
	}, nil
}
