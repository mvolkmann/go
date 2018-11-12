package main

import (
	"fmt"

	"github.com/wesovilabs/koazee"
)

var numbers = []int{1, 2, 3, 4, 5}
var isOdd = func(value int) bool { return value%2 == 1 }
var double = func(value int) int { return value * 2 }

//TODO: Change the type of value to float64 to trigger an error.
var sum = func(acc, value int) int { return acc + value }

func main() {
	//logger.Enabled = true
	result := koazee.StreamOf(numbers).
		Filter(isOdd). // [1, 3, 5]
		Map(double).   // [2, 6, 10]
		Add(3).        // [2, 6, 10, 3]
		Reduce(sum)    // 21
	//Int()          // 21
	fmt.Printf("err = %v\n", result.Err())
	fmt.Printf("val = %v\n", result.Val())
	fmt.Printf("result = %#v\n", result) // 21
	/*
		var foo interface{} = 1
		var bar interface{} = 1
		fmt.Println(foo.Int() + bar.Int())
	*/
}
