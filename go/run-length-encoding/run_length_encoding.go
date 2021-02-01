// Package encode implements run-length encoding and decoding.
package encode

import (
	"strconv"
	"strings"
	"unicode"
)

// RunLengthEncode is a simple form of data compression
func RunLengthEncode(runes string) string {
	if len(runes) == 0 {
		return ""
	}

	var result string
	var value rune
	var count int
	for pos, r := range runes {
		if pos == 0 {
			value = r
		}
		if r == value {
			count++
		} else {
			if count > 1 {
				result += strconv.Itoa(count) + string(value)
			} else {
				result += string(value)
			}
			value = r
			count = 1
		}
	}
	if count > 1 {
		result += strconv.Itoa(count) + string(value)
	} else {
		result += string(value)
	}

	return result
}

// RunLengthDecode reconstrucs original data from the compressed data
func RunLengthDecode(runes string) string {
	var result string
	var count string
	for _, r := range runes {

		if unicode.IsDigit(r) {
			count += string(r)
		} else {
			if len(count) > 0 {
				n, err := strconv.Atoi(count)
				if err != nil {
					return ""
				}
				result += strings.Repeat(string(r), n)
				count = ""
			} else {
				result += string(r)
			}
		}
	}
	return result
}
