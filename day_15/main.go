package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "1,0,16,5,17,4"
	fmt.Println("--- Part One ---")
	fmt.Println(part1(input))

	fmt.Println("--- Part Two ---")
	fmt.Println(part2(input))
}

func part1(input string) int {
	return playGameUntil(2020, input)
}

func part2(input string) int {
	return playGameUntil(30000000, input)
}

func playGameUntil(limit int, input string) int {
	inputStrings := strings.Split(input, ",")
	// numbers[k] = t means number k was spoken last at time t
	numbers := make(map[int]int)
	for ind, number := range inputStrings {
		numbers[toInt(number)] = ind
	}

	last := 0
	difference := 0
	isFirstTime := true

	for clock := len(numbers); clock < limit; clock++ {
		if isFirstTime {
			last = 0
			// if 0 has been spoken before
			if _, ok := numbers[0]; ok {
				isFirstTime = false
			} else {
				isFirstTime = true
			}
			difference = clock - numbers[0]
			numbers[0] = clock
		} else {
			last = difference
			// check if last has been spoken before
			if val, ok := numbers[last]; ok {
				isFirstTime = false
				difference = clock - val
			} else {
				isFirstTime = true
			}
			numbers[last] = clock
		}
	}
	return last
}

func toInt(s string) int {
	result, err := strconv.Atoi(s)
	check(err)
	return result
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
