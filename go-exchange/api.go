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

		// Get currency from request url
		currency := r.PathValue("CEX")

		// convert currency code to uppercase
		upCurrency := strings.ToUpper(currency)

		// pass the currency code into the GetCurrency function
		rate, err := GetCurrency(upCurrency)

		// encode writer to the return value
		err = json.NewEncoder(w).Encode(&rate)

		// handle errors
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}

	})

	// chain middleware functions here
	Chain := middlewareChain(
		RequireAuthMiddleware,
		RequestLogger,
	)

	// server set-up
	server := http.Server{
		Addr:    s.Addr,
		Handler: Chain(router),
	}

	log.Print("Server is running at ", s.Addr)

	// boot server
	return server.ListenAndServe()
}

// GetCurrency takes in currency code, passes it to the API and reads the response json.
//
// The response json is passed to a reference of the [model.Response] object.
// The referential object is returned into the [model.Rate] object which we return to the writer
// in the http request
func GetCurrency(currency string) (*model.Rate, error) {

	// we make sure that length of currency is 3
	if len(currency) != 3 {
		return nil, fmt.Errorf("currency code must be 3 characters")
	}

	// return response.json(r)
	// http.Get returns response and error
	r, err := http.Get(fmt.Sprintf(baseUrl, currency))

	// read the response body
	bodyBytes, err := io.ReadAll(r.Body)

	// declare a variable that is a pointer to the response struct
	var resp *model.Response

	// unmarshal the response body into the variable resp
	// we use the address of the variable because it is a pointer to the type Response
	// json.Unmarshal returns an error
	err = json.Unmarshal(bodyBytes, &resp)

	// handle error
	if err != nil {
		return nil, err
	}

	// return the model.Rate and an error
	// assign the response value into the Rate struct
	return &model.Rate{
		Timestamp:        resp.Timestamp,
		CurrencyPair:     resp.Pair,
		Price:            resp.Bid,
		Volume:           resp.Volume,
		Change:           resp.PriceChange,
		ChangePercentage: resp.PriceChangePercentage,
	}, err
}

// RequestLogger is a middleware that logs the request method and url path to the terminal.
//
// Example:
// 2024/06/18 14:52:12 METHOD: @GET
// URL PATH: /crypto/btc
func RequestLogger(next http.Handler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("METHOD: @%s\n URL PATH: %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}
}

// RequireAuthMiddleware mocks an auth.
// It checks every request for authorization.
// If the token is not valid, it returns `Unauthorized` and protects the route
// from being accessed by unauthorized users
func RequireAuthMiddleware(next http.Handler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get("Authorization")

		if token != "Bearer token" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}

// Middleware is a custom type for middleware
// It is a http Handler that returns a http HandlerFunc
type Middleware func(http.Handler) http.HandlerFunc

// middlewareChain is a function that chains middleware together
// making it easy to prevent clutter and reduce the number of wraps
// around the router.
func middlewareChain(middlewares ...Middleware) Middleware {

	return func(next http.Handler) http.HandlerFunc {

		for i := len(middlewares) - 1; i >= 0; i-- {

			next = middlewares[i](next)

		}
		return next.ServeHTTP

	}
}
