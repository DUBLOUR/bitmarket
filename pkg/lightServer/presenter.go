package lightServer

import "fmt"

type PlainTextPresenter struct{}

func (PlainTextPresenter) Str(data interface{}) string {
	return fmt.Sprintf("%v", data)
}
