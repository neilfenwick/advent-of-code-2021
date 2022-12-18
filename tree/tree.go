package tree

// Tree represents a graph of Nodes data elements
type Tree struct {
	nodes map[string]*Node
}

// NewTree creates a new Tree data structure that will store a graph of Nodes
func NewTree(key Key) *Tree {
	node := Node{Key: key}
	tree := Tree{nodes: make(map[string]*Node)}

	tree.nodes[key.Name] = &node
	return &tree
}

// AppendNode creates a new tree node using the given Key
func (t *Tree) AppendNode(key Key) (*Node, bool) {
	n, ok := t.nodes[key.Name]
	if ok {
		return n, false
	}

	n = &Node{
		Key: key,
	}

	t.nodes[key.Name] = n
	return n, true
}

// AppendChild adds a new child Node to the parent that matches the supplied Key
func (t *Tree) AppendChild(parent Key, child Key) *Node {
	parentNode := t.nodes[parent.Name]
	if parentNode == nil {
		parentNode, _ = t.AppendNode(parent)
	}

	childNode := t.nodes[child.Name]

	if childNode == nil {
		childNode = &Node{
			Key:    child,
			parent: parentNode,
		}
	} else {
		childNode.parent = parentNode
	}

	parentNode.children = append(parentNode.children, childNode)
	t.nodes[child.Name] = childNode

	return childNode
}

// GetRoot returns the root Node of the Tree
func (t *Tree) GetRoot() (string, *Node) {
	for name, node := range t.nodes {
		if node.parent == nil {
			return name, node
		}
	}
	panic("No root node")
}

// GetNode returns a Node that has a key with the unique name supplied
func (t *Tree) GetNode(name string) (*Node, bool) {
	node, ok := t.nodes[name]
	return node, ok
}
