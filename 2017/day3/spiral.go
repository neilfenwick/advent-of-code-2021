package main

import (
	"fmt"
	"math"
)

// SpiralGrid represents a set of integers that are mapped around
// the center of a cartesian plane, spiraling outwards in an
// anti-clockwise direction from the positive x-axis
type SpiralGrid struct {
	grid        map[int]point
	pointValues map[point]int
}

type direction int

const (
	up direction = iota * 90
	right
	down
	left
)

// ManhattanDistance returns the sum of the x and y components
// of a diagonal, from the value to the centre of a spiral grid
func (s *SpiralGrid) ManhattanDistance(val int) int {
	if val == 1 {
		return 0
	}

	ring := nearestOddSquareRootCeil(val)
	distToAxis := ring
	ringXAxisValue := int(math.Pow(float64(ring-2), 2)) + ring/2
	for side := 0; side < 4; side++ {
		midpoint := ringXAxisValue + (side * (ring - 1))
		dist := int(math.Abs(float64(val - midpoint)))
		if dist < distToAxis {
			distToAxis = dist
		}
	}

	return distToAxis + (ring-1)/2
}

func nearestOddSquareRootCeil(val int) int {
	sqrt := int(math.Sqrt(float64(val)))

	if sqrt%2 == 0 {
		return sqrt + 1
	}

	return sqrt + 2
}

// CumulativeSumToPosition returns the sum of all preceding
// neighbours on the spiral that are within 1 x or y co-ordinate
// distance away
func (s *SpiralGrid) CumulativeSumToPosition(val int) int {
	s.generateGridToPoint(val)

	var pointCnt int
	if s.pointValues == nil {
		s.pointValues = make(map[point]int, val)
		pointCnt = 0
	} else {
		pointCnt = len(s.pointValues)
	}

	// start at the head of the point values map
	for i := pointCnt + 1; i <= val; i++ {
		pnt, found := s.grid[i]
		if !found {
			panic(fmt.Sprintf("No grid value exists at position %d", i))
		}

		if i == 1 || i == 2 {
			s.pointValues[pnt] = 1
			continue
		}

		neighbours := s.neighboursFor(i)
		cumulativeSum := 0
		for _, neighbour := range neighbours {
			cumulativeSum += s.pointValues[neighbour]
		}
		s.pointValues[pnt] = cumulativeSum
	}

	lastPoint := s.grid[val]
	return s.pointValues[lastPoint]
}

func (s *SpiralGrid) generateGridToPoint(val int) {
	if s.grid == nil {
		s.grid = make(map[int]point, val)
	}

	var dir direction

	// start at the head of the grid and step
	for pos := len(s.grid) + 1; pos <= val; pos++ {
		if pos == 1 {
			s.grid[1] = point{0, 0}
			continue
		}

		head := s.grid[pos-1]

		ring := nearestOddSquareRootCeil(pos)
		innerRing := ring - 2
		innerRingMax := int(math.Pow(float64(innerRing), 2))

		// special case for bottom-right corner of each ring (perfect-square)
		if innerRingMax == pos {
			ring -= 2
			innerRing -= 2
			innerRingMax = int(math.Pow(float64(innerRing), 2))
		}

		if pos == innerRingMax+1 {
			dir = right
		} else {
			dir = up
			nextCorner := innerRingMax + ring
			for nextCorner < pos {
				nextCorner += ring - 1
				dir = (dir + 270) % 360
			}
			for i := innerRingMax + 1; i <= pos; i++ {
				if i == nextCorner {
					dir = (dir + 270) % 360
				}
			}
		}

		next := nextPnt(head, dir)
		s.grid[pos] = next
	}
}

func nextPnt(pnt point, dir direction) point {
	if pnt.x == 0 && pnt.y == 0 {
		return point{1, 0}
	}
	switch dir {
	case up:
		return point{pnt.x, pnt.y + 1}
	case left:
		return point{pnt.x - 1, pnt.y}
	case down:
		return point{pnt.x, pnt.y - 1}
	case right:
		return point{pnt.x + 1, pnt.y}
	}
	panic(fmt.Sprintf("Expected direction 0, 90, 180, 270. Got %d.", dir))
}

// neighboursFor finds all poIntegers whose position occurs earlier
// in the spiral and whose pythagorean distance is <= sqrt(2) away
// from the point at the location of the parameter value
func (s *SpiralGrid) neighboursFor(val int) map[int]point {
	referencePoint := s.grid[val]
	result := make(map[int]point)
	for i := 1; i < val; i++ {
		pnt := s.grid[i]
		dist := math.Sqrt(math.Pow(float64(referencePoint.x-pnt.x), 2) + math.Pow(float64(referencePoint.y-pnt.y), 2))
		if dist <= math.Sqrt(2) {
			result[i] = pnt
		}
	}
	return result
}
