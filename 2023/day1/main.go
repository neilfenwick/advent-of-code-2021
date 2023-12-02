package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	var (
		file *os.File
		err  error
	)

	file, err = os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("Error opening file: %s", os.Args[1])
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	coordsPt1 := parseCoordinates(file, func(i string) string { return i })
	sumPt1 := sumCoords(coordsPt1)
	fmt.Printf("Part1 sum: %d\n", sumPt1)

	file.Seek(0, 0)

	coordsPt2 := parseCoordinates(file, sequentialReplaceNumberWords)
	sumPt2 := sumCoords(coordsPt2)
	fmt.Printf("Part2 sum: %d\n", sumPt2)
}

var wordToDigitMap = map[string]rune{
	"one":   '1',
	"two":   '2',
	"three": '3',
	"four":  '4',
	"five":  '5',
	"six":   '6',
	"seven": '7',
	"eight": '8',
	"nine":  '9',
}

type digitParseFuncStrategy func(string) string

func parseCoordinates(reader io.Reader, parser digitParseFuncStrategy) []string {
	scanner := bufio.NewScanner(reader)
	results := make([]string, 0)
	for scanner.Scan() {
		line := parser(scanner.Text())
		first, err := findFirstDigit(line)
		if err != nil {
			log.Println(err)
			continue
		}

		digits := string(first)
		last, err := findLastDigit(line)
		if err != nil {
			log.Println(err)
			continue
		}

		digits += string(last)
		results = append(results, digits)
	}

	return results
}

func sequentialReplaceNumberWords(line string) string {
	var result, previous string
	previous = ""
    result = line
	for result != previous {
        previous = result
		result = replaceFirstNumberWord(result)
	}

	return result
}

func replaceFirstNumberWord(line string) string {
	indexes := make([]int, 0)
	indexToWordLookup := make(map[int]string)
	result := line

	for k := range wordToDigitMap {
		idx := strings.Index(line, k)
		if idx == -1 {
			continue
		}
		indexToWordLookup[idx] = k
		indexes = append(indexes, idx)
	}

	if len(indexes) == 0 {
		return result
	}

	slices.Sort(indexes)
	firstWord := indexToWordLookup[indexes[0]]
	result = strings.Replace(result, firstWord, string(wordToDigitMap[firstWord]), 1)

	return result
}

func findFirstDigit(line string) (rune, error) {
	for _, char := range line {
		switch char {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return char, nil
		}
	}

	return 0, errors.New("no digit found:" + line)
}

func findLastDigit(line string) (rune, error) {
	for i := len(line) - 1; i >= 0; i-- {
		switch line[i] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return rune(line[i]), nil
		}
	}

	return 0, errors.New("no last digit found")
}

func sumCoords(coords []string) int {
	var result int
	for _, txt := range coords {
		val, err := strconv.Atoi(txt)
		if err != nil {
			fmt.Printf("Could not convert %s to int\n", txt)
		}
		result += val
	}

	return result
}
