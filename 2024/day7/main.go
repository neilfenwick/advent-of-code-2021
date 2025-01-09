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
	matches := findMatchingEquations(equations, []func(uint64, uint64) uint64{add, multiply})
	for _, eq := range matches {
		total += eq.result
	}

	fmt.Printf("Part1: Total of matching equations: %d\n", total)

	// Part 2
	total = 0

	matchesPt2 := findMatchingEquations(equations, []func(uint64, uint64) uint64{add, multiply, concat})
	for _, eq := range matchesPt2 {
		total += eq.result
	}

	fmt.Printf("Part2: Total of matching equations: %d\n", total)
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

func concat(a, b uint64) uint64 {
	multiplier := uint64(1)

	// Little trick to determine the number of digits in b
	// keep dividing b by 10 until it reaches 0
	// and the number of divisions is the number of digits in b
	for temp := b; temp > 0; temp /= 10 {
		multiplier *= 10
	}

	// Multiply a by the number of digits in b and add b
	// to effectively shify a to the left in base 10, and add b
	return a*multiplier + b
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

func findMatchingEquations(equations []equation, operators []func(uint64, uint64) uint64) []equation {
	results := make([]equation, 0, len(equations))

	for _, eq := range equations {
		treeRoot := treeNode{value: func() uint64 { return eq.operands[0] }}
		treeRoot.appendCalculationBranches(eq.operands[1:], operators)

		if depthFirstSearchMatchingEquation(&treeRoot, eq.result) {
			results = append(results, eq)
		}
	}

	return results
}

func (currentNode *treeNode) appendCalculationBranches(operands []uint64, operators []func(uint64, uint64) uint64) {
	if len(operands) == 0 {
		return
	}

	for _, operator := range operators {
		// Create a new apply() function for the result() function
		result := apply(currentNode.value(), operands[0], operator)
		node := &treeNode{value: result}

		currentNode.appendChild(node)

		node.appendCalculationBranches(operands[1:], operators)
	}
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
