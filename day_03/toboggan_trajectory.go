package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputPath := "./day_03/input.txt"
	lines := readLines(inputPath)

	treesMultiplied := 1
	treesMultiplied *= countTreesInPath(lines, 1, 1)
	treesMultiplied *= countTreesInPath(lines, 3, 1)
	treesMultiplied *= countTreesInPath(lines, 5, 1)
	treesMultiplied *= countTreesInPath(lines, 7, 1)
	treesMultiplied *= countTreesInPath(lines, 1, 2)

	fmt.Println("Number of trees:", treesMultiplied)
}

func countTreesInPath(mapOfTrees []string, right, down int) int {
	x := 0
	var lineLength int
	countTrees := 0
	for y := 0; y < len(mapOfTrees); y += down {
		if y != 0 {
			x += right
			if string(mapOfTrees[y][x%lineLength]) == "#" {
				countTrees++
			}
		} else {
			lineLength = len(mapOfTrees[y])
			continue
		}

	}
	return countTrees
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
