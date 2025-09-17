package main

import "fmt"

func main() {
	difficulty := initGame()
	if err := startGameLoop(difficulty); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
