package tree

// Node represents an element of a Tree data graph
type Node struct {
	parent   *Node
	children []*Node
	Key      Key
}

// GetChildren returns the first-level descendants of a Node
func (n *Node) GetChildren() []*Node {
	return n.children
}
