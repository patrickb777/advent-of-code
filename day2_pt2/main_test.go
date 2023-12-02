package main

import "testing"

// Test input file can be correctly read
func TestReadfile(t *testing.T) {
	input := readFile("game_results_test.txt")
	if len(input.InputRow) != 5 {
		t.Errorf("Read of input file was incorrect, got: %d, want: %d.", len(input.InputRow), 5)
	}
}
