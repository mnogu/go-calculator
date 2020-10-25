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
		return nil, errors.New("expected a number")
	}
	p.i++
	return &node{kind: numNode, val: t.val}, nil
}

func (p *parser) constantNode() (*node, error) {
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
	val, ok := constants[strings.ToLower(p.tokens[p.i].str)]
	if !ok {
		return nil, errors.New("unknown constant")
	}
	p.i++
	return &node{kind: numNode, val: val}, nil
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
		n, err := p.constantNode()
		if err != nil {
			return nil, err
		}
		return n, nil
	}
	return p.numberNode()
}
