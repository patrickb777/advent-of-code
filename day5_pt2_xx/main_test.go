package main

import (
	"advent-of-code/readfile"
	"testing"
)

// Validate input file can be opened
func TestReadfile(t *testing.T) {
	expectedResult := 33
	inputFile := readfile.ReadFile("almanac_test.txt")
	if len(inputFile.InputRow) != expectedResult {
		t.Errorf("Read of input file was incorrect, got: %d, want: %d.", len(inputFile.InputRow), expectedResult)
	}
}

// Validate parsing of Seed list works correctly
func TestParseAlmanac1(t *testing.T) {
	expectedResult := 4
	inputFile := readfile.ReadFile("almanac_test.txt")
	seeds, _ := parseAlmanac(inputFile)
	if len(seeds.Seeds) != expectedResult {
		t.Errorf("Almanac parsing did not return correct number of seeds, got: %d, want: %d.", len(seeds.Seeds), expectedResult)
	}
}

// Validate parsing of Seed to Soil map works correctly
func TestParseAlmanac2(t *testing.T) {
	expectedResult := Mapdata{Map: "seed-to-soil", Dest: 52, Source: 50, Range: 48}
	inputFile := readfile.ReadFile("almanac_test.txt")
	_, maps := parseAlmanac(inputFile)
	if maps.Mapdata[1] != expectedResult {
		t.Errorf("Almanac parsing did not return correct number of seeds, got: %v, want: %v.", maps.Mapdata[1], expectedResult)
	}
}

// Validate parsing of Light to Temperature map works correctly
func TestParseAlmanac3(t *testing.T) {
	expectedResult := Mapdata{Map: "light-to-temperature", Dest: 68, Source: 64, Range: 13}
	inputFile := readfile.ReadFile("almanac_test.txt")
	_, maps := parseAlmanac(inputFile)
	if maps.Mapdata[13] != expectedResult {
		t.Errorf("Almanac parsing did not return correct number of seeds, got: %v, want: %v.", maps.Mapdata[13], expectedResult)
	}
}

// Validate seed to soil lookup - source and destination are equal values
func TestSeedSoil1(t *testing.T) {
	seed := 98
	m := "seed-to-soil"
	expectedResult := 50
	inputFile := readfile.ReadFile("almanac_test.txt")
	_, maps := parseAlmanac(inputFile)
	soil := mapLookup(seed, m, maps)
	if soil != expectedResult {
		t.Errorf("Almanac parsing did not return correct number of seeds, got: %v, want: %v.", soil, expectedResult)
	}
}

// Validate seed to soil lookup - number in range
func TestSeedSoil2(t *testing.T) {
	seed := 99
	m := "seed-to-soil"
	expectedResult := 51
	inputFile := readfile.ReadFile("almanac_test.txt")
	_, maps := parseAlmanac(inputFile)
	soil := mapLookup(seed, m, maps)
	if soil != expectedResult {
		t.Errorf("Almanac parsing did not return correct number of seeds, got: %v, want: %v.", soil, expectedResult)
	}
}

// Validate seed to soil lookup - unmapped number
func TestSeedSoil3(t *testing.T) {
	seed := 10
	m := "seed-to-soil"
	expectedResult := 10
	inputFile := readfile.ReadFile("almanac_test.txt")
	_, maps := parseAlmanac(inputFile)
	soil := mapLookup(seed, m, maps)
	if soil != expectedResult {
		t.Errorf("Almanac parsing did not return correct number of seeds, got: %v, want: %v.", soil, expectedResult)
	}
}

// Validate soil-to-fertilizer lookup - destination is 0
func TestSoilFertl1(t *testing.T) {
	soil := 15
	m := "soil-to-fertilizer"
	expectedResult := 0
	inputFile := readfile.ReadFile("almanac_test.txt")
	_, maps := parseAlmanac(inputFile)
	fert := mapLookup(soil, m, maps)
	if fert != expectedResult {
		t.Errorf("Almanac parsing did not return correct number of seeds, got: %v, want: %v.", fert, expectedResult)
	}
}

// Validate soil-to-fertilizer lookup - source is 0
func TestSoilFertl2(t *testing.T) {
	soil := 0
	m := "soil-to-fertilizer"
	expectedResult := 39
	inputFile := readfile.ReadFile("almanac_test.txt")
	_, maps := parseAlmanac(inputFile)
	fert := mapLookup(soil, m, maps)
	if fert != expectedResult {
		t.Errorf("Almanac parsing did not return correct number of seeds, got: %v, want: %v.", fert, expectedResult)
	}
}

// Validate seed location function
func TestSeedLocation1(t *testing.T) {
	seed := 79
	expectedResult := 82
	inputFile := readfile.ReadFile("almanac_test.txt")
	_, maps := parseAlmanac(inputFile)
	loc := getLocation(seed, maps)
	if loc != expectedResult {
		t.Errorf("Almanac parsing did not return correct number of seeds, got: %v, want: %v.", loc, expectedResult)
	}
}

func TestSeedLocation2(t *testing.T) {
	seed := 14
	expectedResult := 43
	inputFile := readfile.ReadFile("almanac_test.txt")
	_, maps := parseAlmanac(inputFile)
	loc := getLocation(seed, maps)
	if loc != expectedResult {
		t.Errorf("Almanac parsing did not return correct number of seeds, got: %v, want: %v.", loc, expectedResult)
	}
}

func TestSeedLocation3(t *testing.T) {
	seed := 55
	expectedResult := 86
	inputFile := readfile.ReadFile("almanac_test.txt")
	_, maps := parseAlmanac(inputFile)
	loc := getLocation(seed, maps)
	if loc != expectedResult {
		t.Errorf("Almanac parsing did not return correct number of seeds, got: %v, want: %v.", loc, expectedResult)
	}
}

// Test for final result calculation
func TestGetNearestLocation(t *testing.T) {
	locations := []int{82, 43, 86, 35}
	expectedResult := 35
	loc := getNearestLoc(locations)
	if loc != expectedResult {
		t.Errorf("Almanac parsing did not return correct number of seeds, got: %v, want: %v.", loc, expectedResult)
	}
}
