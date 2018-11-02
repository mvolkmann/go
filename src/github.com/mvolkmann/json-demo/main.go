package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Person describes a person.
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// Person2 describes a person and includes field tags.
type Person2 struct {
	FirstName string `json:"name"`
	LastName  string `json:"surname"`
	Age       int    `json:"age,omitempty"`
}

func main() {
	p := Person{FirstName: "Mark", LastName: "Volkmann"}
	json1, err := json.Marshal(p)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Println(string(json1)) // {"FirstName":"Mark","LastName":"Volkmann","Age":0}

	people := []Person{
		Person{FirstName: "Mark", LastName: "Volkmann", Age: 57},
		Person{FirstName: "Tami", LastName: "Volkmann"},
	}
	json2, err := json.Marshal(people)
	// skipping err check
	fmt.Println(string(json2))
	// [{"FirstName":"Mark","LastName":"Volkmann","Age":57},{"FirstName":"Tami","LastName":"Volkmann","Age":0}]

	p2 := Person2{FirstName: "Mark", LastName: "Volkmann"}
	json3, err := json.Marshal(p2)
	// skipping err check
	fmt.Println(string(json3)) // {"name":"Mark","surname":"Volkmann"}

	var p3 Person
	err = json.Unmarshal(json1, &p3)
	// skipping err check
	fmt.Printf("%+v\n", p3) // {FirstName:Mark LastName:Volkmann Age:0}

	type PersonSubset struct {
		LastName string
	}
	var pSubset PersonSubset
	err = json.Unmarshal(json1, &pSubset)
	// skipping err check
	fmt.Printf("%+v\n", pSubset) // {LastName:Volkmann}

	type MyMap map[string]interface{}
	mySlice := []MyMap{}
	err = json.Unmarshal(json2, &mySlice)
	// skipping err check
	fmt.Printf("myMap = %+v\n", mySlice)
	// [map[FirstName:Mark LastName:Volkmann Age:57] map[Age:0 FirstName:Tami LastName:Volkmann]]
}
