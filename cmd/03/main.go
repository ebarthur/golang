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
* - for
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
    }else if remainder == 0 {
		fmt.Printf("The result of the integer division is %v", result)
	}else {
        fmt.Printf("The result of the integer division is %v and the remainder is %v", result, remainder)
    }

    // Using Switches: break is not required in Go
    // switch {
    // case err != nil:
    //     fmt.Println(err)
    // case remainder == 0:
    //     fmt.Printf("The result of the integer division is %v", result)
    // default:
    //     fmt.Printf("The result of the integer division is %v and the remainder is %v", result, remainder)
    // }

    extras()
}

func printName(name string) {
    fmt.Println(name)
}

func integerDivision(numerator int, denominator int) (int, int, error){
    if denominator == 0 {
        return 0, 0, fmt.Errorf("Denominator cannot be zero")
    }

    var result int = numerator/denominator
    var remainder int = numerator%denominator

    return result, remainder, nil
}

func extras(){

    if 1 == 1 && 2 == 2 {
		fmt.Println("Passed the check")
	}

    if  1 == 1 || 2 == 2 {
        fmt.Println("Passed the check")
    }
}

func intergerDivisionButWithSwitch(){

}