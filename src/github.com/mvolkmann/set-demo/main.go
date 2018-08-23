package main

import (
	"fmt"

	"github.com/mvolkmann/set"
)

func main() {
	colorSet := set.Strings{}
	colorSet.Add("red")
	colorSet.Add("blue")

	colors := []string{"red", "yellow", "blue"}
	for _, color := range colors {
		if colorSet.Contains(color) {
			fmt.Println("have", color) // prints "red" and "blue"
		}
	}

	fmt.Printf("%+v\n", colorSet) // map[red:{} yellow:{}]
}
