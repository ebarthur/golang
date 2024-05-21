package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	message := "Hello from Go \n"
	price := 10.50

	const pi = 3.14159

	fmt.Println(message, price)
	fmt.Println(pi)

	test()
}

func test() {
	const y = "halloween"

	var s int = len(y)
	var v int = utf8.RuneCountInString(y)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(s) // same results
	fmt.Println(v) // same results

	//x += 1    // Error, cannot assign to x
	//y = "bye" // Error, cannot assign to y
}

const x int64 = 10

const (
	idKey   = "id"
	nameKey = "name"
)

const z = 20 * 10
