package main

import "errors"

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
