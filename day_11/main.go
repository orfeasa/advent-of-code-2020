package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputPath := "./day_11/test_input.txt"
	fmt.Println("--- Part One ---")
	fmt.Println(part1(inputPath))

	fmt.Println("--- Part Two ---")
	fmt.Println(part2(inputPath))
}

func part1(inputPath string) int {
	textLayout := readStrings(inputPath)
	for _, val := range textLayout {
		fmt.Println(val)
	}
	fmt.Println()

	intLayout := convertLayoutToInts(textLayout)
	for _, val := range intLayout {
		fmt.Println(val)
	}
	fmt.Println()

	prevSeats := intLayout
	currSeats := runRoundOfRules(prevSeats)

	for reflect.DeepEqual(prevSeats, currSeats) {
		currSeats := runRoundOfRules(prevSeats)
		prevSeats := currSeats
	}

	return countOccupied(currSeats)
}

func part2(inputPath string) int {
	return 0
}

func countOccupied(seats) int {

}

func runRoundOfRules(seats) seats {
	for indRow,
}

func isEmpty(seat int) bool {
	return seat ==0
}

func isOccupied(seat int) bool {
	return seat == 1
}

func hasAtLeast4AdjacentOccupied(x, y, int, seats) bool {

}

func hasOccupiedSeatsAdjacent(x, y int, seats) bool {

}

func convertLayoutToInts(textLayout []string) [][]int {
	var intLayout [][]int

	for _, val1 := range textLayout {
		var newRow []int
		for _, val2 := range val1 {
			switch string(val2) {
			case "L":
				// empty seat is 0
				newRow = append(newRow, 0)
			case "#":
				// occupied seat is 1
				newRow = append(newRow, 1)
			case ".":
				// floor is -1
				newRow = append(newRow, -1)
			}
		}
		intLayout = append(intLayout, newRow)
	}
	return intLayout

}

func readStrings(filename string) []string {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	Scanner := bufio.NewScanner(file)

	var text []string
	for Scanner.Scan() {
		text = append(text, strings.TrimRight(Scanner.Text(), "\n"))
	}
	return text
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
