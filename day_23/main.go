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

	// store labels in linked list
	l := CircularLinkedList{len: 0}
	for _, val := range labels {
		l.Insert(val)
	}

	currentCup := l.head
	for i := 0; i < 100; i++ {
		currentCup = l.runCrabMove(currentCup)
	}

	node := l.GetAt(l.Search(1)).next
	repr := 0
	for node.value != 1 {
		repr = 10*repr + node.value
		node = node.next
	}
	return repr
}

func part2(input string) int {

	// circle has 1 mil total cups
	// after 10 mil moves, find the 2 cups next to number 1
	// parse input
	labels := make([]int, 0, len(input))
	for _, val := range input {
		labels = append(labels, toInt(string(val)))
	}

	// store labels in linked list
	l := CircularLinkedList{len: 0}
	for _, val := range labels {
		l.Insert(val)
	}

	for i := l.len + 1; i < 1e6; i++ {
		l.Insert(i)
	}

	currentCup := l.head
	for i := 0; i < 1e4; i++ {
		currentCup = l.runCrabMove(currentCup)
	}

	node := l.GetAt(l.Search(1)).next
	fmt.Println(node.value, node.next.value)
	return 0
}

// TODO: always change current cup to head
func (l *CircularLinkedList) runCrabMove(currentCup *Node) *Node {
	// remove and keep value and order of 3 values next to current cup
	var removed []int
	for j := 0; j < 3; j++ {
		val := l.GetAt(l.Search(currentCup.value) + 1).value
		removed = append(removed, val)
		l.DeleteAt(l.Search(currentCup.value) + 1)
	}
	// select destination cup
	destinationVal := currentCup.value - 1
	for destinationVal <= 0 || intInSlice(removed, destinationVal) {
		destinationVal--
		if destinationVal <= 0 {
			// assign to max value
			destinationVal = l.len + len(removed)
		}
	}
	// place cups next to destination
	for ind, val := range removed {
		l.InsertAt(l.Search(destinationVal)+ind+1, val)
	}
	// l.head = currentCup.next
	return currentCup.next
}

func intInSlice(slice []int, target int) bool {
	for _, val := range slice {
		if val == target {
			return true
		}
	}
	return false
}

// Node represents a node of linked list
type Node struct {
	value int
	next  *Node
	prev  *Node
}

// CircularLinkedList represents a linked list
type CircularLinkedList struct {
	head *Node
	len  int
}

// Insert inserts new node at the end of linked list
func (l *CircularLinkedList) Insert(val int) {
	n := Node{value: val, next: l.head}
	if l.len == 0 {
		l.head = &n
		l.head.next = &n
		l.head.prev = &n
		l.len++
		return
	}

	ptr := l.head.prev
	ptr.next = &n
	l.head.prev = &n
	ptr.next.next = l.head
	l.len++
	return
}

// InsertAt inserts new node at given position
func (l *CircularLinkedList) InsertAt(pos int, value int) {
	prevNode := l.GetAt(pos - 1)
	nextNode := l.GetAt(pos)
	newNode := Node{value: value, next: nextNode, prev: prevNode}
	prevNode.next = &newNode
	nextNode.prev = &newNode
	if pos%l.len == 0 {
		l.head = &newNode
	}
	l.len++
}

// GetAt returns node at given position from linked list
func (l *CircularLinkedList) GetAt(pos int) *Node {
	pos %= l.len
	if pos < 0 {
		pos += l.len
	}
	ptr := l.head
	for i := 0; i < pos; i++ {
		ptr = ptr.next
	}
	return ptr
}

// Search returns node position with given value from linked list
func (l *CircularLinkedList) Search(val int) int {
	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.value == val {
			return i
		}
		ptr = ptr.next
	}
	return -1
}

// DeleteAt deletes node at given position from linked list
func (l *CircularLinkedList) DeleteAt(pos int) error {
	prevNode := l.GetAt(pos - 1)
	prevNode.next = l.GetAt(pos).next
	if pos%l.len == 0 {
		l.head = prevNode.next
	}
	l.len--
	return nil
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
