package main

import "fmt"

// LinkedList implements a linked list of integers.
//TODO: Make this work with any data type.
type LinkedList struct {
	head *node
}

type node struct {
	value int
	next  *node
}

func (listPtr *LinkedList) push(value int) {
	node := node{value, listPtr.head}
	listPtr.head = &node
}

func (listPtr *LinkedList) pop() int {
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
	list.push(5)
	list.push(7)
	value := list.pop()
	fmt.Println("value =", value)
	fmt.Println("list length =", list.len())
}
