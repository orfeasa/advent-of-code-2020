package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	inputPath := "./day_10/input.txt"
	fmt.Println("--- Part One ---")
	fmt.Println(part1(inputPath))

	fmt.Println("--- Part Two ---")
	fmt.Println(part2(inputPath))
}

func part1(inputPath string) int {
	joltSteps := readNumbers(inputPath)
	sort.Ints(joltSteps)

	countdifferences := make(map[int]int)
	prev := 0
	for _, adapter := range joltSteps {
		countdifferences[adapter-prev]++
		prev = adapter
	}
	// device's built-in adapter is always 3 higher than the highest adapter
	countdifferences[3]++
	return countdifferences[1] * countdifferences[3]
}

func part2(inputPath string) int {
	joltSteps := readNumbers(inputPath)
	sort.Ints(joltSteps)
	// add first step as 0. Last step is not required as it's always +3 (1 way)
	joltSteps = append([]int{0}, joltSteps...)

	return calcStepsToReach(len(joltSteps)-1, joltSteps)
}

// calcStepsToReach counts the ways to reach the step at index target in steps[]
func calcStepsToReach(target int, steps []int) int {
	memo := make(map[int]int)
	return calcStepsToReachMemo(target, steps, memo)
}

func calcStepsToReachMemo(target int, steps []int, memo map[int]int) int {
	// base case, 1 way to reach the beginning
	if target == 0 {
		memo[target] = 1
		return memo[target]
	}
	// if it's already calculated return it
	if val, ok := memo[target]; ok {
		return val
	}

	// the ways you can reach a step are the sum of the ways
	// you can reach the previous 3, if they exist
	result := 0
	for i := 1; i <= 3; i++ {
		// if not out of range, and the difference is no more than 3
		if target-i >= 0 && steps[target-i] >= steps[target]-3 {
			// call the function for the previous steps
			steps := calcStepsToReachMemo(target-i, steps, memo)
			// store the value
			memo[target-i] = steps
			result += steps
		}
	}
	return result
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
