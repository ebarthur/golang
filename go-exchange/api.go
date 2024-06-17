package main

import (
	"encoding/json"
	"fmt"
	"github.com/ebarthur/go-exchange/model"
	"io"
	"log"
	"net/http"
	"strings"
)

const baseUrl = "https://cex.io/api/ticker/%s/USD"

type APIServer struct {
	Addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{Addr: addr}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()

	router.HandleFunc("GET /crypto/{CEX}", func(w http.ResponseWriter, r *http.Request) {

		currency := r.PathValue("CEX")

		upCurrency := strings.ToUpper(currency)

		rate, err := GetCurrency(upCurrency)

		err = json.NewEncoder(w).Encode(&rate)

		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}

	})

	server := http.Server{
		Addr:    s.Addr,
		Handler: RequestLogger(router),
	}
	log.Print("Server is running at ", s.Addr)

	return server.ListenAndServe()
}

func GetCurrency(currency string) (*model.Rate, error) {

	if len(currency) != 3 {
		return nil, fmt.Errorf("currency length expected 3, got: &v", len(currency))
	}

	r, err := http.Get(fmt.Sprintf(baseUrl, currency))

	bodyBytes, err := io.ReadAll(r.Body)

	var resp *model.Response
	err = json.Unmarshal(bodyBytes, &resp)

	if err != nil {
		return nil, err
	}

	return &model.Rate{
		Price:    resp.Bid,
		Currency: currency,
	}, err
}

func RequestLogger(next http.Handler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("METHOD: @%s\n URL PATH: %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}
}
