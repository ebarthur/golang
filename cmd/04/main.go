/*
* @Author: lasagna
* @Date:   2024-02-12
* @Github: github.com/ebarthur

* Arrays
* Slices
* Maps
* Loops
 */

package main

import "fmt"

 func main() {

	// Arrays
	var intArr [3]int32 // [0, 0, 0]
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	var intArr2 [2]int32
	intArr2[1] = 5
	fmt.Println(intArr2[1]) // 5
	fmt.Println(intArr2[0:]) // [0, 5]

	// print memory location of an erray using &
	fmt.Println(&intArr2[0])
	fmt.Println(&intArr2[1])

	// var intArr3 [3]int32 = [3]int32{-6, 7, 2}
	intArr3 := [3]int32{-6, 7, 2}
	fmt.Println(intArr3) // [-6 7 2]

	/*
	* Slices wrap arrays to give a more general, powerful, and convevient 
	* interface to sequences of data 
	 */

	 intSlice := []int32{4, 5, 6}
	 fmt.Println(intSlice)
	 intSlice = append(intSlice, 7)
	 fmt.Printf("intSlice but appended: %v\n", intSlice)

	/*
	* Maps: {"key": "value"} pairs
	* map[string]int32
	*/

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap) // map[]

	var myMap2 = map[string]uint8{"Adam":23, "Sarah": 45, "Greg": 21, "Jamie": 27}
	fmt.Println(myMap2) // map[Adam:23 Greg:21 Jamie: 27 Sarah:45]
	fmt.Println(myMap2["Adam"]) // 23

	delete(myMap2, "Sarah")
	fmt.Println(myMap2)

	var age, ok = myMap2["Jason"] // ok returns true if value is in the Map
	if ok{
		fmt.Printf("The age is %v", age)
	}else{
		fmt.Println("Invalid Name")
	}

	/*
	* Loops 
	*/
	
	// Iterating over Maps

	for name, age := range myMap2{
		fmt.Printf("Name: %v, Age: %v\n" , name, age)
	}

	// Go does not have while loop but the same function is achieved by for keyword
	var i int = 0
	for i < 10 {
		fmt.Println(i)
		i = i + 1
	}

	for i:=0; i<10; i++ {
		fmt.Printf("Count: %v\n", i)
	}
}