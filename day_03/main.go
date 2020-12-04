package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {

	inputPath := "./day_03/input.txt"
	slopes := [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}

	lines := readLines(inputPath)

	treesMultiplied := 1
	var wg sync.WaitGroup
	wg.Add(len(slopes))
	for _, slope := range slopes {
		go func(slope []int) {
			defer wg.Done()
			treesMultiplied *= countTreesInPath(lines, slope[0], slope[1])
		}(slope)
	}

	wg.Wait()
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

func check(err error) {
	if err != nil {
		panic(err)
	}
}
