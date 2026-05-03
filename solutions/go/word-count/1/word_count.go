// Package wordcount provides a function that counts words in a given phrase.
package wordcount

import (
	"strings"
	"unicode"
)

// Frequency is a map of word frequency of given phrase.
type Frequency map[string]int

// WordCount given a phrase, counts the occurrences of each _word_ in that phrase.
func WordCount(phrase string) Frequency {

	var word string
	var hasQuote bool
	frequency := Frequency{}
	for _, r := range phrase {
		switch {
		case unicode.IsLetter(r) && (len(word) == 0 || !strings.ContainsAny(word, "0123456789")):
			if hasQuote {
				word += "'"
				hasQuote = false
			}
			word += string(unicode.ToLower(r))
		case r == '\'' && len(word) > 0 && !strings.ContainsAny(word, "0123456789"):
			hasQuote = true
		case unicode.IsDigit(r) && (len(word) == 0 || strings.ContainsAny(word, "0123456789")):
			hasQuote = false
			word += string(r)
		default:
			hasQuote = false
			if len(word) > 0 {
				frequency[word]++
				word = ""
			}
		}
	}
	if len(word) > 0 {
		frequency[word]++
	}
	return frequency
}
