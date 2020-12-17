package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputPath := "./day_17/input.txt"
	fmt.Println("--- Part One ---")
	fmt.Println(part1(inputPath))

	fmt.Println("--- Part Two ---")
	fmt.Println(part2(inputPath))
}

type cube struct {
	x, y, z  int
	isActive bool
}

func part1(inputPath string) int {
	input := strings.Split(readRaw(inputPath), "\n")
	cubesActive := make(map[[3]int]bool)
	// x increases right, y increases down
	for y, line := range input {
		for x, char := range line {
			z := 0
			if char == '#' {
				cubesActive[[3]int{x, y, z}] = true
			} else {
				cubesActive[[3]int{x, y, z}] = false
			}

		}
	}

	for i := 0; i < 6; i++ {
		cubesActive = runCycle(cubesActive)
	}
	count := 0
	for _, isActive := range cubesActive {
		if isActive {
			count++
		}
	}

	return count
}

func part2(inputPath string) int {
	return 0
}

func runCycle(cubesActive map[[3]int]bool) (nextCubesActive map[[3]int]bool) {
	nextCubesActive = make(map[[3]int]bool)
	for coord, isActive := range cubesActive {
		activeNeighborsCount := countActiveNeighbors(coord, cubesActive)
		if isActive && (activeNeighborsCount == 2 || activeNeighborsCount == 3) {
			nextCubesActive[coord] = true
		} else if !isActive && activeNeighborsCount == 3 {
			nextCubesActive[coord] = true
		} else {
			nextCubesActive[coord] = false
		}
	}
	return nextCubesActive
}

func countActiveNeighbors(coord [3]int, cubesActive map[[3]int]bool) int {
	neighbors := getNeighbors(coord)
	count := 0
	for _, val := range neighbors {
		if cubesActive[val] {
			count++
		}
	}
	return count
}

func getNeighbors(coord [3]int) (neighbors [][3]int) {
	x0, y0, z0 := coord[0], coord[1], coord[2]
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			for dz := -1; dz <= 1; dz++ {
				if !(dx == 0 && dy == 0 && dz == 0) {
					neighbors = append(neighbors, [3]int{x0 + dx, y0 + dy, z0 + dz})
				}
			}
		}
	}
	return neighbors
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
