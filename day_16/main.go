package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	inputPath := "./day_16/input.txt"
	fmt.Println("--- Part One ---")
	fmt.Println(part1(inputPath))

	fmt.Println("--- Part Two ---")
	fmt.Println(part2(inputPath))
}

type rule struct {
	name                        string
	min1, max1, min2, max2, pos int
}

func part1(inputPath string) int {
	rules, _, tickets := processInput(inputPath)
	min1, max1, min2, max2 := calculateGlobalRanges(rules)

	acc := 0
	for _, ticket := range tickets {
		for _, field := range ticket {
			if field < min1 || (field > max1 && field < min2) || field > max2 {
				acc += field
			}
		}
	}
	return acc
}

func part2(inputPath string) int {
	rules, myTicket, tickets := processInput(inputPath)

	var validTickets [][]int
	for _, ticket := range tickets {
		if isValidTicket(ticket, rules) {
			validTickets = append(validTickets, ticket)
		}
	}

	// field candidates contains for each field index which rules are valid for it
	fieldCandidates := make(map[int][]int)
	for ruleInd, rule := range rules {
		for fieldPos := 0; fieldPos < len(tickets[0]); fieldPos++ {
			fieldIsValid := true
			for _, ticket := range validTickets {
				if !fieldValidatesRule(ticket[fieldPos], rule) {
					fieldIsValid = false
					break
				}
			}
			if fieldIsValid {
				fieldCandidates[fieldPos] = append(fieldCandidates[fieldPos], ruleInd)
			}
		}
	}

	// until there are no candidates left without a field
	for len(fieldCandidates) != 0 {
		for fieldPos, possibleRules := range fieldCandidates {
			// if there's only one possibility
			if len(possibleRules) == 1 {
				rules[possibleRules[0]].pos = fieldPos
				fieldCandidates = removeKeyFromMap(fieldCandidates, possibleRules[0])
			}
		}
	}

	acc := 1
	for _, rule := range rules {
		if strings.HasPrefix(rule.name, "departure") {
			acc *= myTicket[rule.pos]
		}
	}
	return acc
}

func removeKeyFromMap(m map[int][]int, key int) map[int][]int {
	// remove key from map values slice
	for k, v := range m {
		for ind, i := range v {
			if i == key {
				m[k] = remove(v, ind)
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

// removes element at index i from slice s by replacing element at index i with the last element and
// returning all but the last element. Assumes order is not important
func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	// We do not need to put s[i] at the end, as it will be discarded anyway
	return s[:len(s)-1]
}

func calculateGlobalRanges(rules []rule) (min1, max1, min2, max2 int) {
	min1, min2 = 1000000, 1000000
	for _, rule := range rules {
		if min1 > rule.min1 {
			min1 = rule.min1
		}
		if max1 < rule.max1 {
			max1 = rule.max1
		}
		if min2 > rule.min2 {
			min2 = rule.min2
		}
		if max2 < rule.max2 {
			max2 = rule.max2
		}
	}
	return min1, max1, min2, max2
}

func isValidTicket(ticket []int, rules []rule) bool {
	min1, max1, min2, max2 := calculateGlobalRanges(rules)
	for _, field := range ticket {
		if field < min1 || (field > max1 && field < min2) || field > max2 {
			return false
		}
	}
	return true
}

func fieldValidatesRule(field int, rul rule) bool {
	if field < rul.min1 || (field > rul.max1 && field < rul.min2) || field > rul.max2 {
		return false
	}
	return true
}

func processInput(inputPath string) (rules []rule, myTicket []int, tickets [][]int) {
	input := strings.Split(readRaw(inputPath), "\n\n")

	// process rules
	for _, val := range strings.Split(input[0], "\n") {
		rangesRe := regexp.MustCompile(`^([a-z ]+): (\d+)-(\d+) or (\d+)-(\d+)$`)
		match := rangesRe.FindAllStringSubmatch(val, -1)
		r := rule{name: match[0][1], min1: toInt(match[0][2]), min2: toInt(match[0][4]), max1: toInt(match[0][3]), max2: toInt(match[0][5]), pos: -1}
		rules = append(rules, r)
	}

	// process my ticket
	myTicketStr := strings.Split(input[1], "\n")[1]
	myTicketFieldsStr := strings.Split(myTicketStr, ",")
	for _, field := range myTicketFieldsStr {
		myTicket = append(myTicket, toInt(field))
	}

	// process tickets
	ticketsStr := strings.Split(input[2], "\n")[1:]
	for _, val := range ticketsStr {
		fieldsStr := strings.Split(val, ",")
		var ticket []int
		for _, fieldStr := range fieldsStr {
			field := toInt(fieldStr)
			ticket = append(ticket, field)
		}
		tickets = append(tickets, ticket)
	}

	return rules, myTicket, tickets
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
