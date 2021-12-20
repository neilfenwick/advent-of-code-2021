package main

import (
	"io"
	"strings"
	"testing"
)

func Test_depthIncreasesCount(t *testing.T) {
	type args struct {
		r          io.Reader
		windowSize int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Window size of 1",
			args{
				strings.NewReader(`
					199
					200
					208
					210
					200
					207
					240
					269
					260
					263`),
				1,
			},
			7,
		},
		{"Window size of 5",
			args{
				strings.NewReader(`
					607
					618
					618
					617
					647
					716
					769
					792`),
				3,
			},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := depthIncreasesCount(tt.args.r, tt.args.windowSize); got != tt.want {
				t.Errorf("depthIncreasesCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
