package calculator

import (
	"errors"
	"math"
	"strings"
)

type nodeKind string

const (
	addNode nodeKind = "+"
	subNode nodeKind = "-"
	mulNode nodeKind = "*"
	divNode nodeKind = "/"
	numNode nodeKind = "num"
)

type node struct {
	kind  nodeKind
	left  *node
	right *node
	val   float64
}

func numberNode(tokens []token, i *int) (*node, error) {
	t := tokens[*i]
	if t.kind != numberToken {
		return nil, errors.New("expected a number")
	}
	*i++
	return &node{kind: numNode, val: t.val}, nil
}

func constantNode(tokens []token, i *int) (*node, error) {
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
	val, ok := constants[strings.ToLower(tokens[*i].str)]
	if !ok {
		return nil, errors.New("unknown constant")
	}
	*i++
	return &node{kind: numNode, val: val}, nil
}

func consume(tokens []token, i *int, s string) bool {
	t := tokens[*i]
	if t.kind != reservedToken || t.str != s {
		return false
	}
	*i++
	return true
}

func parse(tokens []token) (*node, error) {
	i := 0
	return add(tokens, &i)

}

func insert(n *node, f func([]token, *int) (*node, error), tokens []token, i *int, kind nodeKind) (*node, error) {
	left := n
	right, err := f(tokens, i)
	if err != nil {
		return n, err
	}
	return &node{kind: kind, left: left, right: right}, err
}

func add(tokens []token, i *int) (*node, error) {
	n, err := mul(tokens, i)
	if err != nil {
		return nil, err
	}

	for *i < len(tokens) {
		if consume(tokens, i, "+") {
			n, err = insert(n, mul, tokens, i, addNode)
			if err != nil {
				return nil, err
			}
		} else if consume(tokens, i, "-") {
			n, err = insert(n, mul, tokens, i, subNode)
			if err != nil {
				return nil, err
			}
		} else {
			return n, nil
		}
	}
	return n, nil
}

func mul(tokens []token, i *int) (*node, error) {
	n, err := unary(tokens, i)
	if err != nil {
		return nil, err
	}

	for *i < len(tokens) {
		if consume(tokens, i, "*") {
			n, err = insert(n, unary, tokens, i, mulNode)
			if err != nil {
				return nil, err
			}
		} else if consume(tokens, i, "/") {
			n, err = insert(n, unary, tokens, i, divNode)
			if err != nil {
				return nil, err
			}
		} else {
			return n, nil
		}
	}
	return n, nil
}

func unary(tokens []token, i *int) (*node, error) {
	if consume(tokens, i, "+") {
		return primary(tokens, i)
	} else if consume(tokens, i, "-") {
		return insert(&node{kind: numNode, val: 0.0}, primary, tokens, i, subNode)
	}
	return primary(tokens, i)
}

func primary(tokens []token, i *int) (*node, error) {
	if consume(tokens, i, "(") {
		n, err := add(tokens, i)
		if err != nil {
			return nil, err
		}
		consume(tokens, i, ")")
		return n, nil

	}
	if tokens[*i].kind == identToken {
		n, err := constantNode(tokens, i)
		if err != nil {
			return nil, err
		}
		return n, nil
	}
	return numberNode(tokens, i)
}
