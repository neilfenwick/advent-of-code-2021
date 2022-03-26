package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Could not open input for reading: %v", err)
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	day6(f)
}

func day6(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	scanner.Split(emptyLineSplitFunc)

	sumUniqueAnswers, sumUnanimousAnswers := 0, 0
	for scanner.Scan() {
		uniqueAnswers := countUniqueAnswersForGroup(scanner.Text())
		unanimousAnswers := countUnanimousAnswersForGroup(scanner.Text())
		sumUniqueAnswers += uniqueAnswers
		sumUnanimousAnswers += unanimousAnswers
	}

	fmt.Printf("Total unique answers: %d\n", sumUniqueAnswers)
	fmt.Printf("Total unanimous answers: %d\n", sumUnanimousAnswers)
}

func countUniqueAnswersForGroup(group string) int {
	oneLine := strings.ReplaceAll(group, "\n", "")
	answers := make(map[rune]int)

	for _, a := range oneLine {
		cnt, _ := answers[a]
		answers[a] = cnt + 1
	}

	return len(answers)
}

func countUnanimousAnswersForGroup(group string) int {
	scanner := bufio.NewScanner(strings.NewReader(group))
	unanimousAnswers := make(map[rune]int)

	first := true
	for scanner.Scan() {
		if first {
			first = false
			for _, a := range scanner.Text() {
				unanimousAnswers[a] = 1
			}
		} else {
			personAnswers := make(map[rune]int)
			for _, a := range scanner.Text() {
				personAnswers[a] = 1
			}
			for answer := range unanimousAnswers {
				_, exist := personAnswers[answer]
				if !exist {
					delete(unanimousAnswers, answer)
				}
			}
		}
	}

	return len(unanimousAnswers)
}

func emptyLineSplitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	if atEOF {
		return len(data), data, nil
	}

	// Find two newlines in a row and slice out the data
	if i := strings.Index(string(data), "\n\n"); i >= 0 {
		return i + 2, data[0:i], nil
	}

	return
}
