package main

import (
	"advent-of-code/readfile"
	"flag"
	"fmt"
	"log"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Scratchcards")
	f := flag.String("f", "none", "Input file")
	flag.Parse()
	inputFile := readfile.ReadFile(*f) // moved the reading of the input file to an external module
	totalWorth := 0
	for _, v := range inputFile.InputRow {
		totalWorth = totalWorth + parseScratchcard(v)
	}
	fmt.Printf("Total scratchcard worth: %d\n", totalWorth)

	// Output execution time
	elapsed := time.Since(start)
	log.Printf("Execution time %s\n", elapsed)
}

func parseScratchcard(sc string) int {
	rgx := regexp.MustCompile(`[0-9]+`)
	card := strings.Split(strings.SplitAfter(sc, ": ")[1], " | ")
	winNums := convNumbers(rgx.FindAllString(card[0], -1))
	cardNums := convNumbers(rgx.FindAllString(card[1], -1))
	worth := 0
	for _, v := range cardNums {
		if slices.Contains(winNums, v) {
			if worth == 0 {
				worth = 1
			} else {
				worth = worth * 2
			}
		}
	}
	return worth
}

func convNumbers(str []string) []int {
	var numList []int
	for _, v := range str {
		num, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		numList = append(numList, num)
	}
	return numList
}
