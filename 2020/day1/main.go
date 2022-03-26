package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type pair struct {
	item1 int
	item2 int
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Could not read input data")
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	data := make([]int, 0, 10)

	s := bufio.NewScanner(f)

	for s.Scan() {
		var n int
		_, err := fmt.Sscanf(s.Text(), "%d", &n)
		if err != nil {
			log.Fatalf("Could not read %s: %v", s.Text(), err)
		}
		data = append(data, n)
	}

	pair := findFirstPairTargetSum(data, 2020)

	fmt.Printf("Pair: %v\n", pair)
	fmt.Printf("Pair Result: %d\n", pair.item1*pair.item2)

	triple := findFirstNItemsTargetSum(data, 3, 2020)

	agg := 1
	for _, v := range triple {
		agg = agg * v
	}
	fmt.Printf("Triple: %v\n", triple)
	fmt.Printf("Triple Result: %d\n", agg)
}

func findFirstPairTargetSum(data []int, target int) pair {
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i]+data[j] == target {
				return pair{item1: data[i], item2: data[j]}
			}
		}
	}

	return pair{}
}

func findFirstNItemsTargetSum(data []int, n int, target int) []int {
	var result []int
	if n == 2 {
		pair := findFirstPairTargetSum(data, target)
		return []int{pair.item1, pair.item2}
	}

	for i := 0; i < len(data); i++ {
		v := data[i]
		result = findFirstNItemsTargetSum(data[i+1:], n-1, target-v)
		sum := v
		for _, val := range result {
			sum += val
		}
		if sum == target {
			result := append(result, v)
			return result
		}
	}

	return result
}
