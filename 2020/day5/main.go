package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
)

type seat struct {
	id     int
	row    int
	column int
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Could not open input for reading")
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	seats := make([]seat, 0, 10)
	s := bufio.NewScanner(f)
	for s.Scan() {
		seat := newSeat(s.Text(), 128, 8)
		seats = append(seats, seat)
	}

	sort.Slice(seats, func(i, j int) bool {
		return seats[i].id < seats[j].id
	})

	lastID := 0
	missingSeatID := 0
	maxSeatID := 0
	for _, seat := range seats {
		//fmt.Printf("Seat: %+v\n", seat)
		if seat.id == lastID+2 {
			missingSeatID = seat.id - 1
		} else {
			lastID = seat.id
		}

		if seat.id > maxSeatID {
			maxSeatID = seat.id
		}
	}

	fmt.Printf("Highest seat id: %d\n", maxSeatID)
	fmt.Printf("My seat id: %d\n", missingSeatID)
}

func newSeat(code string, totalRowCount int, totalColCount int) seat {
	if match, err := regexp.MatchString("([FB]){7}([RL]){3}", code); !match || err != nil {
		err2 := fmt.Errorf("dode '%s' does not match expected format", code)
		fmt.Println(err2.Error())
		return seat{}
	}

	row := 0
	upperRow, lowerRow := totalRowCount-1, 0
	for _, r := range code[:7] {
		switch r {
		case 'F':
			upperRow = lowerRow + (upperRow-lowerRow)/2
			if upperRow-lowerRow == 0 {
				row = lowerRow
			}
		case 'B':
			lowerRow = lowerRow + (upperRow-lowerRow)/2 + 1
			if upperRow-lowerRow == 0 {
				row = upperRow
			}
		}
	}

	col := 0
	upperCol, lowerCol := totalColCount-1, 0
	for _, r := range code[7:] {
		switch r {
		case 'L':
			upperCol = lowerCol + (upperCol-lowerCol)/2
			if upperCol-lowerCol == 0 {
				col = lowerCol
			}
		case 'R':
			lowerCol = lowerCol + (upperCol-lowerCol)/2 + 1
			if upperCol-lowerCol == 0 {
				col = upperCol
			}
		}
	}

	return seat{row: row, column: col, id: row*totalColCount + col}
}
