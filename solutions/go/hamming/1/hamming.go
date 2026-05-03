/*
Package hamming provides a function to calculate the Hamming Distance between
two DNA strands.
*/
package hamming

import "errors"

/*
Distance calculates the Hamming Distance between two DNA sequences of equal
length.
*/
func Distance(a, b string) (int, error) {
	hammingDistance := 0

	if len(a) != len(b) {
		return 0, errors.New("sequences must be of equal length")
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			hammingDistance++
		}
	}
	return hammingDistance, nil
}
