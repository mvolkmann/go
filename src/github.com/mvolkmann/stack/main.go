package main

import "fmt"

// IntStack is a stack of int values.
type IntStack []int

func (s *IntStack) push(value int) {
	*s = append(*s, value)
}

func (s *IntStack) pop() int {
	stack := *s
	l := len(stack)
	top := stack[l-1]
	*s = stack[:l-1]
	return top
}

func main() {
	myStack := IntStack{1, 3, 5}
	fmt.Println(myStack.pop())             // 5, now contains [1, 3]
	fmt.Println(myStack.pop())             // 3, now contains [1]
	myStack.push(7)                        // now contains [1, 7]
	myStack.push(9)                        // now contains [1, 7, 9]
	fmt.Printf("myStack = %+v\n", myStack) // [1 7 9]
}
