package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {

	done := make(chan bool)

	wg.Add(1)
	go printMessage("Hola!", done)

	tb := time.Now()
	wg.Add(1)
	go countDown(10, done)

	wg.Wait()
	fmt.Println("Total execution time: ", time.Since(tb))

	time.Sleep(time.Duration(2000) * time.Millisecond)
	fmt.Println("I just played with Goroutines")

}

func printMessage(text string, done chan bool) {
	defer wg.Done()
	delay := 500
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(delay) * time.Millisecond)
		fmt.Println(text)
	}
	close(done) // Signal that printMessage is done
}

func countDown(num int, done chan bool) {
	<-done // Wait for printMessage to finish
	defer wg.Done()
	delay := 1000
	for i := num; i >= 0; i-- {
		time.Sleep(time.Duration(delay) * time.Millisecond)
		fmt.Println(i)
	}
}
