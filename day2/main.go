package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
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
	defer file.Close()

	depth, distance := calcPositionWithAim(file)

	fmt.Println(depth * distance)
}

func calcPosition(r io.Reader) (int, int) {
	var (
		direction                      string
		distance, totalDistance, depth int
	)
	lineScanner := bufio.NewScanner(r)

	for lineScanner.Scan() {
		line := lineScanner.Text()
		fmt.Sscanf(strings.TrimSpace(line), "%s %d", &direction, &distance)
		switch direction {
		case "forward":
			totalDistance += distance
		case "down":
			depth += distance
		case "up":
			depth -= distance
		}
	}

	return depth, totalDistance
}

func calcPositionWithAim(r io.Reader) (int, int) {
	var (
		direction                           string
		distance, totalDistance, depth, aim int
	)
	lineScanner := bufio.NewScanner(r)

	for lineScanner.Scan() {
		line := lineScanner.Text()
		fmt.Sscanf(strings.TrimSpace(line), "%s %d", &direction, &distance)
		switch direction {
		case "forward":
			totalDistance += distance
			depth += distance * aim
		case "down":
			aim += distance
		case "up":
			aim -= distance
		}
	}

	return depth, totalDistance
}
