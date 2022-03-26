package main

import "testing"

func Test_autoCompleteScore(t *testing.T) {
	type args struct {
		chars []rune
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example from day 10 part 2", args{[]rune("])}>")}, 294},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := autoCompleteScore(tt.args.chars); got != tt.want {
				t.Errorf("autoCompleteScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
