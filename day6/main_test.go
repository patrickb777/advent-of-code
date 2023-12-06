package main

import (
	"advent-of-code/readfile"
	"testing"
)

// Validate input file can be opened
func TestReadfile(t *testing.T) {
	expectedResult := 2
	inputFile := readfile.ReadFile("races.txt")
	if len(inputFile.InputRow) != expectedResult {
		t.Errorf("Read of input file was incorrect, got: %d, want: %d.", len(inputFile.InputRow), expectedResult)
	}
}

// Validate parsing of Seed list works correctly
// func TestParseAlmanac1(t *testing.T) {
// 	expectedResult := 4
// 	inputFile := readfile.ReadFile("almanac_test.txt")
// 	seeds, _ := parseAlmanac(inputFile)
// 	if len(seeds.Seeds) != expectedResult {
// 		t.Errorf("Almanac parsing did not return correct number of seeds, got: %d, want: %d.", len(seeds.Seeds), expectedResult)
// 	}
// }
