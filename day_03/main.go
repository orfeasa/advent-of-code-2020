package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"sync"
)

func main() {

	inputPath := "./day_03/input.txt"
	lines := strings.Split(readRaw(inputPath), "\n")

	slopes := [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}

	treesMultiplied := 1
	var wg sync.WaitGroup

	for _, slope := range slopes {
		wg.Add(1)
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

// readRaw returns the content of a text file as a string
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
