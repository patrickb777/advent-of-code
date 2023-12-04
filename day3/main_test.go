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

func TestFullResult1(t *testing.T) {
	inputFile := readfile.ReadFile("engine-test-1.txt")
	symbols, numbers := parseSchematic(inputFile)
	partNos := getPartNumbers(symbols, numbers)
	total := getSumOfParts(partNos)
	if total != 4361 {
		t.Errorf("Sum of parts calculation incorrect, got: %d, want: %d.", total, 4361)
	}
}

func TestFullResult2(t *testing.T) {
	inputFile := readfile.ReadFile("engine-test-2.txt")
	symbols, numbers := parseSchematic(inputFile)
	partNos := getPartNumbers(symbols, numbers)
	total := getSumOfParts(partNos)
	if total != 413 {
		t.Errorf("Sum of parts calculation incorrect, got: %d, want: %d.", total, 413)
	}
}

func TestFullResult3(t *testing.T) {
	inputFile := readfile.ReadFile("engine-test-3.txt")
	symbols, numbers := parseSchematic(inputFile)
	partNos := getPartNumbers(symbols, numbers)
	total := getSumOfParts(partNos)
	if total != 925 {
		t.Errorf("Sum of parts calculation incorrect, got: %d, want: %d.", total, 925)
	}
}

func TestFullResult4(t *testing.T) {
	inputFile := readfile.ReadFile("engine-test-4.txt")
	symbols, numbers := parseSchematic(inputFile)
	partNos := getPartNumbers(symbols, numbers)
	total := getSumOfParts(partNos)
	if total != 8082 {
		t.Errorf("Sum of parts calculation incorrect, got: %d, want: %d.", total, 8082)
	}
}
func TestFullResult41(t *testing.T) {
	inputFile := readfile.ReadFile("engine-test-4-1.txt")
	symbols, numbers := parseSchematic(inputFile)
	partNos := getPartNumbers(symbols, numbers)
	total := getSumOfParts(partNos)
	if total != 1351 {
		t.Errorf("Sum of parts calculation incorrect, got: %d, want: %d.", total, 1351)
	}
}

func TestFullResult42(t *testing.T) {
	inputFile := readfile.ReadFile("engine-test-4-2.txt")
	symbols, numbers := parseSchematic(inputFile)
	partNos := getPartNumbers(symbols, numbers)
	total := getSumOfParts(partNos)
	if total != 1134 {
		t.Errorf("Sum of parts calculation incorrect, got: %d, want: %d.", total, 1134)
	}
}

func TestFullResult5(t *testing.T) {
	inputFile := readfile.ReadFile("engine-test-5.txt")
	symbols, numbers := parseSchematic(inputFile)
	partNos := getPartNumbers(symbols, numbers)
	total := getSumOfParts(partNos)
	if total != 1931 {
		t.Errorf("Sum of parts calculation incorrect, got: %d, want: %d.", total, 1931)
	}
}
