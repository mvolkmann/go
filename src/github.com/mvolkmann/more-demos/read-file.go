package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// Read entire file into a newly created byte array.
	bytes, err := ioutil.ReadFile("haiku.txt")
	if err != nil {
    log.Fatal(err)
	} else {
    fmt.Println(string(bytes))
	}
}