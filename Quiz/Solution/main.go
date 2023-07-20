package main

import (
	"fmt"
	"strconv"
)

func main() {
	questions := readCSV(filepath)

	correctAns := askQuestions(questions)

	fmt.Println("Correctly answered: " + strconv.Itoa(correctAns) + "/" + strconv.Itoa(totalQuestions))
}
