package main

type point struct {
	col int
	row int
}

type octopus struct {
	location   point
	energy     int
	hasFlashed bool
}

func (o *octopus) Step() bool {
	if o.hasFlashed {
		return false
	}
	o.energy++
	if o.energy > 9 {
		o.hasFlashed = true
	}
	return o.hasFlashed
}

func (o *octopus) Reset() {
	o.hasFlashed = false
	o.energy = 0
}
