package main

import (
	"advent-of-code/readfile"
	"testing"
)

// Validate input file can be opened
func TestReadfile(t *testing.T) {
	expectedResult := 3
	inputFile := readfile.ReadFile("input_test_1.txt")
	result := len(inputFile.InputRow)
	if result != expectedResult {
		t.Errorf("Read of input file was incorrect, got: %d, want: %d.", result, expectedResult)
	}
}

// // Validate 1 Pair
// func TestCardStrength1P(t *testing.T) {
// 	expectedResult := 2
// 	input := "J56KK"
// 	result := calcStrength(input)
// 	if result != expectedResult {
// 		t.Errorf("Parsing of Times did not return correct value, got: %d, want: %d.", result, expectedResult)
// 	}
// }
