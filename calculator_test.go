package calculator

import (
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
