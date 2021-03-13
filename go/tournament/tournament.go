package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

type match struct {
	name   string
	played int
	won    int
	drawn  int
	lost   int
	points int
}

type results map[string]*match

// Tally the results of a small football competition.
func Tally(r io.Reader, w io.Writer) error {

	matches := results{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if len(scanner.Text()) == 0 || strings.HasPrefix(scanner.Text(), "#") {
			continue
		}
		result := strings.Split(scanner.Text(), ";")
		if len(result) != 3 || (result[2] != "win" && result[2] != "loss" && result[2] != "draw") {
			return fmt.Errorf("erroneous input")
		}

		matches.updateResults(result[0], result[1], result[2])
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	tally := []*match{}
	for _, v := range matches {
		tally = append(tally, v)
	}
	sort.Slice(tally, func(i, j int) bool {
		return tally[i].points > tally[j].points ||
			(tally[i].points == tally[j].points && tally[i].name < tally[j].name)
	})

	fmt.Fprintf(w, "%-31s|%3s |%3s |%3s |%3s |%3s\n", "Team", "MP", "W", "D", "L", "P")
	for _, v := range tally {
		fmt.Fprintf(w, "%-31s|%3d |%3d |%3d |%3d |%3d\n", v.name, v.played, v.won, v.drawn, v.lost, v.points)
	}
	return nil
}

func (m results) updateResults(p1, p2, result string) {
	switch result {
	case "win":
		m.update(p1, "win")
		m.update(p2, "loss")
	case "draw":
		m.update(p1, "draw")
		m.update(p2, "draw")
	case "loss":
		m.update(p1, "loss")
		m.update(p2, "win")
	}
}

func (m results) update(team, result string) {
	var r *match
	if v, ok := m[team]; ok {
		r = v
	} else {
		r = &match{name: team}
	}

	r.played++
	switch result {
	case "win":
		r.points += 3
		r.won++
	case "draw":
		r.points++
		r.drawn++
	case "loss":
		r.lost++
	}

	m[team] = r
}
