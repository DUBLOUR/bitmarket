package server

type Server struct {
	Log       ILog
	Model     IModel
	Presenter IPresenter
	Auth      IAuth
	port      string
}

func (*Server) Run() {

}

type ILog interface {
	Info(...interface{})
	Warn(...interface{})
}

type IPresenter interface {
	Format(interface{}) string
}

type IModel interface {
	CurrentRate() (interface{}, error)
}

type IAuth interface {
	Register(TmpUser) error
	Login(TmpUser) (IToken, error)
	LoginByToken(IToken) error
}

type TmpUser struct {
	Email string
	Pass  string
}

type IToken interface {
	Str() string
}
