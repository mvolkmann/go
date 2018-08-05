package main

import "fmt"

type person struct {
	name string
	age  int8
}

func (p *person) birthday() {
	p.age++
}

func main() {
	p := person{name: "Mark", age: 57}
	(&p).birthday()        // invoked on a pointer to a struct
	p.birthday()           // invoked on a struct
	fmt.Printf("%#v\n", p) // main.Person{name:"Mark", age:58}
}
