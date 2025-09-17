package main

import (
	"encoding/json"
	"os"
)

type Score struct {
	Attempts int     `json:"attempts"`
	Seconds  float64 `json:"seconds"`
}

func createFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.Write([]byte("{\"attempts\":-1, \"seconds\":-1.0}")); err != nil {
		return err
	}

	return nil
}

func rewriteFile(filename string, data []byte) error {
	if err := os.WriteFile(filename, data, 0755); err != nil {
		return err
	}
	return nil
}

func updateHighScore(filename string, attempts int, seconds float64) error {
	if _, err := os.Stat(filename); err != nil {
		if err := createFile(filename); err != nil {
			return err
		}
	}

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	var score Score
	if err := json.Unmarshal(data, &score); err != nil {
		return err
	}

	newScore := score
	if attempts < score.Attempts || score.Attempts == -1 {
		newScore.Attempts = attempts
	}

	if seconds < score.Seconds || score.Seconds < 0 {
		newScore.Seconds = seconds
	}

	if score.Attempts != newScore.Attempts || score.Seconds != newScore.Seconds {
		newScoreData, err := json.Marshal(newScore)
		if err != nil {
			return err
		}

		if err := rewriteFile(filename, newScoreData); err != nil {
			return err
		}
	}

	return nil
}

func getScore(filename string) (int, float64, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return 0, 0.0, err
	}

	var score Score
	if err := json.Unmarshal(data, &score); err != nil {
		return 0, 0.0, err
	}

	return score.Attempts, score.Seconds, nil
}
