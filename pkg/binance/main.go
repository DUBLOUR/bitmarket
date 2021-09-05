package binance

import (
	"bitmarket/pkg/generalApiReader"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type Binance struct {
	Endpoint string
}

func New(endpoint string) Binance {
	return Binance{endpoint}
}

func Default() Binance {
	return Binance{DefaultEndpoint}
}

func (b Binance) createRequest(currency string) (*http.Request, error) {
	//docs: https://binance-docs.github.io/apidocs/
	baseURL, err := url.Parse(b.Endpoint)
	if err != nil {
		return &http.Request{}, err
	}
	params := url.Values{}
	params.Add("symbol", currency)
	baseURL.RawQuery = params.Encode()

	r, err := http.NewRequest("GET", baseURL.String(), nil)
	if err != nil {
		return &http.Request{}, err
	}
	return r, nil
}

func (b Binance) CurrentRate(currency string) (float64, error) {
	log.Println("Binance:", currency)
	req, err := b.createRequest(currency)
	if err != nil {
		return 0.0, err
	}

	response := new(struct {
		Symbol string `json:"symbol"`
		Price  string `json:"price"`
	})

	if err := generalApiReader.JsonRequest(req, &response); err != nil {
		return 0.0, err
	}

	price, err := strconv.ParseFloat(response.Price, 64)
	if err != nil {
		return 0.0, fmt.Errorf("incorrect market response")
	}

	return price, nil
}
