package statistics

// Max returns the maximum of slice of float64 values.
func Max(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}

	max := numbers[0]
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}

	return max
}
