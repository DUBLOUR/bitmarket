package lightServer

import (
	"log"
	"net/http"
)

type IMoney interface{}

type IMarket interface {
	CurrentRate(string) (IMoney, error)
}

type IMarketManager interface {
	AddMarket(string, IMarket)
	SetDefaultMarket(string) error
	CurrentRate(string) (IMoney, error)
	CurrentRateInMarket(string, string) (IMoney, error)
}

type IPresenter interface {
	Str(interface{}) string
}

type Server struct {
	Port          string
	MarketManager IMarketManager
	Presenter     IPresenter
}

func (s *Server) Run() {
	server := &http.Server{
		Addr:    s.Port,
		Handler: Handlers(s.MarketManager, s.Presenter),
	}

	log.Fatal(server.ListenAndServe())
}
