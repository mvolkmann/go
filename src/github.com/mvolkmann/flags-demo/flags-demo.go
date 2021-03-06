// To build this, enter "go build".
// To run this,
// enter ./flags -min 3 -max 5 -prefix foo
// or ./flags -min=3 -max=5 -prefix=foo
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
	//prefix := *prefixPtr
	for i := *minPtr; i <= *maxPtr; i++ {
		//fmt.Printf("%s%d\n", prefix, i)
		fmt.Printf("%s%d\n", *prefixPtr, i)
	}
}
