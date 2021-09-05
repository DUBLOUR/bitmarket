package lightServer

import (
	"bitmarket/pkg/binance"
	"fmt"
)

type MarketManager struct {
	markets       map[string]IMarket
	kek 		  map[int]int
	defaultMarket string
}

func (manager *MarketManager) AddMarket(name string, entity IMarket) {
	if manager.markets == nil {
		manager.markets = make(map[string]IMarket)
		manager.defaultMarket = name
	}
	manager.markets[name] = entity
}

func (manager *MarketManager) SetDefaultMarket(name string) error {
	if _, has := manager.markets[name]; has == false {
		return fmt.Errorf("SetDefaultMarket: NotImplemented")
	}
	manager.defaultMarket = name
	return nil
}

func (manager MarketManager) CurrentRateInMarket(coin string, marketName string) (IMoney, error) {
	m, has := manager.markets[marketName]
	if has == true {
		return m.CurrentRate(coin)
	}
	return IMoney(0.0), fmt.Errorf("CurrentRateInMarket: NotImplemented")
}

func (manager MarketManager) CurrentRate(coin string) (IMoney, error) {
	return manager.CurrentRateInMarket(coin, manager.defaultMarket)
}

type BinanceAdapter struct {
	adaptee binance.Binance
}

func (adapter BinanceAdapter) CurrentRate(currency string) (IMoney, error) {
	m, err := adapter.adaptee.CurrentRate(currency)
	return IMoney(m), err
}

func NewBinanceAdapter(adaptee binance.Binance) BinanceAdapter {
	return BinanceAdapter{adaptee}
}
