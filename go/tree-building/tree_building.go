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

		if !addNode(&r, tree) {
			return nil, errors.New("add node failed")
		}
	}

	return tree, nil
}

func addNode(r *Record, n *Node) bool {
	if r.Parent == n.ID {
		n.Children = append(n.Children, &Node{ID: r.ID})
		return true
	}

	for _, c := range n.Children {
		if addNode(r, c) {
			return true
		}
	}
	return false
}
