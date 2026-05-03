package leap

// IsLeapYear given a year, reports if it is a leap year.
func IsLeapYear(year int) bool {
	if year%4 != 0 {
		return false
	}

	if year%100 == 0 {
		return year%400 == 0
	}

	return true
}
