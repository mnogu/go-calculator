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
		{"ln10", math.Ln10},
		{"log10e", math.Log10E},
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
