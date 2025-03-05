package client

import (
	"fmt"
	"testing"
)

func TestSupportAsset(t *testing.T) {
	client := NewMarketPriceClient("http://127.0.0.1:9092")
	result, err := client.GetSupportAsset("all")
	if err != nil {
		fmt.Println("get support asset fail")
		return
	}
	fmt.Println("Support Chain Res:", result)
}

func TestMarketPrice(t *testing.T) {
	client := NewMarketPriceClient("http://127.0.0.1:9092")
	result, err := client.GetMarketPrice("all")
	if err != nil {
		fmt.Println("get market Price fail")
		return
	}
	fmt.Println("Market Price Result:", result)
}
