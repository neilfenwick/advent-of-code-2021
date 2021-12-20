package main

import (
	"io"
	"strings"
	"testing"
)

func Test_calcPosition(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{"test", args{strings.NewReader(
			`forward 5
			down 5
			forward 8
			up 3
			down 8
			forward 2`)}, 10, 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := calcPosition(tt.args.r)
			if got != tt.want {
				t.Errorf("calcPosition() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("calcPosition() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
