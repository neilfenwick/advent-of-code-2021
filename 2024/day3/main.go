package main

import (
	"bufio"
	"fmt"
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
}

type operands struct {
	first  int
	second int
}

func parseInput(file *os.File) []operands {
	operandsList := make([]operands, 0)

	// pattern is a regex that matches the format "mul(x,y)" where x and y are numbers.
	// - (?:mul\()   : Non-capturing group that matches the literal string "mul(".
	// - (\d+)       : Capturing group that matches one or more digits (represents the first number).
	// - (?:,)       : Non-capturing group that matches a comma ",".
	// - (\d+)       : Capturing group that matches one or more digits (represents the second number).
	// - (?:\))      : Non-capturing group that matches the closing parenthesis ")".
	pattern := `(?:mul\()(\d+)(?:,)(\d+)(?:\))`
	re := regexp.MustCompile(pattern)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
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
	}

	return operandsList
}
