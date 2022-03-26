package main

import (
	"testing"
)

func TestRebalance(t *testing.T) {
	type testData struct {
		iterCount int
		memory    []int
	}

	table := []testData{
		{0, []int{}},
		{0, []int{1}},
		{0, []int{4}},
		{2, []int{1, 1}},
		{5, []int{0, 2, 7, 0}},
	}

	for _, data := range table {
		actual := Rebalance(data.memory)
		if actual != data.iterCount {
			t.Errorf("Unexpected result: %d. Expected %d. Final memory state: %v.", actual, data.iterCount, data.memory)
		}
	}
}

func BenchmarkRebalance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := []int{0, 5, 10, 0, 11, 14, 13, 4, 11, 8, 8, 7, 1, 4, 12, 11}
		Rebalance(input)
	}
}
