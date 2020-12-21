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

	fmt.Println("--- Part Two ---")
	fmt.Println(part2(inputPath))
}

func part1(inputPath string) int {
	rules, messages, maxLength := processInput(inputPath)

	// TODO: hide rule memo from this part
	exprMemo := make(map[string][]string)
	validMessages, exprMemo := computeMessagesThatMatch(rules[0], rules, exprMemo, maxLength)
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
	rules, messages, maxLength := processInput(inputPath)
	rules[8] = "42 | 42 8"         // 42 | 42 42 | 42 42 42 | 42 42 42 42 |...
	rules[11] = "42 31 | 42 11 31" // 42 31 | 42 42 31 31 | 42 42 42 31 31 31 | ...
	// 0: 8 11 so 0 is at least one 42 followed by more 31s than 42s
	exprMemo := make(map[string][]string)
	var validMessages42, validMessages31 []string
	validMessages42, exprMemo = computeMessagesThatMatch(rules[42], rules, exprMemo, maxLength)
	validMessages31, exprMemo = computeMessagesThatMatch(rules[31], rules, exprMemo, maxLength)

	// iterate over all messages
	count := 0
	for _, message := range messages {
		count42s := 0
		count31s := 0
		low := 0
		// match as many 42s match
		for {
			foundMatch := false
			for _, validMessage42 := range validMessages42 {
				if strings.HasPrefix(message[low:], validMessage42) {
					low += len(validMessage42)
					foundMatch = true
					count42s++
					break
				}
			}
			if !foundMatch {
				break
			}
		}
		// match as many 31s match
		for {
			foundMatch := false
			for _, validMessage31 := range validMessages31 {
				if strings.HasPrefix(message[low:], validMessage31) {
					low += len(validMessage31)
					foundMatch = true
					count31s++
					break
				}
			}
			if !foundMatch {
				break
			}
		}
		if low == len(message) && count31s != 0 && count42s > count31s {
			count++
		}
	}
	return count
}

func computeMessagesThatMatch(expression string, rules map[int]string, exprMemo map[string][]string, maxLength int) ([]string, map[string][]string) {
	// if already computed
	if _, ok := exprMemo[expression]; ok {
		return exprMemo[expression], exprMemo
	}

	// if rule is matching character return it
	if strings.HasPrefix(expression, `"`) {
		// remove quotation marks
		exprMemo[expression] = []string{strings.ReplaceAll(expression, `"`, ``)}
		return exprMemo[expression], exprMemo
	}

	if strings.Contains(expression, " | ") {
		listsOfSubRulesStr := strings.Split(expression, " | ")
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
				validMessages, exprMemo = computeMessagesThatMatch(rules[ruleID], rules, exprMemo, maxLength)
				sideMessages = combineStringSlices(sideMessages, validMessages)
			}

			allMessages = append(allMessages, sideMessages...)
		}
		exprMemo[expression] = allMessages
		return exprMemo[expression], exprMemo
	}
	sidesRulesStr := strings.Split(expression, " ")
	// convert all side rules
	var sideRules []int
	for _, val := range sidesRulesStr {
		sideRules = append(sideRules, toInt(val))
	}
	var allMessages []string
	for _, ruleID := range sideRules {
		var validMessages []string
		validMessages, exprMemo = computeMessagesThatMatch(rules[ruleID], rules, exprMemo, maxLength)
		allMessages = combineStringSlices(allMessages, validMessages)
	}
	exprMemo[expression] = allMessages
	return exprMemo[expression], exprMemo
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
