// Package etl Extracts-Transforms-Loads data
package etl

import "unicode"

// Transform given a map of integer key and string values returns a
// corresponding map of string keys and integer values.
func Transform(input map[int][]string) map[string]int {
	res := make(map[string]int)
	for key, value := range input {
		for _, str := range value {
			for _, char := range str {
				res[string(unicode.ToLower(char))] = key
			}
		}
	}

	return res
}
