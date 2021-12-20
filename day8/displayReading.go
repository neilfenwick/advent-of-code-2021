package main

type displayReading struct {
	numberMap map[string]int
	notes     []string
}

func NewDisplayReading(signalPatterns []string, notes []string) *displayReading {
	var (
		input   []IntersectableString = make([]IntersectableString, len(signalPatterns))
		numbers []IntersectableString = make([]IntersectableString, len(signalPatterns))
	)

	numberMap := make(map[string]int, len(numbers))

	for i, v := range signalPatterns {
		input[i] = IntersectableString(stringSort(v))
	}

	search := searchableStringSlice(input)

	numbers[1] = search.First(func(s IntersectableString) bool { return len(s) == 2 })
	numbers[7] = search.First(func(s IntersectableString) bool { return len(s) == 3 })
	numbers[4] = search.First(func(s IntersectableString) bool { return len(s) == 4 })
	numbers[8] = search.First(func(s IntersectableString) bool { return len(s) == 7 })
	numbers[9] = search.First(func(s IntersectableString) bool { return len(s) == 6 && len(s.Intersect(numbers[4])) == 4 })
	numbers[0] = search.First(func(s IntersectableString) bool {
		return len(s) == 6 && s != numbers[9] && len(s.Intersect(numbers[1])) == 2
	})
	numbers[6] = search.First(func(s IntersectableString) bool {
		return len(s) == 6 && s != numbers[9] && s != numbers[0]
	})
	numbers[3] = search.First(func(s IntersectableString) bool { return len(s) == 5 && len(s.Intersect(numbers[1])) == 2 })
	numbers[5] = search.First(func(s IntersectableString) bool {
		return len(s) == 5 && s != numbers[3] && len(s.Intersect(numbers[9])) == 5
	})
	numbers[2] = search.First(func(s IntersectableString) bool { return len(s) == 5 && s != numbers[3] && s != numbers[5] })

	for i := 0; i < len(notes); i++ {
		notes[i] = stringSort(notes[i])
	}

	for i, v := range numbers {
		numberMap[string(v)] = i
	}

	return &displayReading{numberMap: numberMap, notes: notes}
}
