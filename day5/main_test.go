package main

import (
	"io"
	"strings"
	"testing"
)

func Test_countVentDensity(t *testing.T) {
	type args struct {
		r         io.Reader
		threshold int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Example",
			args{
				strings.NewReader(`
					0,9 -> 5,9
					8,0 -> 0,8
					9,4 -> 3,4
					2,2 -> 2,1
					7,0 -> 7,4
					6,4 -> 2,0
					0,9 -> 2,9
					3,4 -> 1,4
					0,0 -> 8,8
					5,5 -> 8,2`),
				1,
			},
			12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countVentDensity(tt.args.r, tt.args.threshold); got != tt.want {
				t.Errorf("countVentDensity() = %v, want %v", got, tt.want)
			}
		})
	}
}
