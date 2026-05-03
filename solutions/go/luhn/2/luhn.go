// Package luhn provides a function to determine whether a given number is valid
// or not per the Luhn formula.
package luhn

import (
	"strings"
	"unicode"
	"strconv"
)

func reverseString(input string) string {
	result := ""
	for _, r := range input {
		result = string(r) + result
	}
	return result
}

func isAllDigits(input string) bool {
	for _, r := range input {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

// Valid given a number determines whether or not it is valid per the Luhn
// formula.
func Valid(input string) bool {
	input = strings.ReplaceAll(input, " ", "")

	if len(input) <= 1 {
		return false
	}

	if !isAllDigits(input) {
		return false
	}

	input = reverseString(input)
	sum := 0
	for pos, r := range input {
		// digit := int(r) - '0'
		digit, err := strconv.Atoi(string(r))
		if err != nil {
			return false
		}
		if pos%2 == 1 {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}
	return sum%10 == 0
}
