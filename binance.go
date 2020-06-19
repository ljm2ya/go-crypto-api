package main

import (
	"github.com/adshao/go-binance"
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
