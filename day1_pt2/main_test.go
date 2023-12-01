package main

import "testing"

// Test input file can be correctly read
func TestReadfile(t *testing.T) {
	input := readFile("calibration_test.txt")
	if len(input.Input) != 6 {
		t.Errorf("Read of input file was incorrect, got: %d, want: %d.", len(input.Input), 6)
	}
}

// Test coordinate extraction
func TestGetCoord(t *testing.T) {
	c := Coordinates{Input: []string{"shrzvdcghblt21", "sixdddkcqjdnzzrgfourxjtwosevenhg9", "threevt1onegxgvc9flk", "7dmqzksnlcpbsqkzqlfour1four"}}
	coord := getCoord(c)
	if coord.Coord[2] != 39 {
		t.Errorf("Coordinate extraction incorrect, got: %d, want: %d.", coord.Coord[2], 39)
	}
}

// Test coordinate extraction with overlapping strings
func TestGetCoordOverlap(t *testing.T) {
	c := Coordinates{Input: []string{"shrzvdcghblt21", "hxtwoneqpmbfgkhnr6three86eight7five", "fivenpblbgfive6moneighttzj"}}
	coord := getCoord(c)
	if coord.Coord[2] != 58 {
		t.Errorf("Coordinate extraction incorrect, got: %d, want: %d.", coord.Coord[2], 58)
	}
}
