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
	pointCount, _ := countGuardPathPointsVisited(grid)
	fmt.Printf("Guard visited %d points\n", pointCount)

	// I did this the brute force way by iterating over all points and adding an obstacle to each point
	// and checking if the guard is stuck in a loop. Is there a more efficient way to do this?
	loopObstructionCount := countLoopObstructions(grid)
	fmt.Printf("There are %d points that cause the guard to loop\n", loopObstructionCount)
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

func countGuardPathPointsVisited(grid *obstacleGrid) (int, error) {
	guardPos := grid.guardStartPos
	guardDirection := grid.guardStartDirection
	pointsVisited := make(map[point]vector, grid.sizeX*grid.sizeY)

	for guardPos.x >= 0 && guardPos.x < grid.sizeX && guardPos.y >= 0 && guardPos.y < grid.sizeY {
		if prevDirection, found := pointsVisited[guardPos]; !found {
			pointsVisited[guardPos] = guardDirection
		} else {
			if prevDirection == guardDirection {
				return 0, fmt.Errorf("guard is stuck in a loop")
			}
		}

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
			guardPos = nextPos
		}
	}

	return len(pointsVisited), nil
}

func countLoopObstructions(grid *obstacleGrid) int {
	loopCount := 0

	for x := range grid.sizeX {
		for y := range grid.sizeY {
			// Skip if there is already an obstacle at this point, or it is the guard start position,
			// or directly in front of the guard start position
			if grid.obstacles[point{x: x, y: y}] ||
				(x == grid.guardStartPos.x && y == grid.guardStartPos.y) ||
				(x == grid.guardStartPos.x+grid.guardStartDirection.x && y == grid.guardStartPos.y+grid.guardStartDirection.y) {
				continue
			}

			// add an obstacle and check if the guard is stuck in a loop
			grid.obstacles[point{x: x, y: y}] = true
			if _, err := countGuardPathPointsVisited(grid); err != nil {
				loopCount++
			}

			// remove the obstacle because this same instance of the map is used for all iterations
			delete(grid.obstacles, point{x: x, y: y})

		}
	}

	return loopCount
}
