package ledger

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

type numbers struct {
	decimalSeperator string
	digitGrouping    string
	negativeStart    string
	negativeEnd      string
	symbolSeparator  string
}

type literals struct {
	date        string
	description string
	change      string
	number      *numbers
}

var langs = map[string]*literals{
	"en-US": {"Date", "Description", "Change", &numbers{".", ",", "(", ")", ""}},
	"nl-NL": {"Datum", "Omschrijving", "Verandering", &numbers{",", ".", "", "-", " "}},
}

var symbols = map[string]rune{
	"EUR": '€',
	"USD": '$',
}

// Entry is a typical ledger input entry
type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

func formatDate(input, locale *string) (string, error) {

	// input date format is "Y-m-d"
	date := strings.SplitN(*input, "-", 3)
	if len(date) != 3 || len(date[0]) != 4 || len(date[1]) != 2 || len(date[2]) != 2 {
		return "", errors.New("date string splitting failed")
	}

	switch *locale {
	case "nl-NL":
		return date[2] + "-" + date[1] + "-" + date[0], nil
	case "en-US":
		return date[1] + "/" + date[2] + "/" + date[0], nil
	}

	return "", errors.New("date conversion failed")
}

func formatChange(cents int, number *numbers, symbol rune) string {
	a, negative := "", false
	if cents < 0 {
		cents = -cents
		negative = true
		a += number.negativeStart
	}

	wholeStr := fmt.Sprintf("%d", cents/100)
	for i := len(wholeStr) - 3; i > 0; i -= 3 {
		wholeStr = wholeStr[:i] + number.digitGrouping + wholeStr[i:]
	}

	a += string(symbol) + number.symbolSeparator +
		wholeStr + number.decimalSeperator +
		fmt.Sprintf("%02d", cents%100)

	if negative {
		a += number.negativeEnd
	} else {
		a += " "
	}

	return a
}

// FormatLedger outputs a beautifully formatted ledger
func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	var lang *literals
	var ok bool
	if lang, ok = langs[locale]; !ok {
		return "", errors.New("language mapping failed")
	}

	var symbol rune
	if symbol, ok = symbols[currency]; !ok {
		return "", errors.New("currency symbol mapping failed")
	}

	if len(entries) == 0 {
		return "Date       | Description               | Change\n", nil
	}

	var entriesCopy []Entry
	entriesCopy = append(entriesCopy, entries...)
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
	}, 10)
	for i, et := range entriesCopy {
		go func(i int, entry Entry) {
			date, ok := formatDate(&entry.Date, &locale)
			if ok != nil {
				co <- struct {
					i int
					s string
					e error
				}{e: ok}
				return
			}

			desc := entry.Description
			if len(desc) > 25 {
				desc = desc[:22] + "..."
			}

			amount := formatChange(entry.Change, lang.number, symbol)

			co <- struct {
				i int
				s string
				e error
			}{
				i: i,
				s: date + strings.Repeat(" ", 10-len(date)) + " | " +
					desc + strings.Repeat(" ", 25-len(desc)) + " | " +
					strings.Repeat(" ", 13-len([]rune(amount))) + amount + "\n",
			}
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

	return lang.date + strings.Repeat(" ", 10-len(lang.date)) + " | " +
		lang.description + strings.Repeat(" ", 25-len(lang.description)) + " | " +
		lang.change + "\n" +
		strings.Join(ss, ""), nil
}
