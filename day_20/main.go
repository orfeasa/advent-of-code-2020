package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	inputPath := "./day_20/input.txt"

	fmt.Println("--- Part One ---")
	fmt.Println(part1(inputPath))

	fmt.Println("--- Part Two ---")
	fmt.Println(part2(inputPath))
}

func part1(inputPath string) int {
	_, cornerTilesIDs, _, _ := processInput(inputPath)
	acc := 1
	for _, id := range cornerTilesIDs {
		acc *= id
	}
	return acc
}

func part2(inputPath string) int {
	tiles, cornerTilesIDs, borderTileIDs, internalTileIDs := processInput(inputPath)

	// elect the top-left corner
	gridSize := int(math.Sqrt(float64(len(tiles))))
	grid := make([][]int, gridSize)
	for ind := range grid {
		grid[ind] = make([]int, gridSize)
	}
	grid[0][0] = cornerTilesIDs[0]

	fmt.Println(len(borderTileIDs))
	fmt.Println(len(internalTileIDs))

	// corner tiles are the tiles that 2 of their border IDs exist in only 1 tile

	// rotate and flip tile so that left and top tile are the corner ones

	// find the tile that connects to the right of it
	// this is the only one that has the right id, rotated and flipped to be in the correct way and the top to be the border one

	// assembledImage := make([][]int)

	return 0
}

type tile struct {
	id        int
	image     []string
	borderIDs []int
}

// IDs are calculated by converting # to 1 and . to 0, and taking the resulting number or its reverse, whichever is smaller
func (t *tile) calculateBorderIDs() {
	t.borderIDs = make([]int, 0, 4)
	for i := 0; i < 4; i++ {
		lineID := strings.ReplaceAll(t.image[0], "#", "1")
		lineID = strings.ReplaceAll(lineID, ".", "0")
		i, _ := strconv.ParseInt(lineID, 2, 32)
		ID1 := int(i)
		i, _ = strconv.ParseInt(reverse(lineID), 2, 32)
		ID2 := int(i)
		if ID1 < ID2 {
			t.borderIDs = append(t.borderIDs, ID1)
		} else {
			t.borderIDs = append(t.borderIDs, ID2)
		}
		t.rotateImage()
	}
}

// reverses a string
func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return result
}

func (t *tile) flipImage() {
	n := len(t.image)
	flipped := make([][]rune, n)
	// flip
	for ind, line := range t.image {
		flipped[ind] = make([]rune, n)
		for i, j := 0, len(line)-1; i < j; i, j = i+1, j-1 {
			flipped[ind][i], flipped[ind][j] = rune(line[j]), rune(line[i])
		}
	}
	// convert to string slice
	var newImage []string
	for _, line := range flipped {
		newImage = append(newImage, string(line))
	}
	t.image = newImage
}

func (t *tile) rotateImage() {
	n := len(t.image)
	rotated := make([][]rune, n)
	for i := range rotated {
		rotated[i] = make([]rune, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			rotated[i][j] = rune(t.image[n-j-1][i])
		}
	}
	t.image = nil
	for _, line := range rotated {
		t.image = append(t.image, string(line))
	}
}

func processInput(inputPath string) (tiles []tile, cornerTilesIDs []int, borderTileIDs []int, internalTileIDs []int) {
	input := readRaw(inputPath)
	tilesRaw := strings.Split(input, "\n\n")
	for _, tileRaw := range tilesRaw {
		tileRawSplit := strings.Split(tileRaw, "\n")
		id := toInt(strings.TrimRight(strings.TrimLeft(tileRawSplit[0], "Tile "), ":\n"))
		image := tileRawSplit[1:]
		newTile := tile{id: id, image: image}
		newTile.calculateBorderIDs()
		tiles = append(tiles, newTile)
	}

	// tileBorderIDs[123] = [a, b, c] means the border id 123 is present for tiles with IDs a, b c
	tileBorderIDs := make(map[int][]int)
	for _, tile := range tiles {
		for _, borderID := range tile.borderIDs {
			tileBorderIDs[borderID] = append(tileBorderIDs[borderID], tile.id)
		}
	}

	// countTileOccurrencesInBorder[123] = 2 means tile with id 123 is present in the border twice
	countTileOccurrencesInBorder := make(map[int]int)
	for _, tiles := range tileBorderIDs {
		if len(tiles) == 1 {
			countTileOccurrencesInBorder[tiles[0]]++
		}
	}

	for _, tile := range tiles {
		if countTileOccurrencesInBorder[tile.id] == 2 {
			cornerTilesIDs = append(cornerTilesIDs, tile.id)
		} else if countTileOccurrencesInBorder[tile.id] == 1 {
			borderTileIDs = append(borderTileIDs, tile.id)
		} else {
			internalTileIDs = append(internalTileIDs, tile.id)
		}
	}
	return tiles, cornerTilesIDs, borderTileIDs, internalTileIDs
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
