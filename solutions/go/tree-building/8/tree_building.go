package tree

import (
	"errors"
	"sort"
)

// Record represents post data stored in database
type Record struct {
	ID     int
	Parent int
}

// Node is used for tree presentation of data
type Node struct {
	ID       int
	Children []*Node
}

// Build constructs tree structures from unsorted set of records.
func Build(records []Record) (*Node, error) {

	if len(records) == 0 {
		return nil, nil
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	if records[0].ID != 0 || records[0].Parent != 0 {
		return nil, errors.New("erroneous root node")
	}

	tree := &Node{ID: 0}
	for i, r := range records[1:] {
		if r.ID <= r.Parent || r.ID != records[i].ID+1 {
			return nil, errors.New("non-continuous")
		}

		if r.Parent == tree.ID {
			tree.Children = append(tree.Children, &Node{ID: r.ID})
			continue
		}

		for _, c := range tree.Children {
			if r.Parent == c.ID {
				c.Children = append(c.Children, &Node{ID: r.ID})
				break
			}
		}
	}

	return tree, nil
}
