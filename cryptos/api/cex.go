package api

import (
	"encoding/json"
	"fmt"
	m "github.com/ebarthur/golang/cryptos/model"
	"io"
	"net/http"
	"strings"
)

// API_URL is the URL of the cex API.
// It uses the %s verb to inject the currency
const API_URL = "https://cex.io/api/ticker/%s/USD"

// GetRate gets the exchange rate of a currency
func GetRate(currency string) (*m.Rate, error) {

	// check for entry errors
	if len(currency) != 3 {
		return nil, fmt.Errorf("currency code must be 3 characters")
	}

	// upCurrency converts the currency to uppercase before making the request
	// For example, if currency is "btc", upCurrency will be "BTC"
	upCurrency := strings.ToUpper(currency)

	r, err := http.Get(fmt.Sprintf(API_URL, upCurrency))

	if err != nil {
		return nil, err
	}

	// Handle http errors
	if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error: %s", r.Status)
	}

	// Read the body of the response
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshal the response into a CEXResponse object
	var response m.CEXResponse
	err = json.Unmarshal(bodyBytes, &response)

	if err != nil {
		return nil, err
	}

	// Create a new rate object
	rate := &m.Rate{Currency: upCurrency, Price: response.Bid}

	// return rate and nil error
	return rate, nil

}
