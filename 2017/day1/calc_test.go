package main

import (
	"fmt"
	"testing"
)

func TestSumConsecutivePairs(t *testing.T) {
	data := []int{1, 1, 2, 2}
	result := SumConsecutiveIntegers(data)
	assertExpected(3, result, t)
}

func assertExpected(expected, result int, t *testing.T) {
	if result != expected {
		t.Error(fmt.Sprintf("Expected: %d, got %d", expected, result))
	}
}

func TestSumConsecutiveUniformSlice(t *testing.T) {
	data := []int{1, 1, 1, 1}
	result := SumConsecutiveIntegers(data)
	assertExpected(4, result, t)
}

func TestSumConsecutiveZeroForAllDifferent(t *testing.T) {
	data := []int{1, 2, 3, 4}
	result := SumConsecutiveIntegers(data)
	assertExpected(0, result, t)
}

func TestSumConsecutiveCircularBufferFirstLastSame(t *testing.T) {
	data := []int{9, 1, 2, 1, 2, 1, 2, 9}
	result := SumConsecutiveIntegers(data)
	assertExpected(9, result, t)
}

func TestSumConsecutiveEmptySliceReturnsZero(t *testing.T) {
	var data []int
	result := SumConsecutiveIntegers(data)
	assertExpected(0, result, t)
}

func TestSumOppositePairs(t *testing.T) {
	data := []int{1, 2, 1, 2}
	result := SumOppositeIntegers(data)
	assertExpected(6, result, t)
}

func TestSumOppositeNoDirectOppositePairs(t *testing.T) {
	data := []int{1, 2, 2, 1}
	result := SumOppositeIntegers(data)
	assertExpected(0, result, t)
}

func TestSumOppositeSinglePair(t *testing.T) {
	data := []int{1, 2, 3, 4, 2, 5}
	result := SumOppositeIntegers(data)
	assertExpected(4, result, t)
}

func TestSumOppositeSymmetricalPairs(t *testing.T) {
	data := []int{1, 2, 3, 1, 2, 3}
	result := SumOppositeIntegers(data)
	assertExpected(12, result, t)
}

func TestSumOppositeOnesOnly(t *testing.T) {
	data := []int{1, 2, 1, 3, 1, 4, 1, 5}
	result := SumOppositeIntegers(data)
	assertExpected(4, result, t)
}

func TestSumOppositeEmptySliceReturnsZero(t *testing.T) {
	var data []int
	result := SumOppositeIntegers(data)
	assertExpected(0, result, t)
}
