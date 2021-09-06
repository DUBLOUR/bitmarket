package ratemodel

type MicroMarketConnector struct {
	DefaultCoin   string
	DefaultMarket string
}

func (m MicroMarketConnector) CurrentRate() (interface{}, error) {
	return get(m.DefaultCoin, m.DefaultMarket)
}

func get(coin, market string) (interface{}, error) {
	return 0.0, nil
}
