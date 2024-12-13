package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"strings"

	data "github.com/neilfenwick/advent-of-code/data_structures"
	"github.com/neilfenwick/advent-of-code/files"
)

func main() {
	file, err := files.ReadInputStream()
	if err != nil {
		log.Fatal("Could not open input for reading")
	}
	defer func(file io.ReadCloser) {
		_ = file.Close()
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	input := scanner.Text()
	scanner.Scan()
	polymerMap := buildPolymerMap(scanner)
	min, max := countMinMaxElementOccurrences(input, polymerMap)
	fmt.Printf("Quantities: most common %d, least common %d. Difference: %d\n", max, min, max-min)
}

func buildPolymerMap(scanner *bufio.Scanner) map[string]rune {
	polymerMap := make(map[string]rune)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		parts := strings.Split(line, " -> ")
		polymerMap[parts[0]] = []rune(parts[1])[0]
	}
	return polymerMap
}

func countMinMaxElementOccurrences(input string, polymers map[string]rune) (int, int) {
	elementsLinkedList := data.NewRuneLinkedList([]rune(input))
	for i := 0; i < 10; i++ {
		expandElements(elementsLinkedList, polymers)
	}
	elementCounts := make(map[rune]int)
	current := elementsLinkedList.Head
	for {
		count, _ := elementCounts[current.Value]
		elementCounts[current.Value] = count + 1
		if current.Next == nil {
			break
		}
		current = current.Next
	}
	min, max := math.MaxInt, 0
	for _, v := range elementCounts {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}

func expandElements(elementsList *data.RuneLinkedList, polymers map[string]rune) {
	currentElement := elementsList.Head
	newElements := data.NewRuneLinkedList([]rune{})
	for {
		if currentElement.Next == nil {
			break
		}
		pair := []rune{currentElement.Value, currentElement.Next.Value}
		newElementValue := polymers[string(pair)]
		newElements.AppendValue(newElementValue)
		currentElement = currentElement.Next
	}
	mergeLinkedLists(elementsList, newElements)
}

func mergeLinkedLists(first *data.RuneLinkedList, second *data.RuneLinkedList) {
	currentElement := first.Head
	currentSecondListElement := second.Head
	var tempNext *data.RuneLinkedListNode
	for {
		tempNext = currentElement.Next
		if tempNext == nil || currentSecondListElement == nil {
			return
		}
		currentElement.Next = currentSecondListElement
		currentElement = currentSecondListElement
		currentSecondListElement = currentSecondListElement.Next
		currentElement.Next = tempNext
		currentElement = tempNext
	}
}
