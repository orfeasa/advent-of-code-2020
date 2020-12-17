package main

import (
	"fmt"
	"io/ioutil"
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
	const Dimensions = 3
	cubesActive := make(map[[Dimensions]int]bool)
	// x increases right, y increases down
	for y, line := range input {
		for x, char := range line {
			coord := [Dimensions]int{x, y}
			if char == '#' {
				cubesActive[coord] = true
			} else {
				cubesActive[coord] = false
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
	input := strings.Split(readRaw(inputPath), "\n")
	const Dimensions = 4
	cubesActive := make(map[[Dimensions]int]bool)
	// x increases right, y increases down
	for y, line := range input {
		for x, char := range line {
			z := 0
			w := 0
			if char == '#' {
				cubesActive[[4]int{x, y, z, w}] = true
			} else {
				cubesActive[[4]int{x, y, z, w}] = false
			}

		}
	}

	for i := 0; i < 6; i++ {
		cubesActive = runCycle4d(cubesActive)
	}
	count := 0
	for _, isActive := range cubesActive {
		if isActive {
			count++
		}
	}

	return count
}

func runCycle(cubesActive map[[3]int]bool) (nextCubesActive map[[3]int]bool) {
	nextCubesActive = make(map[[3]int]bool)
	var neighborsQueue [][3]int
	for coord := range cubesActive {
		neighborsQueue = append(neighborsQueue, getNeighbors(coord)...)
		nextCubesActive[coord] = calcNextState(coord, cubesActive)
	}
	for _, neighborCoords := range neighborsQueue {
		// if the next state is not already calculated before
		if _, ok := nextCubesActive[neighborCoords]; !ok {
			nextCubesActive[neighborCoords] = calcNextState(neighborCoords, cubesActive)
		}
	}
	return nextCubesActive
}

func runCycle4d(cubesActive map[[4]int]bool) (nextCubesActive map[[4]int]bool) {
	nextCubesActive = make(map[[4]int]bool)
	var neighborsQueue [][4]int
	for coord := range cubesActive {
		neighborsQueue = append(neighborsQueue, getNeighbors4d(coord)...)
		nextCubesActive[coord] = calcNextState4d(coord, cubesActive)
	}
	for _, neighborCoords := range neighborsQueue {
		// if the next state is not already calculated before
		if _, ok := nextCubesActive[neighborCoords]; !ok {
			nextCubesActive[neighborCoords] = calcNextState4d(neighborCoords, cubesActive)
		}
	}
	return nextCubesActive
}

func calcNextState(coord [3]int, cubesActive map[[3]int]bool) bool {
	isActive := cubesActive[coord]
	activeNeighborsCount := countActiveNeighbors(coord, cubesActive)
	if isActive && (activeNeighborsCount == 2 || activeNeighborsCount == 3) {
		return true
	} else if !isActive && activeNeighborsCount == 3 {
		return true
	}
	return false
}

func calcNextState4d(coord [4]int, cubesActive map[[4]int]bool) bool {
	isActive := cubesActive[coord]
	activeNeighborsCount := countActiveNeighbors4d(coord, cubesActive)
	if isActive && (activeNeighborsCount == 2 || activeNeighborsCount == 3) {
		return true
	} else if !isActive && activeNeighborsCount == 3 {
		return true
	}
	return false
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

func countActiveNeighbors4d(coord [4]int, cubesActive map[[4]int]bool) int {
	neighbors := getNeighbors4d(coord)
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

func getNeighbors4d(coord [4]int) (neighbors [][4]int) {
	x0, y0, z0, w0 := coord[0], coord[1], coord[2], coord[3]
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			for dz := -1; dz <= 1; dz++ {
				for dw := -1; dw <= 1; dw++ {
					if !(dx == 0 && dy == 0 && dz == 0 && dw == 0) {
						neighbors = append(neighbors, [4]int{x0 + dx, y0 + dy, z0 + dz, w0 + dw})
					}
				}
			}
		}
	}
	return neighbors
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
