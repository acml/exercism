// Package cryptosquare provides a square code encoder
package cryptosquare

import (
	"strings"
	"unicode"
)

// Encode implements the classic method of composing secret messages which is
// called square code.
func Encode(plain string) string {
	transform := func(r rune) rune {
		if unicode.IsDigit(r) || unicode.IsLetter(r) {
			return unicode.ToLower(r)
		}
		return -1
	}
	normalized := strings.Map(transform, plain)

	if len(normalized) < 2 {
		return normalized
	}

	var row, column int
	for row = 1; ; row++ {
		if row*row >= len(normalized) {
			column = row
			break
		} else if row*(row+1) >= len(normalized) {
			column = row + 1
			break
		}
	}

	if len(normalized) < row*column {
		normalized += strings.Repeat(" ", row*column-len(normalized))
	}

	runes := []rune{}
	for _, r := range normalized {
		runes = append(runes, r)
	}

	normalized = ""
	for i, pos := 0, 0; i < column; i++ {
		for j := 0; j < row; j++ {
			if pos > 0 && pos%row == 0 {
				normalized += " "
			}
			normalized += string(runes[i+j*column])
			pos++
		}
	}
	return normalized
}
