package main

import (
	"io"
	"strings"
	"testing"
)

func Test_setupBoards(t *testing.T) {
	type args struct {
		r       io.Reader
		rows    int
		columns int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"Read example boards",
			args{
				strings.NewReader(`
					22 13 17 11  0
					8  2 23  4 24
					21  9 14 16  7
					6 10  3 18  5
					1 12 20 15 19

					3 15  0  2 22
					9 18 13 17  5
					19  8  7 25 23
					20 11 10 24  4
					14 21 16 12  6

					14 21 17 24  4
					10 16 15  9 19
					18  8 23 26 20
					22 11 13  6  5
					2  0 12  3  7`),
				5,
				5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setupBoards(tt.args.r, tt.args.rows, tt.args.columns)
			b := boards
			if len(b) != 3 {
				t.Errorf("Expected 3 boards, got %d", len(b))
			}
		})
	}
}

func Test_findBingo(t *testing.T) {
	type args struct {
		r        io.Reader
		strategy WinOrLoseStrategy
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			"Find bingo for one board",
			args{
				strings.NewReader(`
					7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

					14 21 17 24 4
					10 16 15 9 19
					18 8 23 26 20
					22 11 13 6 5
					2 0 12 3 7
					`),
				Win,
			},
			188,
			24,
		},
		{
			"Find bingo for example",
			args{
				strings.NewReader(`
					7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

					22 13 17 11  0
					8  2 23  4 24
					21  9 14 16  7
					6 10  3 18  5
					1 12 20 15 19

					3 15  0  2 22
					9 18 13 17  5
					19  8  7 25 23
					20 11 10 24  4
					14 21 16 12  6

					14 21 17 24  4
					10 16 15  9 19
					18  8 23 26 20
					22 11 13  6  5
					2  0 12  3  7	
				`),
				Win,
			},
			188,
			24,
		},
		{
			"Find losing bingo board for example",
			args{
				strings.NewReader(`
						7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

						22 13 17 11  0
						8  2 23  4 24
						21  9 14 16  7
						6 10  3 18  5
						1 12 20 15 19

						3 15  0  2 22
						9 18 13 17  5
						19  8  7 25 23
						20 11 10 24  4
						14 21 16 12  6

						14 21 17 24  4
						10 16 15  9 19
						18  8 23 26 20
						22 11 13  6  5
						2  0 12  3  7	
					`),
				Lose,
			},
			148,
			13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sum, lastNumber := findBingoSum(tt.args.r, tt.args.strategy)
			if sum != tt.want {
				t.Errorf("findBingoSum() sum = %v, want %v", sum, tt.want)
			}
			if lastNumber != tt.want1 {
				t.Errorf("findBingoSum() lastNumber = %v, want %v", lastNumber, tt.want1)
			}
		})
	}
}
