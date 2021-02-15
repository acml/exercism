package bob

import (
	"strings"
	"unicode"
)

func isQuestion(remark string) bool {
	return strings.HasSuffix(strings.TrimSpace(remark), "?")
}

func isYelling(remark string) bool {
	isLetter, isUpper := 0, 0

	for _, c := range remark {
		if unicode.IsLetter(c) {
			isLetter++
			if unicode.IsUpper(c) {
				isUpper++
			}
		}
	}

	return isLetter > 0 && isLetter == isUpper
}

func isSilence(remark string) bool {
	for _, c := range remark {
		if !unicode.IsSpace(c) {
			return false
		}
	}
	return true
}

// Hey simulates a lackadaisical teenager conversation.
func Hey(remark string) string {
	if isQuestion(remark) {
		if isYelling(remark) {
			return "Calm down, I know what I'm doing!"
		}
		return "Sure."
	} else if isYelling(remark) {
		return "Whoa, chill out!"
	} else if isSilence(remark) {
		return "Fine. Be that way!"
	}
	return "Whatever."
}
