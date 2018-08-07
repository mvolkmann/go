package main

import "fmt"

type person struct {
	name string
	age int8
}

func (p person) String() string {
	return fmt.Sprintf("%v is %v years old.", p.name, p.age)
}

func main() {
	me := person{"Mark", 57}
	fmt.Println(me)
}