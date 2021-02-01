// Package romannumerals provides a function that converts normal numbers to
// Roman Numerals.
package romannumerals

import "errors"

// ToRomanNumeral converts normal numbers to Roman Numerals.
func ToRomanNumeral(input int) (string, error) {
	var digits [4]int
	roman := map[int][]string{
		0: {"M", "MM", "MMM"},
		1: {"C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
		2: {"X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
		3: {"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
	}

	if input < 1 || input > 3000 {
		return "", errors.New("invalid input value")
	}

	for i := 0; i < 4 && input > 0; i++ {
		digits[4-i-1] = input % 10
		input /= 10
	}

	result := ""
	for i := 0; i < 4; i++ {
		if digits[i] == 0 {
			continue
		}
		result += roman[i][digits[i]-1]
	}
	return result, nil
}
