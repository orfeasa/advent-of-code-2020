package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputPath := "./day_09/input.txt"
	fmt.Println("--- Part One ---")
	fmt.Println(part1(inputPath))

	fmt.Println("--- Part Two ---")
	fmt.Println(part2(inputPath))
}

func part1(inputPath string) int {
	xmas_data := readNumbers(inputPath)
	preamble := 25

	for ind, val := range xmas_data {
		if ind >= preamble {
			if !isValidXmasNumber(val, xmas_data[ind-preamble:ind]) {
				return val
			}
		}
	}

	return 0
}

func part2(inputPath string) int {
	return 0
}

// isValidXmasNumber checks if the last number in the slice is
// equal to the sum of any two numbers in that slice
func isValidXmasNumber(target int, rangeToCheck []int) bool {

	// iterate over slice
	for ind1, val1 := range rangeToCheck {

		// check if target - val1 is in the rest of the slice
		for _, val2 := range rangeToCheck[ind1+1:] {
			if val1+val2 == target {
				return true
			}
		}
	}
	return false
}

func readNumbers(filename string) []int {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	Scanner := bufio.NewScanner(file)

	var numbers []int
	for Scanner.Scan() {
		numbers = append(numbers, toInt(Scanner.Text()))
	}
	return numbers
}

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

func toInt(s string) int {
	result, err := strconv.Atoi(s)
	check(err)
	return result
}
