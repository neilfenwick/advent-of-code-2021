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

var (
	grid              = make(map[point]*octopus)
	cumulativeFlashed = make(map[point]bool, 100)
)

type flashStats struct {
	totalIterations, numberOfFlashes int
	iterationsWhereAllFlashed        []int
}

func main() {
	var (
		file               *os.File
		err                error
		numberOfIterations = 100
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
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	buildOctopusMap(file)
	result := iterateSteps(numberOfIterations)
	fmt.Printf("%+v\n", *result)
}

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

func iterateSteps(count int) *flashStats {
	var (
		cumulativeFlashCount int
		result               = flashStats{totalIterations: count, iterationsWhereAllFlashed: make([]int, 0)}
		mapPoIntegers        = make([]point, 0, len(grid))
	)
	for point := range grid {
		mapPoIntegers = append(mapPoIntegers, point)
	}
	for i := 0; i < count; i++ {
		step(mapPoIntegers)
		cumulativeFlashCount += len(cumulativeFlashed)
		if len(cumulativeFlashed) == len(grid) {
			result.iterationsWhereAllFlashed = append(result.iterationsWhereAllFlashed, i+1)
		}
		reset()
	}
	result.numberOfFlashes = cumulativeFlashCount
	return &result
}

func step(poIntegers []point) {
	var (
		flashedPoIntegers = make([]point, 0, 20)
	)
	for _, point := range poIntegers {
		if octopus, found := grid[point]; found && octopus.Step() {
			flashedPoIntegers = append(flashedPoIntegers, point)
		}
	}
	for _, fp := range flashedPoIntegers {
		cumulativeFlashed[fp] = true
	}
	if len(flashedPoIntegers) > 0 {
		poIntegersToStep := expandToNeighbours(flashedPoIntegers)
		step(poIntegersToStep)
	}
}

func expandToNeighbours(poIntegers []point) []point {
	var (
		result = make([]point, 0, 8*len(poIntegers))
	)
	for _, p := range poIntegers {
		left := point{col: p.col - 1, row: p.row}
		right := point{col: p.col + 1, row: p.row}
		up := point{col: p.col, row: p.row - 1}
		down := point{col: p.col, row: p.row + 1}

		upLeft := point{col: p.col - 1, row: p.row - 1}
		upRight := point{col: p.col + 1, row: p.row - 1}
		downLeft := point{col: p.col - 1, row: p.row + 1}
		downRight := point{col: p.col + 1, row: p.row + 1}

		result = append(result, p, left, right, up, down, upLeft, upRight, downLeft, downRight)
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
