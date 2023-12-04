package main

import (
	"advent-of-code/readfile"
	"testing"
)

func TestReadfile(t *testing.T) {
	inputFile := readfile.ReadFile("scratchcard_test_input.txt")
	if len(inputFile.InputRow) != 6 {
		t.Errorf("Read of input file was incorrect, got: %d, want: %d.", len(inputFile.InputRow), 6)
	}
}

func TestFullResult1(t *testing.T) {
	inputFile := readfile.ReadFile("scratchcard_test_1.txt")
	worth := parseScratchcard(inputFile.InputRow[0])
	if worth != 8 {
		t.Errorf("Scratchcard worth incorrectly calculated, got: %d, want: %d.", worth, 8)
	}
}

func TestFullResult2(t *testing.T) {
	inputFile := readfile.ReadFile("scratchcard_test_2.txt")
	worth := parseScratchcard(inputFile.InputRow[0])
	if worth != 0 {
		t.Errorf("Scratchcard worth incorrectly calculated,  got: %d, want: %d.", worth, 0)
	}
}
