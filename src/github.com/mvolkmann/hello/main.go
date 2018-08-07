package main

import (
	"fmt"

	"github.com/mvolkmann/statistics"
)

func main() {
	fmt.Println("Hello, World!")

	numbers := []float64{1, 2, 7}
	fmt.Println("maximum =", statistics.Max(numbers))
	fmt.Println("average =", statistics.Avg(numbers))
}
