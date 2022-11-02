package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

// Main Function
func main() {
	// Defining which csv to open
	csvFilename := ("problems.csv")

	//  Handling csv opening & errors
	file, err := os.Open(csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", csvFilename))
	}
	// Creating the csv reader object & read errors
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}

	// Defining the parsed list as a struct of the parsed lines
	problems := parseLines(lines)

	// Iterating through the questions while comparing questions to answer and adding
	// to total score count
	correct := 0
	for i, p := range problems {
		// Printing the question followed by its number and its content
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		// Defining answer string
		var answer string
		// Waiting for user input and binding it to answer
		fmt.Scanf("%s\n", &answer)
		// Comparing database answer to input answer
		if answer == p.a {
			// if they match, add 1 to score counter
			correct++
		}
		// otherwise, keep looping
	}
	// When Looping/quiz done, display the number of correct answers over the total
	// number of questions
	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

// Parse Lines into a 2dimensional slice of strings and return a problem slice made of problem structs
func parseLines(lines [][]string) []problem {
	// Intialize the slice with make, and make it have problem structure, and as many as the number of lines
	// in our input object
	ret := make([]problem, len(lines))
	// for each line in lines, add the question and answer to the returned slice
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	// This is what problems is equal to
	return ret
}

// Defining problem struct as having a question and an answer
type problem struct {
	q string
	a string
}

// Handling Exits nicely
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
