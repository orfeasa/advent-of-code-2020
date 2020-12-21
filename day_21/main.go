package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputPath := "./day_21/input.txt"
	fmt.Println("--- Part One ---")
	fmt.Println(part1(inputPath))

	fmt.Println("--- Part Two ---")
	fmt.Println(part2(inputPath))
}

func part1(inputPath string) int {
	list := readStrings(inputPath)
	var allergens [][]string
	var ingredients [][]string
	for _, line := range list {
		listSplit := strings.Split(line, " (")
		ingredientsOfFood := strings.Split(listSplit[0], " ")
		listSplit[1] = strings.TrimLeft(listSplit[1], "(contains")
		listSplit[1] = strings.TrimLeft(listSplit[1], " ")
		listSplit[1] = strings.TrimRight(listSplit[1], ")")
		allergensOfFood := strings.Split(listSplit[1], ", ")
		allergens = append(allergens, allergensOfFood)
		ingredients = append(ingredients, ingredientsOfFood)
	}

	// create map for each allergen which ingredients it could be caused from
	allergenToIngredients := make(map[string][]string)
	for foodInd, allergensOfFood := range allergens {
		for _, allergen := range allergensOfFood {
			// if not already in the map
			if val, ok := allergenToIngredients[allergen]; !ok {
				allergenToIngredients[allergen] = ingredients[foodInd]
			} else {
				// remove ingredients that are not present in allergenToIngredients[allergen] and ingredients[foodInd]

			}

		}
	}

	return 0
}

func part2(inputPath string) int {
	return 0
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
