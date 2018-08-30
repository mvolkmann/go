package main

import (
	"fmt"
)

// LinkedList implements a linked list of integers.
//TODO: Make this work with any data type.
type LinkedList struct {
	head *node
}

type any interface{}

type node struct {
	value any
	next  *node
}

func (listPtr *LinkedList) clear() {
	listPtr.head = nil
}

func (listPtr *LinkedList) isEmpty() bool {
	return listPtr.head == nil
}

func (listPtr *LinkedList) push(value any) {
	node := node{value, listPtr.head}
	listPtr.head = &node
}

func (listPtr *LinkedList) pop() any {
	node := listPtr.head
	value := node.value
	listPtr.head = node.next
	return value
}

// Why can't the receiver name differ from the other methods
// and be a LinkedList instead of a pointer to one?
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
		sum += value.(int)
		fmt.Println("value =", value)
	}
	fmt.Println("sum =", sum)
}
