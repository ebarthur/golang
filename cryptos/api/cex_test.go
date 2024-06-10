package api_test

import (
	"github.com/ebarthur/golang/cryptos/api"
	"testing"
)

// Entry errors have been handled
func TestGetRate(t *testing.T) {
	_, err := api.GetRate("")

	if err == nil {
		t.Error("Error was not found")
	}
}

// Test for currency codegit
func TestGetBTCRate(t *testing.T) {
	rate, err := api.GetRate("btc")

	if err != nil {
		t.Error("Error was found")
	}

	if rate.Currency != "BTC" {
		t.Error("Currency code is not BTC")
	}
}
