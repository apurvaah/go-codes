package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func readCSV(file string) [][]string {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer f.Close()

	// Create a new CSV reader
	reader := csv.NewReader(f)

	// Read all the CSV records
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return records
}
