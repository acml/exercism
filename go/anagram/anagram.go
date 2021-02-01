// Package anagram provides anagram detection function
package anagram

import (
	"reflect"
	"strings"
	"unicode"
)

func anagramMap(word string) map[rune]int {
	anagramMap := map[rune]int{}
	for _, r := range word {
		anagramMap[unicode.ToLower(r)]++
	}
	return anagramMap
}

// Detect given a word and a list of candidates, selects the sublist of anagrams
// of the given word.
func Detect(word string, candidates []string) []string {

	word = strings.ToLower(word)
	wordMap := anagramMap(word)

	var result []string
	for _, candidate := range candidates {
		if word == strings.ToLower(candidate) {
			continue
		}
		candidateMap := anagramMap(candidate)
		if reflect.DeepEqual(wordMap, candidateMap) {
			result = append(result, candidate)
		}
	}
	return result
}
