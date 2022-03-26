package main

import "testing"

func Test_countUnanimousAnswersForGroup(t *testing.T) {
	type args struct {
		group string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example all", args{group: "abc"}, 3},
		{"example none", args{group: `a
b
c`}, 0},
		{"example some", args{group: `ab
ac
`}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countUnanimousAnswersForGroup(tt.args.group); got != tt.want {
				t.Errorf("countUnanimousAnswersForGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
