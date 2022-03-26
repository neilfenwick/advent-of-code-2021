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
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	s := NewScanner(file)

	diffChecksum := NewChecksum()
	modChecksum := NewChecksum()
	modChecksum.Checksum(ModulusChecksum)

	for s.Scan() {
		row := s.Values()
		diffChecksum.Add(row)
		modChecksum.Add(row)
	}

	if err := s.scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Diff Checksum: %d\n", diffChecksum.Value())
	fmt.Printf("Modulus Checksum: %d\n", modChecksum.Value())
}

// IntScanner decorates a Scanner and returns integer slices as output
type IntScanner struct {
	scanner *bufio.Scanner
}

// NewScanner creates a Scanner and wraps it within an IntScanner
func NewScanner(r io.Reader) *IntScanner {
	scanner := bufio.NewScanner(r)
	return &IntScanner{
		scanner: scanner,
	}
}

// Scan moves the encapsulated Scanner and returns a boolean to indicate
// whether more data is available or whether EOF
func (s *IntScanner) Scan() bool {
	return s.scanner.Scan()
}

// Values parses the current Scanner text buffer to a slice of integers
// Values will panic if the current text buffer cannot be parsed as integers
func (s *IntScanner) Values() []int {
	str := strings.TrimSpace(s.scanner.Text())
	chars := strings.Fields(str)
	integers := make([]int, len(chars))
	var err error
	for pos, str := range chars {
		integers[pos], err = strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
	}
	return integers
}

// Err returns the last error that occurred within the Scanner
func (s *IntScanner) Err() error {
	return s.scanner.Err()
}
