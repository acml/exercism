/*
Package scrabble provides a function to compute the Scrabble score for a given
word.
*/
package scrabble

import "strings"

/*
Score given a word, computes the Scrabble score for that word.
*/
func Score(input string) int {
	var score int
	for _, char := range strings.ToUpper(input) {
		switch char {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			score++
		case 'D', 'G':
			score += 2
		case 'B', 'C', 'M', 'P':
			score += 3
		case 'F', 'H', 'V', 'W', 'Y':
			score += 4
		case 'K':
			score += 5
		case 'J', 'X':
			score += 8
		case 'Q', 'Z':
			score += 10
		}
	}

	return score
}
