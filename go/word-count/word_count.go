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

	words := strings.FieldsFunc(phrase, func(r rune) bool {
		return !(unicode.IsLetter(r) || unicode.IsDigit(r) || r == '\'')
	})

	frequency := Frequency{}
	for _, word := range words {
		word = strings.Trim(strings.ToLower(word), "'")
		frequency[word]++
	}

	return frequency
}
