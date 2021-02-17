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
	var tree *Node
	for i, r := range records {
		if i == 0 {
			switch {
			case r.ID != 0:
				return nil, errors.New("no root mode")
			case r.Parent > 0:
				return nil, errors.New("root node has parent")
			}

			tree = &Node{
				ID: 0,
			}
			continue
		}
		switch {
		case records[i].ID == records[i-1].ID:
			return nil, errors.New("duplicate record")
		case records[i].ID != records[i-1].ID+1:
			return nil, errors.New("non-continuous")
		case records[i].ID == records[i].Parent:
			return nil, errors.New("cycle directly")
		case records[i].ID <= records[i].Parent:
			return nil, errors.New("higher id parent of lower id")
		}

		if err := treeAddNode(&r, tree); err != nil {
			return nil, err
		}
	}
	return tree, nil
}

func treeAddNode(r *Record, n *Node) error {
	if r.Parent == n.ID {
		n.Children = append(n.Children, &Node{
			ID: r.ID,
		})
		return nil
	}

	for _, c := range n.Children {
		if r.Parent == c.ID {
			c.Children = append(c.Children, &Node{
				ID: r.ID,
			})
			return nil
		}
	}
	return errors.New("add node failed")
}
