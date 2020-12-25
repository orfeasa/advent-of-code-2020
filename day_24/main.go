package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputPath := "./day_24/input.txt"
	fmt.Println("--- Part One ---")
	fmt.Println(part1(inputPath))

	fmt.Println("--- Part Two ---")
	fmt.Println(part2(inputPath))
}

func part1(inputPath string) int {
	tilesToFlip := processInput(inputPath)
	coordToBlackTile := make(map[[3]int]bool)
	for _, tile := range tilesToFlip {
		if _, ok := coordToBlackTile[tile]; ok {
			delete(coordToBlackTile, tile)
		} else {
			coordToBlackTile[tile] = true
		}
	}
	return len(coordToBlackTile)
}

func part2(inputPath string) int {
	tilesToFlipArray := processInput(inputPath)
	coordToBlackTile := make(map[[3]int]bool)
	tilesToCheck := make(map[[3]int]bool)

	// do initial flipping (tiles can be flipped twice)
	for _, tile := range tilesToFlipArray {
		if _, ok := coordToBlackTile[tile]; ok {
			coordToBlackTile[tile] = !coordToBlackTile[tile]
		} else {
			coordToBlackTile[tile] = true
		}
		tilesToCheck[tile] = true
		neighbors := getNeighbors(tile)
		for _, neighbor := range neighbors {
			tilesToCheck[neighbor] = true
		}
	}

	// start daily flipping
	tilesToCheckNext := make(map[[3]int]bool)
	tilesToFlip := make(map[[3]int]bool)
	for days := 0; days < 100; days++ {
		// find tiles to flip
		for tile := range tilesToCheck {
			if shouldFlip(tile, coordToBlackTile) {
				tilesToFlip[tile] = true
				tilesToCheckNext[tile] = true
				neighbors := getNeighbors(tile)
				for _, neighbor := range neighbors {
					tilesToCheckNext[neighbor] = true
				}
			}
		}
		// flip them
		for tile := range tilesToFlip {
			if _, ok := coordToBlackTile[tile]; ok {
				coordToBlackTile[tile] = !coordToBlackTile[tile]
			} else {
				coordToBlackTile[tile] = true
			}
		}
		// prepare for the next round
		tilesToFlip = make(map[[3]int]bool)
		tilesToCheck = tilesToCheckNext
		tilesToCheckNext = make(map[[3]int]bool)
	}
	return countBlackTiles(coordToBlackTile)
}

func countBlackTiles(coordToBlackTile map[[3]int]bool) (count int) {
	for _, v := range coordToBlackTile {
		if v {
			count++
		}
	}
	return count
}

func shouldFlip(tile [3]int, coordToBlackTile map[[3]int]bool) bool {
	neighbors := getNeighbors(tile)
	countBlackNeighbors := 0
	for _, neighbour := range neighbors {
		if coordToBlackTile[neighbour] {
			countBlackNeighbors++
		}
	}
	isBlack := false
	if _, ok := coordToBlackTile[tile]; ok {
		isBlack = coordToBlackTile[tile]
	}

	// Any black tile with zero or more than 2 black tiles immediately adjacent to it
	if isBlack && (countBlackNeighbors == 0 || countBlackNeighbors > 2) {
		return true
	}
	// Any white tile with exactly 2 black tiles immediately adjacent to it
	if !isBlack && countBlackNeighbors == 2 {
		return true
	}
	return false
}

func getNeighbors(tile [3]int) (neighbors [][3]int) {
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx != dy {
				// x + y + z = 0 always
				neighbor := [3]int{tile[0] + dx, tile[1] + dy, tile[2] - dx - dy}
				neighbors = append(neighbors, neighbor)
			}
		}
	}
	return neighbors
}

// Using cube tile coordinates as suggested here
// https://www.redblobgames.com/grids/hexagons/#coordinates-cube
func processInput(inputPath string) (tilesToFlip [][3]int) {
	input := readStrings(inputPath)
	tilesToFlip = make([][3]int, 0, len(input))
	for _, line := range input {
		var direction string
		var currentTile [3]int
		isTileFinalised := false
		for _, ch := range line {
			direction += string(ch)
			isTileFinalised = true
			switch direction {
			case "e":
				currentTile[0]++
				currentTile[1]--
			case "se":
				currentTile[1]--
				currentTile[2]++
			case "sw":
				currentTile[0]--
				currentTile[2]++
			case "w":
				currentTile[0]--
				currentTile[1]++
			case "nw":
				currentTile[1]++
				currentTile[2]--
			case "ne":
				currentTile[0]++
				currentTile[2]--
			default:
				isTileFinalised = false
			}
			if isTileFinalised {
				direction = ""
			}
		}
		tilesToFlip = append(tilesToFlip, currentTile)
	}
	return tilesToFlip
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
