package functional

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_Filter(t *testing.T) {
	values := []int{1, 2, 4, 7, 10}
	// Create a new slice containing only the even values.
	actual := Filter(values, func(n int) bool { return n%2 == 0 })
	expected := []int{2, 4, 10}
	//TODO: Is this the best way to compare slices in a test?
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Filter expected %v but got %v", expected, actual)
	}
}

func Test_Map(t *testing.T) {
	values := []int{1, 2, 7}
	// Create a new slice containing the values multiplied by 2.
	actual := Map(values, func(n int) int { return n * 2 })
	expected := []int{2, 4, 14}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Map expected %v but got %v", expected, actual)
	}
}

func Test_Reduce(t *testing.T) {
	values := []int{1, 2, 7}
	// Sum the values.
	actual := Reduce(values, func(acc int, n int) int { return acc + n }, 0)
	expected := 10
	if actual != expected {
		t.Errorf("Reduce expected %v but got %v", expected, actual)
	}
}
func ExampleFilter() {
	values := []int{1, 2, 4, 7, 10}
	fmt.Println(Filter(values, func(n int) bool { return n%2 == 0 }))
	// Output:
	// [2 4 10]
}

func ExampleMap() {
	values := []int{1, 2, 7}
	fmt.Println(Map(values, func(n int) int { return n * 2 }))
	// Output:
	// [2 4 14]
}

func ExampleReduce() {
	values := []int{1, 2, 7}
	fmt.Println(Reduce(values, func(acc int, n int) int { return acc + n }, 0))
	// Output:
	// 10
}

func BenchmarkFilter(b *testing.B) {
	values := []int{1, 2, 7}
	predicate := func(n int) bool { return n%2 == 0 }
	for i := 0; i < b.N; i++ {
		Filter(values, predicate)
	}
}

func BenchmarkMap(b *testing.B) {
	values := []int{1, 2, 7}
	fn := func(n int) int { return n * 2 }
	for i := 0; i < b.N; i++ {
		Map(values, fn)
	}
}

func BenchmarkReduce(b *testing.B) {
	values := []int{1, 2, 7}
	fn := func(acc int, n int) int { return acc + n }
	for i := 0; i < b.N; i++ {
		Reduce(values, fn, 0)
	}
}
