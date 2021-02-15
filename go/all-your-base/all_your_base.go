// Package allyourbase provides a base conversion function
package allyourbase

import "errors"

// ConvertToBase converts a number, represented as a sequence of digits in one
// base, to any other base.
func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	switch {
	case inputBase < 2:
		return []int{}, errors.New("input base must be >= 2")
	case outputBase < 2:
		return []int{}, errors.New("output base must be >= 2")
	}

	var val int
	for i, pow := 0, 1; i < len(inputDigits); i++ {
		if inputDigits[i] < 0 || inputDigits[i] >= inputBase {
			return []int{}, errors.New("all digits must satisfy 0 <= d < input base")
		}
		val += pow * inputDigits[len(inputDigits)-i-1]
		pow = pow * inputBase
	}

	output := []int{}
	for pow := 1; ; {
		digitVal := val % (pow * outputBase)
		output = append(append([]int{}, digitVal/pow), output...)
		val -= digitVal
		pow = pow * outputBase
		if val == 0 {
			break
		}
	}
	return output, nil
}
