package server

import (
	"bitmarket/internal/userModel"
	"net/http"
)

type ILog interface {
	Debug(...interface{})
	Info(...interface{})
	Warn(...interface{})
}

type IPresenter interface {
	Format(interface{}) (string, error)
}

type IModel interface {
	CurrentRate() (interface{}, error)
}

type IAuth interface {
	Register(userModel.ICredentials) error
	Login(userModel.ICredentials) (userModel.IToken, error)
	LoginByToken(userModel.IToken) error
}

type Server struct {
	Model     IModel
	Auth      IAuth
	Presenter IPresenter
	Log       ILog
	Port      string
}

func (s *Server) Run() {
	h := Handlers(s.Auth, s.Presenter, s.Model, s.Log)
	server := &http.Server{
		Addr:    s.Port,
		Handler: h,
	}

	s.Log.Warn(server.ListenAndServe())

}
