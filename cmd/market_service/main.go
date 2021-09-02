package market_service

import (
	"genesis_se/bitmarket/pkg/binance"
	"genesis_se/bitmarket/pkg/lightServer"
)

const DefaultPort string = "16101"

func main() {
	s := lightServer.Server{
		Port: DefaultPort,
		MarketManager: lightServer.MarketManager{
			Market: lightServer.NewBinanceAdapter(binance.Default()),
		},
		Presenter: lightServer.PlainTextPresenter{},
	}

	s.Run()

}
