package main

import (
	"advent-of-code/readfile"
	"flag"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"sync"
	"time"
)

type Races struct {
	Time     int
	Distance int
}

var wg sync.WaitGroup

func main() {

	start := time.Now()
	fmt.Println("Boat Races")
	f := flag.String("f", "none", "Input file")
	flag.Parse()

	// Parse Input
	inputFile := readfile.ReadFile(*f)
	// fmt.Println(inputFile)
	races := parseRaces(inputFile)

	// Processing

	calculateDistance(races)

	// Output execution time
	elapsed := time.Since(start)
	log.Printf("Execution time %s\n", elapsed)
}

func calculateDistance(races Races) {
	winningCombinations := 0
	raceTime := races.Time
	for timer := 1; timer < raceTime; timer++ {
		result := (raceTime - timer) * timer
		if result > races.Distance {
			winningCombinations++
			//fmt.Printf("Wiining race distance: (RaceTime %d - Timer %d)*Timer %d = %d \n", raceTime, timer, timer, result)
		}
	}
	fmt.Println(winningCombinations)
}

func parseRaces(input readfile.InputFile) Races {
	races := Races{}
	rgx := regexp.MustCompile(`[0-9]+`)
	tin := rgx.FindAllString(input.InputRow[0], -1)
	t := ""
	for _, n := range tin {
		t = t + n
	}
	time, err := strconv.Atoi(t)
	if err != nil {
		log.Fatal(err)
	}
	races.Time = time

	din := rgx.FindAllString(input.InputRow[1], -1)
	d := ""
	for _, n := range din {
		d = d + n
	}
	dist, err := strconv.Atoi(d)
	if err != nil {
		log.Fatal(err)
	}
	races.Distance = dist

	return races
}
