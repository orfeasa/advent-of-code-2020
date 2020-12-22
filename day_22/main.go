package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
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
	pl1deck, pl2deck := processInput(inputPath)
	winningDeck := playCombat(pl1deck, pl2deck)
	acc := 0
	for ind, val := range winningDeck {
		acc += (len(winningDeck) - ind) * val
	}
	return acc
}

func part2(inputPath string) int {
	pl1deck, pl2deck := processInput(inputPath)
	winningDeck, _ := playRecursiveCombat(pl1deck, pl2deck)
	acc := 0
	for ind, val := range winningDeck {
		acc += (len(winningDeck) - ind) * val
	}
	return acc
}

// 1 game consists of many rounds
// a round can spark the creation of a sub-game

func playRecursiveCombat(deck1, deck2 []int) (winningDeck []int, winner int) {
	gameDecksHistory := make([][][]int, 0)
	for len(deck1) != 0 && len(deck2) != 0 {
		for _, item := range gameDecksHistory {
			if reflect.DeepEqual(item, [][]int{deck1, deck2}) {
				return deck1, 1
			}
		}
		gameDecksHistory = append(gameDecksHistory, [][]int{deck1, deck2})

		card1 := deck1[0]
		deck1 = deck1[1:]
		card2 := deck2[0]
		deck2 = deck2[1:]

		if len(deck1) >= card1 && len(deck2) >= card2 {
			// recursive combat round initiation
			deck1copy := make([]int, card1)
			deck2copy := make([]int, card2)
			copy(deck1copy, deck1[:card1])
			copy(deck2copy, deck2[:card2])
			_, winner := playRecursiveCombat(deck1copy, deck2copy)
			if winner == 1 {
				deck1 = append(deck1, card1, card2)
			} else {
				deck2 = append(deck2, card2, card1)
			}
		} else {
			if card1 > card2 {
				deck1 = append(deck1, card1, card2)
			} else {
				deck2 = append(deck2, card2, card1)
			}
		}
	}
	if len(deck1) != 0 {
		return deck1, 1
	}
	return deck2, 2
}

func playCombat(deck1, deck2 []int) (winner []int) {
	for len(deck1) != 0 && len(deck2) != 0 {
		card1 := deck1[0]
		deck1 = deck1[1:]
		card2 := deck2[0]
		deck2 = deck2[1:]
		if card1 > card2 {
			deck1 = append(deck1, card1, card2)
		} else {
			deck2 = append(deck2, card2, card1)
		}
	}
	if len(deck1) != 0 {
		return deck1
	}
	return deck2
}

func stringToIntSlice(input []string) []int {
	var output = make([]int, 0, len(input))
	for _, val := range input {
		output = append(output, toInt(val))
	}
	return output
}

func processInput(inputPath string) (deck1, deck2 []int) {
	inputRaw := readRaw(inputPath)
	playersRaw := strings.Split(inputRaw, "\n\n")
	deck1 = stringToIntSlice(strings.Split(playersRaw[0], "\n")[1:])
	deck2 = stringToIntSlice(strings.Split(playersRaw[1], "\n")[1:])
	return deck1, deck2
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
