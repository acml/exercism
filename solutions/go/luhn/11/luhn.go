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

	input = strings.ReplaceAll(input, " ", "")
	if len(input) <= 1 {
		return false
	}

	sum := 0
	isEven := len(input)%2 == 0
	for _, r := range input {
		digit, err := strconv.Atoi(string(r))
		if err != nil {
			return false
		}
		if isEven {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		isEven = !isEven
	}
	return sum%10 == 0
}
