package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var (
		file *os.File
		err  error
	)

	switch len(os.Args) {
	case 1:
		file = os.Stdin
	case 3:
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

	assignments := readAssignments(file)

	enclosedCount := 0
	for _, pair := range assignments {
		_, found := pair.enclosedAssignment()
		if found {
			enclosedCount++
		}
	}

	fmt.Printf("Found %d enclosed assignments\n", enclosedCount)

	overlapCount := 0
	for _, pair := range assignments {
		_, found := pair.overlappedAssignment()
		if found {
			overlapCount++
		}
	}

	fmt.Printf("Found %d overlapped assignments\n", overlapCount)
}

type assignment struct {
	start int
	end   int
}

type assignmentPair struct {
	first  assignment
	second assignment
}

func (pair *assignmentPair) enclosedAssignment() (*assignment, bool) {
	if pair.first.start >= pair.second.start &&
		pair.first.end <= pair.second.end {
		return &pair.first, true
	}

	if pair.second.start >= pair.first.start &&
		pair.second.end <= pair.first.end {
		return &pair.second, true
	}

	return &assignment{}, false
}

func (pair *assignmentPair) overlappedAssignment() (*assignment, bool) {
	if pair.first.start <= pair.second.end &&
		pair.first.end >= pair.second.start {
		return &pair.first, true
	}

	return &assignment{}, false
}

func readAssignments(file io.Reader) []assignmentPair {
	s := bufio.NewScanner(file)
	var rangeOneStart, rangeOneEnd, rangeTwoStart, rangeTwoEnd int
	assignments := []assignmentPair{}

	for s.Scan() {
		fmt.Sscanf(
			s.Text(),
			"%d-%d,%d-%d",
			&rangeOneStart,
			&rangeOneEnd,
			&rangeTwoStart,
			&rangeTwoEnd)

		asssignmentOne := assignment{start: rangeOneStart, end: rangeOneEnd}
		asssignmentTwo := assignment{start: rangeTwoStart, end: rangeTwoEnd}
		pair := assignmentPair{first: asssignmentOne, second: asssignmentTwo}
		assignments = append(assignments, pair)
	}

	return assignments
}
