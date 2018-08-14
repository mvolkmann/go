package mapover

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
