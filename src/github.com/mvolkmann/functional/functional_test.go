package mapover

import (
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

func Test_FilterInts(t *testing.T) {
	values := []int{1, 2, 4, 7, 10}
	expected := []int{2, 4, 10}
	actual := FilterInts(values, func(n int) bool { return n%2 == 0 })
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("FilterInts expected %v but got %v", expected, actual)
	}
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