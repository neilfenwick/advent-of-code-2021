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
		file               *os.File
		err                error
		numberOfIterations int = 100
	)
	switch len(os.Args) {
	case 1:
		file = os.Stdin
	case 3:
		numberOfIterations, err = strconv.Atoi(os.Args[2])
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

	buildOctopusMap(file)
	result := iterateSteps(numberOfIterations)
	fmt.Printf("%+v\n", *result)
}

var grid map[point]*octopus = make(map[point]*octopus)

func buildOctopusMap(r io.Reader) {
	var (
		rowIndex int
	)
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		row := []rune(strings.TrimSpace(s.Text()))
		for i, energy := range row {
			pnt := point{row: rowIndex, col: i}
			energyVal, _ := strconv.Atoi(string(energy))
			octopus := octopus{location: pnt, energy: energyVal}
			grid[pnt] = &octopus
		}
		rowIndex++
	}
}

type flashStats struct {
	totalIterations, numberOfFlashes int
	iterationsWhereAllFlashed        []int
}

func iterateSteps(count int) *flashStats {
	var (
		cumulativeFlashCount int
		result               flashStats = flashStats{totalIterations: count, iterationsWhereAllFlashed: make([]int, 0)}
		mapPoints            []point    = make([]point, 0, len(grid))
	)
	for point := range grid {
		mapPoints = append(mapPoints, point)
	}
	for i := 0; i < count; i++ {
		step(mapPoints)
		cumulativeFlashCount += len(cumulativeFlashed)
		if len(cumulativeFlashed) == len(grid) {
			result.iterationsWhereAllFlashed = append(result.iterationsWhereAllFlashed, i+1)
		}
		reset()
	}
	result.numberOfFlashes = cumulativeFlashCount
	return &result
}

var cumulativeFlashed map[point]bool = make(map[point]bool, 100)

func step(points []point) {
	var (
		flashedPoints []point = make([]point, 0, 20)
	)
	for _, point := range points {
		if octopus, found := grid[point]; found && octopus.Step() {
			flashedPoints = append(flashedPoints, point)
		}
	}
	for _, fp := range flashedPoints {
		cumulativeFlashed[fp] = true
	}
	if len(flashedPoints) > 0 {
		pointsToStep := expandToNeighbours(flashedPoints)
		step(pointsToStep)
	}
}

func expandToNeighbours(points []point) []point {
	var (
		result []point = make([]point, 0, 8*len(points))
	)
	for _, p := range points {
		left := point{col: p.col - 1, row: p.row}
		right := point{col: p.col + 1, row: p.row}
		up := point{col: p.col, row: p.row - 1}
		down := point{col: p.col, row: p.row + 1}

		upleft := point{col: p.col - 1, row: p.row - 1}
		upright := point{col: p.col + 1, row: p.row - 1}
		downleft := point{col: p.col - 1, row: p.row + 1}
		downRight := point{col: p.col + 1, row: p.row + 1}

		result = append(result, p, left, right, up, down, upleft, upright, downleft, downRight)
	}
	return result
}

func reset() {
	for point := range cumulativeFlashed {
		octopus := grid[point]
		octopus.Reset()
	}
	cumulativeFlashed = make(map[point]bool, 100)
}
