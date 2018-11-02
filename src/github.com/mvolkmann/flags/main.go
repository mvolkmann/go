package main

import (
	"flag"
	"fmt"
)

var minPtr = flag.Int("min", 1, "minimum value")
var maxPtr = flag.Int("max", 10, "maximum value")
var prefixPtr = flag.String("prefix", "", "prefix")

func main() {
	flag.Parse()
	prefix := *prefixPtr
	for i := *minPtr; i <= *maxPtr; i++ {
		fmt.Printf("%s%d\n", prefix, i)
	}
}
