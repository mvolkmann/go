package main

import "fmt"

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
}
