package environment

import (
	"fmt"
)

type Market struct {
	Exchanges string
	BaseCurrency string
	MarketCurrency string
}

type MarketSummary struct {

}

type OrderBook struct {
	Ask []Tick
	Bid []Tick
}

type Tick struct {
	price float64
	size float64
}

type Ticker struct {
	BestAsk Tick
	BestBid Tick
}
