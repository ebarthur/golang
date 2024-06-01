/*
* @Author: lasagna
* @Date:   2024-02-12
* @Github: github.com/ebarthur

* Arrays
* Slices
* Make
* Copy
* Maps
* Loops
 */

package main

import (
	"fmt"
)

func main() {

	// Arrays
	var intArr [3]int32 // [0, 0, 0]
	fmt.Println(intArr)

	intArr[1] = 5

	fmt.Println(intArr) // [0, 5, 0]

	fmt.Println(intArr[1])
	fmt.Println(intArr[1:3])

	var intArr2 [2]int32
	intArr2[1] = 5
	fmt.Println(intArr2[1])  // 5
	fmt.Println(intArr2[0:]) // [0, 5]

	// print memory location of an array using & pointer
	fmt.Println(&intArr2[0])
	fmt.Println(&intArr2[1])

	// var intArr3 [3]int32 = [3]int32{-6, 7, 2}
	intArr3 := [3]int32{-6, 7, 2}
	fmt.Println(intArr3) // [-6 7 2]

	/*
	* Slices wrap arrays to give a more general, powerful, and convenient
	* interface to sequences of data
	 */

	intSlice := []int32{4, 5, 6}

	fmt.Println(intSlice)
	fmt.Println(len(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("intSlice but appended: %v\n", intSlice)

	// more than one value can be appended to a slice
	someArray := []int32{11, 56, 42, 3}

	someArray = append(someArray, intSlice...) // similar to the rest operator in JavaScript

	fmt.Println(someArray)

	/*
	* Make: allows to create empty slice that already has a length or capacity specified
	 */

	someSlice := make([]int32, 5)        // empty slice with specified length
	anotherSlice := make([]int32, 5, 10) // this time I specify the length and the capacity

	someSlice = append(someSlice, 5) // number gets added to the end
	anotherSlice = append(anotherSlice, 6, 7, 8, 9, 0, 1)

	fmt.Println("Some Slice: ", someSlice)
	fmt.Println(anotherSlice)

	/*
	* Slicing slices
	 */

	lasagna := []int32{1, 2, 4}

	fmt.Println(lasagna[1:]) // just like in python

	/*
	* Copy - creates a slice that's independent of the original.
	* It takes two parameters; copy(destination, source)
	 */

	a := []int{2, 1, 4}
	y := make([]int, len(a))

	z := copy(y, a) // y gets populated with the elements in 'a'
	fmt.Println(y, z)
	fmt.Printf("Copied: %v", y)

	/*
	* Maps: {"key": "value"} pairs
	* map[string]int32
	 */

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap) // map[]

	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45, "Greg": 21, "Jamie": 27}
	fmt.Println(myMap2)         // map[Adam:23 Greg:21 Jamie: 27 Sarah:45]
	fmt.Println(myMap2["Adam"]) // 23

	delete(myMap2, "Sarah")
	fmt.Println(myMap2)

	var age, ok = myMap2["Jason"] // ok returns true if value is in the Map
	if ok {
		fmt.Printf("The age is %v", age)
	} else {
		fmt.Println("Invalid Name")
	}

	groups := map[int][]string{1: {"Nhyira, Kwarteng"}, 2: {"Samia, Arthur"}, 3: {"Phil, Adom"}}

	// using comma ok idiom: checks if the key exists in the map
	one1, ok := groups[1]

	fmt.Printf("Group 1: %s\n", one1)
	fmt.Println(ok) // true

	delete(groups, 3)

	fmt.Println(groups)

	/*
	* Loops
	 */

	// Iterating over Maps

	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Age: %v\n", name, age)
	}

	// Go does not have while loop but the same function is achieved by for keyword
	var i int = 0
	for i < 10 {
		fmt.Println(i)
		i = i + 1
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("Count: %v\n", i)
	}
}
