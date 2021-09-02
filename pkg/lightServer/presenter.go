package lightServer

type PlainTextPresenter struct {}

func (PlainTextPresenter) Str(data interface{}) string {
	return ""
}

