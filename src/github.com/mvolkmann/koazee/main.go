package main

import (
	"fmt"
	"log"

	"github.com/wesovilabs/koazee"
)

var numbers = []int{1, 2, 3, 4, 5}
var isOdd = func(value int) bool { return value%2 == 1 }
var double = func(value int) int { return value * 2 }

//TODO: Change the type of value to float64 to trigger an error.
var sum = func(acc, value int) int { return acc + value }

func main() {
	//logger.Enabled = true
	out := koazee.StreamOf(numbers).
		Filter(isOdd). // [1, 3, 5]
		Map(double).   // [2, 6, 10]
		Add(3).        // [2, 6, 10, 3]
		Reduce(sum)    // 21
	if out.Err() != nil {
		log.Fatal("got error:", out.Err())
		return
	}
	result := out.Int()
	fmt.Println("result =", result)
}
