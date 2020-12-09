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
	xmasData := readNumbers(inputPath)
	preamble := 25

	for ind, val := range xmasData {
		if ind >= preamble {
			if !isValidXmasNumber(val, xmasData[ind-preamble:ind]) {
				return val
			}
		}
	}

	return 0
}

func part2(inputPath string) int {
	xmasData := readNumbers(inputPath)
	subset := continuousSubsetSum(xmasData, part1(inputPath))
	return min(subset) + max(subset)
}

func continuousSubsetSum(numbers []int, target int) []int {
	// holds the current set
	var currentNumbers []int
	// low, high point to the index in numbers of the lowest and highest value in currentNumbers
	low := 0
	high := 0

	// accumulator and currentNumbers initialized to hold the first value of numbers
	acc := numbers[high]
	currentNumbers = append(currentNumbers, numbers[high])

	for {
		for acc < target {
			// keep adding to current numbers until currentSum >= target
			high++
			// check for out of bounds
			if high >= len(numbers) {
				return nil
			}
			acc += numbers[high]
			currentNumbers = append(currentNumbers, numbers[high])
		}

		if acc == target {
			return numbers[low : high+1]
		}

		for acc > target {
			// keep removing elements from the beginning until currentSum =< target
			acc -= numbers[low]
			currentNumbers = currentNumbers[1:]
			low++
			// TODO check if low == high
		}
	}
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

// Max returns the larger of x or y.
func max(numbers []int) int {
	currMax := numbers[0]
	for _, val := range numbers {
		if val > currMax {
			currMax = val
		}
	}
	return currMax
}

// Max returns the larger of x or y.
func min(numbers []int) int {
	currMin := numbers[0]
	for _, val := range numbers {
		if val < currMin {
			currMin = val
		}
	}
	return currMin
}
