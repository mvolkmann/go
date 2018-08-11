package main

import "fmt"

type person struct {
	name string
	age  int8
}

func (p *person) birthday() {
	p.age++
}

type number int

func (receiver number) double() number {
	return receiver * 2
}

func main() {
	p := person{name: "Mark", age: 57}
	(&p).birthday()            // The method can be invoked on a pointer to a person.
	p.birthday()               // It can also be invoked on a person.
	fmt.Printf("p = %+v\n", p) // main.person{name:"Mark", age:58}

	n := number(3)
	// could also use var n number = 3
	fmt.Println("double result =", n.double()) // 6
}
