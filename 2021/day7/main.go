package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type alignmentPosition struct {
	position int
	fuelCost int64
}

type inputPositions struct {
	positionArray []int
	min           int
	max           int
}

func main() {
	var (
		file *os.File
		err  error
	)

	file, err = os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("Error opening file: %s", os.Args[1])
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	result := calcLeastEditDistance(file)
	fmt.Printf("%+v\n", result)
}

func calcLeastEditDistance(r io.Reader) alignmentPosition {
	inputPositions := readInputPositions(r)
	allPositions := make(map[int]alignmentPosition)
	for i := inputPositions.min; i <= inputPositions.max; i++ {
		var (
			fuelCost int64
		)
		for _, pos := range inputPositions.positionArray {
			distanceToMove := int64(math.Abs(float64(pos) - float64(i)))
			for d := distanceToMove; d > 0; d-- {
				fuelCost = fuelCost + d
			}
		}
		allPositions[i] = alignmentPosition{position: i, fuelCost: fuelCost}
	}

	// find the lowest fuel cost
	result := alignmentPosition{fuelCost: math.MaxInt32}
	for _, alignment := range allPositions {
		if alignment.fuelCost < result.fuelCost {
			result = alignment
		}
	}
	return result
}

func readInputPositions(r io.Reader) *inputPositions {
	var (
		result = inputPositions{}
	)

	result.positionArray = make([]int, 0, 1000)
	result.min = math.MaxInt32
	result.max = math.MinInt32

	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		numbers := strings.Split(s.Text(), ",")
		for _, n := range numbers {
			pos, err := strconv.Atoi(n)
			if err != nil {
				log.Fatalf("Could not convert '%s' to int", n)
			}
			if pos < result.min {
				result.min = pos
			}
			if pos > result.max {
				result.max = pos
			}
			result.positionArray = append(result.positionArray, pos)
		}
	}
	return &result
}
