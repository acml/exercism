package meetup

import "time"

// WeekSchedule is the specific matching criteria of day
type WeekSchedule int

// first, second ... match of day
const (
	// First occurence of day
	First WeekSchedule = iota
	Second
	Third
	Fourth
	Last
	Teenth
)

// Day calculates the date of meetups.
func Day(week WeekSchedule, weekday time.Weekday, month time.Month, year int) int {
	if week == Last {
		for i := 0; i < 7; i++ {
			if time.Date(year, month+1, -i, 0, 0, 0, 0, time.UTC).Weekday() == weekday {
				return time.Date(year, month+1, -i, 0, 0, 0, 0, time.UTC).Day()
			}
		}
	}
	for i := 1; i < 20; {
		if time.Date(year, month, i, 0, 0, 0, 0, time.UTC).Weekday() == weekday {
			switch week {
			case First, Second, Third, Fourth:
				return int(week)*7 + i
			case Teenth:
				if i < 13 {
					i += 7
				} else if (i >= 13) || (i <= 19) {
					return i
				}
			}
		} else {
			i++
		}
	}
	return 1
}
