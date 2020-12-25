package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inputPath := "./day_25/input.txt"
	fmt.Println("--- Part One ---")
	fmt.Println(part1(inputPath))

	fmt.Println("--- Part Two ---")
	fmt.Println(part2(inputPath))
}

func part1(inputPath string) int {
	publicKeys := readNumbers(inputPath)
	loopSizes := make([]int, 0, len(publicKeys))
	for _, publicKey := range publicKeys {
		loopSizes = append(loopSizes, calculateLoopSize(7, publicKey))
	}
	encryptionKey := transformSubjectNumber(publicKeys[0], loopSizes[1])
	return encryptionKey
}

func part2(inputPath string) int {
	return 0
}

func calculateLoopSize(subjectNumber, publicKey int) (loopSize int) {
	calcPublicKey := 1
	for {
		loopSize++
		calcPublicKey *= subjectNumber
		calcPublicKey %= 20201227
		if publicKey == calcPublicKey {
			return loopSize
		}
	}
}

func transformSubjectNumber(subjectNumber int, loopSize int) (publicKey int) {
	publicKey = 1
	for i := 0; i < loopSize; i++ {
		publicKey *= subjectNumber
		publicKey %= 20201227
	}
	return publicKey
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
