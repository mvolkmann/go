package main

import (
	"fmt"
	"reflect"
)

type any interface{}
type callback func(any) any

func mapper(slice []any, fn callback) []any {
	result := make([]any, len(slice))
	for index, value := range slice {
		result[index] = fn(value)
	}
	return result
}

/*
func anyTimes2(n any) any {
	var result any
	switch value := n.(type) {
	//case int, int32, int64, float32, float64:
	case int, float64:
		result = value * 2
	case float32:
		result = value * 2
	}
	//var result any = n.(t) * 2
	return result
}
*/

func anyTimes2(n any) any {
	kind := reflect.TypeOf(n)
	var result any = n.(kind) * 2
	return result
}

func main() {
	nums := []int{1, 2, 4}

	anys := make([]any, len(nums))
	for i, n := range nums {
		anys[i] = n
	}

	newAnys := mapper(anys, anyTimes2)
	fmt.Println(newAnys)

}
