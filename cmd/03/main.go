/*
* @Author: lasagna
* @Date:   2024-02-12
* @Github: github.com/ebarthur

* Functions in Go
* Functions are the building blocks of a Go program

*Control structures in Go
* Go has the following control structures:
* - if
* - else
* - else if
* - for - the complete for, the conditional-only, infinite, for-range
* - switch etc
 */

package main

import "fmt"

func main() {
	var movieStar string = "Liam Leeson"
	printName(movieStar)

	var num1, num2 int = 10, 0
	result, remainder, err := integerDivision(num1, num2)
	if err != nil {
		fmt.Println(err)
	} else if remainder == 0 {
		fmt.Printf("The result of the integer division is %v", result)
	} else {
		fmt.Printf("The result of the integer division is %v and the remainder is %v", result, remainder)
	}

	// Div()
	Div()

}

func printName(name string) {
	fmt.Println(name)
}

func integerDivision(numerator int, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, fmt.Errorf("denominator cannot be zero")
	}

	var result int = numerator / denominator
	var remainder int = numerator % denominator

	return result, remainder, nil
}

func Div() {
	for i := 0; i < 10; i++ {
		switch {
		case i%2 == 0:
			fmt.Println(i, "is even")
		case i%3 == 0:
			fmt.Println(i, "is divisible by 3 but not 2")
		case i%7 == 0:
			fmt.Println(i, "is divisible by 7 but not 2 or 3")
			break
		default:
			fmt.Println(i, "is boring")
		}
	}
}
