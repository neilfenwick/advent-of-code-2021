package main

// List maintains the state for a slice of values
// that are used to calculate jump-list offsets
type List struct {
	values     []int
	offsetCalc JumpOffsetCalc
}

// JumpOffsetCalc returns the amount to adjust
// the jump by depending on the offset value
// in the parameter
type JumpOffsetCalc func(int) int

// NewList creates a new List, with a copy of the
// input slice
func NewList(data []int) List {
	values := make([]int, len(data))
	copy(values, data)
	return List{values: values, offsetCalc: DefaultOffsetCalc}
}

// OffsetCalcFunc sets the OffsetCalcFunc function to be used when calculating
// the jump offset, whilst traversing the list
func (l *List) OffsetCalcFunc(calcFunc JumpOffsetCalc) {
	l.offsetCalc = calcFunc
}

// CalcJumps returns the number of jumps required until an offset
// outside the bounds of the input slice are reached. The values
// of the input slice are used to calculate the jump offsets in
// an iterative sequence.
func (l *List) CalcJumps() int {
	offset := 0
	jumps := 0
	if len(l.values) == 0 {
		return 0
	}

	for {
		read := l.values[offset]
		jumpOffset := l.offsetCalc(read)
		l.values[offset] = read + jumpOffset
		offset = offset + read
		jumps++
		if offset >= len(l.values) {
			return jumps
		}
	}
}

// DefaultOffsetCalc strategy always returns 1
func DefaultOffsetCalc(int) int {
	return 1
}

// NewStrategyOffsetCalc returns -1 if the original
// offset parameter is 3 or more, otherwise returns 1
func NewStrategyOffsetCalc(offset int) int {
	if offset >= 3 {
		return -1
	}
	return 1
}
