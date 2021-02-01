// Package rotationalcipher provides an implementation of the rotational cipher,
// also sometimes called the Caesar cipher.
package rotationalcipher

import "strings"

// RotationalCipher is a simple shift cipher that relies on transposing all the
// letters in the alphabet using an integer key between `0` and `26`
func RotationalCipher(plain string, shift int) string {
	rotate := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return 'A' + (r-'A'+rune(shift+26))%26
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+rune(shift+26))%26
		default:
			return r
		}
	}
	return strings.Map(rotate, plain)
}
