package calculator

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

type tokenKind string

const (
	reservedToken tokenKind = "reserved"
	numberToken   tokenKind = "number"
	eosToken      tokenKind = "eos"
)

type token struct {
	kind tokenKind
	val  float64
	str  string
}

type invalidTokenError struct {
	input    string
	position int
}

func (e *invalidTokenError) Error() string {
	curr := ""
	pos := e.position
	for _, line := range strings.Split(e.input, "\n") {
		len := len(line)
		curr += line + "\n"
		if pos < len {
			return curr + strings.Repeat(" ", pos) + "^ invalid token"
		}
		pos -= len + 1
	}
	return ""
}

const operators = "+-*/()"

func isOperator(char rune) bool {
	for _, op := range operators {
		if char == op {
			return true
		}
	}
	return false
}

func numberPrefix(runes []rune, i *int, n int) (float64, error) {
	val := 0.0
	len := 0
	for *i < n {
		curr, err := strconv.ParseFloat(string(runes[*i-len:*i+1]), 64)
		if err != nil {
			break
		}
		val = curr
		len++
		*i++
	}
	if len > 0 {
		return val, nil
	}
	return 0, errors.New("expected a number")
}

func tokenize(input string) ([]token, error) {
	runes := []rune(input)
	i := 0
	n := len(runes)
	tokens := []token{}
	for i < n {
		char := runes[i]
		if unicode.IsSpace(char) {
			i++
			continue
		}

		if isOperator(char) {
			tokens = append(tokens, token{kind: reservedToken, str: string(char)})
			i++
			continue
		}

		if val, err := numberPrefix(runes, &i, n); err == nil {
			tokens = append(tokens, token{kind: numberToken, val: val})
			continue
		}

		return nil, &invalidTokenError{input: input, position: i}
	}
	return tokens, nil
}
