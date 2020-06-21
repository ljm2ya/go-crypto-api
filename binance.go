package exchanges

import (
	"github.com/adshao/go-binance"
	"github.com/ljm2ya/go-crypto-api/environment"
)

var (
	apiKey = "your api key"
	secretKey = "your secret key"
)

client := binance.NewClient(apiKey, secretKey)
futuresClient := binance.NewFuturesClient(apiKey, secretKey)

func newOrder() {

}

func
