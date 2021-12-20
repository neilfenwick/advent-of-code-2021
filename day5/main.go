package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var ventMap *sparseVentMap = NewVentMap()

func main() {
	var (
		file *os.File
		err  error
	)

	file, err = os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("Error opening file: %s", os.Args[1])
	}

	result := countVentDensity(file, 1)
	fmt.Println(result)
}

func countVentDensity(r io.Reader, threshold int) int {
	vectors := parseVentLocations(r)
	buildVentMap(vectors)
	return ventMap.CountVentIntersectionsOverThreshold(threshold)
}

func parseVentLocations(r io.Reader) []*ventVector {
	var (
		result []*ventVector = make([]*ventVector, 0, 100)
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
		fmt.Sscanf(s.Text(), "%d,%d -> %d,%d", &start.x, &start.y, &end.x, &end.y)
		result = append(result, &ventVector{start: start, end: end})
	}

	return result
}

func buildVentMap(vectors []*ventVector) {
	for _, v := range vectors {
		ventMap.AddVector(v)
	}
}
