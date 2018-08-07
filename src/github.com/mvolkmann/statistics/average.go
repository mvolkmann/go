package statistics

func sum(numbers []float64) float64 {
	result := 0.0
	for _, number := range numbers {
		result += number
	}
	return result
}

// Avg returns the average of slice of float64 values.
func Avg(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}
	return sum(numbers) / float64(len(numbers))
}
