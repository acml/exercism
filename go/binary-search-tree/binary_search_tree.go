package binarysearchtree

// SearchTreeData is building block of a sorted search tree
type SearchTreeData struct {
	left  *SearchTreeData
	data  int
	right *SearchTreeData
}

// Bst initalizes a search tree
func Bst(val int) *SearchTreeData {
	return &SearchTreeData{data: val}
}

// Insert inserts the given value in to the search tree
func (bst *SearchTreeData) Insert(val int) {
	for n := bst; ; {
		if val <= n.data {
			if n.left == nil {
				n.left = Bst(val)
				return
			}
			n = n.left
			continue
		}

		if val > n.data {
			if n.right == nil {
				n.right = Bst(val)
				return
			}
			n = n.right
			continue
		}
	}
}

// MapString prints out string representation of the search tree
func (bst *SearchTreeData) MapString(f func(int) string) []string {
	var res []string

	if bst.left != nil {
		res = bst.left.MapString(f)
	}
	res = append(res, f(bst.data))
	if bst.right != nil {
		res = append(res, bst.right.MapString(f)...)
	}

	return res
}

// MapInt prints out the search tree values in order
func (bst *SearchTreeData) MapInt(f func(int) int) []int {
	var res []int

	if bst.left != nil {
		res = bst.left.MapInt(f)
	}
	res = append(res, f(bst.data))
	if bst.right != nil {
		res = append(res, bst.right.MapInt(f)...)
	}

	return res
}
