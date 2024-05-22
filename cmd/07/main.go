/*
* @Author: lasagna
* @Date:   2024-02-12
* @Github: github.com/ebarthur

* Pointers:
 */

package main

import "fmt"

func square(thing2 [5]float64) [5]float64 {
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return thing2
}
func main() {
	var p *int32 = new(int32)
	var i int32

	var thing1 = [5]float64{1, 2, 3, 4, 5}
	var result [5]float64 = square(thing1)
	fmt.Printf("\nThe result is: %v", result)

	fmt.Printf("\nThe value p points to is: %v", *p)
	fmt.Printf("\nThe value of i is: %v", i)

	*p = 10
	fmt.Printf("\nThe value p points to is: %v", *p)

}
