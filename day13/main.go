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
	var (
		file *os.File
		err  error
	)
	file, err = os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("Error opening file: %s", os.Args[1])
	}
	defer file.Close()

	page := readInputToPage(file)
	firstFold := page.instructions[0]
	firstFold.fold(page.data)
	fmt.Printf("Points after first fold: %d\n\n", len(page.data))
	for _, instruction := range page.instructions[1:] {
		instruction.fold(page.data)
	}
	fmt.Print(page.toString())
}

func readInputToPage(r io.Reader) *page {
	var (
		result = &page{data: make(map[point]bool, 1000), instructions: make([]foldInstruction, 0)}
	)
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		if len(line) == 0 {
			continue
		}
		if strings.Contains(line, "fold along") {
			if instruction, ok := parseFoldInstruction(line); ok {
				result.instructions = append(result.instructions, instruction)
			}
		} else {
			pnt := parsePoint(line)
			result.data[pnt] = true
		}
	}
	return result
}

func parseFoldInstruction(line string) (foldInstruction, bool) {
	instruction := strings.ReplaceAll(line, "fold along ", "")
	parts := strings.Split(instruction, "=")
	value, _ := strconv.Atoi(parts[1])

	switch parts[0] {
	case "y":
		return &verticalFold{y: value}, true
	case "x":
		return &horiztonalFold{x: value}, true
	}
	return nil, false
}

func parsePoint(text string) point {
	parts := strings.Split(text, ",")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[1])
	pnt := point{x: x, y: y}
	return pnt
}
