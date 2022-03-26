package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Could not open input for reading: %v\n", err)
	}

	rows := make([][]rune, 0, 50)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		rows = append(rows, []rune(scanner.Text()))
	}

	treeCount11 := countTreesAlongPath(rows, 1, 1)
	treeCount31 := countTreesAlongPath(rows, 3, 1)
	treeCount51 := countTreesAlongPath(rows, 5, 1)
	treeCount71 := countTreesAlongPath(rows, 7, 1)
	treeCount12 := countTreesAlongPath(rows, 1, 2)

	fmt.Printf("Tree count 3*1: %d\n", treeCount31)
	fmt.Printf("Tree count: %d\n", treeCount11*treeCount31*treeCount51*treeCount71*treeCount12)
}

func countTreesAlongPath(rows [][]rune, horizontalIncrement int, verticalIncrement int) int {
	horizontalPos := 0
	treeCount := 0
	for i := verticalIncrement; i < len(rows); i = i + verticalIncrement {
		horizontalPos += horizontalIncrement
		row := rows[i]
		rowIndex := horizontalPos % len(row)
		if '#' == row[rowIndex] {
			treeCount++
		}
	}

	return treeCount
}
