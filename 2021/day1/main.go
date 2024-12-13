package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	data "github.com/neilfenwick/advent-of-code/data_structures"
)

func main() {
	var (
		file       *os.File
		windowSize = 1
		err        error
	)

	switch len(os.Args) {
	case 1:
		file = os.Stdin
	case 3:
		windowSize, err = strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("Did not understand window size of: %s\n", os.Args[2])
		}

		fallthrough
	case 2:
		file, err = os.Open(os.Args[1])
		if err != nil {
			log.Fatalf("Error opening file: %s", os.Args[1])
		}
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	count := depthIncreasesCount(file, windowSize)
	fmt.Println(count)
}

func depthIncreasesCount(r io.Reader, windowSize int) int {
	var (
		count, line int
		buffer      = data.NewCircularBuffer(windowSize + 1)
		scanner     = bufio.NewScanner(r)
	)

	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		current, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("Unexpected error converting '%s' to int", scanner.Text())
		}

		buffer.Write(current)
		if line > windowSize-1 {
			if sumWindow(buffer.Read(-windowSize, windowSize)) > sumWindow(buffer.Read(-windowSize-1, windowSize)) {
				count++
			}
		}
		line++
	}

	return count
}

func sumWindow(numbers []interface{}) int {
	var total int
	for _, v := range numbers {
		total += v.(int)
	}
	return total
}
