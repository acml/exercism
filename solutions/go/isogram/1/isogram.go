/*
Package isogram provides a function to determine if a word or phrase is an isogram.
*/
package isogram

import "unicode"

/*
IsIsogram determines if a word or phrase is an isogram.
*/
func IsIsogram(input string) bool {
	occurrence := make(map[rune]int)

	for _, char := range input {
		if char != ' ' && char != '-' {
			char = unicode.ToUpper(char)
			if occurrence[char] > 0 {
				return false
			}
			occurrence[char]++
		}
	}
	return true
}
