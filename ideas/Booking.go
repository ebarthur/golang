package main

import (
	"fmt"
)

var seating [13][6]byte = [13][6]byte{
	{'*', '*', 'x', '*', 'x', 'x'},
	{'*', 'x', '*', 'x', '*', 'x'},
	{'*', '*', 'x', 'x', '*', 'x'},
	{'x', '*', 'x', '*', 'x', 'x'},
	{'*', 'x', '*', 'x', '*', '*'},
	{'*', 'x', '*', '*', '*', 'x'},
	{'x', '*', '*', '*', 'x', 'x'},
	{'*', 'x', '*', 'x', 'x', '*'},
	{'x', '*', 'x', 'x', '*', 'x'},
	{'*', 'x', '*', 'x', 'x', 'x'},
	{'*', '*', 'x', '*', 'x', '*'},
	{'*', '*', 'x', 'x', '*', 'x'},
	{'*', '*', '*', '*', 'x', '*'},
}

func main() {
	var choice, row, attempts int
	var seat string
	attempts = 2

	for {
		displayMenu()
		fmt.Print("Enter your choice(1-3): ")
		fmt.Scan(&choice)
		if choice < 1 || choice > 3 {
			fmt.Println("Invalid choice. Please enter a number between 1 and 3")
			attempts--
		} else {
			break
		}

		if attempts == 0 {
			fmt.Println("Too many wrong choices. Aborting program")
			return
		}
	}

	showSeatingArrangement()

	for attempts > 0 {
		fmt.Print("Enter the desired row (1-13) and seat letter (A-F): ")
		fmt.Scan(&row, &seat)

		if row < 1 || row > 13 || seat < "A" || seat > "F" {
			fmt.Println("Invalid row or seat. Please enter a valid combination")
			attempts--
		} else if (choice == 1 && row > 2) || (choice == 2 && (row < 3 || row > 7)) || (choice == 3 && row < 8) {
			fmt.Println("Invalid seat for your ticket type. Please select another seat")
			attempts--
		} else {
			if bookSeat(row, seat) {
				break
			}
			attempts--
		}
	}

	if attempts == 0 {
		fmt.Println("Too many wrong choices. Aborting program.")
		return
	}
}

func bookSeat(row int, seat string) bool {
	if isSeatAvailable(row, seat) {
		seating[row-1][seat[0]-'A'] = 'X'
		fmt.Println("Seat booked successfully!")
		return true
	} else {
		fmt.Println("Seat already booked. Please select another seat.")
		return false
	}
}

func isSeatAvailable(row int, seat string) bool {
	return seating[row-1][seat[0]-'A'] == '*'
}

func displayMenu() {
	fmt.Println("Select Ticket Type:")
	fmt.Println("1. First Class")
	fmt.Println("2. Business Class")
	fmt.Println("3. Economy Class")
}

func showSeatingArrangement() {
	fmt.Printf("%7s", " ")
	for i := 0; i < 6; i++ {
		fmt.Printf("%2c ", 'A'+i)
	}
	fmt.Println()

	for i := 0; i < 13; i++ {
		fmt.Printf("Row %2d ", i+1)
		for j := 0; j < 6; j++ {
			fmt.Printf("%2c ", seating[i][j])
		}
		fmt.Println()
	}
}
