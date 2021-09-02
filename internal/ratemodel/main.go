package ratemodel

type RateModel struct {
	DefaultCoin   string
	DefaultMarket string
	market        MarketAdapter
}

func (r RateModel) Current() (interface{}, error) {
	return r.market.Request(r.DefaultCoin, r.DefaultMarket)
}

type MarketAdapter struct{}

func (MarketAdapter) Request(coin, market string) (interface{}, error) {
	return 0.0, nil
}
