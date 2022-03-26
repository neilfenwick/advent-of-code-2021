package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic("Could not read input.txt in current directory")
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	values := make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			panic("Error converting value in input to an integer")
		}
		values = append(values, num)
	}

	jumpList := NewList(values)
	jumpListNewStrategy := NewList(values)
	jumpListNewStrategy.OffsetCalcFunc(NewStrategyOffsetCalc)

	fmt.Printf("Took %d jumps to exit the list\n", jumpList.CalcJumps())
	fmt.Printf("New Strategy took %d jumps to exit the list\n", jumpListNewStrategy.CalcJumps())
}
