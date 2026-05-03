// Package isbn provide a ISBN validation function
package isbn

import (
	"strconv"
	"strings"
	"unicode"
)

// IsValidISBN checks if the provided string is a valid ISBN-10.
func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")

	if len(isbn) != 10 {
		return false
	}

	sum := 0
	for pos, r := range isbn {

		if !(unicode.IsDigit(r) || (pos == 9 && r == 'X')) {
			return false
		}

		digit, err := strconv.Atoi(string(r))
		if err != nil {
			if pos == 9 && r == 'X' {
				digit = 10
			} else {
				return false
			}
		}
		sum += digit * (10 - pos)
	}
	return sum%11 == 0
}
