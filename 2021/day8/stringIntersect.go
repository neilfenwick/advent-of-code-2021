package main

type intersectingString string

func (s intersectingString) Intersect(other intersectingString) []rune {
	var (
		result = make([]rune, 0, len(other))
		lookup = make(map[rune]bool, len(s))
	)

	for _, c := range s {
		lookup[c] = true
	}

	for _, c := range other {
		if _, found := lookup[c]; found {
			result = append(result, c)
		}
	}
	return result
}

type searchableStringSlice []intersectingString

func (s searchableStringSlice) First(f func(s intersectingString) bool) intersectingString {
	for _, v := range s {
		if f(v) {
			return v
		}
	}
	return ""
}
