/*
* @Author: lasagna
* @Date:   2024-02-12
* @Github: github.com/ebarthur

* Arrays
* Strings
* Runes
* Bytes
 */

package main

import "fmt"

func main(){
	var myString = []rune("rèsumè")
	var indexed = myString[1]

	fmt.Printf("%v, %T\n", indexed, indexed) // 232 int32

	fmt.Println(indexed) // 232
	fmt.Println(myString) // [114 232 115 117 109 232]

	for i, v := range myString{
		fmt.Println(i, v)
	}

	var char rune = 'a'
	var myRune = 'a'
	fmt.Println(char) // 97
	fmt.Println(myRune) // 97

	// String Building
	var strSlice = []string {"s","o","f","t","w","a","r","e"}
	var catStr = ""
	for i := range strSlice{
	catStr += strSlice[i]
	}
	fmt.Printf("\n%v", catStr) // software
}