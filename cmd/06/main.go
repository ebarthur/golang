/*
* @Author: lasagna
* @Date:   2024-02-12
* @Github: github.com/ebarthur

* Arrays
* Structs
* Interfaces
 */

package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
	// owner
}

type person struct {
	name string
	age  uint32
	job  []string
}

type owner struct {
	name string
}

func (e gasEngine) milesleft() uint8 {
	return e.gallons * e.mpg
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesleft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("Need to fuel up first!")
	}
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

// Interface
type engine interface {
	milesleft() uint8
}

func (e electricEngine) milesleft() uint8 {
	return e.mpkwh * e.kwh
}

func main() {

	Kofi := person{"Kofi", 25, []string{"Software Engineer", "Data Scientist"}}

	fmt.Println(Kofi.age) // 25
	// var myEngine gasEngine
	// var myEngine2 gasEngine = gasEngine{25, 15, owner{"Vin"}}
	// myEngine.mpg = 20

	// fmt.Println(myEngine.mpg, myEngine.gallons)
	// fmt.Println(myEngine2.mpg, myEngine2.gallons, myEngine2.owner)

	// fmt.Printf("Total miles left in tank: %v", myEngine2.milesleft())

	var myEngine = gasEngine{25, 15}
	canMakeIt(myEngine, 50)

	var myEngine1 = electricEngine{25, 15}
	canMakeIt(myEngine1, 50)
}
