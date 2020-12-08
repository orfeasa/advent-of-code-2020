package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	inputPath := "./day_08/input.txt"
	fmt.Println("--- Part One ---")
	fmt.Println(Part1(inputPath))

	fmt.Println("--- Part Two ---")
	fmt.Println(Part2(inputPath))
}

func Part1(inputPath string) int {
	instructions := strings.Split(readRaw(inputPath), "\n")

	instructionsExecuted := make(map[int]bool)

	programCounter := 0
	acc := 0

	for !instructionsExecuted[programCounter] {
		instructionsExecuted[programCounter] = true

		instruction := instructions[programCounter][:3]
		argument := toInt(instructions[programCounter][4:])

		programCounter, acc = runInstruction(programCounter, acc, instruction, argument)
	}
	return acc
}

func Part2(inputPath string) int {
	instructions := strings.Split(readRaw(inputPath), "\n")

	// iterate over codebase to change nop to jmp and inverse
	for ind := range instructions {
		instructionsExecuted := make(map[int]bool)
		programCounter := 0
		acc := 0

		// run the program if the instruction at the current index is nop or jmp
		if instructions[ind][:3] == "nop" || instructions[ind][:3] == "jmp" {
			// execute program while there's no infinite loop
			for !instructionsExecuted[programCounter] {
				// mark instruction as ran
				instructionsExecuted[programCounter] = true

				// execute command
				instruction := instructions[programCounter][:3]
				argument := toInt(instructions[programCounter][4:])

				// make the swap
				if programCounter == ind {
					switch instruction {
					case "nop":
						instruction = "jmp"
					case "jmp":
						instruction = "nop"
					}

				}

				programCounter, acc = runInstruction(programCounter, acc, instruction, argument)

				// check if program terminated
				if programCounter >= len(instructions) {
					return acc
				}
			}
		}
	}
	return 0
}

func runInstruction(programCounter, acc int, instruction string, argument int) (int, int) {

	switch instruction {
	case "acc":
		return programCounter + 1, acc + argument
	case "nop":
		return programCounter + 1, acc
	case "jmp":
		return programCounter + argument, acc
	default:
		return 0, 0
	}

}

// readRaw returns the content of a text file as a string
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
