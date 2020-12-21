package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	ingredients, allergenToIngredientsCandidates := processInput(inputPath)

	// ingredientToOccurrences["mxmxvkd"] = 3 means that ingredient mxmxvkd is present in 3 foods
	ingredientToOccurrences := make(map[string]int)
	for _, ingredientsOfFood := range ingredients {
		for _, allergen := range ingredientsOfFood {
			ingredientToOccurrences[allergen]++
		}
	}

	for _, ingredients := range allergenToIngredientsCandidates {
		for _, ingredient := range ingredients {
			delete(ingredientToOccurrences, ingredient)
		}
	}

	acc := 0
	for _, v := range ingredientToOccurrences {
		acc += v
	}

	return acc
}

func part2(inputPath string) string {
	_, allergenToIngredientsCandidates := processInput(inputPath)

	allergenToIngredient := make(map[string]string)

	// until there are no allergens without an ingredient connection
	for len(allergenToIngredientsCandidates) != 0 {
		for allergen, ingredientCandidates := range allergenToIngredientsCandidates {
			// if there's only one possibility for a rule to validate a field
			if len(ingredientCandidates) == 1 {
				allergenToIngredient[allergen] = ingredientCandidates[0]
				allergenToIngredientsCandidates = removeKeyFromMap(allergenToIngredientsCandidates, ingredientCandidates[0])
			}
		}
	}

	fmt.Println(allergenToIngredient)

	sortedAllergens := make([]string, 0, len(allergenToIngredient))
	for allergen := range allergenToIngredient {
		sortedAllergens = append(sortedAllergens, allergen)
	}
	sort.Strings(sortedAllergens)
	fmt.Println(sortedAllergens)

	sortedIngredientsByAllergen := make([]string, 0, len(allergenToIngredient))
	for _, allergen := range sortedAllergens {
		sortedIngredientsByAllergen = append(sortedIngredientsByAllergen, allergenToIngredient[allergen])
	}

	return strings.Join(sortedIngredientsByAllergen, ",")
}

func removeKeyFromMap(m map[string][]string, key string) map[string][]string {
	// remove key from map values slice
	for k, v := range m {
		for _, i := range v {
			if i == key {
				m[k] = removeFromSlice(v, i)
				break
			}
		}
	}
	// completely remove key if it's now empty
	for k, v := range m {
		if len(v) == 0 {
			delete(m, k)
		}
	}
	return m
}

// deletes the first occurrence of item in array, assuming order is not important
func removeFromSlice(array []string, item string) []string {
	for ind, val := range array {
		if val == item {
			array[ind] = array[len(array)-1]
			return array[:len(array)-1]
		}
	}
	return array
}

func intersect(array1 []string, array2 []string) (result []string) {
	array1map := make(map[string]bool)
	for _, val1 := range array1 {
		array1map[val1] = true
	}
	for _, val2 := range array2 {
		if _, ok := array1map[val2]; ok {
			result = append(result, val2)
		}
	}
	return result
}

func processInput(inputPath string) (ingredients [][]string, allergenToIngredientsCandidates map[string][]string) {
	list := readStrings(inputPath)
	var allergens [][]string
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
	allergenToIngredientsCandidates = make(map[string][]string)
	for foodInd, allergensOfFood := range allergens {
		for _, allergen := range allergensOfFood {
			// if not already in the map
			if _, ok := allergenToIngredientsCandidates[allergen]; !ok {
				allergenToIngredientsCandidates[allergen] = ingredients[foodInd]
			} else {
				// remove ingredients that are not present in allergenToIngredientsCandidates[allergen] and ingredients[foodInd]
				allergenToIngredientsCandidates[allergen] = intersect(allergenToIngredientsCandidates[allergen], ingredients[foodInd])
			}
		}
	}
	return ingredients, allergenToIngredientsCandidates
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
