package calculator

import (
	"fmt"
	"math"
)

var functions = map[string]interface{}{
	"abs":         math.Abs,
	"acos":        math.Acos,
	"acosh":       math.Acosh,
	"asin":        math.Asin,
	"asinh":       math.Asinh,
	"atan":        math.Atan,
	"atan2":       math.Atan2,
	"atanh":       math.Atanh,
	"cbrt":        math.Cbrt,
	"ceil":        math.Ceil,
	"copysign":    math.Copysign,
	"cos":         math.Cos,
	"cosh":        math.Cosh,
	"dim":         math.Dim,
	"erf":         math.Erf,
	"erfc":        math.Erfc,
	"erfcinv":     math.Erfcinv, // Go 1.10+
	"erfinv":      math.Erfinv,  // Go 1.10+
	"exp":         math.Exp,
	"exp2":        math.Exp2,
	"expm1":       math.Expm1,
	"fma":         math.FMA, // Go 1.14+
	"floor":       math.Floor,
	"gamma":       math.Gamma,
	"hypot":       math.Hypot,
	"j0":          math.J0,
	"j1":          math.J1,
	"log":         math.Log,
	"log10":       math.Log10,
	"log1p":       math.Log1p,
	"log2":        math.Log2,
	"logb":        math.Logb,
	"max":         math.Max,
	"min":         math.Min,
	"mod":         math.Mod,
	"nan":         math.NaN,
	"nextafter":   math.Nextafter,
	"pow":         math.Pow,
	"remainder":   math.Remainder,
	"round":       math.Round,       // Go 1.10+
	"roundtoeven": math.RoundToEven, // Go 1.10+
	"sin":         math.Sin,
	"sinh":        math.Sinh,
	"sqrt":        math.Sqrt,
	"tan":         math.Tan,
	"tanh":        math.Tanh,
	"trunc":       math.Trunc,
	"y0":          math.Y0,
	"y1":          math.Y1,
}

func call(funcName string, args []float64) (float64, error) {
	f, ok := functions[funcName]
	if !ok {
		return 0, fmt.Errorf("unknown function %s", funcName)
	}
	switch f := f.(type) {
	case func() float64:
		return f(), nil
	case func(float64) float64:
		return f(args[0]), nil
	case func(float64, float64) float64:
		return f(args[0], args[1]), nil
	case func(float64, float64, float64) float64:
		return f(args[0], args[1], args[2]), nil
	default:
		return 0, fmt.Errorf("invalid function %s", funcName)
	}
}

func calculate(n *node) (float64, error) {
	switch n.kind {
	case addNode:
		left, err := calculate(n.left)
		if err != nil {
			return 0, err
		}
		right, err := calculate(n.right)
		if err != nil {
			return 0, err
		}
		return left + right, nil
	case subNode:
		left, err := calculate(n.left)
		if err != nil {
			return 0, err
		}
		right, err := calculate(n.right)
		if err != nil {
			return 0, err
		}
		return left - right, nil
	case mulNode:
		left, err := calculate(n.left)
		if err != nil {
			return 0, err
		}
		right, err := calculate(n.right)
		if err != nil {
			return 0, err
		}
		return left * right, nil
	case divNode:
		left, err := calculate(n.left)
		if err != nil {
			return 0, err
		}
		right, err := calculate(n.right)
		if err != nil {
			return 0, err
		}
		return left / right, nil
	case numNode:
		return n.val, nil
	case funcNode:
		args := []float64{}
		for _, arg := range n.args {
			val, err := calculate(arg)
			if err != nil {
				return 0, err
			}
			args = append(args, val)
		}
		return call(n.funcName, args)
	}
	return 0, fmt.Errorf("unknown node type: %s", n.kind)
}

// Calculate calculates expr
func Calculate(expr string) (float64, error) {
	tokens, err := tokenize(expr)
	if err != nil {
		return 0, err
	}
	p := newParser(tokens)
	n, err := p.parse()
	if err != nil {
		return 0, err
	}
	return calculate(n)
}
