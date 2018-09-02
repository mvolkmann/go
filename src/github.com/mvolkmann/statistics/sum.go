package statistics

// Sum returns the sum of slice of float64 values.
func Sum(numbers []float64) float64 {
	result := 0.0
	for _, number := range numbers {
		result += number
	}
	return result
}
