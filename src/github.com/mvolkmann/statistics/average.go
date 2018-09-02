package statistics

// Avg returns the average of slice of float64 values.
func Avg(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}
	return Sum(numbers) / float64(len(numbers))
}
