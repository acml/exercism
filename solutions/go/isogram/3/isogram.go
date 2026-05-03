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

	for _, char := range input {
		if char == ' ' || char == '-' {
			continue
		}
		char = unicode.ToUpper(char)
		if seen[char] {
			return false
		}
		seen[char] = true
	}
	return true
}
