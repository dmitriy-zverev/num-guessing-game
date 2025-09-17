package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	difficulty := initGame()
	startGameLoop(difficulty)
}

func initGame() int {
	fmt.Print(HELLO_MESSAGE)
	fmt.Print(CHOOSE_DIFFICULTY_MESSAGE)

	fmt.Print("\nEnter your choice: ")

	var difficulty int
	fmt.Scan(&difficulty)

	difficultyMessage, err := difficultyToString(difficulty)
	if err != nil {
		fmt.Printf("\nError: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\nGreat! You have selected the %s difficulty level.\nLet's start the game!\n", difficultyMessage)

	return difficulty
}

func startGameLoop(difficulty int) {
	randNum := rand.Intn(MAX_RAND_NUMBER) + 1
	chances := difficultyToChances(difficulty)

	var userNum int
	for i := 0; i < chances; i++ {
		fmt.Printf("\n%d. Enter your guess: ", i+1)
		fmt.Scan(&userNum)

		if err := validateInputGuess(userNum); err != nil {
			fmt.Printf("\nError: %v\n", err)
			i--
		}

		if checkIfWon(userNum, randNum, i) {
			return
		}
	}

	fmt.Printf("\nYou have lost. The correct number was %d\n", randNum)
	endGame()
}

func endGame() {
	fmt.Print("\nWant to try again?\n1. Yes\n2. No\n\nStart again: ")

	var input int
	fmt.Scan(&input)
	if input == 1 {
		fmt.Println()
		difficulty := initGame()
		startGameLoop(difficulty)
	}

	fmt.Print("\nOk. It was nice to see you. Bye.")
	os.Exit(0)
}

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

func difficultyToChances(difficulty int) int {
	switch difficulty {
	case 1:
		return 10
	case 2:
		return 5
	case 3:
		return 3
	default:
		return -1
	}
}

func difficultyToString(difficulty int) (string, error) {
	switch difficulty {
	case 1:
		return "Easy", nil
	case 2:
		return "Medium", nil
	case 3:
		return "Hard", nil
	default:
		return "", errors.New("please choose correct difficulty")
	}
}
