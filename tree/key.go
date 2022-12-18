package tree

// Key represents key-value pair that contains a unique Name, and a Value to be stored
// or retrieved from the tree
type Key struct {
	Name  string
	Value interface{}
}
