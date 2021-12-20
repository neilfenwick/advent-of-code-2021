package main

type IntersectableString string

func (s IntersectableString) Intersect(other IntersectableString) []rune {
	var (
		result []rune        = make([]rune, 0, len(other))
		lookup map[rune]bool = make(map[rune]bool, len(s))
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

type searchableStringSlice []IntersectableString

func (s searchableStringSlice) First(f func(s IntersectableString) bool) IntersectableString {
	for _, v := range s {
		if f(v) {
			return v
		}
	}
	return IntersectableString("")
}
