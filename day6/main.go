package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		file         *os.File
		numberOfDays int = 80
		err          error
	)

	switch len(os.Args) {
	case 1:
		file = os.Stdin
	case 3:
		numberOfDays, err = strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("Did not understand window size of: %s\n", os.Args[2])
		}

		fallthrough
	case 2:
		file, err = os.Open(os.Args[1])
		if err != nil {
			log.Fatalf("Error opening file: %s", os.Args[1])
		}
	}
	defer file.Close()

	result := countFishAfterDays(file, numberOfDays)
	fmt.Println(result)
}

func countFishAfterDays(r io.Reader, numberOfDays int) int64 {
	var (
		result int64
	)
	fish := readFish(r)

	for i := 0; i < numberOfDays; i++ {
		fish = runGeneration(fish)
	}

	for _, v := range fish {
		result += v
	}
	return result
}

func readFish(r io.Reader) map[int]int64 {
	fish := make(map[int]int64)
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		numbers := strings.Split(s.Text(), ",")
		for _, n := range numbers {
			f, err := strconv.Atoi(n)
			if err != nil {
				log.Fatalf("Could not convert '%s' to int", n)
			}
			fish[f] = fish[f] + 1
		}
	}
	return fish
}

func runGeneration(fish map[int]int64) map[int]int64 {
	result := make(map[int]int64)
	for age := 0; age < 9; age++ {
		count := fish[age]
		switch age {
		case 0:
			result[6] = result[6] + count
			result[8] = result[8] + count
		case 7:
			result[6] = result[6] + count
		default:
			result[age-1] = count
		}
	}
	return result
}
