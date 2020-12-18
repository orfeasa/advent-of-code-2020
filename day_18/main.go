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

// https://www.geeksforgeeks.org/expression-evaluation/
func evaluateExpression(expression string, precedence precedenceEvaluator) int {
	// remove all spaces
	expression = strings.ReplaceAll(expression, " ", "")
	var valueStack []int
	var operatorStack []rune
	// 1. While there are still tokens to be read in,
	for _, token := range expression {
		// 	1.1 Get the next token.
		// 	1.2 If the token is:
		// 		1.2.1 A number: push it onto the value stack.
		// 		1.2.2 A variable: get its value, and push onto the value stack.
		if unicode.IsDigit(token) {
			valueStack = append(valueStack, toInt(string(token)))
			// 		1.2.3 A left parenthesis: push it onto the operator stack.
		} else if token == '(' {
			operatorStack = append(operatorStack, token)
			// 		1.2.4 A right parenthesis:
		} else if token == ')' {
			// 		1 While the thing on top of the operator stack is not a left parenthesis,
			for operatorStack[len(operatorStack)-1] != '(' {
				// 			1 Pop the operator from the operator stack.
				operator := operatorStack[len(operatorStack)-1]
				operatorStack = operatorStack[:len(operatorStack)-1]
				// 			2 Pop the value stack twice, getting two operands.
				operands := [2]int{valueStack[len(valueStack)-1], valueStack[len(valueStack)-2]}
				valueStack = valueStack[:len(valueStack)-2]
				// 			3 Apply the operator to the operands, in the correct order.
				result := applyOp(operands, operator)
				// 			4 Push the result onto the value stack.
				valueStack = append(valueStack, result)
			}
			// 		2 Pop the left parenthesis from the operator stack, and discard it.
			operatorStack = operatorStack[:len(operatorStack)-1]
			// 		1.2.5 An operator (call it thisOp):
		} else {
			thisOp := token
			// 		1 While the operator stack is not empty, and the top thing on the
			// 			operator stack has the same or greater precedence as thisOp,
			for len(operatorStack) != 0 && precedence(operatorStack[len(operatorStack)-1]) >= precedence(thisOp) {
				// 			1 Pop the operator from the operator stack.
				operator := operatorStack[len(operatorStack)-1]
				operatorStack = operatorStack[:len(operatorStack)-1]
				// 			2 Pop the value stack twice, getting two operands.
				operands := [2]int{valueStack[len(valueStack)-1], valueStack[len(valueStack)-2]}
				valueStack = valueStack[:len(valueStack)-2]
				// 			3 Apply the operator to the operands, in the correct order.
				result := applyOp(operands, operator)
				// 			4 Push the result onto the value stack.
				valueStack = append(valueStack, result)
			}
			// 		2 Push thisOp onto the operator stack.
			operatorStack = append(operatorStack, thisOp)

		}

	}
	// 2. While the operator stack is not empty,
	for len(operatorStack) != 0 {
		// 	1 Pop the operator from the operator stack.
		operator := operatorStack[len(operatorStack)-1]
		operatorStack = operatorStack[:len(operatorStack)-1]
		// 	2 Pop the value stack twice, getting two operands.
		operands := [2]int{valueStack[len(valueStack)-1], valueStack[len(valueStack)-2]}
		valueStack = valueStack[:len(valueStack)-2]
		// 	3 Apply the operator to the operands, in the correct order.
		result := applyOp(operands, operator)
		// 	4 Push the result onto the value stack.
		valueStack = append(valueStack, result)
	}
	// 3. At this point the operator stack should be empty, and the value
	// 	stack should have only one value in it, which is the final result.
	return valueStack[0]
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
