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
