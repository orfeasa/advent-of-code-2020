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
	finalSeats := runRulesUntilEquilibrium(textLayout, adjacentSeatVisibilityRule)
	return countAllOccupied(finalSeats)
}

func part2(inputPath string) int {
	textLayout := readStrings(inputPath)
	finalSeats := runRulesUntilEquilibrium(textLayout, directionalSeatVisibilityRule)
	return countAllOccupied(finalSeats)
}

type ruleset func(rowNum, colNum int, seats [][]byte) byte

func runRulesUntilEquilibrium(textLayout []string, rules ruleset) [][]byte {
	var current, next [][]byte
	for _, line := range textLayout {
		current = append(current, []byte(line))
		next = append(next, []byte(line))
	}

	for {
		for rowNum, row := range current {
			for colNum, _ := range row {
				next[rowNum][colNum] = rules(rowNum, colNum, current)
			}
		}
		current, next = next, current
		if reflect.DeepEqual(current, next) {
			break
		}
	}
	return current
}

func adjacentSeatVisibilityRule(rowNum, colNum int, seats [][]byte) byte {
	seat := seats[rowNum][colNum]
	numOfAdjacentOccupied := countAdjacentOccupied(rowNum, colNum, seats)
	if isEmpty(seat) && numOfAdjacentOccupied == 0 {
		// seat becomes occupied
		return '#'
	} else if isOccupied(seat) && numOfAdjacentOccupied >= 4 {
		// seat becomes empty
		return 'L'
	}
	return seat
}

func directionalSeatVisibilityRule(rowNum, colNum int, seats [][]byte) byte {
	seat := seats[rowNum][colNum]
	numOfVisibleOccupied := countVisibleOccupied(rowNum, colNum, seats)
	if isEmpty(seat) && numOfVisibleOccupied == 0 {
		// seat becomes occupied
		return '#'
	} else if isOccupied(seat) && numOfVisibleOccupied >= 5 {
		// seat becomes empty
		return 'L'
	}
	return seat
}

func countVisibleOccupied(y, x int, seats [][]byte) (count int) {
	maxY := len(seats) - 1
	maxX := len(seats[0]) - 1

	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if !(dx == 0 && dy == 0) {
				currX := x
				currY := y
				for {
					currX += dx
					currY += dy
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

func countAdjacentOccupied(y, x int, seats [][]byte) (count int) {
	maxY := len(seats) - 1
	maxX := len(seats[0]) - 1

	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if !(dx == 0 && dy == 0) {
				newX := x + dx
				newY := y + dy
				if (newX >= 0 && newX <= maxX) && (newY >= 0 && newY <= maxY) {
					if isOccupied(seats[newY][newX]) {
						count++
					}
				}
			}
		}
	}
	return count
}

func countAllOccupied(seats [][]byte) (count int) {
	for _, row := range seats {
		for _, seat := range row {
			if isOccupied(seat) {
				count++
			}
		}
	}
	return count
}

func isEmpty(seat byte) bool { return seat == 'L' }

func isOccupied(seat byte) bool { return seat == '#' }

func isFloor(seat byte) bool { return seat == '.' }

func readStrings(filename string) []string {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var text []string
	for scanner.Scan() {
		text = append(text, strings.TrimRight(scanner.Text(), "\n"))
	}
	return text
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
