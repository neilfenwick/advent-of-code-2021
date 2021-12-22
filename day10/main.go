package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/neilfenwick/advent-of-code-2021/data/stack"
)

type chunk struct {
	openChar, closeChar           rune
	errorScore, autoCompleteScore int
}

var chunkMap = map[rune]chunk{
	'(': {'(', ')', 3, 1},
	'[': {'[', ']', 57, 2},
	'{': {'{', '}', 1197, 3},
	'<': {'<', '>', 25137, 4},
}

type errorType int

const (
	CorruptLine errorType = iota
	IncompleteLine
)

type navigationLine struct {
	errorType         errorType
	corruptChar       rune
	autoCompleteScore int
}

func main() {
	var (
		file *os.File
		err  error
	)
	file, err = os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("Error opening file: %s", os.Args[1])
	}

	navigationLines := findCorruptClosingChars(file)
	corruptChars := make([]rune, 50)
	for _, line := range navigationLines {
		if line.errorType == CorruptLine {
			corruptChars = append(corruptChars, line.corruptChar)
		}
	}
	syntaxScore := errorSyntaxScore(corruptChars)
	fmt.Printf("syntax-error-score: %d\n", syntaxScore)

	scores := make([]int, 0, 10)
	for _, line := range navigationLines {
		if line.errorType == IncompleteLine {
			scores = append(scores, line.autoCompleteScore)
		}
	}

	sort.Ints(scores)

	completeScore := scores[len(scores)/2]

	fmt.Printf("autocomplete-score: %d\n", completeScore)
}

func findCorruptClosingChars(r io.Reader) []navigationLine {
	var (
		result = make([]navigationLine, 0, 100)
	)
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		stack := stack.NewStack()
		line := strings.TrimSpace(s.Text())
		if corruptChar, isCorrupt := isCorruptLine([]rune(line), stack, 0); isCorrupt {
			result = append(result, navigationLine{errorType: CorruptLine, corruptChar: corruptChar})
		} else {
			missingCloseChars := make([]rune, 0, 10)
			for {
				char, found := stack.Pop()
				if !found {
					break
				}
				if chunk, found := chunkMap[char.(rune)]; found {
					missingCloseChars = append(missingCloseChars, chunk.closeChar)
				}
			}
			completeScore := autoCompleteScore(missingCloseChars)
			result = append(result, navigationLine{errorType: IncompleteLine, autoCompleteScore: completeScore})
		}
	}
	return result
}

func isCorruptLine(chars []rune, stack *stack.Stack, currentPosition int) (rune, bool) {
	if currentPosition > len(chars)-1 {
		return 0, false
	}
	currentChar := chars[currentPosition]
	if isOpeningChar(currentChar) {
		stack.Push(currentChar)
		return isCorruptLine(chars, stack, currentPosition+1)
	} else {
		if matchingOpenChar, found := findMatchingOpeningChar(currentChar); found {
			if peekChar, found := stack.Peek(); found && peekChar == matchingOpenChar {
				stack.Pop()
				return isCorruptLine(chars, stack, currentPosition+1)
			} else {
				return currentChar, true
			}
		} else {
			log.Fatalf("Could not find matching opening char for closing char: %v", currentChar)
			return 0, false
		}
	}
}

func isOpeningChar(char rune) bool {
	_, found := chunkMap[char]
	return found
}

func findMatchingOpeningChar(closingChar rune) (rune, bool) {
	for _, chunk := range chunkMap {
		if closingChar == chunk.closeChar {
			return chunk.openChar, true
		}
	}
	return 0, false
}

func errorSyntaxScore(chars []rune) int {
	var (
		result int
	)
	for _, r := range chars {
		for _, chunk := range chunkMap {
			if r == chunk.closeChar {
				result += chunk.errorScore
			}
		}
	}
	return result
}

func autoCompleteScore(chars []rune) int {
	var (
		result int
	)
	for _, r := range chars {
		for _, chunk := range chunkMap {
			if r == chunk.closeChar {
				result = result*5 + chunk.autoCompleteScore
			}
		}
	}
	return result
}
