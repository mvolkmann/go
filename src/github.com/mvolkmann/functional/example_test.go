package functional_test

import (
	"fmt"

	"github.com/mvolkmann/functional"
)

func ExampleFilter() {
	values := []int{1, 2, 4, 7, 10}
	fmt.Println(functional.Filter(values, func(n int) bool { return n%2 == 0 }))
	// Output:
	// [2 4 10]
}

func ExampleMap() {
	values := []int{1, 2, 7}
	fmt.Println(functional.Map(values, func(n int) int { return n * 2 }))
	// Output:
	// [2 4 14]
}

func ExampleReduce() {
	values := []int{1, 2, 7}
	fmt.Println(functional.Reduce(values, func(acc int, n int) int { return acc + n }, 0))
	// Output:
	// 10
}
