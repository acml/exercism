/*
Package isogram provides a function to determine if a word or phrase is an isogram.
*/
package isogram

import "unicode"

/*
IsIsogram determines if a word or phrase is an isogram.
*/
func IsIsogram(input string) bool {
	seen := map[rune]bool{}

	for _, r := range input {
		if r == ' ' || r == '-' {
			continue
		}
		r = unicode.ToUpper(r)
		if seen[r] {
			return false
		}
		seen[r] = true
	}
	return true
}
