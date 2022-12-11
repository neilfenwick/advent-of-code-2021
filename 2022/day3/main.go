package main

import (
	"bufio"
	"fmt"
	"io"
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

	ruckSacks := calculatePriorities(file)

	sum := 0
	for _, ruckSack := range ruckSacks {
		sum += ruckSack.getPriorityScore()
	}

	fmt.Printf("Priority sum: %d\n", sum)
}

type ruckSack struct {
	compartment1 []rune
	compartment2 []rune
	priorityItem rune
}

func (r *ruckSack) calculatePriorityItem() {
	/*
	   Start by sorting the items in each compartment and then comparing each
	   compartment by walking them in order, similar to merge sort, to find the first
	   item that exists in both compartments
	*/
	sort.Slice(r.compartment1, func(i, j int) bool { return r.compartment1[i] < r.compartment1[j] })
	sort.Slice(r.compartment2, func(i, j int) bool { return r.compartment2[i] < r.compartment2[j] })

	var currentRune rune
	i, j := 0, 0

	for i < len(r.compartment1) {
		currentRune = r.compartment1[i]

		for j < len(r.compartment2) {
			testRune := r.compartment2[j]
			if currentRune == testRune {
				goto foundPriorityItem
			} else if currentRune > testRune {
				j++
				continue
			} else if currentRune < testRune {
				break
			}
		}
		i++
	}

foundPriorityItem:
	r.priorityItem = currentRune
}

func (r *ruckSack) getPriorityScore() int {
	/*
	  Go runes for 'a-z' run from 97-122
	  Go runes for 'A-Z' run from 65-90
	  This function needs to return 1-26 for 'a-z' and 27-52 for 'A-Z'
	*/
	r.calculatePriorityItem()
	priorityScore := int(r.priorityItem)

	if 97 <= priorityScore && priorityScore <= 122 {
		return priorityScore - 96
	}

	return priorityScore - 64 + 26
}

func (r *ruckSack) String() string {
	return fmt.Sprintf(
		"RuckSack:\nCompartment1: %v\nCompartment2: %v\n",
		string(r.compartment1),
		string(r.compartment2))
}

func calculatePriorities(file io.Reader) []ruckSack {
	s := bufio.NewScanner(file)
	s.Split(bufio.ScanLines)
	ruckSacks := []ruckSack{}

	for s.Scan() {
		pack := ruckSack{}
		items := s.Text()
		pack.compartment1 = []rune(items[0 : len(items)/2])
		pack.compartment2 = []rune(items[len(items)/2:])
		ruckSacks = append(ruckSacks, pack)
	}

	return ruckSacks
}
