package main

import "fmt"

type stringToString func(string) string

func main() {
	var ptr *string         // pointer
	var c chan string       // channel
	var f stringToString    // function
	var i error             // interface
	var m map[string]string // map
	var s []string          // slice

	if ptr == nil {
		fmt.Println("pointer is nil")
	}
	if c == nil {
		fmt.Println("channel is nil")
	}
	if f == nil {
		fmt.Println("function is nil")
	}
	if i == nil {
		fmt.Println("interface is nil")
	}
	if m == nil {
		fmt.Println("map is nil")
	}
	if s == nil {
		fmt.Println("s is nil")
	}
	if ptr == nil {
		fmt.Println("pointer is nil")
	}
}
