package statistics

import (
	"testing"
)

func TestAvg(t *testing.T) {
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
		{"multiple values", args{[]float64{3.1, 7.2, 5.0}}, 5.1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Avg(tt.args.numbers); !close(got, tt.want) {
				t.Errorf("Avg() = %v, want %v", got, tt.want)
			}
		})
	}
}
