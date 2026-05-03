package pov

type arcst struct{ from, to string }

type adjacencyList []string

// Graph represents a general directed graph.
type Graph struct {
	adjacencies map[string]adjacencyList
	arcs        []arcst
}

// TreeNode is the tree
type TreeNode struct {
	label    string
	children []*TreeNode
}

// New will create a graph.
func New() *Graph {
	return &Graph{
		adjacencies: map[string]adjacencyList{},
		arcs:        []arcst{},
	}
}

// AddNode adds leaf nodes to the graph.
func (g *Graph) AddNode(nodeLabel string) {
}

// AddArc constructs the rest of the tree from the bottom up.
func (g *Graph) AddArc(from, to string) {
	g.addAdjacency(from, to)
	g.arcs = append(g.arcs, arcst{from, to})
}

// ArcList is a dump method to let the test program see your graph. Formats each
// arc as a single string like "from -> to".
func (g *Graph) ArcList() []string {
	res := []string{}
	for _, a := range g.arcs {
		res = append(res, a.from+" -> "+a.to)
	}
	return res
}

// ChangeRoot changes root of a tree.
func (g *Graph) ChangeRoot(oldRoot, newRoot string) *Graph {
	tree := g.buildTree(&TreeNode{label: newRoot}, nil)
	g.arcs = []arcst{}
	tree.buildArcList(g)
	return g
}

func (t *TreeNode) buildArcList(g *Graph) {
	for _, c := range t.children {
		g.AddArc(t.label, c.label)
		c.buildArcList(g)
	}
}

// Build tree recursively depth first.
func (g *Graph) buildTree(node, parent *TreeNode) *TreeNode {
	for _, childLabel := range g.adjacencies[node.label] {
		// Avoid adding an edge pointing back to the parent.
		if parent != nil && childLabel == parent.label {
			continue
		}
		child := &TreeNode{label: childLabel}
		node.children = append(node.children, child)
		g.buildTree(child, node)
	}
	return node
}

func (g *Graph) addAdjacency(from, to string) {
	if _, ok := find(g.adjacencies[from], to); !ok {
		g.adjacencies[from] = append(g.adjacencies[from], to)
	}

	if _, ok := find(g.adjacencies[to], from); !ok {
		g.adjacencies[to] = append(g.adjacencies[to], from)
	}
}

func find(slice []string, val string) (int, bool) {
	for pos, item := range slice {
		if item == val {
			return pos, true
		}
	}

	return -1, false
}
