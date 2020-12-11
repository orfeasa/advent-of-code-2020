package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
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
	currSeats := convertLayoutToInts(textLayout)
	prevSeats := make([][]int, len(currSeats))
	fmt.Println("Initial state")
	printSeats(currSeats)

	for {
		// prevSeats = currSeats
		for i := range prevSeats {
			prevSeats[i] = make([]int, len(currSeats[i]))
			copy(prevSeats[i], currSeats[i])
		}
		fmt.Println(2)
		currSeats = runRoundOfRules(prevSeats)
		printSeats(prevSeats)
		printSeats(currSeats)
		if reflect.DeepEqual(prevSeats, currSeats) {
			break
		}
		fmt.Scanln()
	}
	printSeats(currSeats)
	return countAllOccupied(currSeats)
}

func part2(inputPath string) int {
	return 0
}

func runRoundOfRules(seats [][]int) [][]int {
	newSeats := make([][]int, len(seats))
	for i := range seats {
		newSeats[i] = make([]int, len(seats[i]))
		copy(newSeats[i], seats[i])
	}

	for rowNum, row := range seats {
		for colNum, seat := range row {
			if isEmpty(seat) && hasNoOccupiedSeatsAdjacent(rowNum, colNum, seats) {
				// seat becomes occupied
				newSeats[rowNum][colNum] = 1
			} else if isOccupied(seat) && hasAtLeast4AdjacentOccupied(rowNum, colNum, seats) {
				// seat becomes empty
				newSeats[rowNum][colNum] = 0
			}
		}
	}
	return newSeats
}

func hasAtLeast4AdjacentOccupied(x, y int, seats [][]int) bool {
	adjacentSeats := getValidAdjacentSeats(x, y, seats)
	count := 0
	for _, seat := range adjacentSeats {
		if isOccupied(seat) {
			count++
		}
	}
	return count >= 4
}

func hasNoOccupiedSeatsAdjacent(x, y int, seats [][]int) bool {
	adjacentSeats := getValidAdjacentSeats(x, y, seats)
	count := 0
	for _, seat := range adjacentSeats {
		if isOccupied(seat) {
			count++
		}
	}
	return count == 0
}

func getValidAdjacentSeats(x, y int, seats [][]int) []int {
	maxY := len(seats) - 1
	maxX := len(seats[0]) - 1
	var validSeats []int

	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			// should not be the actual
			if !(dx == 0 && dy == 0) {
				newX := x + dx
				newY := y + dy
				// if they are not out of bounds
				if (newX >= 0 && newX <= maxX) && (newY >= 0 && newY <= maxY) {
					validSeats = append(validSeats, seats[newY][newX])
				}
			}
		}
	}
	return validSeats
}

func countAllOccupied(seats [][]int) int {
	count := 0
	for _, row := range seats {
		for _, seat := range row {
			if isOccupied(seat) {
				count++
			}
		}
	}
	return count
}

func isEmpty(seat int) bool {
	return seat == 0
}

func isOccupied(seat int) bool {
	return seat == 1
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

func printSeats(seats [][]int) {
	for ind1, val1 := range seats {
		for ind2 := range val1 {
			switch seats[ind1][ind2] {
			case 0:
				// empty seat is 0
				fmt.Print("L")
			case 1:
				// occupied seat is 1
				fmt.Print("#")
			case -1:
				// floor is -1
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println()

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
