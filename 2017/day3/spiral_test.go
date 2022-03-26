package main

import (
	"os"
	"testing"
)

var input int
var spiral SpiralGrid

func TestMain(m *testing.M) {
	input = 277678
	spiral = SpiralGrid{}
	tests := m.Run()
	os.Exit(tests)
}

func TestManhattanDistanceIdentity(t *testing.T) {
	result := spiral.ManhattanDistance(1)
	if result != 0 {
		t.Errorf("Value 1 returned %d for centre of grid. Expected 0.", result)
	}
}

func TestManhattanDistanceSmallIntegers(t *testing.T) {
	table := map[int]int{
		12:   3,
		23:   2,
		1024: 31,
	}

	for val, expected := range table {
		result := spiral.ManhattanDistance(val)
		if result != expected {
			t.Errorf("Unexpected result %d for value %d. Expected %d.", result, val, expected)
		}
	}
}

// 147  142  133  122   59
// 304    5    4    2   57
// 330   10    1    1   54
// 351   11   23   25   26
// 362  747  806  880  931
func TestSpiralSumVisitedNeighbours(t *testing.T) {
	table := map[int]int{
		1:  1,
		2:  1,
		3:  2,
		4:  4,
		5:  5,
		6:  10,
		7:  11,
		8:  23,
		9:  25,
		12: 57,
		13: 59,
		14: 122,
		17: 147,
		18: 304,
		25: 931,
	}

	for pos, expected := range table {
		result := spiral.CumulativeSumToPosition(pos)
		if result != expected {
			t.Errorf("Error for position %d. Expected %d. Got %d", pos, expected, result)
		}
	}
}

func BenchmarkManhattanDistance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		spiral.ManhattanDistance(input)
	}
}

func BenchmarkCumulativeSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var cumulativeSum int
		for i := 1; cumulativeSum <= input; i++ {
			cumulativeSum = spiral.CumulativeSumToPosition(i)
		}
	}
}
