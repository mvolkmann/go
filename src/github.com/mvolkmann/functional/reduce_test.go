package functional_test

import (
	"testing"

	"github.com/mvolkmann/functional"
)

func Test_Reduce(t *testing.T) {
	values := []int{1, 2, 7}
	// Sum the values.
	actual := functional.Reduce(values, func(acc int, n int) int { return acc + n }, 0)
	expected := 10
	if actual != expected {
		t.Errorf("Reduce expected %v but got %v", expected, actual)
	}
}
func BenchmarkReduce(b *testing.B) {
	values := []int{1, 2, 7}
	fn := func(acc int, n int) int { return acc + n }
	for i := 0; i < b.N; i++ {
		functional.Reduce(values, fn, 0)
	}
}
