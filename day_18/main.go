package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	inputPath := "./day_18/input.txt"
	fmt.Println("--- Part One ---")
	fmt.Println(part1(inputPath))
	fmt.Println("--- Part Two ---")
	fmt.Println(part2(inputPath))
}

func part1(inputPath string) int {
	input := readStrings(inputPath)
	acc := 0
	for _, line := range input {
		acc += evaluateExpression(line, samePrecedence)
	}
	return acc
}

func part2(inputPath string) int {
	input := readStrings(inputPath)
	acc := 0
	for _, line := range input {
		acc += evaluateExpression(line, additionPrecedence)
	}
	return acc
}

type precedenceEvaluator func(operator rune) int

// Implementation of the Shunting-yard algorithm as explained here https://www.geeksforgeeks.org/expression-evaluation/
func evaluateExpression(expression string, precedence precedenceEvaluator) int {
	// remove all spaces
	expression = strings.ReplaceAll(expression, " ", "")
	var valueStack []int
	var operatorStack []rune
	for _, token := range expression {
		if unicode.IsDigit(token) {
			valueStack = append(valueStack, toInt(string(token)))
		} else if token == '(' {
			operatorStack = append(operatorStack, token)
		} else if token == ')' {
			// 	While the thing on top of the operator stack is not a left parenthesis
			for operatorStack[len(operatorStack)-1] != '(' {
				operatorStack, valueStack = executeOperation(operatorStack, valueStack)
			}
			// discard left parenthesis from the operator stack
			operatorStack = operatorStack[:len(operatorStack)-1]
		} else { //operator
			thisOp := token
			for len(operatorStack) != 0 && precedence(operatorStack[len(operatorStack)-1]) >= precedence(thisOp) {
				operatorStack, valueStack = executeOperation(operatorStack, valueStack)
			}
			operatorStack = append(operatorStack, thisOp)
		}
	}
	for len(operatorStack) != 0 {
		operatorStack, valueStack = executeOperation(operatorStack, valueStack)
	}
	// The result is the only value left in the value stack
	return valueStack[0]
}

func executeOperation(operatorStack []rune, valueStack []int) ([]rune, []int) {
	// pop operator
	operator := operatorStack[len(operatorStack)-1]
	operatorStack = operatorStack[:len(operatorStack)-1]

	// pop operands
	operands := [2]int{valueStack[len(valueStack)-1], valueStack[len(valueStack)-2]}
	valueStack = valueStack[:len(valueStack)-2]

	// execute apply operator and push result to the value stack
	result := applyOp(operands, operator)
	valueStack = append(valueStack, result)
	return operatorStack, valueStack
}

func applyOp(operands [2]int, operator rune) int {
	switch operator {
	case '+':
		return operands[0] + operands[1]
	case '*':
		return operands[0] * operands[1]
	default:
		panic(fmt.Sprintln("Error: operator", string(operator), "not recognised during applyOp"))
	}
}

func samePrecedence(operator rune) int {
	switch operator {
	case '+':
		return 1
	case '*':
		return 1
	default:
		return 0
	}
}

func additionPrecedence(operator rune) int {
	switch operator {
	case '+':
		return 2
	case '*':
		return 1
	default:
		return 0
	}
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
