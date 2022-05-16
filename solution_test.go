package square

import (
	"math"
	"testing"
)

func TestCalcSquare(t *testing.T) {
	type args struct {
		sideLen  float64
		sidesNum Shape
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"square", args{3, SidesSquare}, 9},
		{"circle", args{2, SidesCircle}, math.Pi * 4},
		{"triangle", args{5, SidesTriangle}, (math.Pow(5, 2) * math.Sqrt(3)) / 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcSquare(tt.args.sideLen, tt.args.sidesNum); got != tt.want {
				t.Errorf("CalcSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
