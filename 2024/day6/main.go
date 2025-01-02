package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("Error opening file: %s", os.Args[1])
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	grid := parseInput(file)
	pointCount := countGuardPathPointsVisited(grid)
	fmt.Printf("Guard visited %d points\n", pointCount)
}

type point struct {
	x, y int
}

type vector point

var up = vector{x: 0, y: -1}
var down = vector{x: 0, y: 1}
var left = vector{x: -1, y: 0}
var right = vector{x: 1, y: 0}

type obstacleGrid struct {
	sizeX, sizeY        int
	obstacles           map[point]bool
	guardStartPos       point
	guardStartDirection vector
}

func parseInput(file *os.File) *obstacleGrid {
	var (
		width, height  int
		guardPos       point
		guardDirection vector
	)
	obstacles := make(map[point]bool)

	s := bufio.NewScanner(file)
	for s.Scan() {
		line := strings.TrimSpace(s.Text())

		for i, c := range line {
			switch c {
			case '#':
				point := point{x: i, y: height}
				obstacles[point] = true
			case '^':
				guardPos = point{x: i, y: height}
				guardDirection = up
			}
		}

		width = len(line)
		height++
	}

	grid := &obstacleGrid{
		sizeX:               width,
		sizeY:               height,
		obstacles:           obstacles,
		guardStartPos:       guardPos,
		guardStartDirection: guardDirection,
	}

	return grid
}

func countGuardPathPointsVisited(grid *obstacleGrid) int {
	guardPos := grid.guardStartPos
	guardDirection := grid.guardStartDirection
	pointsVisited := make(map[point]bool)

	for guardPos.x >= 0 && guardPos.x < grid.sizeX && guardPos.y >= 0 && guardPos.y < grid.sizeY {
		nextPos := point{x: guardPos.x + guardDirection.x, y: guardPos.y + guardDirection.y}
		if grid.obstacles[nextPos] {
			switch guardDirection {
			case up:
				guardDirection = right
			case right:
				guardDirection = down
			case down:
				guardDirection = left
			case left:
				guardDirection = up
			}
		} else {
			if _, ok := pointsVisited[guardPos]; !ok {
				pointsVisited[guardPos] = true
			}
			guardPos = nextPos
		}
	}

	return len(pointsVisited)
}
