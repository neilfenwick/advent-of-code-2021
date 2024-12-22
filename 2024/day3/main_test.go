package main

import (
	"bytes"
	"testing"
)

func TestScanCustom(t *testing.T) {
	tests := []struct {
		name        string
		data        []byte
		atEOF       bool
		wantAdvance int
		wantToken   []byte
		wantErr     error
	}{
		{
			name:        "Stop token present, no start token",
			data:        []byte("some data don't() other data"),
			atEOF:       false,
			wantAdvance: len([]byte("some data ")),
			wantToken:   []byte("some data "),
			wantErr:     nil,
		},
		{
			name:        "Stop token present and start token after stop token",
			data:        []byte("stop don't() do() start"),
			atEOF:       false,
			wantAdvance: len("stop don't() "),
			wantToken:   []byte("stop "),
			wantErr:     nil,
		},
		{
			name:        "Partial start token at end",
			data:        []byte("some data do("),
			atEOF:       false,
			wantAdvance: len("some data "),
			wantToken:   []byte("some data "),
			wantErr:     nil,
		},
		{
			name:        "Multiple tokens",
			data:        []byte("do() valid do() token don't() additional do() data"),
			atEOF:       false,
			wantAdvance: len("do() valid do() token "),
			wantToken:   []byte("do() valid do() token "),
			wantErr:     nil,
		},
		{
			name:        "At EOF with partial token",
			data:        []byte("incomplete do"),
			atEOF:       true,
			wantAdvance: len("incomplete "),
			wantToken:   []byte("incomplete "),
			wantErr:     nil,
		},
		{
			name:        "No tokens present",
			data:        []byte("just some random data without tokens"),
			atEOF:       false,
			wantAdvance: len("just some random data without tokens"),
			wantToken:   []byte("just some random data without tokens"),
			wantErr:     nil,
		},
		{
			name:        "Only start token present",
			data:        []byte("do() start without stop"),
			atEOF:       false,
			wantAdvance: len("do() start without stop"),
			wantToken:   []byte("do() start without stop"),
			wantErr:     nil,
		},
		{
			name:        "Stop token before start token",
			data:        []byte("don't() some data do() more data"),
			atEOF:       false,
			wantAdvance: len("don't() some data "),
			wantToken:   []byte(""),
			wantErr:     nil,
		},
		{
			name:        "Example from AoC",
			data:        []byte("xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"),
			atEOF:       false,
			wantAdvance: len("xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)un"),
			wantToken:   []byte("xmul(2,4)&mul[3,7]!^"),
			wantErr:     nil,
		},
		{
			name:        "When buffer ends with partial mul string",
			data:        []byte("_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5"),
			atEOF:       false,
			wantAdvance: len("_mul(5,5)+mul(32,64](mul(11,8)undo()?"),
			wantToken:   []byte("_mul(5,5)+mul(32,64](mul(11,8)undo()?"),
			wantErr:     nil,
		},
		{
			name:        "When buffer contains partial mul string",
			data:        []byte("_mul(5,5)+mul(32,64](mul(11,8)don't()undo()?"),
			atEOF:       false,
			wantAdvance: len("_mul(5,5)+mul(32,64](mul(11,8)don't()un"),
			wantToken:   []byte("_mul(5,5)+mul(32,64](mul(11,8)"),
			wantErr:     nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotAdvance, gotToken, gotErr := customStopStartTokenScanner(tt.data, tt.atEOF)
			if gotAdvance != tt.wantAdvance {
				t.Errorf("scanCustom() gotAdvance = %v, want %v", gotAdvance, tt.wantAdvance)
			}
			if !bytes.Equal(gotToken, tt.wantToken) {
				t.Errorf("scanCustom() gotToken = %v, want %v", string(gotToken), string(tt.wantToken))
			}
			if gotErr != tt.wantErr {
				t.Errorf("scanCustom() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}
