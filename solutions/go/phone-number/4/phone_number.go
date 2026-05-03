// Package phonenumber provides functions to clean up user-entered phone numbers
// so that they can be sent SMS messages.
package phonenumber

import (
	"errors"
	"regexp"
)

var re = regexp.MustCompile(`(?:\+)?(1)?(?:\s*)\(?([2-9]\d{2})\)?(?:\s*|[-.])([2-9]\d{2})(?:\s*|[-.])(\d{4})(?:\s*)`)

// Number validates the input phone number.
func Number(input string) (string, error) {
	if !re.MatchString(input) {
		return "", errors.New("No match")
	}

	if res := re.ReplaceAllString(input, "$1$2$3$4"); len(res) == 11 {
		if res[0] != '1' {
			return "", errors.New("11 digits does not start with a 1")
		}
	} else if len(res) > 11 {
		return "", errors.New("invalid when more than 11 digits")
	}
	return re.ReplaceAllString(input, "$2$3$4"), nil
}

// AreaCode prints the area code part of the phone number.
func AreaCode(input string) (string, error) {
	number, err := Number(input)
	if err != nil {
		return "", err
	}
	return number[:3], nil
}

// Format outputs phone numbers in a standart format.
func Format(input string) (string, error) {
	number, err := Number(input)
	if err != nil {
		return "", err
	}
	return "(" + number[:3] + ") " + number[3:6] + "-" + number[6:], nil
}
