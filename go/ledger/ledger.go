package ledger

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

type (
	// Entry is a typical ledger input entry
	Entry struct {
		Date        string // "Y-m-d"
		Description string
		Change      int // in cents
	}

	numbers struct {
		decimalSeperator string
		digitGrouping    string
		negativeStart    string
		negativeEnd      string
		symbolSeparator  string
	}

	literals struct {
		date        string
		description string
		change      string
		dateFormat  string
		number      *numbers
	}
)

const (
	dateLayout = "2006-01-02"
	headerFormat = "%-10s | %-25s | %s\n"
	entryFormat = "%-10s | %-25s | %13s\n"
)

var (
	langs = map[string]*literals{
		"en-US": {"Date", "Description", "Change", "01/02/2006", &numbers{".", ",", "(", ")", ""}},
		"nl-NL": {"Datum", "Omschrijving", "Verandering", "02-01-2006", &numbers{",", ".", "", "-", " "}},
	}

	symbols = map[string]rune{
		"EUR": '€',
		"USD": '$',
	}
)

func formatChange(cents int, number *numbers, symbol rune) string {
	a, negative := "", false
	if cents < 0 {
		cents = -cents
		negative = true
		a += number.negativeStart
	}

	wholeStr := strconv.Itoa(cents/100)
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
		l := langs["en-US"]
		return fmt.Sprintf(headerFormat, l.date, l.description, l.change), nil
	}

	entriesCopy := make([]Entry, len(entries))
	copy(entriesCopy, entries)
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
			t, err := time.Parse(dateLayout, entry.Date)
			if err != nil {
				co <- struct {
					i int
					s string
					e error
				}{e: err}
				return
			}

			desc := entry.Description
			if len(desc) > 25 {
				desc = desc[:22] + "..."
			}

			co <- struct {
				i int
				s string
				e error
			}{
				i: i,
				s: fmt.Sprintf(entryFormat, t.Format(lang.dateFormat), desc, formatChange(entry.Change, lang.number, symbol)),
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

	return fmt.Sprintf(headerFormat, lang.date, lang.description, lang.change) +
		strings.Join(ss, ""), nil
}
