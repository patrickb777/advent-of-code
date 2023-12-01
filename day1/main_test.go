package main

import "testing"

// Test input file can be correctly read
func TestReadfile(t *testing.T) {
	input := readFile("calibration_test.txt")
	if len(input.Input) != 4 {
		t.Errorf("Read of input file was incorrect, got: %d, want: %d.", len(input.Input), 4)
	}
}

// Test final calculation
