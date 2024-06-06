package main

import "fmt"

func main() {

	var operation string
	var number1, number2 int

	fmt.Println("CALCULATOR GO 1.0.0")
	fmt.Println("======================")
	fmt.Println("Which operation do you want to perform?(add, subtract, multiply, divide)")
	fmt.Scanf("%s", &operation)
	fmt.Println("Enter the first number: ")
	fmt.Scanf("%d", &number1)
	fmt.Println("Enter the second number: ")
	fmt.Scanf("%d", &number2)

	switch operation {
	case "add":
		fmt.Printf("The sum of %d and %d is %d", number1, number2, number1+number2)

	case "subtract":
		fmt.Printf("The difference of %d and %d is %d", number1, number2, number1-number2)

	case "multiply":
		fmt.Printf("The product of %d and %d is %d", number1, number2, number1*number2)

	case "divide":
		if number2 == 0 {
			fmt.Println("Cannot divide by zero")
		} else {
			fmt.Printf("The division of %d by %d is %d", number1, number2, number1/number2)
		}

	default:
		fmt.Println("Invalid operation")
	}

}
