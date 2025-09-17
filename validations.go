package main

import (
	"errors"
	"fmt"
)

func checkIfWon(userInput, correctNum, attempt int) bool {
	if userInput == correctNum {
		fmt.Printf("\nCongratulations! You guessed the correct number in %d attempts\n", attempt+1)
		return true
	}

	fmt.Print("Incorrect! The number is ")
	if userInput < correctNum {
		fmt.Printf("greater than %d\n", userInput)
	} else {
		fmt.Printf("less than %d\n", userInput)
	}

	return false
}

func validateInputGuess(input int) error {
	if input > 0 && input < 101 {
		return nil
	}

	return errors.New("invalid input, try again")
}
