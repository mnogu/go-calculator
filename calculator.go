package calculator

func calculate(n *node) float64 {
	switch n.kind {
	case addNode:
		return calculate(n.left) + calculate(n.right)
	case subNode:
		return calculate(n.left) - calculate(n.right)
	case mulNode:
		return calculate(n.left) * calculate(n.right)
	case divNode:
		return calculate(n.left) / calculate(n.right)
	case numNode:
		return n.val
	}
	return 0
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
	return calculate(n), nil
}
