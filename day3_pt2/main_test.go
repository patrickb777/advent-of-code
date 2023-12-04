package main

import (
	"advent-of-code/readfile"
	"testing"
)

func TestReadfile(t *testing.T) {
	inputFile := readfile.ReadFile("engine-test-1.txt")
	if len(inputFile.InputRow) != 10 {
		t.Errorf("Read of input file was incorrect, got: %d, want: %d.", len(inputFile.InputRow), 4)
	}
}

// test cases need updating

func TestFullResult1(t *testing.T) {
	inputFile := readfile.ReadFile("engine-test-1.txt")
	symbols, numbers := parseSchematic(inputFile)
	partNos := getPartNumbers(symbols, numbers)
	total := getSumOfParts(partNos)
	if total != 467835 {
		t.Errorf("Gear Ratios error: %d, want: %d.", total, 467835)
	}
}

func TestFullResult2(t *testing.T) {
	inputFile := readfile.ReadFile("engine-test-2.txt")
	symbols, numbers := parseSchematic(inputFile)
	partNos := getPartNumbers(symbols, numbers)
	total := getSumOfParts(partNos)
	if total != 0 {
		t.Errorf("Gear Ratios error: %d, want: %d.", total, 0)
	}
}

func TestFullResult3(t *testing.T) {
	inputFile := readfile.ReadFile("engine-test-3.txt")
	symbols, numbers := parseSchematic(inputFile)
	partNos := getPartNumbers(symbols, numbers)
	total := getSumOfParts(partNos)
	if total != 0 {
		t.Errorf("Gear Ratios error: %d, want: %d.", total, 0)
	}
}
