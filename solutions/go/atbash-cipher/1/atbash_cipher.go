// Package atbash provides an implementation of the atbash cipher, an ancient
// encryption system created in the Middle East.
package atbash

import (
	"strings"
	"unicode"
)

// Atbash cipher is a simple substitution cipher that relies on transposing all
// the letters in the alphabet such that the resulting alphabet is backwards.
func Atbash(plain string) string {
	result := ""
	runes := []rune(plain)
	pos := 0
	for i := 0; i < len(runes); i++ {
		switch {
		case runes[i] >= 'A' && runes[i] <= 'Z':
			runes[i] = unicode.ToLower(runes[i])
			fallthrough
		case runes[i] >= 'a' && runes[i] <= 'z':
			result += string('a' + rune('z'-runes[i]))
			pos++
		case runes[i] >= '0' && runes[i] <= '9':
			result += string(runes[i])
			pos++
		default:
			continue
		}
		if pos > 0 && pos%5 == 0 {
			result += " "
		}
	}

	return strings.Trim(result, " ")
}
