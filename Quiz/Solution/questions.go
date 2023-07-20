package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
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

func askQuestions(questions [][]string) int {
	correctAns := 0

	scanner := bufio.NewReader(os.Stdin)

	selectedQuestions := randomQuestionGenerator(questions)

	for _, questions := range selectedQuestions {
		fmt.Println(questions[0])
		ans, err := scanner.ReadString('\n')
		if err != nil {
			fmt.Println("Error:", err)
		}

		ans = strings.TrimSpace(ans)

		if ans == questions[1] {
			correctAns++
		}
	}

	return correctAns
}
