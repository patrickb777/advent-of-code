package main

import (
	"advent-of-code/readfile"
	"testing"
)

// Validate input file can be opened
func TestReadfile(t *testing.T) {
	expectedResult := 5
	inputFile := readfile.ReadFile("cards_test.txt")
	result := len(inputFile.InputRow)
	if result != expectedResult {
		t.Errorf("Read of input file was incorrect, got: %d, want: %d.", result, expectedResult)
	}
}

// Validate 1 Pair
func TestCardStrength1P(t *testing.T) {
	expectedResult := 2
	input := "J56KK"
	result := calcStrength(input)
	if result != expectedResult {
		t.Errorf("Parsing of Times did not return correct value, got: %d, want: %d.", result, expectedResult)
	}
}

// Validate 2 Pairs
func TestCardStrength2P(t *testing.T) {
	expectedResult := 3
	input := "33699"
	result := calcStrength(input)
	if result != expectedResult {
		t.Errorf("Parsing of Times did not return correct value, got: %d, want: %d.", result, expectedResult)
	}
}

// Validate 3 of a Kind
func TestCardStrength3K(t *testing.T) {
	expectedResult := 4
	input := "Q3777"
	result := calcStrength(input)
	if result != expectedResult {
		t.Errorf("Parsing of Times did not return correct value, got: %d, want: %d.", result, expectedResult)
	}
}

// Validate Full House
func TestCardStrengthFH(t *testing.T) {
	expectedResult := 5
	input := "77QQ7"
	result := calcStrength(input)
	if result != expectedResult {
		t.Errorf("Parsing of Times did not return correct value, got: %d, want: %d.", result, expectedResult)
	}
}

// Validate 4 of a Kind
func TestCardStrength4K(t *testing.T) {
	expectedResult := 6
	input := "66656"
	result := calcStrength(input)
	if result != expectedResult {
		t.Errorf("Parsing of Times did not return correct value, got: %d, want: %d.", result, expectedResult)
	}
}

// Validate 5 of a Kind
func TestCardStrength5k(t *testing.T) {
	expectedResult := 7
	input := "TTTTT"
	result := calcStrength(input)
	if result != expectedResult {
		t.Errorf("Parsing of Times did not return correct value, got: %d, want: %d.", result, expectedResult)
	}
}
