package functional_test

import (
	"reflect"
	"testing"

	"github.com/mvolkmann/functional"
)

func Test_Filter(t *testing.T) {
	values := []int{1, 2, 4, 7, 10}
	// Create a new slice containing only the even values.
	actual := functional.Filter(values, func(n int) bool { return n%2 == 0 })
	expected := []int{2, 4, 10}
	// reflect.DeepEqual is the recommended way to compare many things in tests.
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Filter expected %v but got %v", expected, actual)
	}
}

func BenchmarkFilter(b *testing.B) {
	values := []int{1, 2, 7}
	predicate := func(n int) bool { return n%2 == 0 }
	for i := 0; i < b.N; i++ {
		functional.Filter(values, predicate)
	}
}
