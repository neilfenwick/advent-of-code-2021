package main

// SumConsecutiveIntegers returns the sum of consecutive integers in a circular array
func SumConsecutiveIntegers(data []int) int {
	_ = sumConsecutiveIntegersFirstImplementation(data) // Used and ignored to avoid IDE "unused" function warning
	return sumConsecutiveIntegersSecondImplementation(data)
}

func sumConsecutiveIntegersFirstImplementation(data []int) int {
	var first, previous, sum int

	lastPos := len(data) - 1

	for pos, num := range data {
		if pos == 0 {
			first = num
			previous = num
			continue
		}

		if previous == num {
			sum += previous
		}

		if pos == lastPos && first == num {
			sum += first
		}

		previous = num
	}

	return sum
}

func sumConsecutiveIntegersSecondImplementation(data []int) int {
	var sum int

	for pos, num := range data {
		if num == numberAtNextPosition(data, pos) {
			sum += num
		}
	}

	return sum
}

func numberAtNextPosition(data []int, pos int) int {
	if pos < len(data)-1 {
		return data[pos+1]
	}
	return data[0]
}

// SumOppositeIntegers returns the sum of integers that are opposite each other,
// half-way around a circular array
func SumOppositeIntegers(data []int) int {
	var sum int
	for pos, num := range data {
		if num == numberAtCircularOppositePosition(data, pos+len(data)/2) {
			sum += num
		}
	}
	return sum
}

func numberAtCircularOppositePosition(data []int, pos int) int {
	modPos := pos % len(data)
	return data[modPos]
}
