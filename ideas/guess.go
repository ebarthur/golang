package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rules := "The System will select a number between 0 to 10, and you have to guess what number it is. You have 3 chances"
	fmt.Println(rules)

	repeat()
}

func repeat() {
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(11)
	guessCount := 0
	guessLimit := 3

	for guessCount < guessLimit {
		var guess int
		fmt.Print("Guess: ")
		fmt.Scan(&guess)
		guessCount++

		if guess == secretNumber {
			fmt.Println("Congrats. You Won!")
			var tryAgain string
			fmt.Print("Would you like to try again? 'Y' or 'N': ")
			fmt.Scan(&tryAgain)
			tryAgain = strings.ToLower(tryAgain)

			if tryAgain == "y" {
				repeat()
			} else if tryAgain == "n" {
				var toExit string
				fmt.Print("Enter 'E' to exit: ")
				fmt.Scan(&toExit)
				toExit = strings.ToLower(toExit)
				if toExit == "e" {
					return
				}
			}
		} else if guessCount < guessLimit && guess != secretNumber {
			fmt.Println("Sorry. Try again.")
		} else if guessCount == guessLimit && guess != secretNumber {
			fmt.Println("Sorry, You Failed")
			var tryAgain string
			fmt.Print("Would you like to try again? 'Y' or 'N': ")
			fmt.Scan(&tryAgain)
			tryAgain = strings.ToLower(tryAgain)

			if tryAgain == "y" {
				repeat()
			} else if tryAgain == "n" {
				var toExit string
				fmt.Print("Enter 'E' to exit: ")
				fmt.Scan(&toExit)
				toExit = strings.ToLower(toExit)
				if toExit == "e" {
					return
				}
			}
		}
	}
}
