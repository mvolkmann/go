package functional

import (
	"fmt"
	"reflect"
)

var boolType = reflect.TypeOf(false)

func panicF(format string, values ...interface{}) {
	msg := fmt.Sprintf(format, values...)
	panic(msg)
}

// AssertFunc asserts that a given value is a func and
// the parameter and return types are the expected types.
func AssertFunc(fn interface{}, in []reflect.Type, out []reflect.Type) {
	AssertKind(fn, reflect.Func)

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
		expectedType := in[i]
		actualType := fnType.In(i)
		if actualType != expectedType {
			panicF("expected parameter %d to have type %s but was %s\n", i+1, expectedType, actualType)
		}
	}

	for i := 0; i < expectedNumOut; i++ {
		expectedType := out[i]
		actualType := fnType.Out(i)
		if actualType != expectedType {
			panicF("expected result type %d to have type %s but was %s\n", i+1, expectedType, actualType)
		}
	}
}

// AssertKind asserts the kind of a given value.
func AssertKind(value interface{}, kind reflect.Kind) {
	valueKind := reflect.TypeOf(value).Kind()
	if valueKind != kind {
		panicF("expected kind %s but got %s\n", kind, valueKind)
	}
}

// AssertType asserts the type of a given value.
// Note that this does not currently detect mismatched struct types.
// It just verifies that some kind of struct is received if one is required.
//TODO: Define an assertStruct function that verifies all the struct fields!
func AssertType(value interface{}, typ reflect.Type) {
	valueType := reflect.TypeOf(value)
	if valueType != typ {
		panicF("expected type %s but got %s\n", typ, valueType)
	}
}

// Filter creates a new slice from the elements in an existing slice
// that pass a given predicate function.
func Filter(slice interface{}, predicate interface{}) interface{} {
	AssertKind(slice, reflect.Slice)
	sliceType := reflect.TypeOf(slice)
	elementType := sliceType.Elem()
	AssertFunc(predicate, []reflect.Type{elementType}, []reflect.Type{boolType})

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
	AssertKind(slice, reflect.Slice)
	sliceType := reflect.TypeOf(slice)
	elementType := sliceType.Elem()
	AssertFunc(fn, []reflect.Type{elementType}, []reflect.Type{elementType})

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
	AssertKind(slice, reflect.Slice)

	initialType := reflect.TypeOf(initial)

	sliceType := reflect.TypeOf(slice)
	elementType := sliceType.Elem()
	AssertFunc(fn, []reflect.Type{initialType, elementType}, []reflect.Type{elementType})

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
