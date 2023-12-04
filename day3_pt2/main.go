package main

import (
	"advent-of-code/readfile"
	"flag"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type Numbers struct {
	Metadata []NumMetadata
}

type NumMetadata struct {
	Num   string
	Len   int
	Row   int
	Pos   int
	ProxX [2]int
	ProxY [2]int
	GID   string
}

type Symbols struct {
	Metadata []SymMetadata
}

type SymMetadata struct {
	GID   string
	Char  string
	Row   int
	Pos   int
	ProxX [2]int
	ProxY [2]int
}

func main() {
	start := time.Now()
	fmt.Println("Gear Ratios - Part Two")
	f := flag.String("f", "none", "Input file")
	flag.Parse()
	inputFile := readfile.ReadFile(*f) // moved the reading of the input file to an external module
	gears, cogs := parseSchematic(inputFile)
	cogs = getGearAttachments(gears, cogs)
	total := calcGearRatios(gears, cogs)
	//total := getSumOfParts(ratios)

	fmt.Printf("Sum of Gear Ratios = %d\n", total)

	// Output execution time
	elapsed := time.Since(start)
	log.Printf("Execution time %s\n", elapsed)
}

func calcGearRatios(gears Symbols, cogs Numbers) int {
	gearings := make(map[string][2]int)

	// Add all the gears into the gearings map
	var cogPair = [2]int{0, 0}
	for _, g := range gears.Metadata {
		gearings[g.GID] = cogPair
	}
	//fmt.Println(cogs)
	// Add cogs to the gearings map
	for _, c := range cogs.Metadata {
		_, ok := gearings[c.GID]

		if ok {
			cog, err := strconv.Atoi(c.Num)
			if err != nil {
				log.Fatal(err)
			}
			if gearings[c.GID][0] == 0 {
				cogPair[0] = cog
				cogPair[1] = gearings[c.GID][1]
			} else {
				cogPair[0] = gearings[c.GID][0]
				cogPair[1] = cog
			}
			gearings[c.GID] = cogPair
			//fmt.Println(c.GID, cogPair)
		}

	}

	// Calculate the Gear Ratios
	total := 0
	for k := range gearings {
		total = total + (gearings[k][0] * gearings[k][1])
	}
	return total
}

func getGearAttachments(s Symbols, n Numbers) Numbers {
	for numInd, nmd := range n.Metadata {
		for _, smd := range s.Metadata {
			if smd.Row >= nmd.ProxY[0] && smd.Row <= nmd.ProxY[1] && smd.Pos >= nmd.ProxX[0] && smd.Pos <= nmd.ProxX[1] {
				// Print SymbolMetadata
				//fmt.Printf("ID:%s, Type:%s, Row:%d, Position:%d, X-Range:%d:%d, Y-Range:%d:%d \n", smd.GID, smd.Char, smd.Row, smd.Pos, smd.ProxX[0], smd.ProxX[1], smd.ProxY[0], smd.ProxY[1])
				n.Metadata[numInd].GID = smd.GID
			}
		}
		// if n.Metadata[numInd].GID != "" {
		// 	// Print Number Metadata
		// 	fmt.Printf("Number:%s:, Length:%d, Postion:%d, X-Range:%d:%d, Y-Range:%d:%d, Gear:%s \n", n.Metadata[numInd].Num, n.Metadata[numInd].Len, n.Metadata[numInd].Pos, n.Metadata[numInd].ProxX[0], n.Metadata[numInd].ProxX[1], n.Metadata[numInd].ProxY[0], n.Metadata[numInd].ProxY[1], n.Metadata[numInd].GID)
		// }
	}
	return n
}

func parseSchematic(input readfile.InputFile) (Symbols, Numbers) {
	var symMetadata SymMetadata
	var numMetadata NumMetadata
	symbols := Symbols{}
	numbers := Numbers{}
	symRX := regexp.MustCompile(`[*]`)
	numRX := regexp.MustCompile(`[0-9]+`)
	for r, v := range input.InputRow {
		// Find Symbols and Positions
		runes := []rune(v)
		for p, c := range runes {
			sym := symRX.FindAllString(string(c), -1)
			if len(sym) > 0 {
				symMetadata.Char = sym[0]
				symMetadata.Row = r
				symMetadata.Pos = p
				symMetadata.ProxX[0] = symMetadata.Pos - 1
				symMetadata.ProxX[1] = symMetadata.Pos + 1
				symMetadata.ProxY[0] = r - 1
				symMetadata.ProxY[1] = r + 1
				uuid := uuid.New()
				symMetadata.GID = uuid.String()
				symbols.Metadata = append(symbols.Metadata, symMetadata)
			}
		}
		// Find Numbers and Positions
		num := numRX.FindAllString(v, -1)
		numInd := numRX.FindAllStringIndex(v, -1)
		//fmt.Println(v)
		for i, n := range num {
			numMetadata.Num = n
			//fmt.Printf("Number:%s <<>> %s \n", numMetadata.Num, num[i])
			//fmt.Printf("Number Index:%v \n", numInd[i][0])
			numMetadata.Len = len(n)
			numMetadata.Row = r
			numMetadata.Pos = numInd[i][0]
			numMetadata.ProxX[0] = numMetadata.Pos - 1
			numMetadata.ProxX[1] = numMetadata.Pos + numMetadata.Len
			numMetadata.ProxY[0] = r - 1
			numMetadata.ProxY[1] = r + 1
			numbers.Metadata = append(numbers.Metadata, numMetadata)
		}
	}
	return symbols, numbers
}
