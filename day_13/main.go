package main

import (
	"bufio"
	"fmt"
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
	busIDs := strings.Split(input[1], ",")
	earliestTime := toInt(input[0])

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
	var divisors []int
	var remainders []int
	for ind, bus := range busIDs {
		if bus != "x" {
			divisors = append(divisors, toInt(bus))
			remainders = append(remainders, toInt(bus)-ind)
		}
	}
	return chineseRemainderTheorem(remainders, divisors)
}

// https://brilliant.org/wiki/chinese-remainder-theorem/ and https://www.dave4math.com/mathematics/chinese-remainder-theorem/
func chineseRemainderTheorem(remainders, divisors []int) int {
	N := 1
	for _, divisor := range divisors {
		N *= divisor
	}

	solution := 0
	for i := range divisors {
		y := N / divisors[i]
		// inverse modulo y
		z := modularInverse(y, divisors[i])
		solution += remainders[i] * y * z
	}
	return solution % N
}

// https://www.geeksforgeeks.org/multiplicative-inverse-under-modulo-m/
func modularInverse(a, m int) int {
	if gcd(a, m) == 1 {
		return power(a, m-2, m)
	}
	return -1
}

func power(x, y, m int) int {
	if y == 0 {
		return 1
	}
	p := power(x, y/2, m) % m
	p = (p * p) % m
	if y%2 == 0 {
		return p
	}
	return ((x * p) % m)
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
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
