package main

import (
	"fmt"
)

func main() {
	input := []int{0, 5, 10, 0, 11, 14, 13, 4, 11, 8, 8, 7, 1, 4, 12, 11}
	c := make([]int, len(input))
	copy(c, input)
	result := Rebalance(input)
	fmt.Printf("Rebalanced %v in %d iterations.\n Final state %v\n", c, result, input)

	loopCycle := Rebalance(input)
	fmt.Printf("Inifnite loop cycle is %d iterations.\n", loopCycle)
}
