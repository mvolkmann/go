package statistics

import (
	"fmt"
	"testing"
)

func TestMax(t *testing.T) {
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
		{"multiple values", args{[]float64{3.1, 7.2, 5.0}}, 7.2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.numbers); got != tt.want {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleMax() {
	fmt.Println(Max([]float64{}))
	fmt.Println(Max([]float64{3.1}))
	fmt.Println(Max([]float64{3.1, 7.2, 5.0}))
	// Output:
	// 0
	// 3.1
	// 7.2
}

func BenchmarkMax(b *testing.B) {
	values := []float64{3.1, 7.2, 5.0}

	for i := 0; i < b.N; i++ {
		Max(values)
	}
}
