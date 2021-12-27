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

type point struct {
	x, y int
}

type page struct {
	data         map[point]bool
	instructions []foldInstruction
}

func (p *page) toString() string {
	var (
		maxX, maxY int
	)
	for pnt := range p.data {
		if pnt.x > maxX {
			maxX = pnt.x
		}
		if pnt.y > maxY {
			maxY = pnt.y
		}
	}
	builder := strings.Builder{}
	builder.Grow(maxX * maxY)
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if _, found := p.data[point{x: x, y: y}]; found {
				builder.WriteRune('#')
			} else {
				builder.WriteRune(' ')
			}
		}
		builder.WriteString("\n")
	}
	return builder.String()
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
