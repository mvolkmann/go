package mapover

// Floats takes a slice of float64 values and
// a function that takes that type and returns that type.
// It returns a new slice that is the result of
// applying the function to each element in the slice.
func Floats(arr []float64, fn func(float64) float64) []float64 {
	result := make([]float64, len(arr))
	for i, v := range arr {
		result[i] = fn(v)
	}
	return result
}

// Ints takes a slice of int values and
// a function that takes that type and returns that type.
// It returns a new slice that is the result of
// applying the function to each element in the slice.
func Ints(arr []int, fn func(int) int) []int {
	result := make([]int, len(arr))
	for i, v := range arr {
		result[i] = fn(v)
	}
	return result
}
