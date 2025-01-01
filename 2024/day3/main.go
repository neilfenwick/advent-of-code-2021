package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
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

	operandsList := parseInput(file)

	sum := 0
	for _, operands := range operandsList {
		sum += operands.first * operands.second
	}
	fmt.Printf("Sum of all multiplications: %d\n", sum)

	/* Part 2: Wow did I do this the hard way!!!
	* Because part 1 was done with a relatively simple regex,
	* and because I wanted to stick with built-in funcionality, I decided to use bufio.Scanner
	* to try to make up for the fact that the built-in Go regex does not support negative lookbehinds.
	* This part of the puzzle requires parsing a file with start and stop tokens, and to consume data
	* until a stop token is found, and only resume consuming data when a start token is found.

	 */
	file.Seek(0, 0)
	operandsListPart2 := parseInputPart2(file)

	sumPart2 := 0
	for _, operands := range operandsListPart2 {
		sumPart2 += operands.first * operands.second
	}
	fmt.Printf("Sum of all multiplications: %d\n", sumPart2)
}

func init() {
	re = regexp.MustCompile(pattern)
}

// pattern is a regex that matches the format "mul(x,y)" where x and y are numbers.
// - (?:mul\()   : Non-capturing group that matches the literal string "mul(".
// - (\d+)       : Capturing group that matches one or more digits (represents the first number).
// - (?:,)       : Non-capturing group that matches a comma ",".
// - (\d+)       : Capturing group that matches one or more digits (represents the second number).
// - (?:\))      : Non-capturing group that matches the closing parenthesis ")".
var pattern string = `(?s)(?:mul\()(\d+)(?:,)(\d+)(?:\))`
var re *regexp.Regexp

type operands struct {
	first  int
	second int
}

func parseInput(file io.Reader) []operands {
	operandsList := make([]operands, 0)

	// Open part1.txt for writing
	outputFile, err := os.Create("part1_" + os.Args[1])
	if err != nil {
		log.Fatalf("Error creating output file: %s", err)
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		operandsList = parseOperands(line, operandsList)

	}

	return operandsList
}

func parseOperands(line string, operandsList []operands) []operands {
	matches := re.FindAllStringSubmatch(line, -1)

	for _, lineMatches := range matches {
		for index := range lineMatches {
			if index == 0 || index%2 > 0 {
				continue
			}

			previousMatch := lineMatches[index-1]
			currentMatch := lineMatches[index]
			firstOperand, _ := strconv.Atoi(previousMatch)
			secondOperand, _ := strconv.Atoi(currentMatch)
			operandsList = append(operandsList, operands{firstOperand, secondOperand})
		}
	}
	return operandsList
}

func parseInputPart2(file *os.File) []operands {
	operandsList := make([]operands, 0)

	// Open part2.txt for writing
	outputFile, err := os.Create("part2_" + os.Args[1])
	if err != nil {
		log.Fatalf("Error creating output file: %s", err)
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(customStopStartTokenScanner)

	for scanner.Scan() {
		token := scanner.Text()
		operandsList = parseOperands(token, operandsList)

	}

	return operandsList
}

var stopToken []byte = []byte("don't()")
var startToken []byte = []byte("do()")

/*
I had a little bit of trouble with the scanner and learned that some of my assumptions about how
the buffer works were incorrect. The buffer is not always kept full or automatically topped up
until no data is consumed and a nil token is returned. And so because I am alwways effectively
parsing only a fragment of a the total next, I had to be mindful when NOT to try to consume the
entire buffer, and only consume up to the stop token, and then to request more data until a
start token is found.

(Because if you advance past the stop token, without a start token in
sight, then on a later pass, how do you know if you can consume the data at the beginning of
the buffer, because it might have been preceded by a stop token in a previous pass.)
*/
func customStopStartTokenScanner(data []byte, atEOF bool) (advance int, token []byte, err error) {
	stopIndex := bytes.Index(data, stopToken)
	startIndex := bytes.Index(data, startToken)

	switch {

	case len(data) < len(stopToken):

		if atEOF && len(data) > 0 {
			// If the data is less than the length of the stop token, we cannot possibly have a match,
			// so just return the data
			return len(data), data, nil
		}

		// Don't advance, wait for more data
		return 0, nil, nil

	case stopIndex > -1 && startIndex == -1:
		// When there is a stop token but no start token, we can consume data up until the stopToken
		// but no further because we cannot pass over the stop token until we know where the next start
		// token is
		if stopIndex > 0 {
			return stopIndex, data[:stopIndex], nil
		}

		// If there is a stop token and no start token, and we have already consumed the data up to the
		// stop token, then the only remaining data is all following a stop token. Just advance past it.
		if atEOF {
			return len(data), nil, nil
		}

		// Don't advance, wait for more data because the stop token is at the beginning of the buffer
		// and we have to wait for a start token to know if we can consume data
		return 0, nil, nil

	case stopIndex > -1 && startIndex > -1 && startIndex < stopIndex:
		// When there is a stop token and a start token before it, we need to consume the data before
		// the stop token only
		return stopIndex, data[:stopIndex], nil

	case stopIndex > -1 && startIndex > stopIndex:
		// When there is a stop token but the start token is after the stop token, we need
		// to consume the data before the stop token and advance to the start token
		return startIndex, data[:stopIndex], nil

	default:
		rewind := rewindForPartialTokenAtEndOfBuffer(data)
		if len(data)-rewind < 0 {
			return 0, nil, fmt.Errorf("rewind is negative")
		}
		if len(data)-rewind == 0 {
			return len(data), nil, nil
		}

		return len(data) - rewind, data[:len(data)-rewind], nil
	}
}

func rewindForPartialTokenAtEndOfBuffer(data []byte) int {
	rewind := rewindForPartialMulMatch(data, 0)
	rewind = rewindForPartialTokenMatch(data[:len(data)-rewind], rewind)
	return rewind
}

var partialMulMatches [][]byte = [][]byte{
	[]byte(","), []byte("0"), []byte("1"), []byte("2"), []byte("3"), []byte("4"), []byte("5"),
	[]byte("6"), []byte("7"), []byte("8"), []byte("9"),
}

func rewindForPartialMulMatch(data []byte, rewind int) int {
	for _, partialMulMatch := range partialMulMatches {
		index := bytes.LastIndex(data, partialMulMatch)
		if index > -1 && index == len(data)-len(partialMulMatch) {
			return rewindForPartialMulMatch(data[:index], rewind+len(partialMulMatch))
		}
	}
	return rewind
}

var partialTokenMatches [][]byte = [][]byte{
	[]byte("d"), []byte("do"), []byte("do("),
	[]byte("don"), []byte("don'"), []byte("don't"), []byte("don't("),
	[]byte("m"), []byte("mu"), []byte("mul"), []byte("mul("),
}

func rewindForPartialTokenMatch(data []byte, rewind int) int {
	for _, partialTokenMatch := range partialTokenMatches {
		index := bytes.LastIndex(data, partialTokenMatch)
		if index > -1 && index == len(data)-len(partialTokenMatch) {
			return rewindForPartialTokenMatch(data[:index], rewind+len(partialTokenMatch))
		}
	}
	return rewind
}
