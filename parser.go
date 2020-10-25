package calculator

import (
	"fmt"
	"math"
	"strings"
)

type nodeKind string

const (
	addNode  nodeKind = "+"
	subNode  nodeKind = "-"
	mulNode  nodeKind = "*"
	divNode  nodeKind = "/"
	funcNode nodeKind = "func"
	numNode  nodeKind = "num"
)

type node struct {
	kind  nodeKind
	left  *node
	right *node

	funcName string
	args     []*node

	val float64
}

type parser struct {
	tokens []token
	i      int
}

func newParser(tokens []token) *parser {
	return &parser{tokens: tokens, i: 0}
}

func (p *parser) numberNode() (*node, error) {
	t := p.tokens[p.i]
	if t.kind != numberToken {
		return nil, fmt.Errorf("expected a number: %s", t.str)
	}
	p.i++
	return &node{kind: numNode, val: t.val}, nil
}

func (p *parser) constantNode(str string) (*node, error) {
	constants := map[string]float64{
		"e":   math.E,
		"pi":  math.Pi,
		"phi": math.Phi,

		"sqrt2":   math.Sqrt2,
		"sqrte":   math.SqrtE,
		"sqrtpi":  math.SqrtPi,
		"sqrtphi": math.SqrtPhi,

		"ln2":    math.Ln2,
		"log2e":  math.Log2E,
		"ln10":   math.Ln10,
		"log10e": math.Log10E,
	}
	val, ok := constants[strings.ToLower(str)]
	if !ok {
		return nil, fmt.Errorf("unknown constant: %s", str)
	}
	p.i++
	return &node{kind: numNode, val: val}, nil
}

func (p *parser) functionNode(str string) (*node, error) {
	functions := map[string]int{
		"abs":         1,
		"acos":        1,
		"acosh":       1,
		"asin":        1,
		"asinh":       1,
		"atan":        1,
		"atan2":       2,
		"atanh":       1,
		"cbrt":        1,
		"ceil":        1,
		"copysign":    2,
		"cos":         1,
		"cosh":        1,
		"dim":         2,
		"erf":         1,
		"erfc":        1,
		"erfcinv":     1,
		"erfinv":      1,
		"exp":         1,
		"exp2":        1,
		"expm1":       1,
		"fma":         3,
		"floor":       1,
		"gamma":       1,
		"hypot":       2,
		"j0":          1,
		"j1":          1,
		"log":         1,
		"log10":       1,
		"log1p":       1,
		"log2":        1,
		"logb":        1,
		"max":         2,
		"min":         2,
		"mod":         2,
		"nan":         0,
		"nextafter":   2,
		"pow":         2,
		"remainder":   2,
		"round":       1,
		"roundtoeven": 1,
		"sin":         1,
		"sinh":        1,
		"sqrt":        1,
		"tan":         1,
		"tanh":        1,
		"trunc":       1,
		"y0":          1,
		"y1":          1,
	}
	funcName := strings.ToLower(str)
	num, ok := functions[funcName]
	if !ok {
		return nil, fmt.Errorf("unknown function: %s", funcName)
	}
	if p.consume(")") {
		if num != 0 {
			return nil, fmt.Errorf("%s should have argument(s)", funcName)
		}
		return &node{kind: funcNode, funcName: funcName}, nil
	}

	args := []*node{}

	n, err := p.add()
	if err != nil {
		return nil, err
	}
	args = append(args, n)

	for p.consume(",") {
		n, err := p.add()
		if err != nil {
			return nil, err
		}
		args = append(args, n)
	}
	if len(args) != num {
		return nil, fmt.Errorf("%s should have %d argument(s) but has %d arguments(s)",
			funcName, num, len(args))
	}
	p.consume(")")
	return &node{kind: funcNode, funcName: funcName, args: args}, nil
}

func (p *parser) consume(s string) bool {
	t := p.tokens[p.i]
	if t.kind != reservedToken || t.str != s {
		return false
	}
	p.i++
	return true
}

func (p *parser) parse() (*node, error) {
	return p.add()

}

func (p *parser) insert(n *node, f func() (*node, error), kind nodeKind) (*node, error) {
	left := n
	right, err := f()
	if err != nil {
		return n, err
	}
	return &node{kind: kind, left: left, right: right}, err
}

func (p *parser) add() (*node, error) {
	n, err := p.mul()
	if err != nil {
		return nil, err
	}

	for p.i < len(p.tokens) {
		if p.consume("+") {
			n, err = p.insert(n, p.mul, addNode)
			if err != nil {
				return nil, err
			}
		} else if p.consume("-") {
			n, err = p.insert(n, p.mul, subNode)
			if err != nil {
				return nil, err
			}
		} else {
			return n, nil
		}
	}
	return n, nil
}

func (p *parser) mul() (*node, error) {
	n, err := p.unary()
	if err != nil {
		return nil, err
	}

	for p.i < len(p.tokens) {
		if p.consume("*") {
			n, err = p.insert(n, p.unary, mulNode)
			if err != nil {
				return nil, err
			}
		} else if p.consume("/") {
			n, err = p.insert(n, p.unary, divNode)
			if err != nil {
				return nil, err
			}
		} else {
			return n, nil
		}
	}
	return n, nil
}

func (p *parser) unary() (*node, error) {
	if p.consume("+") {
		return p.primary()
	} else if p.consume("-") {
		return p.insert(&node{kind: numNode, val: 0.0}, p.primary, subNode)
	}
	return p.primary()
}

func (p *parser) primary() (*node, error) {
	if p.consume("(") {
		n, err := p.add()
		if err != nil {
			return nil, err
		}
		p.consume(")")
		return n, nil
	}

	if p.tokens[p.i].kind == identToken {
		str := p.tokens[p.i].str
		p.i++
		if p.i < len(p.tokens) && p.consume("(") {
			return p.functionNode(str)
		}
		p.i--
		return p.constantNode(str)
	}
	return p.numberNode()
}
