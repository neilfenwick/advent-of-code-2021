package graph

// Node represents an element of a Graph
type Node struct {
	Links map[string]*Node
	Name  string
	Value interface{}
}

// Graph represents a graph of Nodes data elements
type Graph struct {
	Root  *Node
	Nodes map[string]*Node
}

// NewGraph creates a new data structure that will store a graph of Nodes
func NewGraph() *Graph {
	g := Graph{Nodes: make(map[string]*Node)}
	return &g
}

// GetNode returns a Node that has a key with the unique name supplied
func (g *Graph) GetNode(name string) (*Node, bool) {
	node, ok := g.Nodes[name]
	return node, ok
}

// AppendNode creates a new node using the given Key and adds the node
func (g *Graph) NewNode(name string, value interface{}) *Node {
	node := &Node{
		Name:  name,
		Value: value,
		Links: map[string]*Node{},
	}
	g.Nodes[node.Name] = node
	return node
}

// LinkNodes adds a new child Node to the parent, and attaches both nodes.
func (g *Graph) LinkNodes(first string, second string) bool {
	firstNode, found := g.Nodes[first]
	if !found {
		return false
	}
	secondNode, found := g.Nodes[second]
	if !found {
		return false
	}
	secondNode.Links[first] = firstNode
	firstNode.Links[second] = secondNode
	return true
}
