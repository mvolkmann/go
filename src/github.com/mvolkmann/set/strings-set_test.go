package set

import (
	"fmt"
	"testing"
)

const (
	red    = "red"
	yellow = "yellow"
	blue   = "blue"
)

func TestSet(t *testing.T) {
	colorSet := Strings{}

	colorSet.Add(red)
	if len(colorSet) != 1 {
		t.Errorf("Add did not add correctly")
	}

	if !colorSet.Contains(red) {
		t.Errorf("Add or Contains failed")
	}

	colorSet.Remove(red)

	if len(colorSet) != 0 {
		t.Errorf("Remove did not remove correctly")
	}

	if colorSet.Contains(red) {
		t.Errorf("Remove or Contains failed")
	}

}

func ExampleStrings() {
	colorSet := Strings{}

	colorSet.Add(red)
	colorSet.Add(yellow)
	colorSet.Add(blue)

	colorSet.Remove(yellow)

	testColorSet := []string{red, yellow, blue}

	for _, color := range testColorSet {
		if colorSet.Contains(color) {
			fmt.Println("have", color)
		}
	}

	// hashes are not deterministic	fmt.Printf("%+v\n", colorSet)
	fmt.Println("length =", len(colorSet))
	// Output:
	// have red
	// have blue
	// length = 2
}
