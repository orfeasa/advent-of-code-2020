package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	inputPath := "./day_22/input.txt"
	fmt.Println("--- Part One ---")
	fmt.Println(part1(inputPath))

	fmt.Println("--- Part Two ---")
	fmt.Println(part2(inputPath))
}

func part1(inputPath string) int {
	inputRaw := readRaw(inputPath)
	playersRaw := strings.Split(inputRaw, "\n\n")

	pl1deck := stringToIntSlice(strings.Split(playersRaw[0], "\n")[1:])
	pl2deck := stringToIntSlice(strings.Split(playersRaw[1], "\n")[1:])

	for len(pl1deck) != 0 && len(pl2deck) != 0 {
		pl1card := pl1deck[0]
		pl1deck = pl1deck[1:]
		pl2card := pl2deck[0]
		pl2deck = pl2deck[1:]

		if pl1card > pl2card {
			pl1deck = append(pl1deck, pl1card, pl2card)
		} else {
			pl2deck = append(pl2deck, pl2card, pl1card)
		}
	}
	var winningDeck []int
	if len(pl1deck) != 0 {
		winningDeck = pl1deck
	} else {
		winningDeck = pl2deck
	}

	acc := 0
	for ind, val := range winningDeck {
		acc += (len(winningDeck) - ind) * val
	}

	return acc
}

func part2(inputPath string) int {
	return 0
}
func stringToIntSlice(input []string) []int {
	var output = make([]int, 0, len(input))
	for _, val := range input {
		output = append(output, toInt(val))
	}
	return output
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

func max(numbers []int) int {
	currMax := numbers[0]
	for _, val := range numbers {
		if val > currMax {
			currMax = val
		}
	}
	return currMax
}

func min(numbers []int) int {
	currMin := numbers[0]
	for _, val := range numbers {
		if val < currMin {
			currMin = val
		}
	}
	return currMin
}
