// Package cipher provides Caesar, Shift and Vigenere encoding and decoding
// functions.
package cipher

import (
	"strings"
	"unicode"
)

type ceasar struct{}
type shift struct{ distance int }
type vigenere struct{ key string }

func (c ceasar) Encode(s string) string {
	encode := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			r = unicode.ToLower(r)
			fallthrough
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+3)%26
		default:
			return -1
		}
	}
	return strings.Map(encode, s)
}

func (c ceasar) Decode(s string) string {
	decode := func(r rune) rune {
		switch {
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+23)%26
		default:
			return -1
		}
	}
	return strings.Map(decode, s)
}

func (c shift) Encode(s string) string {
	encode := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			r = unicode.ToLower(r)
			fallthrough
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+rune(c.distance+26))%26
		default:
			return -1
		}
	}
	return strings.Map(encode, s)
}

func (c shift) Decode(s string) string {
	decode := func(r rune) rune {
		switch {
		case r >= 'a' && r <= 'z':
			return 'a' + (r-'a'+rune(26+(26-c.distance)))%26
		default:
			return -1
		}
	}
	return strings.Map(decode, s)
}

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
		res += string('a' + (r-'a'+rune(26+(26-distance)))%26)
	}
	return res
}

// NewCaesar is an implementation of the Caesar Cipher which was used for some
// messages from Julius Caesar that were sent afield.
func NewCaesar() Cipher { return ceasar{} }

// NewShift is a variant of Caesar Cipher with variable distance.
func NewShift(distance int) Cipher {
	if distance < -25 || distance > 25 || distance == 0 {
		return nil
	}
	return shift{distance}
}

// NewVigenere is a more complex cipher using a string as key value.
func NewVigenere(key string) Cipher {
	var countOfa int
	for _, r := range key {
		if !unicode.IsLower(r) {
			return nil
		}
		if r == 'a' {
			countOfa++
		}
	}
	if len(key) == countOfa {
		return nil
	}
	return vigenere{key}
}
