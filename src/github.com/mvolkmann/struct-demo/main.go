package main

import (
	"fmt"
	"time"
)

type address struct {
	street string
	city   string
	state  string
	zip    string
}

type person struct {
	name        string
	address     // home
	workAddress address
}

func unnamedFieldDemo() {
	type age int
	type myType struct {
		name       string // named field
		age               // get field name from a primitive type
		time.Month        // get field name from a library type
	}
	//myStruct := myType{name: "Mark", int: 7, Month: time.April}
	myStruct := myType{"Mark", 57, time.April}
	fmt.Printf("name = %v\n", myStruct.name)
	fmt.Printf("age = %v\n", myStruct.age)
	fmt.Printf("Month = %v\n", myStruct.Month)
}

func main() {
	me := person{
		name: "Mark Volkmann",
		address: address{
			street: "123 Some Street",
			city:   "St. Charles",
			state:  "MO",
			zip:    "63304",
		},
		workAddress: address{
			street: "123 Woodcrest Executive Drive",
			city:   "Creve Coeur",
			state:  "MO",
			zip:    "63141",
		},
	}
	fmt.Printf("me = %+v\n", me)

	zip := me.address.zip
	fmt.Println("zip =", zip)

	unnamedFieldDemo()
}
