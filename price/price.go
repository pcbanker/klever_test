package price

import (
	"encoding/json"
	"fmt"

	"gopkg.in/resty.v1"
)

type CoinGecko struct {
	MarketData struct {
		CurrentPrice struct {
			Usd float64 `json:"usd"`
		} `json:"current_price"`
	} `json:"market_data"`
}

func GetPriceCoin(coin string) (float64, error) {
	res, err := resty.R().Get(fmt.Sprintf("https://api.coingecko.com/api/v3/coins/%s", coin))
	if err != nil {
		return 0, err
	}
	var data CoinGecko
	err = json.Unmarshal(res.Body(), &data)
	if err != nil {
		return 0, err
	}
	return data.MarketData.CurrentPrice.Usd, nil
}
