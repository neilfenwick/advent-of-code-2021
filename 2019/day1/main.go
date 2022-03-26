package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type fuelCalc func(int) int

func main() {
	start()
}

func start() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	fuel := calculateRocketFuel(file, calculationFuelForWeight)
	fmt.Printf("Part 1: %d\n", fuel)

	_, _ = file.Seek(0, 0)
	fuel = calculateRocketFuel(file, calculateModuleInclusiveFuel)
	fmt.Printf("Part 2: %d\n", fuel)
}

func calculateRocketFuel(r io.Reader, fn fuelCalc) int {
	var (
		fuel int
	)
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		var n int
		_, err := fmt.Sscanf(s.Text(), "%d", &n)
		if err != nil {
			log.Fatalf("Could not read %s: %v", s.Text(), err)
		}
		fuel += fn(n)
	}
	return fuel
}

func calculationFuelForWeight(weight int) int {
	return weight/3 - 2
}

func calculateModuleInclusiveFuel(weight int) int {
	var totalFuel int
	fuel := calculationFuelForWeight(weight)
	for fuel > 0 {
		totalFuel += fuel
		fuel = calculationFuelForWeight(fuel)
	}
	return totalFuel
}
