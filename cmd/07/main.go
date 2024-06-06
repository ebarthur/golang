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

func birthday(age *int32) {
	*age += 1
}

func readUser(id int) (user string, err error) {
	// ... we proceed with the reading and see a bool ok value
	ok := true

	if ok {
		return user, nil
	} else {
		return "", fmt.Errorf("User not found", err)
	}
}

func main() {
	var p *int32 = new(int32)
	var i int32

	var thing1 = [5]float64{1, 2, 3, 4, 5}
	var result [5]float64 = square(thing1)

	fmt.Println(thing1)
	fmt.Printf("The result is: %v", result)

	fmt.Printf("\nThe value p points to is: %v", *p)
	fmt.Printf("\nThe value of i is: %v", i)

	*p = int32(result[3])
	fmt.Printf("\nThe value p points to is: %v", *p)

	age := int32(25)
	birthday(&age)
	fmt.Printf("\nAge is: %v\n", age)

	user, err := readUser(1)
	fmt.Println(user, err)

	// panic & defer
	// panic: panic("a problem")
	// defer: defer fmt.Println("This was deferred")

}
