package main

import (
	"sort"
	"strings"
)

// PassValidator validates passphrases for sufficient entropy
type PassValidator struct {
	entropyFunc EntropyTest
}

// EntropyTest is the algorithm used to determine
// sufficient uniqueness in the pass-phrase
type EntropyTest func(phrase string) bool

// NewPassValidator constructs a PassValidator with default
// entropy function to test IsUniqueWords
func NewPassValidator() PassValidator {
	val := PassValidator{
		entropyFunc: IsUniqueWordsInPhrase,
	}
	return val
}

// EntropyFunc sets the function to use that determines
// the uniqueness requirement for the passphrase
func (p *PassValidator) EntropyFunc(test EntropyTest) {
	p.entropyFunc = test
}

// IsValid calls the entropyFunc and returns the result
func (p *PassValidator) IsValid(phrase string) bool {
	return p.entropyFunc(phrase)
}

// IsUniqueWordsInPhrase indicates that each token
// occurs only once throughout the phrase
func IsUniqueWordsInPhrase(phrase string) bool {
	words := strings.Fields(phrase)
	return isUniqueWords(words)
}

func isUniqueWords(words []string) bool {
	uniqMap := make(map[string]bool, len(words))
	var dup bool
	for _, word := range words {
		_, exist := uniqMap[word]
		if !exist {
			uniqMap[word] = true
		} else {
			dup = true
		}
	}

	return !dup
}

// IsAnagramsInPhrase indicates that each token
// occurs only once throughout the phrase, including anagrams
// of that word
func IsAnagramsInPhrase(phrase string) bool {
	words := strings.Fields(phrase)
	for pos, word := range words {
		letters := strings.Split(word, "")
		sort.Strings(letters)
		words[pos] = strings.Join(letters, "")
	}

	return isUniqueWords(words)
}
