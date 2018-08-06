package main

import "fmt"

type any interface{}
type callback func(any) any

func mapper(slice []any, fn callback) []any {
	result := make([]any, len(slice))
	for index, value := range slice {
		result[index] = fn(value)
	}
	return result
}

func anyTimes2(n any) any {
	var result any = n.(int) * 2
	return result
}
func intTimes2(n int) int { return n * 2 }

func main() {
	nums := []int{1, 2, 4}

	anys := make([]any, len(nums))
	for i, n := range nums {
		anys[i] = n
	}

	newAnys := mapper(anys, anyTimes2)
	fmt.Println(newAnys)

}
