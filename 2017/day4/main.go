package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic("Could not read input file")
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	wordsValidator := NewPassValidator()
	anagramsValidator := NewPassValidator()
	anagramsValidator.EntropyFunc(IsAnagramsInPhrase)

	numValidUniques, numValidAnagrams := 0, 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		passPhrase := scanner.Text()
		if wordsValidator.IsValid(passPhrase) {
			numValidUniques++
		}

		if anagramsValidator.IsValid(passPhrase) {
			numValidAnagrams++
		}
	}

	fmt.Printf("Number of valid phrases with unique words: %d\n", numValidUniques)
	fmt.Printf("Number of valid phrases with unique anagrams: %d\n", numValidAnagrams)
}
