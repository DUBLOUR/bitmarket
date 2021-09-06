package main

import (
	"bitmarket/internal/authentificator"
	"bitmarket/internal/ratemodel"
	"bitmarket/internal/server"
)

const ServerPort string = ":16100"
const LogFile string = "data/log"

func main() {

	s := &server.Server{
		Model:     ratemodel.MicroMarketConnector{},
		Auth:      authentificator.Auth{},
		Presenter: server.JsonPresenter{},
		Log:       server.NewFileLogger(LogFile),
		Port:      ServerPort,
	}
	s.Run()

}
