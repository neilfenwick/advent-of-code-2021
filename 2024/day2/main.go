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
	analyzeReports(reports, undampedReportAnaylyzer)
	analyzeReports(reports, dampedReportAnaylyzer)
}

type reportAnalyzer func([]int) bool

func parseReports(file *os.File) [][]int {
	reports := make([][]int, 0, 1000)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
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

func analyzeReports(reports [][]int, analyzer reportAnalyzer) {
	safeCount := 0

	for _, report := range reports {
		if analyzer(report) {
			safeCount++
		}
	}

	fmt.Printf("Safe reports: %d\n", safeCount)
}

func undampedReportAnaylyzer(report []int) bool {
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

	return stableIncreasing || stableDecreasing
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func dampedReportAnaylyzer(report []int) bool {
	if undampedReportAnaylyzer(report) {
		return true
	}

	// Try removing one element at a time to test if the report is is valid without it
	for i := 0; i < len(report); i++ {
		newReport := make([]int, len(report)-1)
		copy(newReport, report[:i])
		copy(newReport[i:], report[i+1:])
		if undampedReportAnaylyzer(newReport) {
			return true
		}
	}

	return false
}
