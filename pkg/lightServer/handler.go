package lightServer

import (
	"fmt"
	"log"
	"net/http"
)

func Handlers(manager IMarketManager, presenter IPresenter) http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var marketName string = r.URL.Query().Get("market")
		var currency string = r.URL.Query().Get("currency")
		log.Println(r.URL.Query())

		if currency == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var response string

		if marketName == "" {
			value, err := manager.CurrentRate(currency)
			log.Println(value, err)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			response = presenter.Str(value)
		} else {
			value, err := manager.CurrentRateInMarket(currency, marketName)
			log.Println(value, err)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			response = presenter.Str(value)
		}

		log.Println(response)
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprintf(w, response)
	})

	return r
}
