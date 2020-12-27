package main

import (
	"fmt"
	"strconv"
)

func main() {
	input := "487912365"
	fmt.Println("--- Part One ---")
	fmt.Println(part1(input))

	fmt.Println("--- Part Two ---")
	fmt.Println(part2(input))
}

func part1(input string) int {
	// parse input
	labels := make([]int, 0, len(input))
	for _, val := range input {
		labels = append(labels, toInt(string(val)))
	}

	// store values
	valToCup := make(map[int]*cup)
	var prevLabel int
	for _, label := range labels {
		newCup := cup{val: label}
		valToCup[label] = &newCup
		if prevLabel != 0 {
			valToCup[prevLabel].next = &newCup
		}
		prevLabel = label
	}
	valToCup[prevLabel].next = valToCup[labels[0]]

	// play game as many times as required
	currentCup := labels[0]
	for round := 0; round < 100; round++ {
		currentCup, valToCup = runCrabMove(currentCup, valToCup)
	}

	// generate representation
	cup := valToCup[1].next
	repr := 0
	for cup.val != 1 {
		repr = 10*repr + cup.val
		cup = cup.next
	}
	return repr
}

func part2(input string) int {
	// parse input
	labels := make([]int, 0, len(input))
	for _, val := range input {
		labels = append(labels, toInt(string(val)))
	}

	// store values
	valToCup := make(map[int]*cup)
	var prevLabel int
	for _, label := range labels {
		newCup := cup{val: label}
		valToCup[label] = &newCup
		if prevLabel != 0 {
			valToCup[prevLabel].next = &newCup
		}
		prevLabel = label
	}
	for i := len(valToCup) + 1; i <= 1e6; i++ {
		newCup := cup{val: i}
		valToCup[i] = &newCup
		valToCup[prevLabel].next = &newCup
		prevLabel = i
	}
	valToCup[prevLabel].next = valToCup[labels[0]]

	// play game as many times as required
	currentCup := labels[0]
	for round := 0; round < 10e6; round++ {
		currentCup, valToCup = runCrabMove(currentCup, valToCup)
	}

	// calculate output
	return valToCup[1].next.val * valToCup[1].next.next.val
}

func runCrabMove(currentCup int, valToCup map[int]*cup) (int, map[int]*cup) {
	// remove the 3 cups next to the current cup
	firstRemovedCup := valToCup[currentCup].next
	// assign the current cup's next cup to the cup after 4 cups
	valToCup[currentCup].next = valToCup[currentCup].next.next.next.next

	// find destination cup
	destination := currentCup - 1
	for destination <= 0 || destination == firstRemovedCup.val || destination == firstRemovedCup.next.val || destination == firstRemovedCup.next.next.val {
		destination--
		if destination <= 0 {
			destination = len(valToCup)
		}
	}

	// place 3 cups next to destination cup
	oldNext := valToCup[destination].next
	valToCup[destination].next = firstRemovedCup
	valToCup[destination].next.next.next.next = oldNext

	currentCup = valToCup[currentCup].next.val
	return currentCup, valToCup
}

type cup struct {
	val  int
	next *cup
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
