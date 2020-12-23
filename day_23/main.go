package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
)

func main() {
	// input := "487912365"
	testInput := "389125467"
	fmt.Println("--- Part One ---")
	fmt.Println(part1(testInput))

	fmt.Println("--- Part Two ---")
	fmt.Println(part2(testInput))
}

// TODO: check https://golang.org/pkg/container/ring/ and https://github.com/atedja/gring/blob/master/ring.go
func part1(input string) int {
	// parse input
	labels := make([]int, 0, len(input))
	for _, val := range input {
		labels = append(labels, toInt(string(val)))
	}

	var labelsSorted = make([]int, len(input))
	copy(labelsSorted[:], labels[:])
	sort.Sort(sort.Reverse(sort.IntSlice(labelsSorted[:])))
	fmt.Println(labelsSorted)

	// store labels in linked list
	ll := CircularLinkedList{len: 0}
	for _, val := range labels {
		ll.Insert(val)
	}

	// simulate 100 moves
	currentCup := ll.head
	for i := 0; i < 10; i++ {
		currentCupInd := ll.Search(currentCup.value)
		fmt.Println("-- move", i+1, "--")
		fmt.Println("cups:", ll.ListToArray(), ", current cup index:", currentCupInd)
		// remove and keep value and order of 3 values next to current cup
		var removed []int
		for j := 0; j < 3; j++ {
			val := ll.GetAt(currentCupInd + 1).value
			removed = append(removed, val)
			ll.DeleteAt(currentCupInd + 1)
		}
		fmt.Println("pick up:", removed)

		// select destination cup
		destinationVal := getNext(currentCup.value, labelsSorted)
		destinationPos := ll.Search(destinationVal)
		for destinationPos == -1 {
			destinationVal = getNext(destinationVal, labelsSorted)
			destinationPos = ll.Search(destinationVal)
		}
		fmt.Println("destination:", destinationVal, "\n")

		// place cups next to destination
		for ind, val := range removed {
			ll.InsertAt(destinationPos+ind+1, val)
		}
		currentCup = currentCup.next
	}

	return ll.GenerateIntReprStartingWithVal(1)
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

// GenerateIntReprStartingWithVal generates an int representation of the list that starts after val and goes through the entire list (except for val)
func (l *CircularLinkedList) GenerateIntReprStartingWithVal(val int) (repr int) {
	node := l.GetAt(l.Search(val))
	node = node.next
	for node.value != val {
		repr = 10*repr + node.value
		node = node.next
	}
	return repr
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
	// create a new node
	newNode := Node{}
	newNode.value = value
	// validate the position
	if pos < 0 {
		return
	}
	if pos == 0 {
		l.head = &newNode
		l.len++
		return
	}
	if pos > l.len {
		return
	}
	n := l.GetAt(pos)
	newNode.next = n
	prevNode := l.GetAt(pos - 1)
	prevNode.next = &newNode
	l.len++
}

// GetAt returns node at given position from linked list
func (l *CircularLinkedList) GetAt(pos int) *Node {
	ptr := l.head
	if pos < 0 || pos > (l.len-1) {
		pos = pos % l.len
	}
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
	// validate the position
	if pos < 0 {
		fmt.Println("position can not be negative")
		return errors.New("position can not be negative")
	}
	if l.len == 0 {
		fmt.Println("No nodes in list")
		return errors.New("No nodes in list")
	}
	prevNode := l.GetAt(pos - 1)
	if prevNode == nil {
		fmt.Println("Node not found")
		return errors.New("Node not found")
	}
	prevNode.next = l.GetAt(pos).next
	l.len--
	return nil
}

// DeleteVal deletes node having given value from linked list
func (l *CircularLinkedList) DeleteVal(val int) error {
	ptr := l.head
	if l.len == 0 {
		fmt.Println("List is empty")
		return errors.New("List is empty")
	}
	for i := 0; i < l.len; i++ {
		if ptr.value == val {
			if i > 0 {
				prevNode := l.GetAt(i - 1)
				prevNode.next = l.GetAt(i).next
			} else {
				l.head = ptr.next
			}
			l.len--
			return nil
		}
		ptr = ptr.next
	}
	fmt.Println("Node not found")
	return errors.New("Node not found")
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
