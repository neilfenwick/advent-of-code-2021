package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var ventMap = NewVentMap()

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

	result := countVentDensity(file)
	fmt.Println(result)
}

func countVentDensity(r io.Reader) int {
	vectors := parseVentLocations(r)
	buildVentMap(vectors)
	return ventMap.CountVentIntersectionsOverThreshold()
}

func parseVentLocations(r io.Reader) []*ventVector {
	var (
		result = make([]*ventVector, 0, 100)
	)
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		if len(strings.TrimSpace(s.Text())) == 0 {
			continue
		}

		var (
			start, end point
		)
		_, _ = fmt.Sscanf(s.Text(), "%d,%d -> %d,%d", &start.x, &start.y, &end.x, &end.y)
		result = append(result, &ventVector{start: start, end: end})
	}

	return result
}

func buildVentMap(vectors []*ventVector) {
	for _, v := range vectors {
		ventMap.AddVector(v)
	}
}
