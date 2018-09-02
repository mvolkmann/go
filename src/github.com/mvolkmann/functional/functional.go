package functional

import (
	"log"
	"reflect"
)

// FilterAny tries to work.
func FilterAny(slice interface{}, fn interface{}) interface{} {
	// Verify the first argument.
	sliceType := reflect.TypeOf(slice)
	if sliceType.Kind() != reflect.Slice {
		log.Fatal("FilterAny first argument must be a slice")
	}
	elementType := sliceType.Elem()

	// Verify the second argument.
	fnType := reflect.TypeOf(fn)
	if fnType.Kind() != reflect.Func {
		log.Fatal("FilterAny second argument must be a function")
	}
	if fnType.NumIn() != 1 {
		log.Fatal("FilterAny second argument must be a function that has one parameter")
	}
	if elementType.Kind() != fnType.In(0).Kind() {
		log.Fatal("FilterAny slice element type must match fn first parameter type")
	}
	if fnType.NumOut() != 1 {
		log.Fatal("FilterAny second argument must be a function that has one return type")
	}
	if fnType.Out(0).Kind() != reflect.Bool {
		log.Fatal("FilterAny second argument must be a function that returns a bool")
	}

	sliceValue := reflect.ValueOf(slice)
	fnValue := reflect.ValueOf(fn)

	result := reflect.New(sliceType).Elem()

	for i := 0; i < sliceValue.Len(); i++ {
		element := sliceValue.Index(i).Interface()
		elementValue := reflect.ValueOf(element)
		in := make([]reflect.Value, 1)
		in[0] = elementValue
		out := fnValue.Call(in)
		if out[0].Bool() == true {
			result = reflect.Append(result, elementValue)
		}
	}

	return result
}

// FilterInts takes a slice of int values and
// a function that takes that type and returns a boolean.
// It returns a new slice that only contains the slice elements
// for which the function returns true when passed the element.
func FilterInts(arr []int, fn func(int) bool) []int {
	result := make([]int, 0, len(arr))
	for _, v := range arr {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// MapAny tries to work.
func MapAny(arr []interface{}, fn func(interface{}) interface{}) []interface{} {
	result := make([]interface{}, len(arr))
	for i, v := range arr {
		result[i] = fn(v)
	}
	return result
}

// MapFloats takes a slice of float64 values and
// a function that takes that type and returns that type.
// It returns a new slice that is the result of
// applying the function to each element in the slice.
func MapFloats(arr []float64, fn func(float64) float64) []float64 {
	result := make([]float64, len(arr))
	for i, v := range arr {
		result[i] = fn(v)
	}
	return result
}

// MapInts takes a slice of int values and
// a function that takes that type and returns that type.
// It returns a new slice that is the result of
// applying the function to each element in the slice.
func MapInts(arr []int, fn func(int) int) []int {
	result := make([]int, len(arr))
	for i, v := range arr {
		result[i] = fn(v)
	}
	return result
}

// first returns the first value in slice of ints
// or zero if the slice is empty.
func first(values []int) int {
	if len(values) == 0 {
		return 0
	}
	return values[0]
}

// ReduceInts takes a slice of int values,
// a function that acts as an accumulator,
// and an initial value.
// It returns the accumulated result.
func ReduceInts(arr []int, fn func(int, int) int, initial ...int) int {
	result := first(initial)
	for _, v := range arr {
		result = fn(result, v)
	}
	return result
}
