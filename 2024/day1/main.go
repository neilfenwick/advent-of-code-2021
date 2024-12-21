package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	var (
		file *os.File
		err  error
	)

	switch len(os.Args) {
	case 1:
		file = os.Stdin
	case 2:
		file, err = os.Open(os.Args[1])
		if err != nil {
			log.Fatalf("Error opening file: %s", os.Args[1])
		}
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	left, right := parseLocations(file)
	sort.Ints(left)
	sort.Ints(right)

	part1(left, right)
	part2(left, right)
}

func parseLocations(file *os.File) ([]int, []int) {
	left := make([]int, 1000)
	right := make([]int, 1000)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		loc1, loc2 := parseLine(line)
		left = append(left, loc1)
		right = append(right, loc2)
	}

	return left, right
}

func parseLine(line string) (int, int) {
	var id1, id2 int

	_, err := fmt.Sscanf(line, "%d %d", &id1, &id2)
	if err != nil {
		log.Fatalf("Error parsing line: %s", line)
	}

	return id1, id2
}

func part1(left, right []int) {
	total := 0

	for i := range left {
		total += abs(left[i] - right[i])
	}

	fmt.Printf("Part 1 Total: %d\n", total)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func part2(left, right []int) {
	total := 0
	rightGroup := make(map[int]int)

	for i := range right {
		rightGroup[right[i]]++
	}

	for i := range left {
		leftValue := left[i]
		total += leftValue * rightGroup[leftValue]
	}

	fmt.Printf("Part 2 Total: %d\n", total)
}
