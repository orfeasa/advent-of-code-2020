package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputPath := "./day_04/input.txt"
	lines := readLines(inputPath)

	fmt.Println(countValidPassports(lines))
}

func countValidPassports(lines []string) int {
	fields := make(map[string]string)

	countValid := 0

	passports := strings.Split(lines, "\n\n")

	for _, passport := range passports {
		passportData := strings.ReplaceAll(passport, "\n", " ")
		fieldsAndValues := strings.Split(line, " ")
		for _, fieldAndValue := range fieldsAndValues {
			fieldValue := strings.Split(fieldAndValue, ":")
			fields[fieldValue[0]] = fieldValue[1]
		}
		if isValidPassport(fields) {
			// increase count of valid passports
			countValid++
		}
		// reset fields
		fields = make(map[string]string)
	}

	for _, line := range lines {
		// if line is not empty (we are reading passport data)
		if line != "" {
			fieldsAndValues := strings.Split(line, " ")
			for _, fieldAndValue := range fieldsAndValues {
				fieldValue := strings.Split(fieldAndValue, ":")
				fields[fieldValue[0]] = fieldValue[1]
			}
		} else {
			if isValidPassport(fields) {
				// increase count of valid passports
				countValid++
			}
			// reset fields
			fields = make(map[string]string)
		}
	}
	return countValid
}

func isValidPassport(fields map[string]string) bool {
	// validate birth year
	if val, ok := fields["byr"]; !ok || len(val) != 4 || toInt(val) < 1920 || toInt(val) > 2002 {
		return false
	}

	// validate issue year
	if val, ok := fields["iyr"]; !ok || len(val) != 4 || toInt(val) < 2010 || toInt(val) > 2020 {
		return false
	}

	// validate expiration year
	if val, ok := fields["eyr"]; !ok || len(val) != 4 || toInt(val) < 2020 || toInt(val) > 2030 {
		return false
	}

	// validate height
	if val, ok := fields["hgt"]; !ok {
		return false
	} else {
		unit := val[len(val)-2:]
		// check valid unit
		if unit != "cm" && unit != "in" {
			return false
		}
		height := toInt(val[:len(val)-2])
		if unit == "cm" && (height < 150 || height > 193) {
			return false
		} else if unit == "in" && (height < 59 || height > 76) {
			return false
		}
	}

	// validate hair color
	if val, ok := fields["hcl"]; !ok {
		return false
	} else {
		if string(val[0]) != "#" {
			return false
		}
		color := val[1:]
		if _, err := strconv.ParseInt(color, 16, 32); len(color) != 6 || err != nil {
			return false
		}

	}

	// validate eye color
	// TODO: implement "in" function
	if val, ok := fields["ecl"]; !ok || (val != "amb" && val != "blu" && val != "brn" && val != "gry" && val != "grn" && val != "hzl" && val != "oth") {
		return false
	}

	// validate passport id
	if val, ok := fields["pid"]; !ok || len(val) != 9 {
		return false
	} else if _, err := strconv.Atoi(val); err != nil {
		return false
	}

	return true
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

func toInt(s string) int {
	result, err := strconv.Atoi(s)
	check(err)
	return result
}
