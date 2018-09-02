package functional

import (
	"fmt"
	"reflect"
	"testing"
)

/* This demonstrates how to add setup and teardown steps.
func TestMain(m *testing.M) {
	fmt.Println("doing setup")
	result := m.Run()
	fmt.Println("doing teardown")
	os.Exit(result)
}
*/

func Test_Filter(t *testing.T) {
	values := []int{1, 2, 4, 7, 10}
	predicate := func(n int) bool { return n%2 == 0 }
	actual := Filter(values, predicate)
	expected := []int{2, 4, 10}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Filter expected %v but got %v", expected, actual)
	}
}

func Test_Map(t *testing.T) {
	values := []int{1, 2, 7}
	fn := func(n int) int { return n * 2 }
	actual := Map(values, fn)
	expected := []int{2, 4, 14}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Map expected %v but got %v", expected, actual)
	}
}

func Test_Reduce(t *testing.T) {
	values := []int{1, 2, 7}
	fn := func(acc int, n int) int { return acc + n }
	actual := Reduce(values, fn, 0)
	expected := 10
	if actual != expected {
		t.Errorf("Reduce expected %v but got %v", expected, actual)
	}
}

func Test_FilterInts(t *testing.T) {
	values := []int{1, 2, 4, 7, 10}
	expected := []int{2, 4, 10}
	actual := FilterInts(values, func(n int) bool { return n%2 == 0 })
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("FilterInts expected %v but got %v", expected, actual)
	}
}

func ExampleFilterInts() {
	values := []int{1, 2, 4, 7, 10}
	fmt.Println(FilterInts(values, func(n int) bool { return n%2 == 0 }))
	// Output:
	// [2 4 10]
}

func Test_MapFloats(t *testing.T) {
	values := []float64{1.2, 2.3, 4.5}
	expected := []float64{2.4, 4.6, 9.0}
	actual := MapFloats(values, func(n float64) float64 { return n * 2.0 })
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("MapFloats expected %v but got %v", expected, actual)
	}
}

func Test_MapInts(t *testing.T) {
	values := []int{1, 2, 4}
	expected := []int{2, 4, 8}
	actual := MapInts(values, func(n int) int { return n * 2 })
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("MapInts expected %v but got %v", expected, actual)
	}
}

func Test_ReduceInts(t *testing.T) {
	values := []int{1, 2, 4, 7, 10}
	expected := 24
	actual := ReduceInts(values, func(acc int, n int) int { return acc + n })
	if actual != expected {
		t.Errorf("ReduceInts expected %v but got %v", expected, actual)
	}
}
