package main

import "sort"

type sortableRuneSlice []rune

func stringSort(s string) string {
	runes := []rune(s)
	sort.Sort(sortableRuneSlice(runes))
	return string(runes)
}

func (s sortableRuneSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortableRuneSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortableRuneSlice) Len() int {
	return len(s)
}
