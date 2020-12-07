package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	inputPath := "./day_07/input.txt"
	fmt.Println("--- Part One ---")
	fmt.Println(Part1(inputPath))

	fmt.Println("--- Part Two ---")
	fmt.Println(Part2(inputPath))
}

func Part1(inputPath string) int {
	bagContains := parseRules(readRaw(inputPath))

	bagIsContainedBy := make(map[string][]string)

	// reverse map
	for k1, v1 := range bagContains {
		for k2, _ := range v1 {
			bagIsContainedBy[k2] = append(bagIsContainedBy[k2], k1)
		}
	}

	bagsFound := make(map[string]bool)
	queue := []string{"shiny gold"}

	for len(queue) > 0 {
		currentColor := queue[0]
		for _, color := range bagIsContainedBy[currentColor] {
			queue = append(queue, color)
			bagsFound[color] = true
		}

		// dequeue
		queue = queue[1:]
	}

	return len(bagsFound)
}

func Part2(inputPath string) int {
	bagContains := parseRules(readRaw(inputPath))
	return countBagsContainedIn("shiny gold", bagContains) - 1
}

func countBagsContainedIn(bagColor string, bagContains map[string]map[string]int) int {
	if len(bagContains[bagColor]) == 0 {
		return 1
	}

	count := 0
	for color, qty := range bagContains[bagColor] {
		count += qty * countBagsContainedIn(color, bagContains)
	}
	return count + 1
}

func parseRules(rules string) map[string]map[string]int {
	ruleMap := make(map[string]map[string]int)
	ruleLines := strings.Split(rules, "\n")

	for _, ruleLine := range ruleLines {
		lineProcessed := ruleLine
		// remove trailing whitespaces
		lineProcessed = strings.TrimRight(lineProcessed, " ")
		// remove trailing dot
		lineProcessed = strings.TrimRight(lineProcessed, ".")
		// remove reduntant bag reference
		lineProcessed = strings.ReplaceAll(lineProcessed, " bags", "")
		lineProcessed = strings.ReplaceAll(lineProcessed, " bag", "")
		lineProcessed = strings.ReplaceAll(lineProcessed, "no other", "")

		// split rule to containing and contained bags
		ruleSplit := strings.Split(lineProcessed, " contain ")
		containing := ruleSplit[0]
		contained := strings.Split(ruleSplit[1], ", ")
		ruleMap[containing] = make(map[string]int)
		for _, value := range contained {
			if value != "" {
				// regex to split numbers and color (\d+) ([a-zA-Z ]+)
				re := regexp.MustCompile(`^(\d+) ([a-zA-Z ]+)$`)
				match := re.FindAllStringSubmatch(value, -1)
				qty := toInt(match[0][1])
				color := match[0][2]
				ruleMap[containing][color] = qty
			}
		}

	}

	return ruleMap

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
