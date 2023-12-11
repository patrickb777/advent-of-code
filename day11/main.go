package main

import (
	"advent-of-code/readfile"
	"flag"
	"fmt"
	"log"
	"time"
)

func main() {
	// Init
	start := time.Now()
	fmt.Printf("\n.   *  .  .     .  * ..   [Cosmic Expansion] .        .        .   *    ..\n")
	f := flag.String("f", "none", "Input file")
	flag.Parse()

	// Parse Input
	inputFile := readfile.ReadFile(*f)
	fmt.Sprintln(inputFile) //Does not output to stdout
	universe := parseData(inputFile)
	for i := range universe {
		fmt.Println(universe[i])
	}

	// Processing
	//
	// Expand the Universe on Y Axis
	for y := 0; y < len(universe); y++ {
		expandY := true
		for x := 0; x < len(universe[y]); x++ {
			//fmt.Println("\n", x, " = ", universe[y][x], " | ", expand)
			if universe[y][x] == 35 {
				expandY = false
				break
			}
		}
		if expandY {
			log.Printf("Found no universe in row %d", y)
			universe = append(universe[:y+1], universe[y:]...)
			universe[y] = universe[y+1]
			y++
		}
	}
	// Expand the Universe on X Axis
	for x := 0; x < len(universe[0]); x++ {
		expandX := true
		for y := 0; y < len(universe); y++ {
			//fmt.Println("\n", x, " = ", universe[y][x], " | ", expand)
			if universe[y][x] == 35 {
				expandX = false
				break
			}
		}
		if expandX {
			log.Printf("Found no universe on X axis %d", x)
			for y := 0; y < len(universe); y++ {
				universe[y] = append(universe[y][:x+1], universe[y][x:]...)
				universe[y][x] = universe[y][x+1]
			}
			x++
		}
	}

	fmt.Println("Expanded universe:")
	for i := range universe {
		fmt.Println(universe[i])
	}

	// Output execution time
	elapsed := time.Since(start)
	fmt.Printf("\n")
	log.Printf("Execution time %s\n", elapsed)
}

func parseData(input readfile.InputFile) [][]byte {
	var universe [][]byte

	for _, row := range input.InputRow {
		var dataRow []byte
		for i := 0; i < len(row); i++ {
			dataPoint := row[i]
			dataRow = append(dataRow, dataPoint)
		}
		universe = append(universe, dataRow)
	}
	return universe
}
