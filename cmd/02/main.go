/*
* @Author: lasagna
* @Date:   2024-02-12
* @Github: github.com/ebarthur

*Constants
* Variables
* Basic Data Types
 */

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int = 32767
	intNum += 1	

	fmt.Println(intNum)

	// float64 stores most precise floating point numbers
	
	var floatNum float64 = 12345.6789
	var floatNum2 float32 = 12345.6789

	fmt.Println(floatNum) // 12345.6789
	fmt.Println(floatNum2) // 12345.679

	// adding two variables of different types is not allowed in Go 
	// fmt.Println(intNum + floatNum) // invalid operation: intNum + floatNum (mismatched types int and float64)
	// Unless you type cast the variables

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var resultfloatNum32 float32 = floatNum32 + float32(intNum32)

	fmt.Println(resultfloatNum32) // 12.1

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2) // 1
	fmt.Println(intNum1 % intNum2) // 1

	// var myString string = "Hello 
	// World!"

	// fmt.println(myString) // syntax error: unexpected newline, expecting comma or )

	var myString string = "Hello\nWorld!"
	var myString2 string = `Hello
	World!`
	var myString3 string = "Hello" + " " + "World!"

	fmt.Println(myString)
	fmt.Println(myString2)
	fmt.Println(myString3)

	var location string = "California"
	var anottherString string = "I live in" + " " + location

	fmt.Println(len(location)) // 10
	fmt.Println(utf8.RuneCountInString(anottherString)) // 20

	var myRune rune = 'a' // rune is an alias for int32
	fmt.Println(myRune) // 97 (ASCII value of 'a')

	var myBoolean bool = true
	fmt.Println(myBoolean) // true

	var myComplex complex64 = 1 + 2i
	fmt.Println(myComplex) // (1+2i)

	// Drop the variable type and let Go infer it
	var myString4 = "Hasta Luego!"
	fmt.Println(myString4)

	// Drop the var keyword and use the short declaration operator

	food := "Pizza"

	fmt.Println(food) // Pizza

	// Multiple assignment
	num1, num2 := 1, 2

	fmt.Println(num1, num2) // 1 2

	// Constants are declared using the const keyword
	// Constants are immutable

	const year int = 2024
	fmt.Println(year) 

	const pi float32 = 3.14159
	fmt.Println(pi)
}
