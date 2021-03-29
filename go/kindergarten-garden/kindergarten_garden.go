package kindergarten

import (
	"fmt"
	"sort"
	"strings"
)

type Garden struct {
	cups     []string
	children []string
}

// NewGarden is a Garden constructor.
func NewGarden(diagram string, children []string) (*Garden, error) {
	g := &Garden{
		cups:     strings.Split(strings.TrimLeft(diagram, "\n"), "\n"),
		children: make([]string, len(children)),
	}
	if diagram[0] != '\n' {
		return nil, fmt.Errorf("wrong diagram format")
	}
	for i := 1; i < len(diagram); i++ {
		if i == len(children)*2+1 && diagram[i] == '\n' {
			continue
		}
		switch diagram[i] {
		case 'C', 'G', 'R', 'V':
		default:
			return nil, fmt.Errorf("wrong diagram format")
		}
	}
	set := map[string]struct{}{}
	for _, c := range children {
		if _, ok := set[c]; ok {
			return nil, fmt.Errorf("duplicate name")
		}
		set[c] = struct{}{}
	}
	copy(g.children, children)
	if len(g.cups) != 2 && len(g.cups[0]) != len(g.cups[1]) || len(g.cups[0]) != len(g.children)*2 {
		return nil, fmt.Errorf("wrong diagram format")
	}
	sort.Slice(g.children, func(i, j int) bool {
		return g.children[i] < g.children[j]
	})
	return g, nil
}

// Plants given a diagram, determines which plants each child in the
// kindergarten class is responsible for.
func (g *Garden) Plants(child string) ([]string, bool) {
	plants := map[rune]string{
		'C': "clover",
		'G': "grass",
		'R': "radishes",
		'V': "violets",
	}
	index := 0
	found := false
	for p, c := range g.children {
		if child == c {
			index = p
			found = true
		}
	}
	if !found {
		return []string{}, false
	}

	res := make([]string, 4, 4)
	for i := 0; i < 2; i++ {
		res[i*2] = plants[[]rune(g.cups[i])[index*2]]
		res[i*2+1] = plants[[]rune(g.cups[i])[index*2+1]]
	}
	return res, true
}
