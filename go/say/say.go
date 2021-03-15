package say

import (
	"math"
)

var groups = []string{"", "thousand", "million", "billion"}
var spell = [][]string{
	{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"},
	{"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"},
	{"twenty"},
	{"thirty"},
	{"forty"},
	{"fifty"},
	{"sixty"},
	{"seventy"},
	{"eighty"},
	{"ninety"},
	{"hundred"},
	{"thousand"},
	{"million"},
}

// Say given a number from 0 to 999,999,999,999, spells out that number in
// English.
func Say(input int64) (string, bool) {

	if input < 0 || input >= int64(math.Pow10(len(groups)*3)) {
		return "", false
	}

	if input == 0 {
		return "zero", true
	}

	s := ""
	for n := len(groups) - 1; n >= 0; n-- {

		group := (input % int64(math.Pow10((n+1)*3))) / int64(math.Pow10(n*3))
		if group == 0 {
			continue
		}

		if s != "" {
			s += " "
		}

		// spell hundreds
		hundreds := group / 100
		if hundreds > 0 {
			s += spell[0][hundreds] + " hundred"
		}

		// spell tens
		switch tens := group % 100; tens / 10 {
		case 0:
			if tens != 0 {
				if hundreds != 0 {
					s += " "
				}
				s += spell[0][tens]
			}
		case 1:
			if hundreds != 0 {
				s += " "
			}
			s += spell[1][tens%10]
		case 2, 3, 4, 5, 6, 7, 8, 9:
			if hundreds != 0 {
				s += " "
			}
			s += spell[tens/10][0]
			if tens%10 != 0 {
				s += "-" + spell[0][tens%10]
			}
		}

		if n > 0 {
			s += " " + groups[n]
		}
	}

	return s, true
}
