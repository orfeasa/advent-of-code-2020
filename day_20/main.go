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
	tiles, borderIDtoTiles := processInput(inputPath)
	// tileIDtoCountBorderOccurrences[123] = 2 means tile with id 123 is present in the border twice
	tileIDtoCountBorderOccurrences := make(map[int]int)
	for _, tiles := range borderIDtoTiles {
		if len(tiles) == 1 {
			tileIDtoCountBorderOccurrences[tiles[0]]++
		}
	}
	acc := 1
	for _, tile := range tiles {
		if tileIDtoCountBorderOccurrences[tile.id] == 2 {
			acc *= tile.id
		}
	}
	return acc
}

func part2(inputPath string) int {
	tileIDToTile, borderIDtoTiles := processInput(inputPath)
	// tileIDtoCountBorderOccurrences[123] = 2 means tile with id 123 is present in the border twice
	tileIDtoCountBorderOccurrences := make(map[int]int)
	for _, tiles := range borderIDtoTiles {
		if len(tiles) == 1 {
			tileIDtoCountBorderOccurrences[tiles[0]]++
		}
	}

	// elect the top-left corner
	gridSize := int(math.Sqrt(float64(len(tileIDToTile))))
	grid := make([][]int, gridSize)
	for ind := range grid {
		grid[ind] = make([]int, gridSize)
	}
	tilesLeft := make(map[int]*tile)
	for id, tile := range tileIDToTile {
		tilesLeft[id] = tile
	}
	var top, left string
	left = ""
	top = ""
	for x := 0; x < gridSize; x++ {
		for y := 0; y < gridSize; y++ {
			if x != 0 {
				left = reverseString(tileIDToTile[grid[y][x-1]].getRightBorder())
			} else {
				left = ""
			}
			if y != 0 {
				top = reverseString(tileIDToTile[grid[y-1][x]].getBottomBorder())
			} else {
				top = ""
			}
			grid[y][x] = findTileThatFits(top, left, tilesLeft, borderIDtoTiles)
			delete(tilesLeft, grid[y][x])
		}
	}
	fmt.Println(grid)

	// create final image
	tileSize := len(tileIDToTile[grid[0][0]].image) - 2 // without borders
	imageSize := gridSize * tileSize
	finalImage := make([]string, 0, imageSize)

	// look for sea monsters
	for y := 0; y < imageSize; y++ {
		for _, id := range grid[y/tileSize] {
			finalImage[y] += tileIDToTile[id].image[y/tileSize][1 : tileSize+1]
		}
	}

	for _, line := range finalImage {
		fmt.Println(line)
	}

	// seaMonster := "                  #"
	// +"#    ##    ##    ###"
	// +"#  #  #  #  #  #   "

	return 0
}

// reverse returns the reverse of a string
func reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

func findTileThatFits(top, left string, tiles map[int]*tile, borderIDtoTiles map[int][]int) (id int) {
	for _, t := range tiles {
		for fl := 0; fl < 2; fl++ {
			for rot := 0; rot < 4; rot++ {
				isTopOk := false
				if top != "" {
					isTopOk = t.getTopBorder() == top
				} else {
					isTopOk = len(borderIDtoTiles[calculateBorderID(t.getTopBorder())]) == 1
				}
				isLeftOk := false
				if left != "" {
					isLeftOk = t.getLeftBorder() == left
				} else {
					isLeftOk = len(borderIDtoTiles[calculateBorderID(t.getLeftBorder())]) == 1
				}
				if isTopOk && isLeftOk {
					return t.id
				}
				t.rotateImage()
			}
			t.flipImage()
		}
	}
	return -1
}

type tile struct {
	id        int
	image     []string
	borderIDs []int
}

// getLeftBorder returns the left border as a string. Does not mutate the original image
func (t tile) getLeftBorder() string {
	t.rotateImage()
	return t.image[0]
}

// getTopBorder returns the top border as a string. Does not mutate the original image
func (t tile) getTopBorder() string {
	return t.image[0]
}

// getRightBorder returns the right border as a string. Does not mutate the original image
func (t tile) getRightBorder() string {
	t.rotateImage()
	t.rotateImage()
	t.rotateImage()
	return t.image[0]
}

// getBottomBorder returns the bottom border as a string. Does not mutate the original image
func (t tile) getBottomBorder() string {
	t.rotateImage()
	t.rotateImage()
	return t.image[0]
}

// calculateBorderIDs calculates and stores border IDs to the tile. IDs are calculated by converting
// # to 1 and . to 0, and taking the resulting number or its reverse, whichever is smaller
func (t *tile) calculateBorderIDs() {
	t.borderIDs = make([]int, 0, 4)
	for i := 0; i < 4; i++ {
		t.borderIDs = append(t.borderIDs, calculateBorderID(t.image[0]))
		t.rotateImage()
	}
}

func calculateBorderID(line string) int {
	lineID := strings.ReplaceAll(line, "#", "1")
	lineID = strings.ReplaceAll(lineID, ".", "0")
	i, _ := strconv.ParseInt(lineID, 2, 32)
	id1 := int(i)
	i, _ = strconv.ParseInt(reverse(lineID), 2, 32)
	id2 := int(i)
	if id1 < id2 {
		return id1
	}
	return id2
}

// reverses a string
func reverseString(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return result
}

// flipImage flips the image of a tile horizontally. Mutates the original image
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

// rotateImage rotates the image of a tile clockwise. Mutates the original image
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

func processInput(inputPath string) (tileIDToTile map[int]*tile, borderIDtoTiles map[int][]int) {
	input := readRaw(inputPath)
	tilesRaw := strings.Split(input, "\n\n")
	tileIDToTile = make(map[int]*tile)
	for _, tileRaw := range tilesRaw {
		tileRawSplit := strings.Split(tileRaw, "\n")
		id := toInt(strings.TrimRight(strings.TrimLeft(tileRawSplit[0], "Tile "), ":\n"))
		image := tileRawSplit[1:]
		newTile := tile{id: id, image: image}
		newTile.calculateBorderIDs()
		tileIDToTile[id] = &newTile
	}

	// borderIDtoTiles[123] = [a, b, c] means the border id 123 is present for tiles with IDs a, b c
	borderIDtoTiles = make(map[int][]int)
	for _, tile := range tileIDToTile {
		for _, borderID := range tile.borderIDs {
			borderIDtoTiles[borderID] = append(borderIDtoTiles[borderID], tile.id)
		}
	}

	return tileIDToTile, borderIDtoTiles
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
