package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

var boards []*board

type WinOrLoseStrategy int

const (
	Win WinOrLoseStrategy = iota
	Lose
)

func main() {
	var (
		file     *os.File
		err      error
		strategy WinOrLoseStrategy
	)

	switch len(os.Args) {
	case 1:
		file = os.Stdin
	case 3:
		switch os.Args[2] {
		case "Win":
			fallthrough
		case "win":
			strategy = Win
		case "Lose":
			fallthrough
		case "lose":
			strategy = Lose
		default:
			log.Fatalf("Did not understand win/lose strategy of : %s\n", os.Args[2])
		}

		fallthrough
	case 2:
		file, err = os.Open(os.Args[1])
		if err != nil {
			log.Fatalf("Error opening file: %s", os.Args[1])
		}
	}

	sum, number := findBingoSum(file, strategy)
	result := sum * number
	fmt.Println(result)
}

func findBingoSum(r io.Reader, strategy WinOrLoseStrategy) (int, int) {
	var (
		lastSum, lastNumber int
	)

	s := bufio.NewScanner(r)
	s.Scan()
	for len(strings.TrimSpace(s.Text())) == 0 {
		s.Scan()
	}

	numbers := parseIntegers(s.Text())

	builder := new(strings.Builder)
	for s.Scan() {
		builder.Write([]byte(s.Text()))
		builder.Write([]byte("\n"))
	}
	boardsReader := strings.NewReader(builder.String())
	setupBoards(boardsReader, 5, 5)

	for _, number := range numbers {
		for boardIndex := len(boards) - 1; boardIndex >= 0; boardIndex-- {
			b := boards[boardIndex]
			if b.playNumber(number) {
				lastSum, lastNumber = b.sumUnmarked(), number
				boards = append(boards[:boardIndex], boards[boardIndex+1:]...)
				if strategy == Win {
					return lastSum, lastNumber
				}
			}
		}
	}

	return lastSum, lastNumber
}

func parseIntegers(numbers string) []int {
	var (
		result []int = make([]int, 0, 10)
	)

	numberStrings := strings.Split(strings.TrimSpace(numbers), ",")
	for _, str := range numberStrings {
		i, _ := strconv.ParseInt(str, 10, 32)
		result = append(result, int(i))
	}
	return result
}

func setupBoards(r io.Reader, rows int, columns int) {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords)

	boards = make([]*board, 0, 10)

	boardValues := make([]int, rows*columns)
	i := 0
	for s.Scan() {
		text := s.Text()
		number, _ := strconv.ParseInt(text, 10, 32)
		pos := i % (rows * columns)
		boardValues[pos] = int(number)
		if pos == (rows*columns)-1 {
			board := NewBoard(5, 5, boardValues)
			boards = append(boards, board)
			boardValues = make([]int, rows*columns)
		}
		i++
	}
}
