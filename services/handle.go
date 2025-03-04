package services

import (
	"context"
	"github.com/ethereum/go-ethereum/log"
	"github.com/google/uuid"
	"github.com/the-web3/market-services/database"
	"time"

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
	var marketPriceWrite []database.MarketPrice
	var coinRateWrite []database.OfficialCoinRate
	marketPriceBTC := database.MarketPrice{
		GUID:      uuid.New(),
		AssetName: "BTC",
		PriceUsdt: "80000",
		Volume:    "8000000000",
		Rate:      "10",
		Timestamp: uint64(time.Now().Unix()),
	}
	marketPriceETH := database.MarketPrice{
		GUID:      uuid.New(),
		AssetName: "ETH",
		PriceUsdt: "2000",
		Volume:    "2000000000",
		Rate:      "5",
		Timestamp: uint64(time.Now().Unix()),
	}
	marketPriceWrite = append(marketPriceWrite, marketPriceBTC)
	marketPriceWrite = append(marketPriceWrite, marketPriceETH)

	coinRateItem := database.OfficialCoinRate{
		GUID:      uuid.New(),
		AssetName: "Cny",
		BaseAsset: "USD",
		Price:     "7.3",
		Timestamp: uint64(time.Now().Unix()),
	}

	coinRateWrite = append(coinRateWrite, coinRateItem)

	err := ms.db.MarkerPrice.StoreMarketPrice(marketPriceWrite)
	if err != nil {
		log.Error("store market price fail", "err", err)
		return nil, err
	}

	err = ms.db.OfficialCoinRate.StoreOfficialCoinRate(coinRateWrite)
	if err != nil {
		log.Error("store coin rate fail", "err", err)
		return nil, err
	}

	var marketPriceReturns []*market.MarketPrice
	var coinRateReturns []*market.OfficialCoinRate

	assetPriceList, err := ms.db.MarkerPrice.QueryMarketPriceByAsset("all")
	if err != nil {
		return nil, err
	}
	for _, value := range assetPriceList {
		mItem := &market.MarketPrice{
			AssetName:   value.AssetName,
			AssetPrice:  value.PriceUsdt,
			AssetVolume: value.Volume,
			AssetRate:   value.Rate,
		}
		marketPriceReturns = append(marketPriceReturns, mItem)
	}
	coinRateList, err := ms.db.OfficialCoinRate.QueryOfficialCoinRateByAsset("all")
	if err != nil {
		return nil, err
	}
	for _, value := range coinRateList {
		mItem := &market.OfficialCoinRate{
			Name: value.AssetName,
			Rate: value.Price,
		}
		coinRateReturns = append(coinRateReturns, mItem)
	}
	return &market.MarketPriceResponse{
		ReturnCode:       100,
		Message:          "get asset market success",
		MarketPrice:      marketPriceReturns,
		OfficialCoinRate: coinRateReturns,
	}, nil
}
