package mapover

import (
	"reflect"
	"testing"
)

func Test_Floats(t *testing.T) {
	values := []float64{1.2, 2.3, 4.5}
	expected := []float64{2.4, 4.6, 9.0}
	actual := Floats(values, func(n float64) float64 { return n * 2.0 })
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("mapOverInts expected %v but got %v", expected, actual)
	}
}

func Test_Ints(t *testing.T) {
	values := []int{1, 2, 4}
	expected := []int{2, 4, 8}
	actual := Ints(values, func(n int) int { return n * 2 })
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("mapOverInts expected %v but got %v", expected, actual)
	}
}
