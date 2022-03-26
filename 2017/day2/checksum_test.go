package main

import (
	"fmt"
	"testing"
)

func TestDiffChecksumReturnsZeroForEmptyInput(t *testing.T) {
	var data []int
	sut := NewChecksum()
	sut.Add(data)
	result := sut.RowSums[0]
	assertExpected(0, result, t)
}

func TestDiffChecksumReturnsExpectedRowValues(t *testing.T) {
	expected := []int{8, 4, 6}
	sut := givenChecksumWithDataRows(
		[][]int{{5, 1, 9, 5}, {7, 5, 3}, {2, 4, 6, 8}},
		DiffChecksum)

	for pos, result := range sut.RowSums {
		assertExpected(expected[pos], result, t)
	}
}

func givenChecksumWithDataRows(data [][]int, strategy ChecksumFunc) *Checksum {
	checksum := NewChecksum()
	checksum.Checksum(strategy)

	for _, row := range data {
		checksum.Add(row)
	}
	return checksum
}

func TestModulusChecksumReturnsExpectedRowValues(t *testing.T) {
	expected := []int{4, 3, 2}
	sut := givenChecksumWithDataRows(
		[][]int{{5, 9, 2, 8}, {9, 4, 7, 3}, {3, 8, 6, 5}},
		ModulusChecksum)

	for pos, result := range sut.RowSums {
		assertExpected(expected[pos], result, t)
	}
}

func TestDiffChecksumReturnsExpectedTotalValue(t *testing.T) {
	expected := 18
	sut := givenChecksumWithDataRows(
		[][]int{{5, 1, 9, 5}, {7, 5, 3}, {2, 4, 6, 8}},
		DiffChecksum)

	assertExpected(expected, sut.Value(), t)
}

func assertExpected(expected, result int, t *testing.T) {
	if result != expected {
		t.Error(fmt.Sprintf("Expected: %d, got %d", expected, result))
	}
}

func TestModulusChecksumReturnsExpectedTotalValue(t *testing.T) {
	expected := 9
	sut := givenChecksumWithDataRows(
		[][]int{{5, 9, 2, 8}, {9, 4, 7, 3}, {3, 8, 6, 5}},
		ModulusChecksum)

	assertExpected(expected, sut.Value(), t)
}

func TestChecksumPanicsIfRowsAlreadyAdded(t *testing.T) {
	defer assertPanicOccurred(t)

	sut := NewChecksum()
	sut.Add([]int{})
	sut.Checksum(ModulusChecksum)
}

func assertPanicOccurred(t *testing.T) {
	// if we recover, then a panic has occurred to recover from
	if r := recover(); r == nil {
		t.Error("The function did not panic")
	}
}

func TestModulusChecksumPanicsIfNoData(t *testing.T) {
	defer assertPanicOccurred(t)
	sut := NewChecksum()
	sut.Checksum(ModulusChecksum)
	sut.Add([]int{})
}

func TestModulusChecksumPanicsIfNoEvenlyDivisibleNumbers(t *testing.T) {
	defer assertPanicOccurred(t)
	sut := NewChecksum()
	sut.Checksum(ModulusChecksum)
	sut.Add([]int{7, 11, 3, 13, 19})
}
