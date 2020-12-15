package main

import (
	"fmt"
)

func main() {
	input := []int{1, 0, 16, 5, 17, 4}
	fmt.Println("--- Part One ---")
	fmt.Println(part1(input))

	fmt.Println("--- Part Two ---")
	fmt.Println(part2(input))
}

func part1(input []int) int {
	return playGameUntil(2020, input)
}

func part2(input []int) int {
	return playGameUntil(30000000, input)
}

func playGameUntil(limit int, input []int) int {
	numbers := make(map[int]int)

	for ind, num := range input {
		numbers[num] = ind
	}

	last := 0
	difference := 0
	isFirstTime := true
	for clock := len(numbers); clock < limit; clock++ {
		if isFirstTime {
			last = 0
			// if 0 has been spoken before
			_, ok := numbers[0]
			isFirstTime = !ok
			difference = clock - numbers[0]
			numbers[0] = clock
		} else {
			last = difference
			// check if last has been spoken before
			_, ok := numbers[last]
			isFirstTime = !ok
			if !isFirstTime {
				difference = clock - numbers[last]
			}
			numbers[last] = clock
		}
	}
	return last
}
