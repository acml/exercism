/*
Package reverse provides a string reversal function.
*/
package reverse

/*
Reverse given a string, returns the reversed string.
*/
func Reverse(input string) string {
	var reverseString string

	for _, char := range input {
		reverseString = string(char) + reverseString
	}

	return reverseString
}
