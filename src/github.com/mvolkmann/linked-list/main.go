package main

import (
	"fmt"
)

// LinkedList implements a linked list with values of any type.
type LinkedList struct {
	head *node
}

type any interface{}

type node struct {
	value any
	next  *node
}

// Clear removes all nodes.
func (listPtr *LinkedList) Clear() {
	listPtr.head = nil
}

// IsEmpty determines if the LinkedList is empty.
func (listPtr *LinkedList) IsEmpty() bool {
	return listPtr.head == nil
}

// Pop removes the first node and returns its value.
func (listPtr *LinkedList) Pop() any {
	if listPtr.IsEmpty() {
		return nil
	}
	node := listPtr.head
	listPtr.head = node.next
	return node.value
}

// Push adds a node to the front.
func (listPtr *LinkedList) Push(value any) {
	node := node{value, listPtr.head}
	listPtr.head = &node
}

// Len returns the length of the LinkedList.
func (listPtr *LinkedList) Len() int {
	len := 0
	node := listPtr.head
	for node != nil {
		len++
		node = node.next
	}
	return len
}

func main() {
	list := LinkedList{}
	list.Push(1)
	list.Push(3)
	list.Clear()
	list.Push(5)
	list.Push(7)

	sum := 0
	for !list.IsEmpty() {
		value := list.Pop()
		fmt.Println("value =", value)
		sum += value.(int)
	}
	fmt.Println("sum =", sum)
}
