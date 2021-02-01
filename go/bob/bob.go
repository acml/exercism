package bob

import (
	"strings"
	"unicode"
)

func is_question(remark string) bool {
	return strings.HasSuffix(strings.TrimSpace(remark), "?")
}

func is_yelling(remark string) bool {
	is_letter, is_upper := 0, 0

	for _, c := range remark {
		if unicode.IsLetter(c) {
			is_letter++
			if unicode.IsUpper(c) {
				is_upper++
			}
		}
	}

	return is_letter > 0 && is_letter == is_upper
}

func is_silence(remark string) bool {
	for _, c := range remark {
		if !unicode.IsSpace(c) {
			return false
		}
	}
	return true
}

// Lackadaisical teenager conversation.
func Hey(remark string) string {
	if is_question(remark) {
		if is_yelling(remark) {
			return "Calm down, I know what I'm doing!"
		}
		return "Sure."
	} else if is_yelling(remark) {
		return "Whoa, chill out!"
	} else if is_silence(remark) {
		return "Fine. Be that way!"
	}
	return "Whatever."
}
