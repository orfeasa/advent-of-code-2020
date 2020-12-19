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
	rules, messages, maxLength := processInput(inputPath)

	// TODO: hide rule memo from this part
	ruleMemo := make(map[int][]string)
	validMessages, ruleMemo := computeMessagesThatMatch(0, rules, ruleMemo, maxLength)
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
	rules, _, maxLength := processInput(inputPath)
	rules[8] = "42 | 42 8"
	rules[11] = "42 31 | 42 11 31"

	fmt.Println(maxLength)

	return 0
}

func computeMessagesThatMatch(ruleID int, rules map[int]string, ruleMemo map[int][]string, maxLength int) ([]string, map[int][]string) {
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
			for _, val := range sidesRulesStr {
				sideRules = append(sideRules, toInt(val))
			}
			var sideMessages []string
			for _, ruleID := range sideRules {
				var validMessages []string
				validMessages, ruleMemo = computeMessagesThatMatch(ruleID, rules, ruleMemo, maxLength)
				sideMessages = combineStringSlices(sideMessages, validMessages)
			}

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
	for _, ruleID := range sideRules {
		var validMessages []string
		validMessages, ruleMemo = computeMessagesThatMatch(ruleID, rules, ruleMemo, maxLength)
		allMessages = combineStringSlices(allMessages, validMessages)
	}
	ruleMemo[ruleID] = allMessages
	return ruleMemo[ruleID], ruleMemo
}

func combineStringSlices(slices ...[]string) (result []string) {
	if len(slices) > 2 {
		var temp [][]string
		temp = append(temp, combineStringSlices(slices[0], slices[1]))
		temp2 := append(temp, slices[2:]...)
		return combineStringSlices(temp2...)
	}
	if len(slices[0]) == 0 {
		return slices[1]
	}
	for _, val1 := range slices[0] {
		for _, val2 := range slices[1] {
			result = append(result, val1+val2)
		}
	}
	return result
}

func processInput(inputPath string) (rules map[int]string, messages []string, maxLength int) {
	input := strings.Split(readRaw(inputPath), "\n\n")
	rulesRaw := strings.Split(input[0], "\n")
	messages = strings.Split(input[1], "\n")

	rules = make(map[int]string)
	for _, rule := range rulesRaw {
		rulesRe := regexp.MustCompile(`^(\d+): (.*)$`)
		match := rulesRe.FindAllStringSubmatch(rule, -1)
		rules[toInt(match[0][1])] = match[0][2]
	}

	maxLength = 0
	for _, message := range messages {
		if len(message) > maxLength {
			maxLength = len(message)
		}
	}
	return rules, messages, maxLength
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
