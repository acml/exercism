package wordy

import (
	"strconv"
	"strings"
)

type state int

const (
	initial state = iota
	getOperand
	getOperation
)

// Answer parses and evaluate simple math word problems returning the answer as
// an integer.
func Answer(input string) (int, bool) {
	if !strings.HasPrefix(input, "What is") {
		return 0, false
	}

	op, res, operation := initial, 0, ""
	s := strings.Split(strings.TrimSuffix(input, "?"), " ")
	for i := 2; i < len(s); {
		switch op {
		case initial:
			fallthrough
		case getOperand:
			n, _ := strconv.Atoi(s[i])
			i++
			switch operation {
			case "plus":
				res += n
			case "minus":
				res -= n
			case "divided":
				res /= n
			case "multiplied":
				res *= n
			default:
				res = n
			}
			operation = ""
			op = getOperation
		case getOperation:
			operation = s[i]
			switch s[i] {
			case "multiplied", "divided":
				if s[i+1] != "by" {
					return 0, false
				}
				i++
			case "plus", "minus":
			default:
				return 0, false
			}
			i++
			op = getOperand
		}
	}
	if op == initial || op == getOperand {
		return 0, false
	}
	return res, true
}
