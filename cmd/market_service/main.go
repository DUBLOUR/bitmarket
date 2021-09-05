package main

import (
	"bitmarket/pkg/binance"
	"bitmarket/pkg/lightServer"
)

const DefaultPort string = ":16101"

func main() {
	manager := &lightServer.MarketManager{}
	manager.AddMarket("binance", lightServer.NewBinanceAdapter(binance.Default()))
	//manager.SetDefaultMarket("binance")

	s := lightServer.Server{
		Port:          DefaultPort,
		MarketManager: manager,
		Presenter:     lightServer.PlainTextPresenter{},
	}

	s.Run()

}
