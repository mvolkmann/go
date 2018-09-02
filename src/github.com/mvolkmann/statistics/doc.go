// Package statistics provides functions that compute statistical values
// for a slice of float64 values.
package statistics

import "math"

// This function is used by multiple tests.
func close(n1 float64, n2 float64) bool {
	return math.Abs(n1-n2) < 1e-9
}
