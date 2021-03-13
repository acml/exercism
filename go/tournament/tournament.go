package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

type team struct {
	name   string
	won    int
	drawn  int
	lost   int
	points int
}

type teams map[string]*team

// Tally the results of a small football competition.
func Tally(r io.Reader, w io.Writer) error {

	competition := teams{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if scanner.Text() == "" || strings.HasPrefix(scanner.Text(), "#") {
			continue
		}
		r := strings.Split(scanner.Text(), ";")
		if len(r) != 3 || (r[2] != "win" && r[2] != "loss" && r[2] != "draw") {
			return fmt.Errorf("erroneous input")
		}

		switch r[2] {
		case "win":
			competition.update(r[0], "win")
			competition.update(r[1], "loss")
		case "loss":
			competition.update(r[0], "loss")
			competition.update(r[1], "win")
		default:
			competition.update(r[0], "draw")
			competition.update(r[1], "draw")
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	fmt.Fprintf(w, "%-31s|%3s |%3s |%3s |%3s |%3s\n", "Team", "MP", "W", "D", "L", "P")
	for _, v := range competition.toSortedSlice() {
		fmt.Fprintf(w, "%-31s|%3d |%3d |%3d |%3d |%3d\n",
			v.name,
			v.won+v.drawn+v.lost,
			v.won, v.drawn, v.lost, v.points)
	}
	return nil
}

func (competition teams) update(teamName, result string) {
	var t *team
	if v, ok := competition[teamName]; ok {
		t = v
	} else {
		t = &team{name: teamName}
	}

	switch result {
	case "win":
		t.points += 3
		t.won++
	case "draw":
		t.points++
		t.drawn++
	case "loss":
		t.lost++
	}

	competition[teamName] = t
}

func (competition teams) toSortedSlice() []*team {
	tally := make([]*team, len(competition))
	idx := 0
	for _, t := range competition {
		tally[idx] = t
		idx++
	}
	sort.Slice(tally, func(i, j int) bool {
		return tally[i].points > tally[j].points ||
			(tally[i].points == tally[j].points && tally[i].name < tally[j].name)
	})

	return tally
}
