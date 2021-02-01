// Package proverb provides a proverb generation function.
package proverb

import "fmt"

// Proverb given a list of inputs, generates the relevant proverb.
func Proverb(rhyme []string) []string {
	result := []string{}

	if len(rhyme) == 0 {
		return result
	}

	for i := 0; i < len(rhyme)-1; i++ {
		result = append(result, fmt.Sprintf("For want of a %v the %v was lost.", rhyme[i], rhyme[i+1]))
	}
	result = append(result, fmt.Sprintf("And all for the want of a %v.", rhyme[0]))

	return result
}
