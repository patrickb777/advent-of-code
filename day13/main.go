package main

import (
	"advent-of-code/readfile"
	"flag"
	"fmt"
	"log"
	"time"
)

type LavaIsland struct {
	Pattern []Pattern
}

type Pattern struct {
	Pattern [][]byte
}

func main() {
	// Init
	start := time.Now()
	fmt.Printf("\n[Point of Incidence]\n")
	f := flag.String("f", "none", "Input file")
	flag.Parse()

	// Parse Input
	inputFile := readfile.ReadFile(*f)
	fmt.Sprintln(inputFile) //Does not output to stdout
	lavaIsland := parseData(inputFile)

	// for i, pat := range lavaIsland.Pattern {
	// 	fmt.Println("Pattern:", i)
	// 	for _, p := range pat.Pattern {
	// 		fmt.Println(p)
	// 	}
	// }

	// Processing

	// look into INPUT TEST 3
	reflectionPoints := []int{}
	// Look for horizontal reflection points
	for _, pat := range lavaIsland.Pattern {
		//fmt.Printf("\nProcessing pattern %d\n", i)
		match := 1
		for p := 0; p < len(pat.Pattern)-1; p++ {
			// fmt.Println("Pattern p: ", pat.Pattern[p])
			// fmt.Println("Pattern p+1: ", pat.Pattern[p+1])
			// fmt.Print("\n")
			match = compare(pat.Pattern[p], pat.Pattern[p+1])
			if match == 1 {
				match = checkReflectionPoint(pat.Pattern, p)
			}
			if match == 1 {
				value := (p + 1) * 100
				// append reflection point and value as pairs for debugging, will remove point once confident with results.
				//reflectionPoints = append(reflectionPoints, i+1)
				//reflectionPoints = append(reflectionPoints, p)
				reflectionPoints = append(reflectionPoints, value)
				break
			}
		}
		// If no horzontal matches are found, transpose the vertical patterns to horizontal and repeat checks
		if match == 0 {
			trn := Pattern{}
			for xp := 0; xp < len(pat.Pattern[0]); xp++ {
				transpose := []byte{}
				for yp := range pat.Pattern {
					transpose = append(transpose, pat.Pattern[yp][xp])
				}
				trn.Pattern = append(trn.Pattern, transpose)
			}
			for p := 0; p < len(trn.Pattern)-1; p++ {
				match = compare(trn.Pattern[p], trn.Pattern[p+1])
				if match == 1 {
					match = checkReflectionPoint(trn.Pattern, p)
				}
				if match == 1 {
					value := (p + 1)
					// append reflection point and value as pairs for debugging, will remove point once confident with results.
					//reflectionPoints = append(reflectionPoints, i+1)
					//reflectionPoints = append(reflectionPoints, p)
					reflectionPoints = append(reflectionPoints, value)
					break
				}
			}

		}

	}
	total := 0
	for t := range reflectionPoints {
		total = total + reflectionPoints[t]
	}
	//fmt.Println("Reflection Points", reflectionPoints)
	fmt.Println("Reflection Point Values", total)

	// Output execution time
	elapsed := time.Since(start)
	fmt.Printf("\n")
	log.Printf("Execution time %s\n", elapsed)
}

func checkReflectionPoint(pattern [][]byte, p int) int {
	reflectionPoint := 1
	for x := 1; x < (len(pattern)-1)-p; x++ { // set upper range limit
		if p-x < 0 || reflectionPoint == 0 { // break if lower range limit would be exceeded
			break
		} else {
			reflectionPoint = compare(pattern[p-x], pattern[p+1+x]) // compare row pairs from potential reflection point
		}
	}
	return reflectionPoint
}

func compare(patternA []byte, patternB []byte) int {
	match := 1
	// fmt.Println("Pattern A: ", patternA)
	// fmt.Println("Pattern B: ", patternB)
	for i := 0; i < len(patternA); i++ {
		if patternA[i] != patternB[i] {
			match = 0
		}
	}
	// fmt.Println("Match Flag: ", match)
	// fmt.Print("\n")
	return match
}

func parseData(input readfile.InputFile) LavaIsland {
	lavaIsland := LavaIsland{}
	pattern := Pattern{}
	for _, row := range input.InputRow {
		var dataRow []byte
		if row == "EOF" {
			break
		}
		if row != "" {
			for i := 0; i < len(row); i++ {
				dataRow = append(dataRow, row[i])
			}
			pattern.Pattern = append(pattern.Pattern, dataRow)
		} else {
			lavaIsland.Pattern = append(lavaIsland.Pattern, pattern)
			pattern = Pattern{}
		}
	}
	return lavaIsland
}
