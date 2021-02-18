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

	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	if len(records) == 0 {
		return nil, nil
	}
	if records[0].ID != 0 || records[0].Parent != 0 {
		return nil, errors.New("erroneous root node")
	}

	tree := &Node{ID: 0}
	for i, r := range records[1:] {
		if r.ID <= r.Parent || r.ID != records[i].ID+1 {
			return nil, errors.New("non-continuous")
		}

		if err := treeAddNode(&r, tree); err != nil {
			return nil, err
		}
	}
	return tree, nil
}

func treeAddNode(r *Record, n *Node) error {
	if r.Parent == n.ID {
		n.Children = append(n.Children, &Node{ID: r.ID})
		return nil
	}

	for _, c := range n.Children {
		if r.Parent == c.ID {
			c.Children = append(c.Children, &Node{ID: r.ID})
			return nil
		}
	}
	return errors.New("add node failed")
}
