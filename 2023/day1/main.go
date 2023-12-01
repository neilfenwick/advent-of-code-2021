package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
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

	coords := parseCoordinates(file)
    sum := sumCoords(coords)
	fmt.Println(sum)
}

func parseCoordinates(reader io.Reader) []string {
	scanner := bufio.NewScanner(reader)
	results := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		var digits string
		for _, char := range line {
			switch char {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				digits = string(char)
				goto nextchar
			}
		}

	nextchar:
		for i := len(line) - 1; i >= 0; i-- {
			switch line[i] {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				digits += string(line[i])
				goto done
			}
		}

	done:
		results = append(results, digits)
	}

	return results
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
