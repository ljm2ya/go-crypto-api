package exchanges

import (
	"errors"
	"github.com/ljm2ya/go-crypto-api/environment"
)

type ExchangeWrapper interface {
	Name() string
	GetMarketSummary(market *enviornment.Market) (*environment.MarketSummary, error)
	GetOrderBook(market *environment.Market) (*environment.OrderBook, error)
	GetTicker(market *environment.Market) (*environment.Ticker, error)

	LimitOrder(market *environment.Market, side *environment.OrderSide, size float64, price float64) (*environment.OrderResult, error)
	MarketOrder(market *environment.Market, side *environment.OrderSide, size float64) (*environment.OrderResult, error)

	ImmediateOrder(market *environment.Market, side *environment.OrderSide, size float64) error
	RenewalLimitOrder(market *environment.Market, side *environment.OrderSide, size float64) error

	GetTradingFee(market *environment.Market, orderType *environment.OrderType) float64
	GetBalance(symbol string) (float64, error)

	String() string
}
