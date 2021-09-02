package server

type JsonPresenter struct{}

func (*JsonPresenter) Format(interface{}) string {
	return ""
}
