package routes

import (
	"fmt"

	"net/http"

	"github.com/the-web3/market-services/services/rest/model"
)

func (h Routes) GetSupportAsset(w http.ResponseWriter, r *http.Request) {
	assetName := r.URL.Query().Get("asset_name")
	cr := &model.SupportAssetRequest{
		AssetName: assetName,
	}
	supRet, err := h.srv.GetSupportAsset(cr)
	if err != nil {
		return
	}
	err = jsonResponse(w, supRet, http.StatusOK)
	if err != nil {
		fmt.Println("Error writing response", "err", err.Error())
	}
}

func (h Routes) GetMarketPrice(w http.ResponseWriter, r *http.Request) {
	assetName := r.URL.Query().Get("chain")
	cr := &model.MarketPriceRequest{
		AssetName: assetName,
	}

	addrRet, err := h.srv.GetMarketPrice(cr)
	if err != nil {
		return
	}

	err = jsonResponse(w, addrRet, http.StatusOK)
	if err != nil {
		fmt.Println("Error writing response", "err", err.Error())
	}
}
