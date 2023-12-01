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
func TestGetCalibrationNum(t *testing.T) {
	c := Calibrations{Input: []string{"shrzvdcghblt21", "sixdddkcqjdnzzrgfourxjtwosevenhg9", "threevt1onegxgvc9flk", "7dmqzksnlcpbsqkzqlfour1four"}}
	calibrations := getCalibrationNum(c)
	if calibrations.CalibrationNum[2] != 39 {
		t.Errorf("Coordinate extraction incorrect, got: %d, want: %d.", calibrations.CalibrationNum[2], 39)
	}
}

// Test coordinate extraction with overlapping strings
func TestGetCalibrationNumOverlaps(t *testing.T) {
	c := Calibrations{Input: []string{"shrzvdcghblt21", "hxtwoneqpmbfgkhnr6three86eight7five", "fivenpblbgfive6moneighttzj"}}
	calibrations := getCalibrationNum(c)
	if calibrations.CalibrationNum[2] != 58 {
		t.Errorf("Coordinate extraction incorrect, got: %d, want: %d.", calibrations.CalibrationNum[2], 58)
	}
}
