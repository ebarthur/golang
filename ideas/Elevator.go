package main

import (
	"fmt"
)

// Lift represents an individual lift
type Lift struct {
	currentLevel int // Current level of the lift
	maxLevel     int // Maximum level of the building
}

// NewLift initializes a new Lift object with the maximum level
func NewLift(maxLevel int) *Lift {
	return &Lift{currentLevel: 1, maxLevel: maxLevel}
}

// MoveToLevel moves the lift to the specified level
func (l *Lift) MoveToLevel(level int) {
	// Check if the requested level is valid
	if level < 1 || level > l.maxLevel {
		fmt.Println("Invalid level.")
		return
	}

	l.currentLevel = level
	fmt.Printf("Lift moved to level %d.\n", l.currentLevel)
}

// Operate simulates lift operation
func (l *Lift) Operate() {
	fmt.Println("Lift started.")
	// Simulate lift operation, e.g., by moving to requested levels.
	fmt.Println("Lift reached all destinations.")
}

// Building represents a building with multiple lifts
type Building struct {
	lifts []Lift // Slice of Lift objects representing all lifts in the building
}

// NewBuilding creates a new Building object with the specified number of lifts and maximum level
func NewBuilding(numLifts, maxLevel int) *Building {
	building := &Building{}
	for i := 0; i < numLifts; i++ {
		building.lifts = append(building.lifts, *NewLift(maxLevel))
	}
	return building
}

// RequestLift requests a lift to a specific level for a given lift index
func (b *Building) RequestLift(liftIndex, level int) {
	// Check if the lift index is valid
	if liftIndex < 0 || liftIndex >= len(b.lifts) {
		fmt.Println("Invalid lift index.")
		return
	}

	// Move the requested lift to the specified level
	b.lifts[liftIndex].MoveToLevel(level)
}

// OperateLifts operates all lifts in the building
func (b *Building) OperateLifts() {
	for i := range b.lifts {
		b.lifts[i].Operate()
	}
}

func main() {
	var numLifts, maxLevel int

	// User input: number of lifts and maximum level
	fmt.Print("Enter the number of lifts in the building: ")
	fmt.Scan(&numLifts)
	fmt.Print("Enter the maximum level of the building: ")
	fmt.Scan(&maxLevel)

	// Create a Building object based on user input
	building := NewBuilding(numLifts, maxLevel)

	var liftIndex, destination int

	// User input loop for lift requests
	for {
		// Prompt user to enter lift index and destination level or -1 to operate lifts
		fmt.Printf("Enter lift index (0 to %d) and destination level (-1 to operate lifts): ", numLifts-1)
		fmt.Scan(&liftIndex)

		// Break the loop if user enters -1
		if liftIndex == -1 {
			break
		}

		fmt.Scan(&destination)

		// Request a lift to the specified level
		building.RequestLift(liftIndex, destination)
	}

	// Operate all lifts in the building
	building.OperateLifts()
}
