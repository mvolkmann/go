package main

import (
	"fmt"
	"strings"
)

func main() {
	names := []string{"Mark", "Tami", "Amanda", "Jeremy"}
	text := strings.Join(names, " - ")
	fmt.Printf("text = %+v\n", text)
}
