package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	data "github.com/neilfenwick/advent-of-code/data_structures"
)

func main() {
	var (
		file *os.File
		err  error
	)

	strategy = processMoveInstructions
	switch len(os.Args) {
	case 1:
		file = os.Stdin
	case 3:
		inputStrategy, _ := strconv.Atoi(os.Args[2])
		if inputStrategy > 0 {
			strategy = processMoveInstructionsPart2
		}
		fallthrough
	case 2:
		file, err = os.Open(os.Args[1])
		if err != nil {
			log.Fatalf("Error opening file: %s", os.Args[1])
		}
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	/*
	   This is a stacking problem. Parse the input into N stacks
	   Each input header is fixed width, so it is possible to determine the number
	   of stacks from the width of the input.

	   # of stacks = line width + 1 / 4

	   The input is parsed 'top down' and read first-in, first-out so that implies
	   using a queue data structure.

	   But then when we switch to moving crates around, we need the behaviour of a
	   stack 🏗️.

	   I already have a stack data structure from previous years, so can use that as
	   a template for my own queue. These are fairly trivial data structures, so I don't
	   really need to go looking for a collections package, and will just stick with
	   something simple. There is no multi-threading going on here, so nothing to worry
	   about with synchronising access to copies of slices.
	*/

	stacks := processStacks(file)
	for _, s := range stacks {
		var result string

		crate, found := s.Pop()
		if !found {
			result = " "
		} else {
			result = fmt.Sprintf("%s", string(crate.(rune)))
		}
		fmt.Print(result)
	}

	fmt.Println()
}

type craneStrategy func(s *bufio.Scanner, crates []*data.Stack)

var strategy craneStrategy

func processStacks(file *os.File) []*data.Stack {
	s := bufio.NewScanner(file)
	var queues []*data.Queue

	for s.Scan() {
		if strings.HasPrefix(s.Text(), " 1") {
			s.Scan()
			break
		}

		if s.Err() == io.EOF || len(s.Text()) == 0 {
			break
		}

		line := []rune(s.Text())

		if queues == nil {
			queues = setupCrateQueues(len(line))
		}
		for p, v := range line {
			if v == '[' {
				val := line[p+1]
				queues[p/4].Push(val)
			}
		}
	}

	stacks := convertToStacks(queues)
	strategy(s, stacks)
	return stacks
}

func convertToStacks(queues []*data.Queue) []*data.Stack {
	stacks := make([]*data.Stack, len(queues))

	for p, q := range queues {
		items := q.Items()
		// reverse the items
		sort.SliceStable(items, func(i, j int) bool {
			return i > j
		})
		stacks[p] = data.NewStackFromItems(q.Items())
	}

	return stacks
}

func setupCrateQueues(lineLength int) []*data.Queue {
	numberOfStacks := (lineLength + 1) / 4
	result := make([]*data.Queue, numberOfStacks)

	for i := 0; i < numberOfStacks; i++ {
		result[i] = data.NewQueue()
	}

	return result
}

func processMoveInstructions(s *bufio.Scanner, crates []*data.Stack) {
	printCrates(crates)

	for s.Scan() {
		if s.Err() == io.EOF {
			break
		}

		count, fromIndex, toIndex := 0, 0, 0
		fmt.Sscanf(s.Text(), "move %d from %d to %d", &count, &fromIndex, &toIndex)

		for i := 0; i < count; i++ {
			crate, found := crates[fromIndex-1].Pop()
			if !found {
				log.Printf("Expected to find a crate in stack %d, but was empty", fromIndex-1)
			}
			crates[toIndex-1].Push(crate)
		}
	}

	printCrates(crates)
}

func processMoveInstructionsPart2(s *bufio.Scanner, crates []*data.Stack) {
	fmt.Println("Processing strategy part2")
	printCrates(crates)

	for s.Scan() {
		if s.Err() == io.EOF {
			break
		}

		count, fromIndex, toIndex := 0, 0, 0
		fmt.Sscanf(s.Text(), "move %d from %d to %d", &count, &fromIndex, &toIndex)

		subStack := data.NewStack()
		for i := 0; i < count; i++ {
			crate, found := crates[fromIndex-1].Pop()
			if !found {
				log.Printf("Expected to find a crate in stack %d, but was empty", fromIndex-1)
			}
			subStack.Push(crate)
		}
		for i := 0; i < count; i++ {
			crate, _ := subStack.Pop()
			crates[toIndex-1].Push(crate)
		}
	}

	printCrates(crates)
}

func printCrates(crates []*data.Stack) {
	longest := 0
	stacks := make([]*data.Stack, len(crates))

	for p, s := range crates {
		stacks[p] = s.Copy()
	}

	for _, s := range stacks {
		if s.Size() > longest {
			longest = s.Size()
		}
	}

	for i := longest; i > 0; i-- {
		builder := strings.Builder{}
		for _, q := range stacks {
			if q.Size() >= i {
				val, _ := q.Pop()
				builder.WriteRune('[')
				builder.WriteRune(val.(rune))
				builder.WriteString("] ")
			} else {
				builder.WriteString("    ")
			}
		}
		fmt.Println(builder.String())
	}

	fmt.Println()
}
