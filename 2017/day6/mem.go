package main

import (
	"crypto/md5"
	"encoding/binary"
	"math"
)

// Rebalance iterates over the input and distributes
// the values evenly amongst each of the slots in the
// input slice, returning the number of iterations required
// to rebalance
func Rebalance(memory []int) int {
	if len(memory) < 2 {
		return 0
	}

	iterations := 0
	hashes := make(map[int]bool)

	for {
		hsh := hash(memory)
		if _, exists := hashes[hsh]; exists {
			break
		}

		hashes[hsh] = false
		iterations++
		pos, max := maxVal(memory)
		memory[pos] = 0
		for i := max; i > 0; i-- {
			pos = pos + 1
			if pos >= len(memory) {
				pos = 0
			}

			val := memory[pos]
			memory[pos] = val + 1
		}
	}
	return iterations
}

func maxVal(values []int) (int, int) {
	var pos, val int
	for currPos, currVal := range values {
		if val < currVal {
			pos = currPos
			val = currVal
		}
	}

	return pos, val
}

func hash(values []int) int {
	hash := md5.New()
	for pos, val := range values {
		b := make([]byte, 8)
		binary.LittleEndian.PutUint64(b, uint64(math.Pow(float64(pos), 2)))
		binary.LittleEndian.PutUint64(b, uint64(val))
		hash.Write(b)
	}

	result := hash.Sum(nil)
	return int(binary.LittleEndian.Uint64(result))
}
