package leap

// Given a year, report if it is a leap year.
func IsLeapYear(year int) bool {
	if year%4 != 0 {
		return false
	}

	if year%100 == 0 {
		if year%400 == 0 {
			return true
		}

		return false
	}

	return true
}
