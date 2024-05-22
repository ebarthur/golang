/*
* @Author: lasagna
* @Date:   2024-02-12
* @Github: github.com/ebarthur'


* Generics
 */

package main

import (
	"fmt"
)

// Swap function to swap two elements of any type
func Swap[T any](a, b T) (T, T) {
	return b, a
}

func main() {
	// Swap two integers
	x, y := 10, 20
	fmt.Printf("Before swap: x = %d, y = %d\n", x, y)
	x, y = Swap(x, y)
	fmt.Printf("After swap: x = %d, y = %d\n", x, y)

	// Swap two strings
	str1, str2 := "hello", "world"
	fmt.Printf("Before swap: str1 = %s, str2 = %s\n", str1, str2)
	str1, str2 = Swap(str1, str2)
	fmt.Printf("After swap: str1 = %s, str2 = %s\n", str1, str2)

	// Swap two slices
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	fmt.Printf("Before swap: slice1 = %v, slice2 = %v\n", slice1, slice2)
	slice1, slice2 = Swap(slice1, slice2)
	fmt.Printf("After swap: slice1 = %v, slice2 = %v\n", slice1, slice2)
}
