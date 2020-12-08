package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	inputPath := "./day_02/input.txt"
	fmt.Println("--- Part One ---")
	fmt.Println(part1(inputPath))

	fmt.Println("--- Part Two ---")
	fmt.Println(part2(inputPath))
}

func part1(inputPath string) int {
	lines := readLines(inputPath)

	countOld := 0
	regex := regexp.MustCompile(`^(\d+)-(\d+) ([a-zA-Z]+): ([a-zA-Z]+)$`)

	for _, line := range lines {
		match := regex.FindStringSubmatch(line)
		pos1, pos2, char, password := toInt(match[1]), toInt(match[2]), match[3], match[4]
		if isValidOldPolicy(pos1, pos2, char, password) {
			countOld++
		}
	}
	return countOld
}

func part2(inputPath string) int {
	lines := readLines(inputPath)

	countNew := 0
	regex := regexp.MustCompile(`^(\d+)-(\d+) ([a-zA-Z]+): ([a-zA-Z]+)$`)

	for _, line := range lines {
		match := regex.FindStringSubmatch(line)
		pos1, pos2, char, password := toInt(match[1]), toInt(match[2]), match[3], match[4]
		if isValidNewPolicy(pos1, pos2, char, password) {
			countNew++
		}
	}
	return countNew
}

func isValidOldPolicy(minOcc, maxOcc int, char string, password string) bool {
	count := strings.Count(password, char)
	if count >= minOcc && count <= maxOcc {
		return true
	}
	return false
}

func isValidNewPolicy(pos1, pos2 int, char string, password string) bool {
	isInFirstPos := string(password[pos1-1]) == char
	isInSecondPos := string(password[pos2-1]) == char
	// equivalent to XOR, one exactly must be true
	if isInFirstPos != isInSecondPos {
		return true
	}
	return false
}

func readLines(filename string) []string {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func toInt(s string) int {
	result, err := strconv.Atoi(s)
	check(err)
	return result
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
