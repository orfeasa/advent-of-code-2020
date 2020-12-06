package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	inputPath := "./day_06/input.txt"
	fmt.Println("--- Part One ---")
	fmt.Println(Part1(inputPath))

	fmt.Println("--- Part Two ---")
	fmt.Println(Part2(inputPath))
}

func Part1(inputPath string) int {
	allAnswers := readRaw(inputPath)

	groups := strings.Split(allAnswers, "\n\n")
	sumCounts := 0
	for _, group := range groups {
		persons := strings.Split(group, "\n")

		questions := make(map[string]bool)

		for _, person := range persons {
			for _, answer := range person {
				questions[string(answer)] = true
			}
		}
		sumCounts += len(questions)
	}
	return sumCounts
}

func Part2(inputPath string) int {
	allAnswers := readRaw(inputPath)

	groups := strings.Split(allAnswers, "\n\n")

	sumCounts := 0
	// iterate over each group
	for _, group := range groups {
		personsInGroup := strings.Split(group, "\n")
		commonAnswers := make(map[string]bool)

		// initialize map with 1st person's answers
		for _, answer := range personsInGroup[0] {
			commonAnswers[string(answer)] = true
		}

		// iterate over each person's answers
		for _, personAnswers := range personsInGroup {
			// iterate over the answers in each question
			for answer, _ := range commonAnswers {
				if !strings.Contains(personAnswers, answer) {
					delete(commonAnswers, answer)
				}
			}
		}
		sumCounts += len(commonAnswers)
	}

	return sumCounts
}

// readRaw returns the content of a text file as a string
func readRaw(filename string) string {
	content, err := ioutil.ReadFile(filename)
	check(err)
	return strings.TrimRight(string(content), "\n")
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
