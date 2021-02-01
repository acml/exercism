package acronym

import (
	"strings"
)

// Convert a phrase to its acronym.
func Abbreviate(s string) string {
	acronym := ""
	s = strings.ReplaceAll(s, "-", " ")
	s = strings.ReplaceAll(s, "_", " ")
	substrings := strings.Fields(s)
	for i := range substrings {
		acronym += substrings[i][0:1]
	}
	return strings.ToUpper(acronym)
}
