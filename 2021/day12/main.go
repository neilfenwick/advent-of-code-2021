package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/neilfenwick/advent-of-code/data"
)

var (
	caveGraph = data.NewGraph()
	allPaths  = make([]data.Stack, 0)
)

type (
	canVisitCaveFunc func(node, start *data.Node, visited []*data.Node) bool
	Strategy         int
)

const (
	SmallCavesOnce       Strategy = 1
	SingleSmallCaveTwice Strategy = 2
)

func main() {
	var (
		file     *os.File
		err      error
		strategy = 1
	)
	switch len(os.Args) {
	case 1:
		file = os.Stdin
	case 3:
		strategy, err = strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("Did not understand window size of: %s\n", os.Args[2])
		}

		fallthrough
	case 2:
		file, err = os.Open(os.Args[1])
		if err != nil {
			log.Fatalf("Error opening file: %s", os.Args[1])
		}
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	startName, endName := "start", "end"
	populateCaveSystemGraph(file)
	fmt.Printf("Walking from %s to %s with strategy %d\n", startName, endName, strategy)
	findPaths(startName, endName, Strategy(strategy))
	fmt.Printf("Found %d paths\n", len(allPaths))
}

func populateCaveSystemGraph(r io.Reader) {
	var (
		startNode, endNode *data.Node
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

func findPaths(startName, endName string, strategy Strategy) {
	var (
		strategyFunc canVisitCaveFunc
	)
	start, _ := caveGraph.GetNode(startName)
	end, _ := caveGraph.GetNode(endName)

	switch strategy {
	case SmallCavesOnce:
		strategyFunc = canVisitSmallCavesOnlyOnce
	case SingleSmallCaveTwice:
		strategyFunc = canVisitSingleSmallCaveTwice
	}
	currentPath := *data.NewStack()
	walkDepthFirst(start, end, start, currentPath, []*data.Node{}, strategyFunc)
}

func walkDepthFirst(current, end, start *data.Node, currentPathDepthFirst data.Stack, visitedCurrentTraverse []*data.Node, canVisitFunc canVisitCaveFunc) {
	currentPathDepthFirst.Push(current)
	if current == end {
		allPaths = append(allPaths, *currentPathDepthFirst.Copy())
		currentPathDepthFirst.Pop()
		return
	}
	visitedCurrentTraverse = append(visitedCurrentTraverse, current)
	for _, child := range current.Links {
		if canVisitFunc(start, child, visitedCurrentTraverse) {
			walkDepthFirst(child, end, start, currentPathDepthFirst, visitedCurrentTraverse, canVisitFunc)
		}
	}
}

// canVisitSmallCavesOnlyOnce returns false if the name is lowercase, and it has already been visited
func canVisitSmallCavesOnlyOnce(start, node *data.Node, visited []*data.Node) bool {
	if node == start {
		return false
	}
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

// canVisitSingleSmallCaveTwice returns false if the name is lowercase, and it has already been visited
// except that one lowercase cave may be visited twice
func canVisitSingleSmallCaveTwice(start, node *data.Node, visited []*data.Node) bool {
	var (
		smallCaveVisitCount   = make(map[string]int)
		smallCaveLimitReached bool
	)
	if node == start {
		return false
	}
	if strings.ToUpper(node.Name) == node.Name {
		return true
	}
	for _, seen := range visited {
		if strings.ToUpper(seen.Name) == seen.Name {
			continue
		}
		visitCount := smallCaveVisitCount[seen.Name] + 1
		smallCaveVisitCount[seen.Name] = visitCount
		if visitCount == 2 && node.Name != seen.Name {
			smallCaveLimitReached = true
		}
	}
	visitCount := smallCaveVisitCount[node.Name]
	if smallCaveLimitReached && visitCount > 0 {
		return false
	} else {
		return visitCount < 2
	}
}
