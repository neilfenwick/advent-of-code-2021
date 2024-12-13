package data

import (
	"reflect"
	"testing"
)

func TestIntBuffer_Write(t *testing.T) {
	type fields struct {
		buffer         []interface{}
		WriteCursorPos int
	}
	type args struct {
		val int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"Overwrite value increments write cursor by one",
			fields{[]interface{}{1, 2, 3}, 1},
			args{123},
		},
		{
			"Write wraps back to beginning when write to end",
			fields{[]interface{}{1, 2, 3}, 2},
			args{123},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &CircularBuffer{
				buffer:         tt.fields.buffer,
				WriteCursorPos: tt.fields.WriteCursorPos,
			}
			before := b.WriteCursorPos
			b.Write(tt.args.val)
			if b.WriteCursorPos != (before+1)%len(tt.fields.buffer) {
				t.Error("Write cursor pos was not incremented by one")
			}
		})
	}
}

func TestIntBuffer_Read(t *testing.T) {
	type fields struct {
		buffer         []interface{}
		WriteCursorPos int
	}
	type args struct {
		offset int
		count  int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []interface{}
	}{
		{
			"Read whole buffer with cursor at start",
			fields{[]interface{}{1, 2, 3, 4, 5}, 0},
			args{0, 5},
			[]interface{}{1, 2, 3, 4, 5},
		},
		{
			"Read first two values with cursor at start",
			fields{[]interface{}{1, 2, 3, 4, 5}, 0},
			args{0, 2},
			[]interface{}{1, 2},
		},
		{
			"Read last and first value with cursor at end",
			fields{[]interface{}{1, 2, 3, 4, 5}, 4},
			args{0, 2},
			[]interface{}{5, 1},
		},
		{
			"Read last value with cursor at beginning and negative offset",
			fields{[]interface{}{1, 2, 3, 4, 5}, 0},
			args{-1, 1},
			[]interface{}{5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &CircularBuffer{
				buffer:         tt.fields.buffer,
				WriteCursorPos: tt.fields.WriteCursorPos,
			}
			if got := b.Read(tt.args.offset, tt.args.count); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntBuffer.Read() = %v, want %v", got, tt.want)
			}
		})
	}
}
