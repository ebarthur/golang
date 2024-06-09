package main

import (
	"fmt"
	"github.com/ebarthur/golang/cryptos/api"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

// Cryptos is a simple API application that provides exchange rates for cryptocurrencies
// It fetches the rates from various exchanges using cex API
func main() {
	currencies := []string{"btc", "sol", "eth", "ltc", "ada"}
	tb := time.Now()

	for _, currency := range currencies {

		wg.Add(1)
		go func(code string) {
			GetCurrencyRate(code)
			wg.Done()
		}(currency)

	}
	wg.Wait()
	fmt.Println("These are today's crypto rates", time.Now().UTC())
	fmt.Println("Total execution time: ", time.Since(tb).Seconds(), "seconds")
}

func GetCurrencyRate(currency string) {
	rate, err := api.GetRate(currency)
	if err != nil {
		fmt.Println("Error getting rate for", currency, ":", err)
		return
	}
	update := fmt.Sprintf(`the rate of %s is %f`, rate.Currency, rate.Price)

	fmt.Println(update)
}
