package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("Error opening file: %s", os.Args[1])
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	mapData := parseMap(file)
	wordCount := countWords(mapData, []rune{'X', 'M', 'A', 'S'})
	log.Printf("Word count: %d", wordCount)

	xmasCount := countX(mapData, []rune{'M', 'A', 'S'})
	log.Printf("XMAS count: %d", xmasCount)
}

type point struct {
	x int
	y int
}

func parseMap(file io.Reader) map[point]rune {

	mapData := make(map[point]rune)

	scanner := bufio.NewScanner(file)
	rowIndex := 0
	for scanner.Scan() {
		line := scanner.Text()
		for columnIndex, char := range line {
			mapData[point{columnIndex, rowIndex}] = char
		}
		rowIndex++
	}

	return mapData
}

type vector point

var up vector = vector{0, -1}
var down vector = vector{0, 1}
var left vector = vector{-1, 0}
var right vector = vector{1, 0}
var upLeft vector = vector{-1, -1}
var upRight vector = vector{1, -1}
var downLeft vector = vector{-1, 1}
var downRight vector = vector{1, 1}

var directions = []vector{up, down, left, right, upLeft, upRight, downLeft, downRight}

func countWords(mapData map[point]rune, searchWord []rune) int {
	wordCount := 0

	for point, char := range mapData {
		if char == searchWord[0] {
			for _, direction := range directions {
				if checkWord(mapData, point, searchWord, direction) {
					wordCount++
				}
			}
		}
	}

	return wordCount
}

func checkWord(mapData map[point]rune, start point, searchWord []rune, direction vector) bool {
	for i, char := range searchWord {
		currentPoint := point{start.x + direction.x*i, start.y + direction.y*i}
		if mapData[currentPoint] != char {
			return false
		}
	}
	return true
}

func countX(mapData map[point]rune, searchWord []rune) int {
	xCount := 0
	var middle int = len(searchWord) / 2 //  "Rounds" down beause division truncates to int
	if len(searchWord) == 0 {
		log.Println("Search word must have an odd number of characters to search for X's")
	}

	// Find all the upleft and downright diagonals that contain the search word
	// Store the location of the middle rune
	backdiagonals := make(map[point]rune)
	for loc, char := range mapData {
		if char == searchWord[0] {
			for _, direction := range []vector{upLeft, downRight} {
				if checkWord(mapData, loc, searchWord, direction) {
					backdiagonals[point{loc.x + direction.x*middle, loc.y + direction.y*middle}] = char
				}
			}
		}
	}

	// Find all of the upRight and downLeft diagonals that contain the search word
	// Store the location of the middle rune
	diagonals := make(map[point]rune)
	for loc, char := range mapData {
		if char == searchWord[0] {
			for _, direction := range []vector{upRight, downLeft} {
				if checkWord(mapData, loc, searchWord, direction) {
					diagonals[point{loc.x + direction.x*middle, loc.y + direction.y*middle}] = char
				}
			}
		}
	}

	// Cross reference the two sets to find where the middle rune in the search word is in both sets
	for loc := range backdiagonals {
		if _, ok := diagonals[loc]; ok {
			xCount++
		}
	}

	return xCount
}
