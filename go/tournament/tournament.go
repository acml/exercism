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
		if len(scanner.Text()) == 0 || strings.HasPrefix(scanner.Text(), "#") {
			continue
		}
		r := strings.Split(scanner.Text(), ";")
		if len(r) != 3 || (r[2] != "win" && r[2] != "loss" && r[2] != "draw") {
			return fmt.Errorf("erroneous input")
		}

		competition.updateResults(r[0], r[1], r[2])
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	tally := []*team{}
	for _, t := range competition {
		tally = append(tally, t)
	}
	sort.Slice(tally, func(i, j int) bool {
		return tally[i].points > tally[j].points ||
			(tally[i].points == tally[j].points && tally[i].name < tally[j].name)
	})

	fmt.Fprintf(w, "%-31s|%3s |%3s |%3s |%3s |%3s\n", "Team", "MP", "W", "D", "L", "P")
	for _, v := range tally {
		fmt.Fprintf(w, "%-31s|%3d |%3d |%3d |%3d |%3d\n",
			v.name,
			v.won+v.drawn+v.lost,
			v.won, v.drawn, v.lost, v.points)
	}
	return nil
}

func (competition teams) updateResults(team1, team2, result string) {
	switch result {
	case "win":
		competition.update(team1, "win")
		competition.update(team2, "loss")
	case "draw":
		competition.update(team1, "draw")
		competition.update(team2, "draw")
	case "loss":
		competition.update(team1, "loss")
		competition.update(team2, "win")
	}
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
