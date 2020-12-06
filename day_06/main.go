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
	for _, group := range groups {
		persons := strings.Split(group, "\n")
		questionsCount := make(map[string]int)

		for _, person := range persons {
			for _, answer := range person {
				questionsCount[string(answer)] += 1
			}
		}
		noOfQuestionsAnsweredByAll := 0
		for k, _ := range questionsCount {
			if questionsCount[k] == len(persons) {
				noOfQuestionsAnsweredByAll++
			}
		}
		sumCounts += noOfQuestionsAnsweredByAll
	}
	return sumCounts
}

// readRaw returns the content of a text file as a string
func readRaw(filename string) string {
	content, err := ioutil.ReadFile(filename)
	check(err)
	return string(content)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
