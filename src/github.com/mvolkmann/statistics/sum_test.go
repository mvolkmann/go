package statistics

import (
	"fmt"
	"testing"
)

func Test_sum(t *testing.T) {
	type args struct {
		numbers []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"empty slice", args{[]float64{}}, 0.0},
		{"single value", args{[]float64{3.1}}, 3.1},
		{"multiple values", args{[]float64{3.1, 7.2, 5.0}}, 15.3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.numbers); !close(got, tt.want) {
				t.Errorf("sum() = %v, want %v", got, tt.want)
			}
		})
	}

}
func ExampleSum() {
	nums := []float64{1, 2, 3}
	fmt.Println("sum is", Sum(nums))
	// Output:
	// sum is 6
}
