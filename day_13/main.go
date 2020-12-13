package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputPath := "./day_13/input.txt"
	fmt.Println("--- Part One ---")
	fmt.Println(part1(inputPath))

	fmt.Println("--- Part Two ---")
	fmt.Println(part2(inputPath))
}

func part1(inputPath string) int {
	input := readStrings(inputPath)
	earliestTime := toInt(input[0])

	busIDs := strings.Split(input[1], ",")

	minTime := 0
	minBus := 0
	initialized := false
	for _, bus := range busIDs {
		if bus != "x" {
			busID := toInt(bus)
			currTime := int(math.Ceil(float64(earliestTime)/float64(busID)) * float64(busID))
			if !initialized {
				minTime = currTime
				minBus = busID
				initialized = true
			} else if currTime < minTime {
				minTime = currTime
				minBus = busID
			}
		}
	}
	return minBus * (minTime - earliestTime)
}

func part2(inputPath string) int {
	input := readStrings(inputPath)
	busIDs := strings.Split(input[1], ",")
	// need to find t such as for all buses
	// (t + ind) % bus = 0
	t := 0
	for {
		t++
		found := true
		for ind, bus := range busIDs {
			if bus != "x" && (t+ind)%toInt(bus) != 0 {
				found = false
				break
			}
		}
		if found {
			break
		}
	}

	return t
}

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

func readNumbers(filename string) []int {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	Scanner := bufio.NewScanner(file)

	var numbers []int
	for Scanner.Scan() {
		numbers = append(numbers, toInt(Scanner.Text()))
	}
	return numbers
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
