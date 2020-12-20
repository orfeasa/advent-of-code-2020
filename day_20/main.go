package main

import (
	"fmt"
	"io/ioutil"
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
	input := readRaw(inputPath)
	tilesRaw := strings.Split(input, "\n\n")
	var tiles []tile
	for _, tileRaw := range tilesRaw {
		tileRawSplit := strings.Split(tileRaw, "\n")
		id := toInt(strings.TrimRight(strings.TrimLeft(tileRawSplit[0], "Tile "), ":\n"))
		image := tileRawSplit[1:]
		borderIDs := calculateBorderIDs(image)
		newTile := tile{
			id: id, image: image, borderIDs: borderIDs,
		}
		tiles = append(tiles, newTile)
	}

	// tileBorderIDs[123] = [a, b, c] means the border id 123 is present for tiles with IDs a, b c
	tileBorderIDs := make(map[int][]int)
	for _, tile := range tiles {
		for _, borderID := range tile.borderIDs {
			tileBorderIDs[borderID] = append(tileBorderIDs[borderID], tile.id)
		}
	}

	// countTileOccurencesInBorder[123] = 2 means tile with id 123 is present in the border twice
	countTileOccurencesInBorder := make(map[int]int)
	for _, val := range tileBorderIDs {
		if len(val)%2 == 1 {
			countTileOccurencesInBorder[val[0]]++
		}
	}

	acc := 1
	for k, v := range countTileOccurencesInBorder {
		if v%2 == 0 {
			acc *= k
		}
	}
	return acc
}

func part2(inputPath string) int {
	return 0
}

// IDs are calculated by converting # to 1 and . to 0, and taking the resulting number or its reverse, whichever is smaller
func calculateBorderIDs(image []string) (IDs []int) {
	for i := 0; i < 4; i++ {
		lineID := strings.ReplaceAll(image[0], "#", "1")
		lineID = strings.ReplaceAll(lineID, ".", "0")
		i, _ := strconv.ParseInt(lineID, 2, 32)
		ID1 := int(i)
		i, _ = strconv.ParseInt(reverse(lineID), 2, 32)
		ID2 := int(i)
		if ID1 < ID2 {
			IDs = append(IDs, ID1)
		} else {
			IDs = append(IDs, ID2)
		}
		image = rotateImage(image)
	}
	return IDs
}

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return result
}

func flipImage(image []string) []string {
	n := len(image)
	flipped := make([][]rune, n)
	for ind, line := range image {
		flipped[ind] = make([]rune, n)
		for i, j := 0, len(line)-1; i < j; i, j = i+1, j-1 {
			flipped[ind][i], flipped[ind][j] = rune(line[j]), rune(line[i])
		}
	}
	var newImage []string
	for _, line := range flipped {
		newImage = append(newImage, string(line))
	}
	return newImage
}

func rotateImage(image []string) []string {
	n := len(image)
	rotated := make([][]rune, n)
	for i := range rotated {
		rotated[i] = make([]rune, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			rotated[i][j] = rune(image[n-j-1][i])
		}
	}
	var newImage []string
	for _, line := range rotated {
		newImage = append(newImage, string(line))
	}
	return newImage
}

type tile struct {
	id        int
	image     []string
	borderIDs []int
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
