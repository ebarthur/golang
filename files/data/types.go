package main

import "fmt"

// type-aliasing
type integer = int
type json = map[string]string

// for semantics
type distance float64 // miles
type distanceKm float64

func main() {

	d := distance(10.5)
	d.toKm()
	fmt.Println("Distance in Km:", d.toKm()) //

	k := distanceKm(16.9)
	k.toMiles()
	fmt.Println("Distance in Miles:", k.toMiles()) //

	var x integer = 10
	var config json = map[string]string{"name": "John Doe", "age": "30"}

	fmt.Println(x, "\n", config)

}

// Defining Methods for the Types
// Method: toKm
func (miles distance) toKm() distanceKm {
	return (distanceKm(miles * 1.609340))
}

// Method: toMiles
func (kilometer distanceKm) toMiles() distance {
	return (distance(kilometer / 1.609340))
}
