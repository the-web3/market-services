package model

type SupportAssetRequest struct {
	ConsumerToken string `json:"consumer_token"`
	AssetName     string `json:"asset_name"`
}

type SupportAssetResponse struct {
	ReturnCode uint64 `json:"return_code"`
	Message    string `json:"message"`
	IsSupport  bool   `json:"is_support"`
}

type OfficialCoinRate struct {
	Name string `json:"name"`
	Rate string `json:"rate"`
}

type MarketPrice struct {
	AssetName   string `json:"asset_name"`
	AssetPrice  string `json:"asset_price"`
	AssetVolume string `json:"asset_volume"`
	AssetRate   string `json:"asset_rate"`
}

type MarketPriceRequest struct {
	ConsumerToken string `json:"consumer_token"`
	AssetName     string `json:"asset_name"`
}

type MarketPriceResponse struct {
	ReturnCode           uint64             `json:"return_code"`
	Message              string             `json:"message"`
	MarketPriceList      []MarketPrice      `json:"market_price_list"`
	OfficialCoinRateList []OfficialCoinRate `json:"official_coin_rate_list"`
}
