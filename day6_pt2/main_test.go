package main

import (
	"advent-of-code/readfile"
	"testing"
)

// Validate input file can be opened
func TestReadfile(t *testing.T) {
	expectedResult := 2
	inputFile := readfile.ReadFile("races.txt")
	if len(inputFile.InputRow) != expectedResult {
		t.Errorf("Read of input file was incorrect, got: %d, want: %d.", len(inputFile.InputRow), expectedResult)
	}
}

// Validate parsing of times
func TestParseRaces(t *testing.T) {
	expectedResult := 15
	inputFile := readfile.ReadFile("races_test.txt")
	races := parseRaces(inputFile)
	if races.Time[1] != expectedResult {
		t.Errorf("Parsing of Times did not return correct value, got: %d, want: %d.", races.Time[1], expectedResult)
	}
}

// Validate parsing of distances
func TestParseRaces2(t *testing.T) {
	expectedResult := 200
	inputFile := readfile.ReadFile("races_test.txt")
	races := parseRaces(inputFile)
	if races.Distance[2] != expectedResult {
		t.Errorf("Parsing of Times did not return correct value, got: %d, want: %d.", races.Distance[2], expectedResult)
	}
}
