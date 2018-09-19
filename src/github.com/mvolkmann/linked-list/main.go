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

// clear removes all nodes.
func (listPtr *LinkedList) clear() {
	listPtr.head = nil
}

// isEmpty determines if the LinkedList is empty.
func (listPtr *LinkedList) isEmpty() bool {
	return listPtr.head == nil
}

// pop removes the first node and returns its value.
func (listPtr *LinkedList) pop() any {
	if listPtr.isEmpty() {
		return nil
	}
	node := listPtr.head
	listPtr.head = node.next
	return node.value
}

// push adds a node to the front.
func (listPtr *LinkedList) push(value any) {
	node := node{value, listPtr.head}
	listPtr.head = &node
}

// len returns the length of the LinkedList.
func (listPtr *LinkedList) len() int {
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
	list.push(1)
	list.push(3)
	list.clear()
	list.push(5)
	list.push(7)

	sum := 0
	for !list.isEmpty() {
		value := list.pop()
		fmt.Println("value =", value)
		sum += value.(int)
	}
	fmt.Println("sum =", sum)
}
