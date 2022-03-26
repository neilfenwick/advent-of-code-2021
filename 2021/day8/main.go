package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strings"
)

var readings = make([]*displayReading, 0, 100)

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

	readings = readInputs(file)
	uniqueCount := countUniqueValues()
	sum := sumReadings()
	fmt.Printf("Unique Count: %d\nSum: %d\n", uniqueCount, sum)
}

func readInputs(r io.Reader) []*displayReading {
	result := make([]*displayReading, 0, 100)
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		line := s.Text()
		parts := strings.Split(line, "|")
		displayReading := NewDisplayReading(strings.Fields(parts[0]), strings.Fields(parts[1]))
		result = append(result, displayReading)
	}
	return result
}

func countUniqueValues() int {
	var (
		result int
	)
	for _, r := range readings {
		for _, n := range r.notes {
			switch len(n) {
			case 2, 3, 4, 7:
				result++
			}
		}
	}
	return result
}

func sumReadings() int {
	var (
		sum int
	)
	for _, r := range readings {
		var reading int
		for i := 0; i < 4; i++ {
			multiplier := int(math.Pow(10, float64(3-i)))
			note := r.notes[i]
			number := r.numberMap[note]
			reading += number * multiplier
		}
		sum += reading
	}
	return sum
}
