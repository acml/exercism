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

type teams map[string]team

// Tally the results of a small football competition.
func Tally(r io.Reader, w io.Writer) error {

	competition := teams{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if scanner.Text() == "" || strings.HasPrefix(scanner.Text(), "#") {
			continue
		}
		r := strings.Split(scanner.Text(), ";")
		if len(r) != 3 {
			return fmt.Errorf("erroneous input")
		}

		t1 := competition[r[0]]
		t1.name = r[0]

		t2 := competition[r[1]]
		t2.name = r[1]

		switch r[2] {
		case "win":
			t1.points += 3
			t1.won++

			t2.lost++

		case "loss":
			t1.lost++

			t2.points += 3
			t2.won++

		case "draw":
			t1.points++
			t1.drawn++

			t2.points++
			t2.drawn++

		default:
			return fmt.Errorf("erroneous input")
		}

		competition[r[0]] = t1
		competition[r[1]] = t2
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	fmt.Fprintf(w, "%-30s | %2s | %2s | %2s | %2s | %2s\n", "Team", "MP", "W", "D", "L", "P")
	for _, v := range competition.toSortedSlice() {
		fmt.Fprintf(w, "%-30s | %2d | %2d | %2d | %2d | %2d\n",
			v.name,
			v.won+v.drawn+v.lost,
			v.won, v.drawn, v.lost, v.points)
	}
	return nil
}

func (competition teams) toSortedSlice() []team {
	tally := make([]team, len(competition))
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
