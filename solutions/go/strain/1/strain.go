// Package strain implements the `keep` and `discard` operation on collections.
package strain

// Ints is a collection of integers
type Ints []int

// Lists is a collection of integer slices
type Lists [][]int

// Strings is a collection of string slices
type Strings []string

// Keep returns a collection of items that evaluate positive per the predicate
// function
func (ints Ints) Keep(predicate func(int) bool) Ints {
	var result Ints
	for _, i := range ints {
		if predicate(i) {
			result = append(result, i)
		}
	}
	return result
}

// Discard returns a collection of items that evaluate negative per the predicate
// function
func (ints Ints) Discard(predicate func(int) bool) Ints {
	var result Ints
	for _, i := range ints {
		if !predicate(i) {
			result = append(result, i)
		}
	}
	return result
}

// Keep returns a collection of items that evaluate positive per the predicate
// function
func (lists Lists) Keep(predicate func([]int) bool) Lists {
	var result Lists
	for _, i := range lists {
		if predicate(i) {
			result = append(result, i)
		}
	}
	return result
}

// Keep returns a collection of items that evaluate positive per the predicate
// function
func (strings Strings) Keep(predicate func(string) bool) Strings {
	var result Strings
	for _, i := range strings {
		if predicate(i) {
			result = append(result, i)
		}
	}
	return result
}
