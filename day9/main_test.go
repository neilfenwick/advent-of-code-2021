package main

import (
	"strings"
	"testing"
)

func Test_findProductOfBasinSizes(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"Example basin calculation", 1134},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := strings.NewReader(`
				2199943210
				3987894921
				9856789892
				8767896789
				9899965678`)

			buildHeightMap(input)

			if got := findProductOfBasinSizes(); got != tt.want {
				t.Errorf("findProductOfBasinSizes() = %v, want %v", got, tt.want)
			}
		})
	}
}
