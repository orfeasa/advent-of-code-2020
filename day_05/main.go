package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputPath := "./day_05/input.txt"
	lines := readLines(inputPath)

	var seatList []int
	maxSeatID := 0
	for _, line := range lines {
		currentSeat := calculateSeatID(line)
		seatList = append(seatList, currentSeat)
		maxSeatID = max(maxSeatID, currentSeat)
	}
	fmt.Println(maxSeatID)
	fmt.Println(findMissingSeat(seatList))
}

func findMissingSeat(seatList []int) int {

	// to find the missing seat we will add up all the seats
	// and find the min value. The missing value is equal to the sum
	// of what the value would have been if all values were present
	// minus the current one
	minSeat := seatList[0]
	maxSeat := seatList[0]
	sum := 0
	for _, seat := range seatList {
		minSeat = min(minSeat, seat)
		maxSeat = max(maxSeat, seat)
		sum += seat
	}

	return calculateContinuousSum(minSeat, maxSeat) - sum
}

func calculateSeatID(boardingPass string) int {
	boardingPass = strings.ReplaceAll(boardingPass, "F", "0")
	boardingPass = strings.ReplaceAll(boardingPass, "B", "1")
	boardingPass = strings.ReplaceAll(boardingPass, "L", "0")
	boardingPass = strings.ReplaceAll(boardingPass, "R", "1")
	seatID, _ := strconv.ParseInt(boardingPass, 2, 16)

	return int(seatID)
}

// calculates the sum of the values from min to max inclusive
func calculateContinuousSum(min, max int) int {
	return (max*(max+1))/2 - ((min-1)*min)/2
}

func readLines(filename string) []string {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, strings.TrimSuffix(scanner.Text(), "\n"))
	}

	return lines
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// Max returns the larger of x or y.
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// Max returns the larger of x or y.
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
