package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic("Cannot read input")
	}
	str := strings.TrimSpace(string(input))
	chars := strings.Split(str, "")
	integers := make([]int, len(chars))
	for pos, str := range chars {
		integers[pos], err = strconv.Atoi(str)
		if err != nil {
			panic(err.Error)
		}
	}
	consecutiveSum := SumConsecutiveIntegers(integers)
	oppositeSum := SumOppositeIntegers(integers)

	fmt.Printf("Consecutive numbers result: %d\n", consecutiveSum)
	fmt.Printf("Opposite numbers result: %d\n", oppositeSum)
}
