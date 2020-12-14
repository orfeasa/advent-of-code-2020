package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
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

	// DEBUG
	currentMask = "000000000000000000000000000000X1001X"
	address := 42
	addresses := maskAddresses(address, currentMask)
	fmt.Println(addresses)

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

	for _, val := range floating {

	}
	for ind, char := range mask {
		if char == 'X' {
			fmt.Println("address=", address, ", ind=", ind, "so appending", setBit(address, len(mask)-ind-1), "and", clearBit(address, len(mask)-ind-1))
			addresses = append(addresses, setBit(address, len(mask)-ind-1))
			addresses = append(addresses, clearBit(address, len(mask)-ind-1))
		}
	}
	return addresses
}

// Sets the bit at pos in the integer n.
func setBit(n int, pos int) int {
	n |= (1 << pos)
	return n
}

// Clears the bit at pos in n.
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
