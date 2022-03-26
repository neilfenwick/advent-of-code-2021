package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal("Unable to open input.txt file")
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	registerMap := make(map[string]int)

	max := 0
	s := bufio.NewScanner(file)
	for s.Scan() {
		var reg, ins, reg2, op string
		var delta, referenceValue int
		_, _ = fmt.Sscanf(s.Text(), "%s %s %d if %s %s %d", &reg, &ins, &delta, &reg2, &op, &referenceValue)

		registerOperation := newInstruction(ins, delta)
		conditionOperation := newCondition(op, referenceValue)

		regVal := registerMap[reg]
		refVal := registerMap[reg2]

		if conditionOperation.TestFunc(refVal) {
			result := registerOperation.OpFunc(regVal)
			registerMap[reg] = result
			if result > max {
				max = result
			}
		}
	}

	fmt.Printf("%v\n\nMax register value: %d\n", registerMap, max)
}

type instruction struct {
	register string
	OpFunc   func(int) int
}

func newInstruction(op string, val int) instruction {
	var f func(int) int
	switch op {
	case "dec":
		f = func(ref int) int { return ref - val }
	case "inc":
		f = func(ref int) int { return ref + val }
	}

	return instruction{OpFunc: f}
}

type condition struct {
	TestFunc func(int) bool
}

func newCondition(op string, val int) condition {
	var f func(int) bool
	switch op {
	case "==":
		f = func(ref int) bool { return val == ref }
	case "!=":
		f = func(ref int) bool { return val != ref }
	case ">=":
		f = func(ref int) bool { return ref >= val }
	case "<=":
		f = func(ref int) bool { return ref <= val }
	case "<":
		f = func(ref int) bool { return ref < val }
	case ">":
		f = func(ref int) bool { return ref > val }
	}
	return condition{TestFunc: f}
}
