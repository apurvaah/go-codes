package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {
	// Define command-line flags with default values
	timer := flag.Float64("timer", 30.0, "Timer for each question")
	filepath := flag.String("filepath", "../Data/problems.csv", "Quiz File Path")

	// Parse the command-line flags
	flag.Parse()

	// Access the values of the flags
	timerValue := *timer
	filepathValue := *filepath

	questions := readCSV(filepathValue)

	correctAns := askQuestions(questions, timerValue)

	fmt.Println("Correctly answered: " + strconv.Itoa(correctAns) + "/" + strconv.Itoa(totalQuestions))
}
