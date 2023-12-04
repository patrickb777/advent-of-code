package main

import (
	"advent-of-code/readfile"
	"flag"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"time"
)

type PartNumbers struct {
	PartNo []int
}

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
}

type Symbols struct {
	Metadata []SymMetadata
}

type SymMetadata struct {
	Char string
	Row  int
	Pos  int
}

func main() {
	start := time.Now()
	fmt.Println("Gear Ratios")
	f := flag.String("f", "none", "Input file")
	flag.Parse()
	inputFile := readfile.ReadFile(*f) // moved the reading of the input file to an external module
	symbols, numbers := parseSchematic(inputFile)
	partNos := getPartNumbers(symbols, numbers)
	total := getSumOfParts(partNos)
	fmt.Printf("Sum of Part Numbers = %d\n", total)

	// Output execution time
	elapsed := time.Since(start)
	log.Printf("Execution time %s\n", elapsed)
}

func getSumOfParts(partNos PartNumbers) int {
	total := 0
	for _, n := range partNos.PartNo {
		total = total + n
	}
	return total
}

func parseSchematic(input readfile.InputFile) (Symbols, Numbers) {
	var symMetadata SymMetadata
	var numMetadata NumMetadata
	symbols := Symbols{}
	numbers := Numbers{}
	symRX := regexp.MustCompile(`[^a-zA-Z0-9\.]`)
	numRX := regexp.MustCompile(`[0-9]+`)

	for r, v := range input.InputRow {
		runes := []rune(v)
		for p, c := range runes {
			sym := symRX.FindAllString(string(c), -1)
			if len(sym) > 0 {
				symMetadata.Char = sym[0]
				symMetadata.Row = r
				symMetadata.Pos = p
				symbols.Metadata = append(symbols.Metadata, symMetadata)
			}
		}
		num := numRX.FindAllString(v, -1)
		for _, n := range num {
			numMetadata.Num = n
			numMetadata.Len = len(n)
			//numMetadata.Pos = strings.Index(v, n) // this doesn't work
			nRX := regexp.MustCompile(n)
			posIndex := nRX.FindAllStringSubmatchIndex(v, -1)
			//fmt.Printf("Number:%s, Regex Match:%s, Regex Index:%v, String Index:%d   \n", numMetadata.Num, test, test2, numMetadata.Pos)
			for _, ind := range posIndex {
				str := v[ind[0]:]
				res := numRX.FindString(str)
				//fmt.Printf("%s  \n", test3)
				//fmt.Printf("%s  \n", test4)
				if res == numMetadata.Num {
					//fmt.Printf("Match @ %d <<>> Pos %d \n", val[0], numMetadata.Pos)
					numMetadata.Pos = ind[0]
				} // else {
				// 	fmt.Println("No Match")
				// }
			}
			numMetadata.Row = r
			numMetadata.ProxX[0] = numMetadata.Pos - 1
			numMetadata.ProxX[1] = numMetadata.Pos + numMetadata.Len
			numMetadata.ProxY[0] = r - 1
			numMetadata.ProxY[1] = r + 1
			numbers.Metadata = append(numbers.Metadata, numMetadata)
		}
	}
	return symbols, numbers
}

func getPartNumbers(s Symbols, n Numbers) PartNumbers {
	var partNos PartNumbers
	for _, nmd := range n.Metadata {
		//fmt.Printf("Number:%s:, Length:%d, Postion:%d, X Range:%d:%d, Y Range:%d:%d \n", nmd.Num, nmd.Len, nmd.Pos, nmd.ProxX[0], nmd.ProxX[1], nmd.ProxY[0], nmd.ProxY[1])
		for _, smd := range s.Metadata {
			if smd.Row >= nmd.ProxY[0] && smd.Row <= nmd.ProxY[1] && smd.Pos >= nmd.ProxX[0] && smd.Pos <= nmd.ProxX[1] {
				//log.Printf("R:%d, N:%s, Proximity of S:%s \n", nmd.Row, nmd.Num, smd.Char)
				p, err := strconv.Atoi(nmd.Num)
				if err != nil {
					log.Fatal(err)
				}
				partNos.PartNo = append(partNos.PartNo, p)

			}
		}
	}
	return partNos
}
