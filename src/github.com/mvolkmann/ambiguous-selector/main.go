package main

import (
	"fmt"
)

type foo struct {
	f int
}

func (f foo) doIt() {
	fmt.Println("foo: doIt entered")
}

type bar struct {
	b bool
}

func (b bar) doIt() {
	fmt.Println("bar: doIt entered")
}

type baz struct {
	foo
	bar
}

func main() {
	myBaz := baz{}
	myBaz.doIt()
}
