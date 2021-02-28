package ledger

import (
	"errors"
	"sort"
	"strconv"
	"strings"
)

type numbers struct {
	decimalSeperator string
	digitGrouping string
	negativeStart string
	negativeEnd   string
	symbolSeparator string
}

type literals struct {
	date string
	description string
	change string
	number numbers
}

var langs = map[string]literals{
	"en-US": { "Date",  "Description",  "Change",      numbers{ ".", ",", "(", ")", "" }, },
	"nl-NL": { "Datum", "Omschrijving", "Verandering", numbers{ ",", ".", "",  "-", " "}, },
}

var symbol = map[string]rune {
	"EUR": '€',
	"USD": '$',
}

// Entry is a typical ledger input entry
type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

func formatDate(input, locale string) (string, error) {
	if len(input) != 10 {
		return "", errors.New("invalid input date length")
	}

	// input date format is "Y-m-d"
	date := strings.SplitN(input, "-", 3)
	if len(date) != 3 || len(date[0]) != 4 || len(date[1]) > 2 || len(date[2]) > 2 {
		return "", errors.New("date string splitting failed")
	}

	switch locale {
	case "nl-NL":
		return date[2] + "-" + date[1] + "-" + date[0], nil
	case "en-US":
		return date[1] + "/" + date[2] + "/" + date[0], nil
	}

	return "", errors.New("date conversion failed")
}

func formatChange(cents int, number numbers, currencySymbol rune) string {
	negative := false
	if cents < 0 {
		cents = cents * -1
		negative = true
	}
	var a string
	if negative {
		a += number.negativeStart
	}
	a += string(currencySymbol)
	a += number.symbolSeparator
	centsStr := strconv.Itoa(cents)
	switch len(centsStr) {
	case 1:
		centsStr = "00" + centsStr
	case 2:
		centsStr = "0" + centsStr
	}
	rest := centsStr[:len(centsStr)-2]
	// Groups the cents in groups of three digits
	var parts []string
	for len(rest) > 3 {
		parts = append(parts, rest[len(rest)-3:])
		rest = rest[:len(rest)-3]
	}
	// After groups of three digits last one is the rest
	if len(rest) > 0 {
		parts = append(parts, rest)
	}
	// Print the whole part with digit grouping seperator
	for i := len(parts) - 1; i >= 0; i-- {
		a += parts[i] + number.digitGrouping
	}
	// delete the last digit group seperator
	a = a[:len(a)-1]
	// append decimal seperator
	a += number.decimalSeperator
	// append decimal amount
	a += centsStr[len(centsStr)-2:]
	// append negative end symbol
	if negative {
		a += number.negativeEnd
	} else {
		a += " "
	}
	return a
}

// FormatLedger outputs a beautifully formatted ledger
func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	if len(entries) == 0 {
		if _, err := FormatLedger(currency, "en-US", []Entry{{Date: "2014-01-01", Description: "", Change: 0}}); err != nil {
			return "", err
		}
	}

	var lang literals
	var ok bool
	if lang, ok = langs[locale]; !ok {
		return "", errors.New("language mapping failed")
	}

	var currencySymbol rune
	if currencySymbol, ok = symbol[currency]; !ok {
		return "", errors.New("currency symbol mapping failed")
	}

	var entriesCopy []Entry
	entriesCopy = append(entriesCopy, entries...)
	// sort entries by date, description and change amount
	sort.Slice(entriesCopy, func(i, j int) bool {
		return entriesCopy[i].Date < entriesCopy[j].Date ||
			entriesCopy[i].Description < entriesCopy[j].Description ||
			entriesCopy[i].Change < entriesCopy[j].Change
	})

	// Parallelism, always a great idea
	co := make(chan struct {
		i int
		s string
		e error
	})
	for i, et := range entriesCopy {
		go func(i int, entry Entry) {
			de := entry.Description
			if len(de) > 25 {
				de = de[:22] + "..."
			} else {
				de = de + strings.Repeat(" ", 25-len(de))
			}
			d, ok := formatDate(entry.Date, locale)
			if ok != nil {
				co <- struct {
					i int
					s string
					e error
				}{e: ok}
			}
			a := formatChange(entry.Change, lang.number, currencySymbol)

			co <- struct {
				i int
				s string
				e error
			}{i: i, s: d + strings.Repeat(" ", 10-len(d)) + " | " + de + " | " +
				strings.Repeat(" ", 13-len([]rune(a))) + a + "\n"}
		}(i, et)
	}
	ss := make([]string, len(entriesCopy))
	for range entriesCopy {
		v := <-co
		if v.e != nil {
			return "", v.e
		}
		ss[v.i] = v.s
	}

	s := lang.date +
		strings.Repeat(" ", 10-len(lang.date)) +
		" | " +
		lang.description +
		strings.Repeat(" ", 25-len(lang.description)) +
		" | " + lang.change + "\n"

	s += strings.Join(ss, "");
	return s, nil
}
