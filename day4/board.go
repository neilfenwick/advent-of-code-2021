package main

type location struct {
	row int
	col int
}

type board struct {
	numbers map[int]location
	marked  map[int]interface{}
	squares [][]bool
}

func NewBoard(rows int, columns int, values []int) *board {
	b := board{numbers: make(map[int]location, len(values)), marked: make(map[int]interface{}, rows)}
	b.squares = make([][]bool, rows)
	for i := range b.squares {
		b.squares[i] = make([]bool, columns)
	}

	for i, v := range values {
		b.numbers[v] = location{row: i / rows, col: i % columns}
	}
	return &b
}

func (b *board) playNumber(number int) bool {
	var (
		rowCheck bool = true
		colCheck bool = true
	)

	l, ok := b.numbers[number]
	if !ok {
		return false
	}

	b.marked[number] = true

	b.squares[l.row][l.col] = true
	for _, colValue := range b.squares[l.row] {
		rowCheck = rowCheck && colValue
	}

	if rowCheck {
		return true
	}

	for _, row := range b.squares {
		colCheck = colCheck && row[l.col]
	}

	return colCheck
}

func (b *board) sumUnmarked() int {
	var (
		sum int
	)

	for k := range b.numbers {
		if _, marked := b.marked[k]; !marked {
			sum += k
		}
	}

	return sum
}
