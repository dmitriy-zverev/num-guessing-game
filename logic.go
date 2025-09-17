package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func initGame() int {
	fmt.Print(HELLO_MESSAGE)
	if err := printTopScores(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

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

func printTopScores() error {
	attempts, seconds, err := getScore(STORAGE_FILENAME)
	if err != nil {
		return err
	}

	if attempts > 0 && seconds > 0 {
		fmt.Printf("Your best scores are:\n\tAttempts: %d\n\tSeconds: %.2f\n\n", attempts, seconds)
	}

	return nil
}

func startGameLoop(difficulty int) error {
	randNum := rand.Intn(MAX_RAND_NUMBER) + 1
	chances := difficultyToChances(difficulty)
	hintGiven := false

	var userNum int

	start := time.Now()
	for i := 0; i < chances; i++ {
		fmt.Printf("\n%d. Enter your guess: ", i+1)
		fmt.Scan(&userNum)

		if err := validateInputGuess(userNum); err != nil {
			fmt.Printf("\nError: %v\n", err)
			i--
		}

		if checkIfWon(userNum, randNum, i) {
			seconds := printElapsedTime(start)
			if err := updateHighScore(STORAGE_FILENAME, i+1, seconds); err != nil {
				return err
			}
			return nil
		}

		if !hintGiven {
			hintGiven = giveHint(randNum, difficulty, i+1)
		}
	}

	fmt.Printf("\nYou have lost. The correct number was %d\n", randNum)
	endGame()
	return nil
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

	fmt.Print("\nOk. It was nice to see you. Bye.\n")
	os.Exit(0)
}

func printElapsedTime(start time.Time) float64 {
	elapsed := time.Since(start)
	fmt.Printf("\nYou have guessed in %.2f seconds\n", elapsed.Seconds())
	return elapsed.Seconds()
}

func giveHint(correctNum, difficulty, attempts int) bool {
	if difficulty < 3 && attempts > 3 {
		switch rand.Intn(2) {
		case 0:
			fmt.Printf("Hint: the number ends with %d\n", correctNum%10)
			return true
		case 1:
			if correctNum == 100 {
				fmt.Printf("Hint: the number starts with 1\n")
			}
			fmt.Printf("Hint: the number starts with %d\n", (correctNum-correctNum%10)/10)
			return true
		}
	}

	return false
}
