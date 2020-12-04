package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputPath := "./day_04/input.txt"
	lines := readLines(inputPath)

	fieldsString := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid", "cid"}
	possibleFields := make(map[string]bool)
	for _, field := range fieldsString {
		possibleFields[field] = true
	}

	fmt.Println(countValidPassports(lines, possibleFields, "cid"))

}

func countValidPassports(lines []string, possibleFields map[string]bool, optionalField string) int {
	missingFields := make(map[string]bool)
	for k, v := range possibleFields {
		missingFields[k] = v
	}

	countValid := 0

	for _, line := range lines {
		if line != "" {
			data := strings.Split(line, " ")
			var lineFields []string
			for _, pair := range data {
				field := strings.Split(pair, ":")
				lineFields = append(lineFields, field[0])
			}

			for _, field := range lineFields {
				delete(missingFields, field)
			}
		} else {
			if len(missingFields) == 0 || (len(missingFields) == 1 && missingFields["cid"]) {
				// increase count of valid passports
				countValid++
			}

			// reset missingFields with all possible fields
			for k := range possibleFields {
				missingFields[k] = true
			}
		}
	}
	return countValid
}

func readLines(filename string) []string {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, strings.TrimSuffix(scanner.Text(), "\n"))
	}

	// add extra new line for the last passport
	lines = append(lines, "")
	return lines
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
