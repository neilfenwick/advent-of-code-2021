package stack

import (
	"reflect"
	"testing"
)

func TestPush(t *testing.T) {
	type args struct {
		items []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{"One item", args{[]interface{}{10}}},
		{"Three items in order", args{[]interface{}{10, "abc", 30}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sut := NewStack()
			for i, item := range tt.args.items {
				sut.Push(item)
				if sut.stack[i] != item {
					t.Errorf("Expected item '%v' to be in position %d", item, i)
				}
			}

		})
	}
}

func TestStack_Pop(t *testing.T) {
	type fields struct {
		stack []interface{}
	}
	tests := []struct {
		name      string
		fields    fields
		wantItem  interface{}
		wantFound bool
	}{
		{"Pop only item",
			fields{stack: []interface{}{"abc"}},
			"abc",
			true,
		},
		{"Pop second item",
			fields{stack: []interface{}{"abc", 20}},
			20,
			true,
		},
		{"Already empty",
			fields{stack: []interface{}{}},
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				stack: tt.fields.stack,
			}
			if got, success := s.Pop(); !reflect.DeepEqual(got, tt.wantItem) || tt.wantFound != success {
				t.Errorf("Stack.Pop() = (%v, %v), want (%v, %v)", got, success, tt.wantItem, tt.wantFound)
			}
		})
	}
}
