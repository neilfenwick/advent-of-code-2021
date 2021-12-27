package main

type foldInstruction interface {
	fold(data map[point]bool)
}

type horiztonalFold struct {
	x int
}

func (hf *horiztonalFold) fold(data map[point]bool) {
	for pnt := range data {
		if pnt.x > hf.x {
			delete(data, pnt)
			delta := pnt.x - hf.x
			editedPoint := point{x: hf.x - delta, y: pnt.y}
			data[editedPoint] = true
		}
	}
}

type verticalFold struct {
	y int
}

func (vf *verticalFold) fold(data map[point]bool) {
	for pnt := range data {
		if pnt.y > vf.y {
			delete(data, pnt)
			delta := pnt.y - vf.y
			editedPoint := point{x: pnt.x, y: vf.y - delta}
			data[editedPoint] = true
		}
	}
}
