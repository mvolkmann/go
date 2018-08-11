package mapover

// IntToIntFn is the type of the functions
// passed to MapOverInts as the second argument.
type IntToIntFn = func(int) int

// Ints takes a slice of int values
// and a function that takes an int and returns an int.
// It returns a new slice that is the result of
// applying the function to each element in the slice.
func Ints(arr []int, fn IntToIntFn) []int {
	result := make([]int, len(arr))
	for i, v := range arr {
		result[i] = fn(v)
	}
	return result
}
