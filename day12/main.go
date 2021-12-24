package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/neilfenwick/advent-of-code-2021/data/graph"
	"github.com/neilfenwick/advent-of-code-2021/data/stack"
)

var caveGraph *graph.Graph = graph.NewGraph()
var allPaths []stack.Stack = make([]stack.Stack, 0)

func main() {
	var (
		file *os.File
		err  error
	)
	file, err = os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("Error opening file: %s", os.Args[1])
	}
	startName, endName := "start", "end"
	populateCaveSystemGraph(file)
	fmt.Printf("Walking from %s to %s\n", startName, endName)
	findPaths(startName, endName)
	fmt.Printf("Found %d paths\n", len(allPaths))
}

func populateCaveSystemGraph(r io.Reader) {
	var (
		startNode, endNode *graph.Node
		found              bool
	)
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		line := strings.Split(strings.TrimSpace(s.Text()), "-")
		if startNode, found = caveGraph.GetNode(line[0]); !found {
			startNode = caveGraph.NewNode(line[0], line[0])
		}
		if endNode, found = caveGraph.GetNode(line[1]); !found {
			endNode = caveGraph.NewNode(line[1], line[1])
		}
		caveGraph.LinkNodes(startNode.Name, endNode.Name)
	}
}

func findPaths(startName, endName string) {
	start, _ := caveGraph.GetNode(startName)
	end, _ := caveGraph.GetNode(endName)
	for _, child := range start.Links {
		currentPath := *stack.NewStack()
		currentPath.Push(start)
		visitedCurrentTraverse := []*graph.Node{start}
		walkDepthFirst(child, end, currentPath, visitedCurrentTraverse)
	}
}

func walkDepthFirst(current, end *graph.Node, currentPathDepthFirst stack.Stack, visitedCurrentTraverse []*graph.Node) {
	currentPathDepthFirst.Push(current)
	if current == end {
		allPaths = append(allPaths, *currentPathDepthFirst.Copy())
		currentPathDepthFirst.Pop()
		return
	}
	visitedCurrentTraverse = append(visitedCurrentTraverse, current)
	for _, child := range current.Links {
		if canVisit(child, visitedCurrentTraverse) {
			walkDepthFirst(child, end, currentPathDepthFirst, visitedCurrentTraverse)
		}
	}
}

// canVisit returns false if the name is lowercase and it has already been visited
func canVisit(node *graph.Node, visited []*graph.Node) bool {
	if strings.ToUpper(node.Name) == node.Name {
		return true
	}
	for _, seen := range visited {
		if seen == node {
			return false
		}
	}
	return true
}
