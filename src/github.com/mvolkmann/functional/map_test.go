package functional

import (
	"reflect"
	"testing"
)

func Test_Map(t *testing.T) {
	values := []int{1, 2, 7}
	// Create a new slice containing the values multiplied by 2.
	actual := Map(values, func(n int) int { return n * 2 })
	expected := []int{2, 4, 14}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Map expected %v but got %v", expected, actual)
	}
}

func BenchmarkMap(b *testing.B) {
	values := []int{1, 2, 7}
	fn := func(n int) int { return n * 2 }
	for i := 0; i < b.N; i++ {
		Map(values, fn)
	}
}
