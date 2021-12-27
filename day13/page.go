package main

import "strings"

type point struct {
	x, y int
}

type page struct {
	data         map[point]bool
	instructions []foldInstruction
}

func (p *page) toString() string {
	var (
		maxX, maxY int
	)
	for pnt := range p.data {
		if pnt.x > maxX {
			maxX = pnt.x
		}
		if pnt.y > maxY {
			maxY = pnt.y
		}
	}
	builder := strings.Builder{}
	builder.Grow(maxX * maxY)
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if _, found := p.data[point{x: x, y: y}]; found {
				builder.WriteRune('#')
			} else {
				builder.WriteRune(' ')
			}
		}
		builder.WriteString("\n")
	}
	return builder.String()
}
