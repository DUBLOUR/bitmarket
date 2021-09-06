package server

import (
	"fmt"
	"net/http"
)

func Handlers(auth IAuth, presenter IPresenter, model IModel, lg ILog) http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprintf(w, "")
	})

	return r
}
