package main

import (
	"advent-of-code/readfile"
	"testing"
)

// Test input file can be correctly read
func TestReadfile(t *testing.T) {
	input := readfile.ReadFile("engine-test.txt")
	if len(input.InputRow) != 10 {
		t.Errorf("Read of input file was incorrect, got: %d, want: %d.", len(input.InputRow), 4)
	}
}
