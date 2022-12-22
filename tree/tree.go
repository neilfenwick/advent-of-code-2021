package tree

import (
	"fmt"
	"strings"
)

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
			Parent: parentNode,
		}
	} else {
		childNode.Parent = parentNode
	}

	parentNode.Children = append(parentNode.Children, childNode)
	t.nodes[child.Name] = childNode

	return childNode
}

// GetRoot returns the root Node of the Tree
func (t *Tree) GetRoot() (string, *Node) {
	for name, node := range t.nodes {
		if node.Parent == nil {
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

// Node represents an element of a Tree data graph
type Node struct {
	Parent   *Node
	Children []*Node
	Key      Key
}

func (n *Node) GetChild(name string) (*Node, bool) {
	children := n.Children
	for _, child := range children {
		if child.Key.Name == name {
			return child, true
		}
	}
	return nil, false
}

func (n *Node) GetPath() string {
	segments := make([]string, 0)
	segments = append(segments, n.Key.Name)
	parent := n.Parent

	for parent != nil {
		segments = append(segments, parent.Key.Name)
		parent = parent.Parent
	}

	builder := strings.Builder{}
	for i := len(segments) - 1; i > 0; i-- {
		builder.WriteRune('/')
		builder.WriteString(segments[i])
	}

    fmt.Println(builder.String())
	return builder.String()
}

// Key represents key-value pair that contains a unique Name, and a Value to be stored
// or retrieved from the tree
type Key struct {
	Name  string
	Value interface{}
}
