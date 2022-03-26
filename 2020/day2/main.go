package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type passwordDefinition struct {
	minOccurrences int
	maxOccurrences int
	requiredChar   rune
	clearText      string
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Could not open input for reading: %v\n", err)
	}

	pwdDefinitions := make([]passwordDefinition, 0, 10)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var minOccurrences, maxOccurrences int
		var requiredChar rune
		var clearText string

		_, _ = fmt.Sscanf(scanner.Text(), "%d-%d %c: %s", &minOccurrences, &maxOccurrences, &requiredChar, &clearText)
		pwdDefinitions = append(pwdDefinitions, passwordDefinition{minOccurrences: minOccurrences, maxOccurrences: maxOccurrences, requiredChar: requiredChar, clearText: clearText})
	}

	validPasswords := 0
	for _, def := range pwdDefinitions {
		if testPasswordMinOccurrences(def) {
			validPasswords++
		}
	}

	fmt.Printf("Valid password min-occurence count: %d\n", validPasswords)

	validPasswords = 0
	for _, def := range pwdDefinitions {
		if testPasswordSpecificPosition(def) {
			validPasswords++
		}
	}

	fmt.Printf("Valid password specific-position count: %d\n", validPasswords)
}

func testPasswordMinOccurrences(def passwordDefinition) bool {
	cnt := 0
	pwd := []rune(def.clearText)
	for i := 0; i < len(pwd); i++ {
		if def.requiredChar == pwd[i] {
			cnt++
		}
	}

	if cnt >= def.minOccurrences && cnt <= def.maxOccurrences {
		return true
	}

	return false
}

func testPasswordSpecificPosition(def passwordDefinition) bool {
	pwd := []rune(def.clearText)
	if (def.requiredChar == pwd[def.minOccurrences-1]) != (def.requiredChar == pwd[def.maxOccurrences-1]) {
		return true
	}

	return false
}
