package main

import (
	"fmt"
	"sort"
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

	var labelsSorted = make([]int, len(input))
	copy(labelsSorted[:], labels[:])
	sort.Sort(sort.Reverse(sort.IntSlice(labelsSorted[:])))

	// store labels in linked list
	l := CircularLinkedList{len: 0}
	for _, val := range labels {
		l.Insert(val)
	}

	currentCup := l.head
	for i := 0; i < 100; i++ {
		// remove and keep value and order of 3 values next to current cup
		var removed []int
		for j := 0; j < 3; j++ {
			val := l.GetAt(l.Search(currentCup.value) + 1).value
			removed = append(removed, val)
			l.DeleteAt(l.Search(currentCup.value) + 1)
		}
		// select destination cup
		destinationVal := getNext(currentCup.value, labelsSorted)
		for l.Search(destinationVal) == -1 {
			destinationVal = getNext(destinationVal, labelsSorted)
		}
		// place cups next to destination
		for ind, val := range removed {
			l.InsertAt(l.Search(destinationVal)+ind+1, val)
		}
		currentCup = currentCup.next
	}

	node := l.GetAt(l.Search(1))
	node = node.next
	repr := 0
	for node.value != 1 {
		repr = 10*repr + node.value
		node = node.next
	}
	return repr
}

func part2(input string) int {
	return 0
}

func getNext(find int, a []int) (next int) {
	for ind, val := range a {
		if val == find {
			return a[(ind+1)%len(a)]
		}
	}
	return -1
}

// Node represents a node of linked list
type Node struct {
	value int
	next  *Node
}

// CircularLinkedList represents a linked list
type CircularLinkedList struct {
	head *Node
	len  int
}

// ListToArray prints the circular list without repetition
func (l *CircularLinkedList) ListToArray() []int {
	array := make([]int, 0, l.len)
	current := l.head
	for i := 0; i < l.len; i++ {
		array = append(array, current.value)
		current = current.next
	}
	return array
}

// Insert inserts new node at the end of linked list
func (l *CircularLinkedList) Insert(val int) {
	n := Node{value: val, next: l.head}
	if l.len == 0 {
		l.head = &n
		l.head.next = &n
		l.len++
		return
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.next == l.head {
			ptr.next = &n
			ptr.next.next = l.head
			l.len++
			return
		}
		ptr = ptr.next
	}
}

// InsertAt inserts new node at given position
func (l *CircularLinkedList) InsertAt(pos int, value int) {
	newNode := Node{value: value, next: l.GetAt(pos)}
	prevNode := l.GetAt(pos - 1)
	prevNode.next = &newNode
	if pos%l.len == 0 {
		l.head = prevNode.next
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
