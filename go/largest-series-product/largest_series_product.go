// Package lsproduct provides a function that gives Largest Series Product
package lsproduct

import (
	"errors"
	"strconv"
	"unicode"
)

// LargestSeriesProduct given a string of digits, calculates the largest product
// for a contiguous substring of digits of length n
func LargestSeriesProduct(digits string, span int) (int64, error) {
	runes := []rune(digits)

	switch {
	case span > len(digits):
		return 0, errors.New("span must be smaller than string length")
	case span == 0:
		return 1, nil
	case span < 0:
		return 0, errors.New("span must be greater than zero")
	}

	for _, r := range digits {
		if !unicode.IsDigit(r) {
			return 0, errors.New("digits input must only contain digits")
		}
	}

	var err error
	var max int64
	var digit int
	for i := 0; i <= len(runes)-span; i++ {
		var p int64
		digit, err = strconv.Atoi(string(runes[i]))
		if err != nil {
			return 0, err
		}
		p = int64(digit)
		for j := i + 1; j-i < span; j++ {
			digit, err = strconv.Atoi(string(runes[j]))
			if err != nil {
				return 0, err
			}
			p *= int64(digit)
			if i == 0 && j == span {
				max = p
			}
		}
		if p > max {
			max = p
		}
	}
	return max, nil
}
