package main

import "testing"

// Test input file can be correctly read
func TestReadfile(t *testing.T) {
	input := readFile("calibration_test.txt")
	if len(input.Input) != 4 {
		t.Errorf("Read of input file was incorrect, got: %d, want: %d.", len(input.Input), 4)
	}
}

// Test coordinate extraction
func TestGetCalibrationNum(t *testing.T) {
	c := Calibrations{Input: []string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"}}
	calibrations := getCalibrationNum(c)
	if calibrations.CalibrationNum[2] != 15 {
		t.Errorf("Coordinate extraction incorrect, got: %d, want: %d.", calibrations.CalibrationNum[2], 15)
	}
}
