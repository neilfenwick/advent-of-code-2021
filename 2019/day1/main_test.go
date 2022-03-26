package main

import (
	"testing"
)

func TestCalculationFuelForWeightRoundsDown(t *testing.T) {
	result := calculationFuelForWeight(100756)
	if result != 33583 {
		t.Errorf("Expected 33583, got %d", result)
	}
}

func TestCalculateModuleInclusiveFuelIncludesRecursiveFuelWeight(t *testing.T) {
	result := calculateModuleInclusiveFuel(1969)
	if result != 966 {
		t.Errorf("Expected 966, got %d", result)
	}
}
