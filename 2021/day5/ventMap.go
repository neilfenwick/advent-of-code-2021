package main

import (
	"math"
)

type sparseVentMap struct {
	vectors []*ventVector
	size    int
}

type point struct {
	x int
	y int
}

func (p *point) Max() int {
	if p.x > p.y {
		return p.x
	}
	return p.y
}

type ventVector struct {
	start point
	end   point
}

func (v *ventVector) Max() int {
	if v.start.Max() > v.end.Max() {
		return v.start.Max()
	}
	return v.end.Max()
}

// PointIsOnVector calculates whether a point is on a line by checking if the
// sum of the distance from each end of the line is the same
// as the length of the line
func (v *ventVector) PointIsOnVector(x int, y int) bool {
	pointToTest := point{x: x, y: y}

	// vector vectorLength
	vectorLength := v.Length()

	// distance from start
	toStart := ventVector{start: v.start, end: pointToTest}
	startSegmentLength := toStart.Length()

	// distance from end
	toEnd := ventVector{start: pointToTest, end: v.end}
	endSegmentLength := toEnd.Length()

	tolerance := 0.001
	diff := math.Abs(vectorLength - startSegmentLength - endSegmentLength)
	return diff < tolerance
}

func (v *ventVector) Length() float64 {
	return math.Sqrt(math.Pow(float64(v.end.x-v.start.x), 2) + math.Pow(float64(v.end.y-v.start.y), 2))
}

func NewVentMap() *sparseVentMap {
	m := sparseVentMap{}
	m.vectors = make([]*ventVector, 0, 100)
	return &m
}

func (m *sparseVentMap) AddVector(vector *ventVector) {

	m.vectors = append(m.vectors, vector)

	if vector.Max() > m.size {
		m.size = vector.Max()
	}
}

func consume(m *sparseVentMap, jobs <-chan *point, result chan<- int) {
	var count int

	for p := range jobs {
		var pointDensity int
		for _, v := range m.vectors {
			if v.PointIsOnVector(p.x, p.y) {
				pointDensity++
				if pointDensity > 1 {
					count++
					goto nextJob
				}
			}
		}
	nextJob:
	}
	result <- count
}

func (m *sparseVentMap) CountVentIntersectionsOverThreshold() int {
	var (
		count int
	)
	const workerCount int = 8
	poIntegersChannel := make(chan *point)
	results := make(chan int)

	go func() {
		for x := 0; x <= m.size; x++ {
			for y := 0; y <= m.size; y++ {
				poIntegersChannel <- &point{x: x, y: y}
			}
		}
		close(poIntegersChannel)
	}()

	for i := 1; i <= workerCount; i++ {
		go consume(m, poIntegersChannel, results)
	}

	for i := 1; i <= workerCount; i++ {
		count += <-results
	}
	return count
}
