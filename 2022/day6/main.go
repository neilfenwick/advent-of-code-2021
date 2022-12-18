package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/neilfenwick/advent-of-code/data"
)

func main() {
	var (
		file *os.File
		err  error
	)

	signalLength := 0

	switch len(os.Args) {
	case 1:
		file = os.Stdin
	case 3:
		signalLength, _ = strconv.Atoi(os.Args[2])
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

	count := searchForMarker(file, signalLength)
	fmt.Printf("Token count until complete signal: %d\n", count)
}

func searchForMarker(file *os.File, signalLength int) int {
	s := bufio.NewScanner(file)
	s.Split(bufio.ScanRunes)

	buffer := data.NewCircularBuffer(signalLength)
	tokenCount := 0

	for s.Scan() {
		tokenCount++
		r := []rune(s.Text())[0]
		buffer.Write(r)
		if isHeterogenousTokens(buffer) {
			return tokenCount
		}
	}
	return 0
}

func isHeterogenousTokens(buffer *data.CircularBuffer) bool {
	dict := make(map[rune]int, buffer.Size())
	data := buffer.Read(0, buffer.Size())
	for _, r := range data {
		if r == nil {
			return false
		}

		_, found := dict[r.(rune)]
		if found {
			return false
		}

		dict[r.(rune)] = 1
	}

	return true
}
