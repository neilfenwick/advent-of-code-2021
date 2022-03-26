package main

import (
	"fmt"
)

func main() {
	spiral := SpiralGrid{}
	input := 277678
	manhattanDistance := spiral.ManhattanDistance(input)
	fmt.Printf("Manhattan distance for %d: %d\n", input, manhattanDistance)

	var cumulativeSum int
	for i := 1; cumulativeSum <= input; i++ {
		cumulativeSum = spiral.CumulativeSumToPosition(i)
	}
	fmt.Printf("Cumulative spiral distance for %d: %d\n", input, cumulativeSum)
}
