package main

import (
	"fmt"
	"strings"
)

func main() {
	s := [...]string{"foo", "bar", "baz"}
	s2 := strings.Join(s[:], "/")
	fmt.Printf("s2 = %+v\n", s2)
}
