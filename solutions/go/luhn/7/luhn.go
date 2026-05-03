// Package luhn provides a function to determine whether a given number is valid
// or not per the Luhn formula.
package luhn

import (
	"strconv"
	"strings"
)

// Valid given a number determines whether or not it is valid per the Luhn
// formula.
func Valid(input string) bool {

	runes := []rune(strings.ReplaceAll(input, " ", ""))
	if len(runes) <= 1 {
		return false
	}

	sum := 0
	for i, isOdd := 0, false; i < len(runes); i, isOdd = i + 1, !isOdd {
		digit, err := strconv.Atoi(string(runes[len(runes)-i-1]))
		if err != nil {
			return false
		}
		if isOdd {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}
	return sum%10 == 0
}
