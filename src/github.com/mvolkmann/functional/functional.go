package functional

import (
	"fmt"
	"reflect"
)

func panicF(format string, values ...interface{}) {
	msg := fmt.Sprintf(format, values...)
	panic(msg)
}

// assertFunc asserts the parameter and return types of a given function.
// Note that this does not currently detect mismatched struct types.
// It just verifies that some kind of struct is received if one is required.
//TODO: Define an assertStruct function that verifies all the struct fields!
func assertFunc(fn interface{}, in []reflect.Kind, out []reflect.Kind) {
	assertKind(fn, reflect.Func)

	fnType := reflect.TypeOf(fn)

	actualNumIn := fnType.NumIn()
	expectedNumIn := len(in)
	if actualNumIn != expectedNumIn {
		panicF("expected func with %d parameters but had %d\n", expectedNumIn, actualNumIn)
	}

	actualNumOut := fnType.NumOut()
	expectedNumOut := len(out)
	if actualNumOut != expectedNumOut {
		panicF("expected func with %d return types but had %d\n", expectedNumOut, actualNumOut)
	}

	for i := 0; i < expectedNumIn; i++ {
		expectedKind := in[i]
		actualKind := fnType.In(i).Kind()
		if actualKind != expectedKind {
			panicF("expected parameter %d to have kind %s but was %s\n", i+1, expectedKind, actualKind)
		}
	}

	for i := 0; i < expectedNumOut; i++ {
		expectedKind := out[i]
		actualKind := fnType.Out(i).Kind()
		if actualKind != expectedKind {
			panicF("expected result type %d to have kind %s but was %s\n", i+1, expectedKind, actualKind)
		}
	}
}

// assertKind asserts the kind of a given value.
func assertKind(value interface{}, kind reflect.Kind) {
	valueType := reflect.TypeOf(value)
	valueKind := valueType.Kind()
	if valueKind != kind {
		panicF("expected %s value but got %s\n", kind, valueKind)
	}
}

// Filter creates a new slice from the elements in an existing slice
// that pass a given predicate function.
func Filter(slice interface{}, predicate interface{}) interface{} {
	assertKind(slice, reflect.Slice)
	sliceType := reflect.TypeOf(slice)
	elementKind := sliceType.Elem().Kind()
	assertFunc(predicate, []reflect.Kind{elementKind}, []reflect.Kind{reflect.Bool})

	// Create result slice with same type as first argument.
	result := reflect.New(sliceType).Elem()

	predicateValue := reflect.ValueOf(predicate)
	sliceValue := reflect.ValueOf(slice)

	// Can "range" be used here?
	for i := 0; i < sliceValue.Len(); i++ {
		element := sliceValue.Index(i).Interface()
		elementValue := reflect.ValueOf(element)

		in := make([]reflect.Value, 1)
		in[0] = elementValue
		out := predicateValue.Call(in)

		if out[0].Bool() {
			result = reflect.Append(result, elementValue)
		}
	}

	return result.Interface()
}

// Map creates a new slice from the elements in an existing slice where
// the new elements are the results of calling fn on each existing element.
func Map(slice interface{}, fn interface{}) interface{} {
	assertKind(slice, reflect.Slice)
	sliceType := reflect.TypeOf(slice)
	elementKind := sliceType.Elem().Kind()
	assertFunc(fn, []reflect.Kind{elementKind}, []reflect.Kind{elementKind})

	// Create result slice with same type as first argument.
	result := reflect.New(sliceType).Elem()

	fnValue := reflect.ValueOf(fn)
	sliceValue := reflect.ValueOf(slice)

	// Can "range" be used here?
	for i := 0; i < sliceValue.Len(); i++ {
		element := sliceValue.Index(i).Interface()
		elementValue := reflect.ValueOf(element)

		in := make([]reflect.Value, 1)
		in[0] = elementValue
		out := fnValue.Call(in)
		result = reflect.Append(result, out[0])
	}

	return result.Interface()
}

// Reduce reduces the elements in a slice to a single value.
func Reduce(slice interface{}, fn interface{}, initial interface{}) interface{} {
	assertKind(slice, reflect.Slice)

	initialKind := reflect.TypeOf(initial).Kind()

	sliceType := reflect.TypeOf(slice)
	elementKind := sliceType.Elem().Kind()
	assertFunc(fn, []reflect.Kind{initialKind, elementKind}, []reflect.Kind{elementKind})

	// Create result slice with same type as first argument.
	result := reflect.ValueOf(initial)

	fnValue := reflect.ValueOf(fn)
	sliceValue := reflect.ValueOf(slice)

	for i := 0; i < sliceValue.Len(); i++ {
		element := sliceValue.Index(i).Interface()
		elementValue := reflect.ValueOf(element)

		in := make([]reflect.Value, 2)
		in[0] = result
		in[1] = elementValue
		out := fnValue.Call(in)
		result = out[0]
	}

	return result.Interface()
}

/*
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
*/
