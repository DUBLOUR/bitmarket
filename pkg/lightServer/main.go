package lightServer


type IMarketManager interface {
	CurrentRate(string) (IMoney, error)
}

type IPresenter interface {
	Str(interface{}) string
}

type IMoney interface{}



type Server struct {
	Port string
	MarketManager IMarketManager
	Presenter IPresenter
}

func (s *Server) Run() {

}







