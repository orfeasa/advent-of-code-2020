package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputPath := "./day_12/input.txt"
	fmt.Println("--- Part One ---")
	fmt.Println(part1(inputPath))

	fmt.Println("--- Part Two ---")
	fmt.Println(part2(inputPath))
}

type ship struct {
	x           int
	y           int
	orientation int // value from 0 to 270, 0 is North
}

func part1(inputPath string) int {
	lines := readStrings(inputPath)
	sh := ship{x: 0, y: 0, orientation: 90}
	for _, command := range lines {
		sh = move(sh, command[0], toInt(command[1:]))
	}
	return abs(sh.x) + abs(sh.y)
}

func part2(inputPath string) int {
	lines := readStrings(inputPath)
	sh := ship{x: 0, y: 0, orientation: 90}
	// waypoint coordinates are relative to the ship
	waypoint := ship{x: 10, y: 1, orientation: 0}

	for _, command := range lines {
		action := command[0]
		value := toInt(command[1:])
		switch action {
		case 'N', 'S', 'E', 'W':
			waypoint = move(waypoint, action, value)
		case 'F':
			sh.x += value * waypoint.x
			sh.y += value * waypoint.y
		case 'R', 'L':
			waypoint.x, waypoint.y = rotatePoint(waypoint.x, waypoint.y, action, value)
		}
	}
	return abs(sh.x) + abs(sh.y)
}

func rotatePoint(x, y int, direction byte, deg int) (int, int) {
	// convert L to R
	if direction == 'L' {
		deg = 360 - deg
	}

	// calculate amount of right rotation and apply
	times := deg / 90
	for i := 0; i < times; i++ {
		x, y = y, -x
	}
	return x, y
}

func move(sh ship, action byte, value int) ship {
	switch action {
	case 'N':
		sh.y += value
	case 'S':
		sh.y -= value
	case 'E':
		sh.x += value
	case 'W':
		sh.x -= value
	case 'L':
		sh.orientation += 360 - value
		sh.orientation %= 360
	case 'R':
		sh.orientation += value
		sh.orientation %= 360
	case 'F':
		var orientations = [4]byte{'N', 'E', 'S', 'W'}
		sh = move(sh, orientations[sh.orientation/90], value)
	}
	return sh
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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
