package acronym

import (
	"strings"
)

// Abbreviate converts a phrase to its acronym.
func Abbreviate(s string) string {
	fieldsPredicate := func(r rune) bool {
		if r == '-' || r == '_' || r == ' ' {
			return true
		}
		return false
	}
	var acronym strings.Builder
	for _, word := range strings.FieldsFunc(s, fieldsPredicate) {
		acronym.WriteString(word[0:1])
	}
	return strings.ToUpper(acronym.String())
}
