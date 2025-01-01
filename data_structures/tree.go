package data

import (
	"fmt"
	"strings"
)

// Tree represents a graph of Nodes data elements
type Tree struct {
	nodes map[string]*TreeNode
}

// NewTree creates a new Tree data structure that will store a graph of Nodes
func NewTree(key TreeKey) *Tree {
	node := TreeNode{Key: key}
	tree := Tree{nodes: make(map[string]*TreeNode)}

	tree.nodes[key.Name] = &node
	return &tree
}

// AppendNode creates a new tree node using the given Key
func (t *Tree) AppendNode(key TreeKey) (*TreeNode, bool) {
	n, ok := t.nodes[key.Name]
	if ok {
		return n, false
	}

	n = &TreeNode{
		Key: key,
	}

	t.nodes[key.Name] = n
	return n, true
}

// AppendChild adds a new child Node to the parent that matches the supplied Key
func (t *Tree) AppendChild(parent TreeKey, child TreeKey) *TreeNode {
	parentNode := t.nodes[parent.Name]
	if parentNode == nil {
		parentNode, _ = t.AppendNode(parent)
	}

	childNode := t.nodes[child.Name]

	if childNode == nil {
		childNode = &TreeNode{
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
func (t *Tree) GetRoot() (string, *TreeNode) {
	for name, node := range t.nodes {
		if node.Parent == nil {
			return name, node
		}
	}
	panic("No root node")
}

// GetNode returns a Node that has a key with the unique name supplied
func (t *Tree) GetNode(name string) (*TreeNode, bool) {
	node, ok := t.nodes[name]
	return node, ok
}

// Node represents an element of a Tree data graph
type TreeNode struct {
	Parent   *TreeNode
	Children []*TreeNode
	Key      TreeKey
}

func (n *TreeNode) GetChild(name string) (*TreeNode, bool) {
	children := n.Children
	for _, child := range children {
		if child.Key.Name == name {
			return child, true
		}
	}
	return nil, false
}

func (n *TreeNode) GetPath() string {
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
type TreeKey struct {
	Name  string
	Value interface{}
}
