package main

import (
	"math"
	"sort"
)

// Checksum represents a series of int-slice checksums that
// are calculated
type Checksum struct {
	RowSums  []int
	checksum ChecksumFunc
}

// ChecksumFunc is the signature of the function used to calculate
// the checksum of each row
type ChecksumFunc func(data []int) int

// NewChecksum initialises a new instance
func NewChecksum() *Checksum {
	return &Checksum{
		RowSums:  make([]int, 0),
		checksum: DiffChecksum,
	}
}

// Add calculates and adds the checksum for the given slice
func (c *Checksum) Add(row []int) {
	rowSum := c.checksum(row)
	c.RowSums = append(c.RowSums, rowSum)
}

// DiffChecksum calculates a checksum by taking the difference between the highest
// and lowest values in the input
//
// DiffChecksum returns zero if the parameter is nil or empty
func DiffChecksum(data []int) int {
	max := 0
	min := math.MaxInt64

	if data == nil || len(data) == 0 {
		return 0
	}

	for _, val := range data {
		if val < min {
			min = val
		}

		if val > max {
			max = val
		}
	}
	return max - min
}

// ModulusChecksum calculates a checksum by finding the first occurrence
// of two numbers that divide evenly and returning the result of dividing
// the greater by the lessor of the two.
//
// ModulusChecksum panics if the parameter is nil or empty, or if there
// are no two numbers that divide without remainder
func ModulusChecksum(data []int) int {
	if data == nil || len(data) == 0 {
		panic("Cannot perform checksum with no data")
	}

	// Make a copy of the data so that we don't create side effects
	// for the caller
	sortedIntegers := make([]int, len(data))
	copy(sortedIntegers, data)
	sort.Ints(sortedIntegers)

	for pos, val := range sortedIntegers {
		for i := len(sortedIntegers) - 1; i > pos; i-- {

			// Note: sortedIntegers are in ascending order, need to divide
			// the later value by the earlier one
			if sortedIntegers[i]%val == 0 {
				result := int(float32(sortedIntegers[i]) / float32(val))
				return result
			}
		}
	}

	panic("Did not find any numbers that divide without remainder")
}

// Checksum sets the check-summing function for the Checksum.
// The default checksum function is DiffChecksum.
//
// Checksum panics if it is called after rows have already been added.
func (c *Checksum) Checksum(checksum ChecksumFunc) {
	if len(c.RowSums) > 0 {
		panic("Checksum called after rows already added")
	}
	c.checksum = checksum
}

// Value returns the total checksum value for all rows
func (c *Checksum) Value() int {
	var sum int
	for _, val := range c.RowSums {
		sum += val
	}
	return sum
}
