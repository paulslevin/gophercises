package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"

	"./types"
)

func main() {

	csvFileName := flag.String(
		"csv",
		"problems.csv",
		"a csv file in the format of 'question,answer'",
	)

	timeLimit := flag.Int(
		"limit",
		30,
		"the time limit for the quiz in seconds",
	)

	flag.Parse()

	game := types.NewGame(*timeLimit)

	csvFile, err := os.Open(*csvFileName)

	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFileName))
	}

	r := csv.NewReader(csvFile)

	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file")
	}

	game.ParseCSVLines(lines)

}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
