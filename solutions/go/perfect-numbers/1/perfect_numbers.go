// Package perfect provides a function which determines if a number is perfect,
// abundant, or deficient based on Nicomachus' (60 - 120 CE) classification
// scheme for natural numbers.
package perfect

import (
	"errors"
	"math"
)

// ErrOnlyPositive error is given when the input is not a positive integer
var ErrOnlyPositive = errors.New("input is not a positive integer")

// Classification type for natural numbers
type Classification uint8

// Classification values for natural numbers
const (
	ClassificationDeficient Classification = iota // aliquot sum < number
	ClassificationPerfect                         // aliquot sum = number
	ClassificationAbundant                        // aliquot sum > number
)

// Classify determines if a number is perfect, abundant, or deficient based on
// Nicomachus' (60 - 120 CE) classification scheme for natural numbers.
func Classify(input int64) (Classification, error) {
	if input < 1 {
		return 0, ErrOnlyPositive
	}

	var sum int64
	var i int64
	for i = 1; i <= int64(math.Sqrt(float64(input))); i++ {
		if input%i == 0 && input != i {
			if (input / i) == i {
				sum += i
			} else {
				sum += i
				if (input / i) != input {
					sum += (input / i)
				}
			}
		}
	}

	switch {
	case sum < input:
		return ClassificationDeficient, nil
	case sum > input:
		return ClassificationAbundant, nil
	default:
		return ClassificationPerfect, nil
	}
}
