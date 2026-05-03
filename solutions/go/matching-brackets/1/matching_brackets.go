package brackets

var match = map[rune]rune{
	')': '(',
	'}': '{',
	']': '[',
}

// Bracket given a string containing brackets `[]`, braces `{}`, parentheses
// `()`, or any combination thereof, verifies that any and all pairs are matched
// and nested correctly.
func Bracket(input string) bool {
	brackets := []rune{}
	for _, r := range input {
		switch r {
		case '(', '{', '[':
			brackets = append(brackets, r)
		case ')', '}', ']':
			if len(brackets) == 0 || brackets[len(brackets)-1] != match[r] {
				return false
			}
			brackets = brackets[:len(brackets)-1]
		}
	}
	return len(brackets) == 0
}
