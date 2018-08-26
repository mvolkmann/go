package main

import (
	"fmt"

	"github.com/mvolkmann/set"
)

func main() {
	colorSet := set.Strings{}
	colorSet.Add("red")
	colorSet.Add("yellow")
	colorSet.Add("blue")
	colorSet.Remove("yellow")

	colors := []string{"red", "yellow", "blue"}
	for _, color := range colors {
		if colorSet.Contains(color) {
			fmt.Println("have", color) // prints "red" and "blue"
		}
	}

	fmt.Printf("%+v\n", colorSet) // map[red:{} blue:{}]
	fmt.Println("length =", len(colorSet))
}
