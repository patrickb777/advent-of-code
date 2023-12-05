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

type Winnings struct {
	Cards []string
}

func main() {
	start := time.Now()
	fmt.Println("Scratchcards")
	f := flag.String("f", "none", "Input file")
	flag.Parse()
	inputFile := readfile.ReadFile(*f) // moved the reading of the input file to an external module

	winnings := Winnings{}
	for _, row := range inputFile.InputRow {
		winnings.Cards = append(winnings.Cards, row)
	}

	l := len(winnings.Cards)
	rgx := regexp.MustCompile(`[0-9]+`)
	for i := 0; i < l; i++ {
		// fmt.Println("__ i = ", i, "l =", l)
		wins := parseScratchcard(winnings.Cards[i])
		//fmt.Println(winnings)
		//fmt.Printf("Playing card:%v, wins: %d\n", winnings.Cards[i], wins)
		c, err := strconv.Atoi(rgx.FindString(winnings.Cards[i]))
		if err != nil {
			log.Fatal(err)
		}
		for ind := 1; ind <= wins; ind++ {
			if wins != 0 {
				//fmt.Println(winnings.Cards[i+ind])
				winnings.Cards = append(winnings.Cards, inputFile.InputRow[(c-1)+ind])
				//fmt.Println("++")
			}
		}
		l = len(winnings.Cards)

		// for _, k := range winnings.Cards {
		// 	fmt.Println(k)
		// }
		// fmt.Println("__")
	}

	// Output result
	totalCards := len(winnings.Cards)
	fmt.Printf("Total scratchcards: %d\n", totalCards)

	// Output execution time
	elapsed := time.Since(start)
	log.Printf("Execution time %s\n", elapsed)
}

func parseScratchcard(sc string) int {
	rgx := regexp.MustCompile(`[0-9]+`)
	card := strings.Split(strings.SplitAfter(sc, ": ")[1], " | ")
	winNums := convNumbers(rgx.FindAllString(card[0], -1))
	cardNums := convNumbers(rgx.FindAllString(card[1], -1))
	wins := 0
	for _, v := range cardNums {
		if slices.Contains(winNums, v) {
			if wins == 0 {
				wins = 1
			} else {
				wins = wins + 1
			}
		}
	}
	return wins
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
