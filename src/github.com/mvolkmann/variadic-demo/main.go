package main

import (
	"fmt"
	"strings"
)

func log(args ...interface{}) {
	fmt.Println(args...)
}

func report(name string, colors ...string) {
	text := strings.Join(colors, " and ") + "."
	fmt.Println(name, "likes the colors", text)
}

func main() {
	log("red", 7, true)
	report("Mark", "yellow", "orange")
}
