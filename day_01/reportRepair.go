package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func sumOfTwoEqualsTo(numbers []string, target int) int {
	for _, expense1 := range numbers {
		for _, expense2 := range numbers {
			exp1, _ := strconv.Atoi(expense1)
			exp2, _ := strconv.Atoi(expense2)
			if exp1+exp2 == target {
				return (exp1 * exp2)
			}
		}
	}
	return 0
}

func sumOfThreeEqualsTo(numbers []string, target int) int {
	for _, expense1 := range numbers {
		for _, expense2 := range numbers {
			for _, expense3 := range numbers {
				exp1, _ := strconv.Atoi(expense1)
				exp2, _ := strconv.Atoi(expense2)
				exp3, _ := strconv.Atoi(expense3)
				if exp1+exp2+exp3 == target {
					return (exp1 * exp2 * exp3)
				}
			}
		}
	}
	return 0
}

func main() {
	expenses, _ := readLines("/home/orfeas/code/advent-of-code-2020/day_01/input.txt")

	fmt.Println(sumOfTwoEqualsTo(expenses, 2020))
	fmt.Println(sumOfThreeEqualsTo(expenses, 2020))
}
