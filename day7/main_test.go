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

// Validate parsing of times
// func TestParseRaces(t *testing.T) {
// 	expectedResult := 15
// 	inputFile := readfile.ReadFile("races_test.txt")
// 	races := parseRaces(inputFile)
// 	if races.Time[1] != expectedResult {
// 		t.Errorf("Parsing of Times did not return correct value, got: %d, want: %d.", races.Time[1], expectedResult)
// 	}
// }
