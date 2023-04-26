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

// The GetPriceCoin function is a function that takes a "coin" string parameter that specifies the cryptocurrency coin and retrieves the current price of the specified coin from the CoinGecko API.
func GetPriceCoin(coin string) (float64, error) {

	////Creates a new GET request to the CoinGecko API with the specified cryptocurrency coin
	res, err := resty.R().Get(fmt.Sprintf("https://api.coingecko.com/api/v3/coins/%s", coin)) //Sends the request using the resty package and retrieves the response.
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	var data CoinGecko

	err = json.Unmarshal(res.Body(), &data) //Unmarshals the JSON response into a CoinGecko struct
	if err != nil {
		return 0, err
	}
	return data.MarketData.CurrentPrice.Usd, nil //Returns the current price of the cryptocurrency coin in USD from the CoinGecko struct and nil error if no error occurs
}
