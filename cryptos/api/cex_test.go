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
