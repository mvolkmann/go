package main

import (
	"reflect"
	"testing"
)

func Test_mapOverInts(t *testing.T) {
	ints := []int{1, 2, 4}
	expected := []int{2, 4, 8}
	actual := mapOverInts(ints, func(n int) int { return n * 2 })
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("mapOverInts expected %v but got %v", expected, actual)
	}
}

/*
func Test_log(t *testing.T) {
	type args struct {
		args []any
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log(tt.args.args...)
		})
	}
}

func Test_logInts(t *testing.T) {
	type args struct {
		args []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logInts(tt.args.args...)
		})
	}
}

func Test_demo(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for range tests {
		t.Run(tt.name, func(t *testing.T) {
			demo()
		})
	}
}

func Test_myAsync(t *testing.T) {
	type args struct {
		done chan<- bool
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			myAsync(tt.args.done)
		})
	}
}

func Test_selectDemo(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for range tests {
		t.Run(tt.name, func(t *testing.T) {
			selectDemo()
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
*/
