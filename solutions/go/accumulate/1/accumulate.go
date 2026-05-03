/*
Package accumulate provides Accumulate function.
*/
package accumulate

/*
Accumulate given a collection and an operation to perform on each element of the
collection, returns a new collection containing the result of applying that
operation to each element of the input collection.
*/
func Accumulate(input []string, converter func(string) string) []string {
	var result []string
	for _, s := range input {
		result = append(result, converter(s))
	}
	return result
}
