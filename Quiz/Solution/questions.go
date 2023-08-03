package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func randomQuestionGenerator(questions [][]string) [][]string {
	if totalQuestions > len(questions) {
		fmt.Println("Error: totalQuestions exceeds the number of available questions.")
		return nil
	}

	rand.Seed(time.Now().UnixNano()) // Initialize random seed

	// Shuffle the indices using the Fisher-Yates algorithm
	indices := rand.Perm(len(questions))

	// Select the first 'totalQuestions' indices
	selectedIndices := indices[:totalQuestions]

	// Collect the randomly selected questions
	selectedQuestions := make([][]string, totalQuestions)
	for i, idx := range selectedIndices {
		selectedQuestions[i] = questions[idx]
	}

	return selectedQuestions
}

func askQuestions(questions [][]string, timerValue float64) int {

	correctAns := 0

	selectedQuestions := randomQuestionGenerator(questions)

	for _, question := range selectedQuestions {
		correct, err := askQuestion(question, timerValue)
		if err != nil {
			fmt.Println(err)
		}
		correctAns += correct
	}

	return correctAns
}

func askQuestion(question []string, timerValue float64) (int, error) {

	timer := time.NewTimer(time.Duration(timerValue) * time.Second)
	done := make(chan string)

	go getInput(done)

	fmt.Println(question[0])
	for {
		select {
		case <-timer.C:
			return 0, fmt.Errorf("Time out")
		case ans := <-done:
			score := 0
			if strings.TrimSpace(strings.ToLower(ans)) == question[1] {
				score = 1
			} else {
				return 0, nil
			}

			return score, nil
		}
	}
}
