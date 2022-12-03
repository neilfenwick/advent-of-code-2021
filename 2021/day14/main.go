package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"strings"

	"github.com/neilfenwick/advent-of-code/data/linkedList"
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
	//min, max := expandMapLinkedListStrategy(input, polymerMap, 10)
	min, max := expandMapHashmapStrategy(input, polymerMap, 10)
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

func expandMapHashmapStrategy(input string, polymers map[string]rune, iterationCount int) (min int, max int) {
	elementsCountMap := make(map[string]int)

	// Add pairs of runes as keys to the map
	for i := 0; i < len(input)-1; i++ {
		pair := input[i : i+2]
		value, found := elementsCountMap[pair]
		if found {
			elementsCountMap[pair] = value + 1
			continue
		}
		elementsCountMap[pair] = 0
	}
	fmt.Printf("%s\n", input)
	fmt.Printf("%v", elementsCountMap)
	for i := 0; i < iterationCount; i++ {
		break
	}
	return 0, 0
}

func expandMapLinkedListStrategy(input string, polymers map[string]rune, iterationCount int) (min int, max int) {
	elementsLinkedList := linkedList.NewRuneLinkedList([]rune(input))
	for i := 0; i < iterationCount; i++ {
		expandPolymerLinkedList(elementsLinkedList, polymers)
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
	min, max = math.MaxInt, 0
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

func expandPolymerLinkedList(elementsList *linkedList.RuneLinkedList, polymers map[string]rune) {
	currentElement := elementsList.Head
	newElements := linkedList.NewRuneLinkedList([]rune{})
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

func mergeLinkedLists(first *linkedList.RuneLinkedList, second *linkedList.RuneLinkedList) {
	currentElement := first.Head
	currentSecondListElement := second.Head
	var tempNext *linkedList.RuneLinkedListNode
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
