package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputPath := "./day_24/test_input.txt"
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
	tilesToFlip := processInput(inputPath)
	coordToBlackTile := make(map[[3]int]bool)
	var tileQueue [][3]int
	for _, tile := range tilesToFlip {
		if _, ok := coordToBlackTile[tile]; ok {
			delete(coordToBlackTile, tile)
		} else {
			coordToBlackTile[tile] = true
			tileQueue = append(tileQueue, getNeighbors(tile)...)
		}
	}
	tilesToFlip = tileQueue
	tileQueue = nil

	newcoordToBlackTile := make(map[[3]int]bool)
	for k, v := range coordToBlackTile {
		newcoordToBlackTile[k] = v
	}
	for days := 0; days < 10; days++ {
		for _, tile := range tilesToFlip {
			if shouldFlip(tile, coordToBlackTile) {
				if _, ok := coordToBlackTile[tile]; ok {
					delete(newcoordToBlackTile, tile)
				} else {
					newcoordToBlackTile[tile] = true
					tileQueue = append(tileQueue, getNeighbors(tile)...)
				}
			}
		}
		tilesToFlip = tileQueue
		tileQueue = nil

		coordToBlackTile := make(map[[3]int]bool)
		for k, v := range newcoordToBlackTile {
			coordToBlackTile[k] = v
		}

		fmt.Println("Day", days+1, ":", len(coordToBlackTile))
	}

	return len(coordToBlackTile)
}

func shouldFlip(tile [3]int, coordToBlackTile map[[3]int]bool) bool {
	neighbors := getNeighbors(tile)
	countBlackNeighbors := 0
	for _, neighbor := range neighbors {
		if coordToBlackTile[neighbor] {
			countBlackNeighbors++
		}
	}

	// Any black tile with zero or more than 2 black tiles immediately adjacent to it
	if coordToBlackTile[tile] && (countBlackNeighbors == 0 || countBlackNeighbors > 2) {
		return true
	}
	// Any white tile with exactly 2 black tiles immediately adjacent to it
	if !coordToBlackTile[tile] && countBlackNeighbors == 2 {
		return true
	}
	return false
}

func getNeighbors(tile [3]int) (neighbors [][3]int) {
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			// x + y + z = 0 always
			neighbor := [3]int{tile[0] + dx, tile[1] + dy, tile[2] - dx - dy}
			neighbors = append(neighbors, neighbor)
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

func toInt(s string) int {
	result, err := strconv.Atoi(s)
	check(err)
	return result
}
