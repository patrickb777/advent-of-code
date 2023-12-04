package main

import (
	"advent-of-code/readfile"
	"fmt"
	"testing"
)

func TestReadfile(t *testing.T) {
	inputFile := readfile.ReadFile("engine-test-1.txt")
	if len(inputFile.InputRow) != 10 {
		t.Errorf("Read of input file was incorrect, got: %d, want: %d.", len(inputFile.InputRow), 4)
	}
}

// test cases need updating

func TestFullResult1(t *testing.T) {
	inputFile := readfile.ReadFile("engine-test-1.txt")
	gears, cogs := parseSchematic(inputFile)
	cogs = getGearAttachments(gears, cogs)
	total := calcGearRatios(gears, cogs)
	if total != 467835 {
		t.Errorf("Gear Ratios error, got: %d, want: %d.", total, 467835)
	}
}

func TestFullResult2(t *testing.T) {
	inputFile := readfile.ReadFile("engine-test-2.txt")
	gears, cogs := parseSchematic(inputFile)
	cogs = getGearAttachments(gears, cogs)
	total := calcGearRatios(gears, cogs)
	if total != 1038734 {
		t.Errorf("Gear Ratios error, got: %d, want: %d.", total, 1038734)
	}
}

func TestFullResult3(t *testing.T) {
	inputFile := readfile.ReadFile("engine-test-3.txt")
	gears, cogs := parseSchematic(inputFile)
	cogs = getGearAttachments(gears, cogs)
	total := calcGearRatios(gears, cogs)
	if total != 3882133 {
		t.Errorf("Gear Ratios error, got: %d, want: %d.", total, 3882133)
	}
}
func TestFullResult4(t *testing.T) {
	inputFile := readfile.ReadFile("engine-test-4.txt")
	gears, cogs := parseSchematic(inputFile)
	cogs = getGearAttachments(gears, cogs)
	total := calcGearRatios(gears, cogs)
	if total != 3340266 {
		t.Errorf("Gear Ratios error, got: %d, want: %d.", total, 3340266)
	}
}

func TestFullResult5(t *testing.T) {
	fmt.Println("Test ....")
	inputFile := readfile.ReadFile("engine-test-5.txt")
	gears, cogs := parseSchematic(inputFile)
	cogs = getGearAttachments(gears, cogs)
	total := calcGearRatios(gears, cogs)
	if total != 3384871 {
		t.Errorf("Gear Ratios error, got: %d, want: %d.", total, 3384871)
	}
}

func TestFullResult6(t *testing.T) {
	fmt.Println("Test ....")
	inputFile := readfile.ReadFile("engine-test-6.txt")
	gears, cogs := parseSchematic(inputFile)
	cogs = getGearAttachments(gears, cogs)
	total := calcGearRatios(gears, cogs)
	if total != 3848758 {
		t.Errorf("Gear Ratios error, got: %d, want: %d.", total, 3848758)
	}
}
