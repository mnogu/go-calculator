package calculator

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculation(t *testing.T) {
	tests := []struct {
		in  string
		out float64
	}{
		{"1.2+2.3", 1.2 + 2.3},
		{"1.2-2.3", 1.2 - 2.3},
		{"2.5*3", 2.5 * 3},
		{"3.0/2.0", 3.0 / 2.0},
		{"-1.2+2.3", -1.2 + 2.3},
		{"+1.2+2.3", +1.2 + 2.3},
		{"1.2+2.5*2.0", 1.2 + 2.5*2.0},
		{"(1.2+2.5)*2.0", (1.2 + 2.5) * 2.0},
		{" (1.2 +  2.5 ) *   2.0  ", (1.2 + 2.5) * 2.0},

		{"e", math.E},
		{"E", math.E},

		{"pi*2.0", math.Pi * 2.0},
		{"Pi*2.0", math.Pi * 2.0},
		{"PI*2.0", math.Pi * 2.0},

		{"sqrt2", math.Sqrt2},
		{"sqrte", math.SqrtE},
		{"sqrtpi", math.SqrtPi},
		{"sqrtphi", math.SqrtPhi},

		{"ln2", math.Ln2},
		{"log2e", math.Log2E},
		{"ln2 * log2e", math.Ln2 * math.Log2E},
		{"ln10", math.Ln10},
		{"log10e", math.Log10E},
		{"ln10 * log10e", math.Ln10 * math.Log10E},

		{"abs(-1.5)", math.Abs(-1.5)},
		{"Abs(-1.5)", math.Abs(1.5)},
		{"ABS(-1.5)", math.Abs(1.5)},
		{"abs( (1.2 +  2.5 ) *   2.0  )", math.Abs((1.2 + 2.5) * 2.0)},

		{"Atan2(1.2,3.4)", math.Atan2(1.2, 3.4)},
		{"Atan2(1.2, 3.4)", math.Atan2(1.2, 3.4)},
		{"Atan2( (1.0 + 0.2) * 0.4, -1.7 * 2)", math.Atan2((1.0+0.2)*0.4, -1.7*2)},

		{"fma(1.2, 2.3, 4.5)", math.FMA(1.2, 2.3, 4.5)},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			val, err := Calculate(tt.in)
			if assert.NoError(t, err) {
				assert.InDelta(t, tt.out, val, 0.001)
			}
		})
	}
}

func TestNan(t *testing.T) {
	val, err := Calculate("nan()")
	if assert.NoError(t, err) {
		assert.True(t, math.IsNaN(val))
	}
}
