package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

var heightmap map[int][]int = make(map[int][]int, 100)

type point struct {
	x, y int
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

	buildHeightMap(file)

	lowPointRiskSum := sumLowPointRisk()
	fmt.Printf("Low point risk sum: %d\n", lowPointRiskSum)

	basinSizeSum := findProductOfBasinSizes()
	fmt.Printf("Basin size sum: %d\n", basinSizeSum)
}

func buildHeightMap(r io.Reader) {
	var (
		rowIndex int
	)

	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)

	for s.Scan() {
		inputRow := []rune(strings.TrimSpace(s.Text()))
		if len(inputRow) == 0 {
			continue
		}
		row := make([]int, 0, len(inputRow))
		for _, d := range inputRow {
			row = append(row, int(d)-48) // minus 48 to get from ascii table index to actual int value
		}
		heightmap[rowIndex] = row
		rowIndex++
	}
}

func sumLowPointRisk() int {
	var (
		sum int
	)
	lowPointValues, _ := scanForLowPoints()
	for _, v := range lowPointValues {
		sum += (1 + v)
	}
	return sum
}

func scanForLowPoints() ([]int, []point) {
	var (
		lowPointHeights = make([]int, 0, 50)
		lowPoints       = make([]point, 0, 50)
	)

	for y := 0; y < len(heightmap); y++ {
		currentRow := heightmap[y]
		for x := 0; x < len(currentRow); x++ {
			location := point{x: x, y: y}
			lowerPoints := findNeighboursHigherOrLower(location, Lower)
			if len(lowerPoints) == 0 {
				lowPointHeights = append(lowPointHeights, currentRow[x])
				lowPoints = append(lowPoints, location)
			}
		}
	}
	return lowPointHeights, lowPoints
}

type HigherOrLowerStrategy int

const (
	Lower  HigherOrLowerStrategy = -1
	Higher HigherOrLowerStrategy = 1
)

func findNeighboursHigherOrLower(p point, strategy HigherOrLowerStrategy) []point {
	var (
		result = make([]point, 0, 4)
	)
	currentRow, found := heightmap[p.y]
	if !found {
		return []point{}
	}
	currentVal := currentRow[p.x]

	up := point{y: p.y - 1, x: p.x}
	if upVal, found := getValue(up); found && (upVal-currentVal)*int(strategy) >= 0 {
		switch {
		case strategy == Higher && upVal != 9:
			fallthrough
		case strategy == Lower:
			result = append(result, up)
		}
	}

	right := point{y: p.y, x: p.x + 1}
	if rightVal, found := getValue(right); found && (rightVal-currentVal)*int(strategy) >= 0 {
		switch {
		case strategy == Higher && rightVal != 9:
			fallthrough
		case strategy == Lower:
			result = append(result, right)
		}
	}

	down := point{y: p.y + 1, x: p.x}
	if downVal, found := getValue(down); found && (downVal-currentVal)*int(strategy) >= 0 {
		switch {
		case strategy == Higher && downVal != 9:
			fallthrough
		case strategy == Lower:
			result = append(result, down)
		}
	}

	left := point{y: p.y, x: p.x - 1}
	if leftVal, found := getValue(left); found && (leftVal-currentVal)*int(strategy) >= 0 {
		switch {
		case strategy == Higher && leftVal != 9:
			fallthrough
		case strategy == Lower:
			result = append(result, left)
		}
	}

	return result
}

func getValue(p point) (int, bool) {
	if p.x < 0 {
		return 0, false
	}

	row, found := heightmap[p.y]
	if !found || p.x >= len(row) {
		return 0, false
	}

	return row[p.x], true
}

func findProductOfBasinSizes() int {
	var (
		sum        int = 1
		basinSizes     = make([]int, 0, 100)
	)

	_, lowPoints := scanForLowPoints()

	for _, lowPoint := range lowPoints {
		neighbourMap := make(map[point]bool)
		neighbourMap[lowPoint] = true
		searchMapForHigherPoints(&neighbourMap)
		basinSizes = append(basinSizes, len(neighbourMap))
	}

	sort.Ints(basinSizes)
	topThree := basinSizes[len(basinSizes)-3:]

	for _, v := range topThree {
		sum *= v
	}
	return sum
}

func searchMapForHigherPoints(points *map[point]bool) {
	currentCount := len(*points)
	floorMap := *points
	for p := range floorMap {
		higherNeighbours := findNeighboursHigherOrLower(p, Higher)
		for _, neighbour := range higherNeighbours {
			floorMap[neighbour] = true
		}
		newCount := len(floorMap)
		if newCount > currentCount {
			searchMapForHigherPoints(points)
		}
	}
}
