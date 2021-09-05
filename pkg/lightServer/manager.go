package lightServer

import "bitmarket/pkg/binance"

type IMarket interface {
	CurrentRate(string) (IMoney, error)
}

type MarketManager struct {
	Market IMarket
}

func (manager MarketManager) CurrentRate(coin string) (IMoney, error) {
	return manager.Market.CurrentRate(coin)
}

type BinanceAdapter struct {
	b binance.Binance
}

func (b BinanceAdapter) CurrentRate(currency string) (IMoney, error) {
	m, err := b.CurrentRate(currency)
	return IMoney(m), err
}

func NewBinanceAdapter(b binance.Binance) BinanceAdapter {
	return BinanceAdapter{b}
}