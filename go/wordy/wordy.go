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

	var op state
	var res int
	var operation string
	s := strings.Split(strings.TrimSuffix(input, "?"), " ")
	for i := 2; i < len(s); {
		switch op {
		case initial:
			fallthrough
		case getOperand:
			n, _ := strconv.Atoi(s[i])
			i++
			switch operation {
			case "add":
				res = res + n
			case "subtract":
				res = res - n
			case "divide":
				res = res / n
			case "multiply":
				res = res * n
			case "":
				res = n
			}
			operation = ""
			op = getOperation
		case getOperation:
			switch s[i] {
			case "plus":
				operation = "add"
			case "minus":
				operation = "subtract"
			case "multiplied":
				if s[i+1] != "by" {
					return 0, false
				}
				operation = "multiply"
				i++
			case "divided":
				if s[i+1] != "by" {
					return 0, false
				}
				operation = "divide"
				i++
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
