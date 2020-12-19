package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	inputPath := "./day_19/input.txt"
	fmt.Println("--- Part One ---")
	fmt.Println(part1(inputPath))

	inputPath = "./day_19/input2.txt"
	fmt.Println("--- Part Two ---")
	fmt.Println(part2(inputPath))
}

func part1(inputPath string) int {
	input := strings.Split(readRaw(inputPath), "\n\n")

	rulesRaw := strings.Split(input[0], "\n")
	messages := strings.Split(input[1], "\n")

	rules := make(map[int]string)
	for _, rule := range rulesRaw {
		rulesRe := regexp.MustCompile(`^(\d+): (.*)$`)
		match := rulesRe.FindAllStringSubmatch(rule, -1)
		rules[toInt(match[0][1])] = match[0][2]
	}

	// TODO: hide rule memo from this part
	ruleMemo := make(map[int][]string)
	validMessages, ruleMemo := computeMessagesThatMatch(0, rules, ruleMemo)
	count := 0
	for _, message := range messages {
		// if message in validMessages
		for _, validMessage := range validMessages {
			if message == validMessage {
				count++
				break
			}
		}
	}
	return count
}

func part2(inputPath string) int {
	return 0
}

func computeMessagesThatMatch(ruleID int, rules map[int]string, ruleMemo map[int][]string) ([]string, map[int][]string) {

	// if already computed
	if _, ok := ruleMemo[ruleID]; ok {
		return ruleMemo[ruleID], ruleMemo
	}

	// if rule is matching character return it
	if strings.HasPrefix(rules[ruleID], `"`) {
		// remove quotation marks
		ruleMemo[ruleID] = []string{strings.ReplaceAll(rules[ruleID], `"`, ``)}
		return ruleMemo[ruleID], ruleMemo
	}

	if strings.Contains(rules[ruleID], " | ") {
		listsOfSubRulesStr := strings.Split(rules[ruleID], " | ")
		// each side of the |
		var allMessages []string
		for _, val := range listsOfSubRulesStr {

			sidesRulesStr := strings.Split(val, " ")
			// convert all side rules
			var sideRules []int
			var sideMessages []string
			for _, val := range sidesRulesStr {
				sideRules = append(sideRules, toInt(val))
			}
			sideMessages, ruleMemo = computeMessagesThatMatchAll(sideRules, rules, ruleMemo)
			allMessages = append(allMessages, sideMessages...)
		}
		ruleMemo[ruleID] = allMessages
		return ruleMemo[ruleID], ruleMemo
	}

	sidesRulesStr := strings.Split(rules[ruleID], " ")
	// convert all side rules
	var sideRules []int
	for _, val := range sidesRulesStr {
		sideRules = append(sideRules, toInt(val))
	}
	var allMessages []string
	allMessages, ruleMemo = computeMessagesThatMatchAll(sideRules, rules, ruleMemo)
	ruleMemo[ruleID] = allMessages
	return ruleMemo[ruleID], ruleMemo
}

func computeMessagesThatMatchAll(ruleIDs []int, rules map[int]string, ruleMemo map[int][]string) ([]string, map[int][]string) {
	var prevMessages []string
	for _, ruleID := range ruleIDs {
		var validMessages []string
		validMessages, ruleMemo = computeMessagesThatMatch(ruleID, rules, ruleMemo)
		prevMessages = combineStringSlices(prevMessages, validMessages)
	}
	return prevMessages, ruleMemo
}

func combineStringSlices(slices ...[]string) (result []string) {
	if len(slices) > 2 {
		var temp [][]string
		temp = append(temp, combineStringSlices(slices[0], slices[1]))
		temp2 := append(temp, slices[2:]...)
		return combineStringSlices(temp2...)
	}

	slice1 := slices[0]
	slice2 := slices[1]

	if len(slice1) == 0 {
		return slice2
	}
	for _, val1 := range slice1 {
		for _, val2 := range slice2 {
			result = append(result, val1+val2)
		}
	}
	return result
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
