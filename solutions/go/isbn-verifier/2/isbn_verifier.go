// Package isbn provide a ISBN validation function
package isbn

import (
	"strconv"
	"strings"
)

// IsValidISBN checks if the provided string is a valid ISBN-10.
func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")

	if len(isbn) != 10 {
		return false
	}

	var sum, digit int
	var err error
	for pos, r := range isbn {
		if pos == 9 && r == 'X' {
			digit = 10
		} else {
			digit, err = strconv.Atoi(string(r))
			if err != nil {
				return false
			}
		}
		sum += digit * (10 - pos)
	}
	return sum%11 == 0
}
