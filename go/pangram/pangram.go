// Package pangram provides a functions which determines if a sentence is a
// pangram.
package pangram

import "unicode"

// IsPangram determines if a sentence is a pangram.
func IsPangram(input string) bool {
	letters := map[rune]bool{}
	for _, r := range input {
		if unicode.IsLetter(r) {
			letters[unicode.ToLower(r)] = true
		}
	}
	return len(letters) == 26
}
