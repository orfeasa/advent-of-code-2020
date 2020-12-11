package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

func main() {
	inputPath := "./day_11/input.txt"
	fmt.Println("--- Part One ---")
	fmt.Println(part1(inputPath))

	fmt.Println("--- Part Two ---")
	fmt.Println(part2(inputPath))
}

func part1(inputPath string) int {
	textLayout := readStrings(inputPath)
	currSeats := convertLayoutToInts(textLayout)
	finalSeats := runRulesUntilEquilibrium(currSeats, adjacentSeatVisibilityRule)

	return countAllOccupied(finalSeats)
}

func part2(inputPath string) int {
	textLayout := readStrings(inputPath)
	currSeats := convertLayoutToInts(textLayout)
	finalSeats := runRulesUntilEquilibrium(currSeats, directionalSeatVisibilityRule)

	return countAllOccupied(finalSeats)
}

type ruleset func(rowNum, colNum int, seats [][]int) int

func runRulesUntilEquilibrium(currSeats [][]int, rules ruleset) [][]int {
	prevSeats := make([][]int, len(currSeats))

	for {
		// prevSeats = currSeats
		for i := range prevSeats {
			prevSeats[i] = make([]int, len(currSeats[i]))
			copy(prevSeats[i], currSeats[i])
		}
		currSeats = runRoundOfRules(prevSeats, rules)
		if reflect.DeepEqual(prevSeats, currSeats) {
			break
		}
	}
	return currSeats
}

func runRoundOfRules(seats [][]int, rules ruleset) [][]int {
	newSeats := make([][]int, len(seats))
	for i := range seats {
		newSeats[i] = make([]int, len(seats[i]))
		copy(newSeats[i], seats[i])
	}

	for rowNum, row := range seats {
		for colNum, _ := range row {
			newSeats[rowNum][colNum] = rules(rowNum, colNum, seats)
		}
	}
	return newSeats
}

func adjacentSeatVisibilityRule(rowNum, colNum int, seats [][]int) int {
	seat := seats[rowNum][colNum]
	numOfAdjacentOccupied := countAdjacentOccupied(rowNum, colNum, seats)
	if isEmpty(seat) && numOfAdjacentOccupied == 0 {
		// seat becomes occupied
		return 1
	} else if isOccupied(seat) && numOfAdjacentOccupied >= 4 {
		// seat becomes empty
		return 0
	}
	return seat
}

func directionalSeatVisibilityRule(rowNum, colNum int, seats [][]int) int {
	seat := seats[rowNum][colNum]
	numOfVisibleOccupied := countVisibleOccupied(rowNum, colNum, seats)
	if isEmpty(seat) && numOfVisibleOccupied == 0 {
		// seat becomes occupied
		return 1
	} else if isOccupied(seat) && numOfVisibleOccupied >= 5 {
		// seat becomes empty
		return 0
	}
	return seat
}

func countVisibleOccupied(y, x int, seats [][]int) int {
	count := 0

	maxY := len(seats) - 1
	maxX := len(seats[0]) - 1

	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			// should not be the actual
			if !(dx == 0 && dy == 0) {
				currX := x
				currY := y
				for {
					currX += dx
					currY += dy
					// out of bounds
					if currX > maxX || currY > maxY || currX < 0 || currY < 0 {
						break
					}
					if isOccupied(seats[currY][currX]) {
						count++
						break
					} else if isEmpty(seats[currY][currX]) {
						break
					}
				}
			}
		}
	}
	return count
}

func countAdjacentOccupied(y, x int, seats [][]int) int {
	adjacentSeats := getValidAdjacentSeats(y, x, seats)
	count := 0
	for _, seat := range adjacentSeats {
		if isOccupied(seat) {
			count++
		}
	}
	return count
}

func getValidAdjacentSeats(y, x int, seats [][]int) []int {
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

// printSeats is used for debug purposes
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
