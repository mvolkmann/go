package mapover

import (
	"reflect"
	"testing"
)

func Test_Ints(t *testing.T) {
	ints := []int{1, 2, 4}
	expected := []int{2, 4, 8}
	actual := Ints(ints, func(n int) int { return n * 2 })
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("mapOverInts expected %v but got %v", expected, actual)
	}
}
