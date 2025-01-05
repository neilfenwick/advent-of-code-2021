package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("Error opening file: %s", os.Args[1])
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	equations := readInput(file)

	// Part 1
	var total uint64
	matches := findMatchingEquations(equations)
	for _, eq := range matches {
		total += eq.result
	}

	fmt.Printf("Total of matching equations: %d\n", total)
}

type equation struct {
	result   uint64
	operands []uint64
}

func add(a, b uint64) uint64 {
	return a + b
}

func multiply(a, b uint64) uint64 {
	return a * b
}

func apply(a, b uint64, f func(uint64, uint64) uint64) func() uint64 {
	return func() uint64 {
		return f(a, b)
	}
}

func readInput(file *os.File) []equation {
	equations := make([]equation, 0, 1000)
	s := bufio.NewScanner(file)
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		equations = append(equations, parseEquation(line))
	}
	return equations
}

func parseEquation(line string) equation {
	eq := equation{}

	// Split the line uint64o result part and operands part
	parts := strings.SplitN(line, ":", 2)
	if len(parts) != 2 {
		log.Fatalf("Invalid equation: %s", line)
	}

	// Parse the result part
	fmt.Sscanf(strings.TrimSpace(parts[0]), "%d", &eq.result)

	// Parse the operands part
	operands := strings.Fields(strings.TrimSpace(parts[1]))
	for _, op := range operands {
		var operand uint64
		fmt.Sscanf(op, "%d", &operand)
		eq.operands = append(eq.operands, operand)
	}

	return eq
}

type treeNode struct {
	value    func() uint64
	children []*treeNode
}

func (currentNode *treeNode) appendChild(value *treeNode) {
	currentNode.children = append(currentNode.children, value)
}

func findMatchingEquations(equations []equation) []equation {
	results := make([]equation, 0, len(equations))

	for _, eq := range equations {
		treeRoot := treeNode{value: func() uint64 { return eq.operands[0] }}
		treeRoot.appendCalculationBranches(eq.operands[1:])

		if depthFirstSearchMatchingEquation(&treeRoot, eq.result) {
			results = append(results, eq)
		}
	}

	return results
}

func (currentNode *treeNode) appendCalculationBranches(operands []uint64) {
	if len(operands) == 0 {
		return
	}

	// Create a new apply() function for the sum() function
	sum := apply(currentNode.value(), operands[0], add)
	sumNode := &treeNode{value: sum}

	currentNode.appendChild(sumNode)

	// Create a new apply() function for the multiply() function
	mul := apply(currentNode.value(), operands[0], multiply)
	mulNode := &treeNode{value: mul}
	currentNode.appendChild(mulNode)

	sumNode.appendCalculationBranches(operands[1:])
	mulNode.appendCalculationBranches(operands[1:])
}

func depthFirstSearchMatchingEquation(currentNode *treeNode, result uint64) bool {
	if len(currentNode.children) == 0 {
		// We are at the leaf node
		return currentNode.value() == result
	}

	for _, child := range currentNode.children {
		if depthFirstSearchMatchingEquation(child, result) {
			return true
		}
	}
	return false
}
