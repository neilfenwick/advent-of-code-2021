package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
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

	pageOrderingRules, pageUpdates := parseInput(file)
	populateIndeces(pageUpdates)

	sum := 0
	for _, update := range pageUpdates {
		pageOrderValid := true

		for _, rule := range pageOrderingRules {
			leftIdx, leftFound := update.index[rule.left]
			rightIdx, rightFound := update.index[rule.right]
			if !leftFound || !rightFound {
				continue
			}

			if leftIdx > rightIdx {
				pageOrderValid = false
				break
			}
		}

		if pageOrderValid {
			middleValue := update.pageUpdates[len(update.pageUpdates)/2]
			sum += middleValue
		}
	}

	fmt.Printf("Sum of all middle values of out of order updates: %d\n", sum)
}

type rule struct {
	left  int
	right int
}

type pageUpdateIndex struct {
	pageUpdates []int
	index       map[int]int
}

func parseInput(file io.Reader) ([]rule, []pageUpdateIndex) {
	rules := make([]rule, 0)
	indeces := make([]pageUpdateIndex, 0)

	isProcessingRulesSection := true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			isProcessingRulesSection = false
			continue
		}

		if isProcessingRulesSection {
			rule := rule{}
			fmt.Sscanf(line, "%d|%d", &rule.left, &rule.right)
			rules = append(rules, rule)
			continue
		}

		index := pageUpdateIndex{pageUpdates: make([]int, 0)}
		for _, page := range strings.Split(line, ",") {
			pageNum, _ := strconv.Atoi(page)
			index.pageUpdates = append(index.pageUpdates, pageNum)
		}
		indeces = append(indeces, index)
	}

	return rules, indeces
}

func populateIndeces(indeces []pageUpdateIndex) {
	for i, update := range indeces {
		index := make(map[int]int, 0)
		for idx, pageNum := range update.pageUpdates {
			index[pageNum] = idx
		}
		indeces[i].index = index
	}
}
