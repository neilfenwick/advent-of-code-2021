package main

import (
	"reflect"
	"testing"
)

func Test_newSeat(t *testing.T) {
	type args struct {
		code          string
		totalRowCount int
		totalColCount int
	}
	tests := []struct {
		name string
		args args
		want seat
	}{
		{"example", args{code: "FBFBBFFRLR", totalRowCount: 128, totalColCount: 8}, seat{id: 357, row: 44, column: 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newSeat(tt.args.code, tt.args.totalRowCount, tt.args.totalColCount); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newSeat() = %v, want %v", got, tt.want)
			}
		})
	}
}
