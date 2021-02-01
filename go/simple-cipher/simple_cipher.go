// Package cipher provides Caesar, Shift and Vigenere encoding and decoding
// functions.
package cipher

import (
	"strings"
	"unicode"
)

type vigenere struct{ key string }

func (c vigenere) Encode(s string) string {
	normalize := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return unicode.ToLower(r)
		case r >= 'a' && r <= 'z':
			return r
		default:
			return -1
		}
	}
	s = strings.Map(normalize, s)
	res := ""
	for pos, r := range s {
		distance := ([]rune(c.key)[pos%len(c.key)] - 'a') % 26
		res += string('a' + (r-'a'+distance)%26)
	}
	return res
}

func (c vigenere) Decode(s string) string {
	res := ""
	for pos, r := range s {
		distance := ([]rune(c.key)[pos%len(c.key)] - 'a') % 26
		res += string('a' + (r-'a'+rune(26-distance))%26)
	}
	return res
}

// NewCaesar is an implementation of the Caesar Cipher which was used for some
// messages from Julius Caesar that were sent afield.
func NewCaesar() Cipher { return NewShift(3) }

// NewShift is a variant of Caesar Cipher with variable distance.
func NewShift(distance int) Cipher {
	if distance < 0 {
		distance += 26
	}
	return NewVigenere(string(rune('a' + distance)))
}

// NewVigenere is a more complex cipher using a string as key value.
func NewVigenere(key string) Cipher {
	var hasNotA bool
	for _, r := range key {
		if !unicode.IsLower(r) {
			return nil
		}
		if r != 'a' {
			hasNotA = true
		}
	}
	if !hasNotA {
		return nil
	}
	return vigenere{key}
}
