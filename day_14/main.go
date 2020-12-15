package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	inputPath := "./day_14/input.txt"
	fmt.Println("--- Part One ---")
	fmt.Println(part1(inputPath))

	fmt.Println("--- Part Two ---")
	fmt.Println(part2(inputPath))

}

func part1(inputPath string) int {
	program := readStrings(inputPath)

	currentMask := ""
	memory := make(map[int]int)
	memRe := regexp.MustCompile(`^mem\[(\d+)\] = (\d+)$`)
	maskRe := regexp.MustCompile(`^mask = ([01X]+)$`)
	for _, line := range program {
		if line[:4] == "mask" {
			match := maskRe.FindStringSubmatch(line)
			currentMask = match[1]
		} else {
			match := memRe.FindAllStringSubmatch(line, -1)
			address := toInt(match[0][1])
			value := toInt(match[0][2])
			maskedVal := applyMask(value, currentMask)
			memory[address] = maskedVal
		}

	}

	acc := 0
	for _, val := range memory {
		acc += val
	}
	return acc
}

func part2(inputPath string) int {
	program := readStrings(inputPath)

	currentMask := ""
	memory := make(map[int]int)
	memRe := regexp.MustCompile(`^mem\[(\d+)\] = (\d+)$`)
	maskRe := regexp.MustCompile(`^mask = ([01X]+)$`)
	for _, line := range program {
		if line[:4] == "mask" {
			match := maskRe.FindStringSubmatch(line)
			currentMask = match[1]
		} else {
			match := memRe.FindAllStringSubmatch(line, -1)
			address := toInt(match[0][1])
			value := toInt(match[0][2])
			addresses := maskAddresses(address, currentMask)
			for _, address := range addresses {
				memory[address] = value
			}
		}
	}

	acc := 0
	for _, val := range memory {
		acc += val
	}
	return acc
}

func applyMask(val int, mask string) int {
	for ind, char := range mask {
		// iteration here goes left to right, but masks are applied right to left
		switch char {
		case '1':
			val = setBit(val, len(mask)-ind-1)
		case '0':
			val = clearBit(val, len(mask)-ind-1)
		}
	}
	return val
}

func maskAddresses(address int, mask string) (addresses []int) {
	var floating []int
	for ind, char := range mask {
		// iteration here goes left to right, but masks are applied right to left
		if char == '1' {
			address = setBit(address, len(mask)-ind-1)
		} else if char == 'X' {
			floating = append(floating, ind)
		}
	}

	nextAddress := address
	for i := 0; i < 1<<len(floating); i++ {
		binary := strconv.FormatInt(int64(i), 2)

		// prepend zeros
		binary = padLeft(binary, "0", len(floating))
		for ind, char := range binary {
			switch char {
			case '1':
				nextAddress = setBit(nextAddress, len(mask)-floating[ind]-1)
			case '0':
				nextAddress = clearBit(nextAddress, len(mask)-floating[ind]-1)
			}
			addresses = append(addresses, nextAddress)
		}
	}

	return addresses
}

func padLeft(str, pad string, length int) string {
	for {
		str = pad + str
		if len(str) > length {
			return str[len(str)-length:]
		}
	}
}

// Sets the bit at pos in the integer n
func setBit(n int, pos int) int {
	n |= (1 << pos)
	return n
}

// Clears the bit at pos in n
func clearBit(n int, pos int) int {
	mask := ^(1 << pos)
	n &= mask
	return n
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
