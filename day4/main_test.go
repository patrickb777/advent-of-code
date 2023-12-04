package main

import (
	"advent-of-code/readfile"
	"testing"
)

func TestReadfile(t *testing.T) {
	inputFile := readfile.ReadFile("scratchcard-test-1.txt")
	if len(inputFile.InputRow) != 6 {
		t.Errorf("Read of input file was incorrect, got: %d, want: %d.", len(inputFile.InputRow), 6)
	}
}

// func TestFullResult1(t *testing.T) {
// 	fmt.Printf("\nTestcase 1")
// 	inputFile := readfile.ReadFile("engine-test-1.txt")
// 	symbols, numbers := parseSchematic(inputFile)
// 	partNos := getPartNumbers(symbols, numbers)
// 	total := getSumOfParts(partNos)
// 	if total != 4361 {
// 		t.Errorf("Sum of parts calculation incorrect, got: %d, want: %d.", total, 4361)
// 	}
// }
