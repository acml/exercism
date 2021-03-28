package twelve

import "fmt"

var day = [12]string{
	"first",
	"second",
	"third",
	"fourth",
	"fifth",
	"sixth",
	"seventh",
	"eighth",
	"ninth",
	"tenth",
	"eleventh",
	"twelfth",
}

var given = [12]string{
	"a Partridge in a Pear Tree",
	"two Turtle Doves",
	"three French Hens",
	"four Calling Birds",
	"five Gold Rings",
	"six Geese-a-Laying",
	"seven Swans-a-Swimming",
	"eight Maids-a-Milking",
	"nine Ladies Dancing",
	"ten Lords-a-Leaping",
	"eleven Pipers Piping",
	"twelve Drummers Drumming",
}

// Song outputs the lyrics to 'The Twelve Days of Christmas'.
func Song() string {
	song := ""
	for d := 1; d <= 12; d++ {
		song += Verse(d)
		if d < 12 {
			song += "\n"
		}
	}
	return song
}

// Verse outputs a single verse for the given day of 'The Twelve Days of
// Christmas'.
func Verse(n int) string {
	if n > 12 {
		n = 12
	}
	verse := fmt.Sprintf("On the %s day of Christmas my true love gave to me: ", day[n-1])
	for d := n; d > 0; d-- {
		if d < n {
			if d == 1 {
				verse += ", and "
			} else {
				verse += ", "
			}
		}
		verse += given[d-1]
	}
	verse += "."
	return verse
}
