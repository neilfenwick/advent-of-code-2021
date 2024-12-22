package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		file *os.File
		err  error
	)

	switch len(os.Args) {
	case 1:
		file = os.Stdin
	case 2:
		file, err = os.Open(os.Args[1])
		if err != nil {
			log.Fatalf("Error opening file: %s", os.Args[1])
		}
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	reports := parseReports(file)
	analyzeReports(reports)
}

func parseReports(file *os.File) [][]int {
	reports := make([][]int, 0, 1000)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// Split the line into words (numbers)
		words := strings.Fields(scanner.Text())
		var report []int

		for _, word := range words {
			num, err := strconv.Atoi(word)
			if err != nil {
				fmt.Println("Error converting to int:", err)
				continue
			}
			report = append(report, num)
		}

		reports = append(reports, report)
	}

	return reports
}

func analyzeReports(reports [][]int) {
	safeCount := 0

	for line, report := range reports {
		stableIncreasing := true
		stableDecreasing := true

		for i := 1; i < len(report); i++ {
			levelDelta := report[i] - report[i-1]
			switch {
			case abs(levelDelta) < 1 || abs(levelDelta) > 3:
				stableIncreasing = false
				stableDecreasing = false
			case levelDelta < 0:
				stableIncreasing = false
			default:
				stableDecreasing = false
			}
		}

		if stableIncreasing || stableDecreasing {
			safeCount++
			fmt.Printf("Safe report on line %d: %v \n", line+1, report)
		}
	}

	fmt.Printf("Safe reports: %d\n", safeCount)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
