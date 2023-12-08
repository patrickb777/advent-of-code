package main

import (
	"advent-of-code/readfile"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"
)

func main() {
	// Init
	start := time.Now()
	fmt.Println("‹’’›(Ͼ˳Ͽ)‹’’›")
	f := flag.String("f", "none", "Input file")
	flag.Parse()

	// Parse Input
	inputFile := readfile.ReadFile(*f)
	//fmt.Println(inputFile)
	directions, navigation := parseDirections(inputFile)

	// Processing
	// rune 65 == A
	// rune 90 == Z

	// What are the possible start positions
	var startPositions []string
	var endPositions []string
	for pos := range navigation {
		switch pos[2] {
		case 65:
			startPositions = append(startPositions, pos)
		case 90:
			endPositions = append(endPositions, pos)
		}
	}

	position := "AAA"
	endFlag := 0
	pedometer := 0

	for {
		for d := 0; d < len(directions); d++ {
			position = nextStep(position, navigation, directions[d])
			pedometer++
		}

		if position == "ZZZ" {
			endFlag = 1
		}
		if endFlag == 1 {
			break
		}
	}

	fmt.Printf("Total steps taken: %d\n", pedometer)

	fmt.Sprintln(directions, endPositions)
	// Output execution time
	elapsed := time.Since(start)
	log.Printf("Execution time %s\n", elapsed)
}

func nextStep(position string, navigation map[string][2]string, directions rune) string {
	fmt.Println("Current Position:", position)
	switch directions {
	case 76: // left
		position = navigation[position][0]
	case 82: // right
		position = navigation[position][1]

	}
	fmt.Println("Moving to:", position)
	//fmt.Println("Moving to:", position)
	return position
}

func parseDirections(input readfile.InputFile) ([]rune, map[string][2]string) {
	// rune 76 == L
	// rune 82 == R
	var directions []rune
	nav := make(map[string][2]string)
	for _, d := range input.InputRow[0] {
		directions = append(directions, d)
	}
	for i, row := range input.InputRow {
		if i >= 2 {
			splitStr := strings.Split(row, " = ")
			key := splitStr[0]
			v := strings.Split(strings.TrimSuffix(strings.TrimPrefix(splitStr[1], "("), ")"), ", ") // trim leading "(", then trailing ")", then split on ", "
			val := [2]string{v[0], v[1]}
			nav[key] = val
		}
	}
	return directions, nav
}
