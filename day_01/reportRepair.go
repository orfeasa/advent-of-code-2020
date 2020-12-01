package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

// ErrNumNotFound is returned if an input list has no two numbers that sum up to the target
var ErrNumNotFound = errors.New("cannot find two numbers from input that sum to the target")

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, num)
	}
	return lines, scanner.Err()
}

func sumOfTwoEqualsTo(numbers []int, target int) (int, error) {
	hasSeen := make(map[int]bool)

	for _, num := range numbers {
		need := target - num
		if hasSeen[need] {
			return num * need, nil
		}
		hasSeen[num] = true
	}
	return 0, ErrNumNotFound
}

func sumOfThreeEqualsTo(numbers []int, target int) (int, error) {
	for ind, num := range numbers {
		need := target - num
		mult, err := sumOfTwoEqualsTo(numbers[ind:], need)
		if err != nil {
			continue
		} else {
			return num * mult, nil
		}
	}
	return 0, ErrNumNotFound
}

func main() {
	expenses, _ := readLines("/home/orfeas/code/advent-of-code-2020/day_01/input.txt")

	res1, _ := sumOfTwoEqualsTo(expenses, 2020)
	fmt.Println(res1)
	res2, _ := sumOfThreeEqualsTo(expenses, 2020)
	fmt.Println(res2)
}
