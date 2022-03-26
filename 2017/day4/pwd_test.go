package main

import (
	"testing"
)

func TestIsUniqueWordsUniquePassphrase(t *testing.T) {
	input := "Unique words\n"
	ok := IsUniqueWordsInPhrase(input)
	if !ok {
		t.Errorf("Expected unique words")
	}
}

func TestIsUniqueWordsDuplicateWord(t *testing.T) {
	input := "The words words appear twice\n"
	ok := IsUniqueWordsInPhrase(input)
	if ok {
		t.Errorf("Expected failure for duplicate word, but was ok")
	}
}

func TestDefaultValidFuncIsUniqueWords(t *testing.T) {
	input := "aabb bbaa"
	passValidator := NewPassValidator()
	ok := passValidator.IsValid(input)
	if !ok {
		t.Errorf("Expect unique words and anagrams to pass. Got failure.")
	}
}

func BenchmarkIsUniqueWords(b *testing.B) {
	input := "rdngdy jde wvgkhto bdvngf mdup eskuvg ezli opibo mppoc mdup zrasc\n"
	for i := 0; i < b.N; i++ {
		IsUniqueWordsInPhrase(input)
	}
}

func TestIsUniqueAnagramUniquePassphrase(t *testing.T) {
	input := "aaaaaaa bbbbbb\n"
	passValidator := NewPassValidator()
	passValidator.EntropyFunc(IsAnagramsInPhrase)
	ok := passValidator.IsValid(input)
	if !ok {
		t.Errorf("Expected unique words")
	}
}

func TestIsUniqueAnagram(t *testing.T) {
	input := "aabb bbaa"
	passValidator := NewPassValidator()
	passValidator.EntropyFunc(IsAnagramsInPhrase)
	ok := passValidator.IsValid(input)
	if ok {
		t.Errorf("Anagrams to fail validation. Input: '%v'", input)
	}
}

func BenchmarkIsUniqueAnagram(b *testing.B) {
	input := "rdngdy jde wvgkhto bdvngf mdup eskuvg ezli opibo mppoc mdup zrasc\n"
	for i := 0; i < b.N; i++ {
		IsAnagramsInPhrase(input)
	}
}
