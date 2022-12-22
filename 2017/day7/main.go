package main

import (
	"bufio"
	"fmt"
	"github.com/neilfenwick/advent-of-code/tree"
	"os"
	"sort"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic("Could not open input.txt")
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	scanner := bufio.NewScanner(file)
	discs := make(map[string]Disc, 0)
	for scanner.Scan() {
		text := scanner.Text()
		d := ParseDisc(text)
		discs[d.Name] = d
	}

	var t *tree.Tree
	for name, disc := range discs {
		var node *tree.Node
		if t == nil {
			t = tree.NewTree(tree.Key{Name: disc.Name, Value: disc})
			_, node = t.GetRoot()
		}

		if node == nil {
			node, _ = t.AppendNode(tree.Key{Name: disc.Name, Value: disc})
		}

		for _, childName := range disc.Children {
			child := discs[childName]
			parent := tree.Key{Name: name, Value: node.Key}
			newChild := tree.Key{Name: child.Name, Value: child}
			t.AppendChild(parent, newChild)
		}
	}

	name, _ := t.GetRoot()
	unbalancedName, weightDiff := GetUnbalanced(t)
	unbalancedNode, _ := t.GetNode(unbalancedName)
	unbalancedDisc := unbalancedNode.Key.Value.(Disc)

	fmt.Printf("The bottom program is called: %s\n", name)
	fmt.Printf("Program '%+v' is unbalanced. Its weight is %d away from what it should be.\n", unbalancedDisc, weightDiff)
}

// GetUnbalanced recursively searches down the tree (depth-first), following branches that do not
// have Discs with a total weight equal to their siblings.
// Returns the name of the lowest-level unbalanced Disc, with the difference in its weight
func GetUnbalanced(t *tree.Tree) (string, int) {

	rootName, _ := t.GetRoot()
	name, weightDelta := searchForUnbalancedChildren(t, rootName)
	return name, weightDelta
}

func searchForUnbalancedChildren(t *tree.Tree, startNodeName string) (string, int) {
	type nodeWeights struct {
		names  []string
		weight int
	}

	startNode, ok := t.GetNode(startNodeName)
	if !ok {
		panic(fmt.Sprintf("Could not recurse into node '%s'", startNodeName))
	}

	children := startNode.Children
	var weights []nodeWeights
	weightMap := make(map[int]int)

	for _, child := range children {
		weight := getNodeWeight(child)
		if pos, ok := weightMap[weight]; ok {
			nodeWeight := weights[pos]
			nodeWeight.names = append(nodeWeight.names, child.Key.Name)
			weights[pos] = nodeWeight
		} else {
			var names []string
			names = append(names, child.Key.Name)
			nodeWeight := nodeWeights{names: names, weight: weight}
			weights = append(weights, nodeWeight)
			weightMap[weight] = len(weights) - 1
		}
	}

	// all weights equal, exit
	if len(weights) == 1 {
		return startNodeName, 0
	}

	sort.Slice(weights, func(i, j int) bool {
		return len(weights[i].names) < len(weights[j].names)
	})

	unbalancedName := weights[0].names[0]
	unbalancedDelta := weights[0].weight - weights[1].weight

	childNameSearch, childDelta := searchForUnbalancedChildren(t, unbalancedName)

	if childDelta == 0 {
		return unbalancedName, unbalancedDelta
	}

	return childNameSearch, childDelta
}

func getNodeWeight(node *tree.Node) int {
	var weight int
	for _, child := range node.Children {
		weight += getNodeWeight(child)
	}

	disc := node.Key.Value.(Disc)
	weight += disc.Weight
	return weight
}
