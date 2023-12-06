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
	Time     []int
	Distance []int
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
	cnt := len(races.Time)

	// Processing
	wg.Add(cnt) // Specify number of go routines that will need to wait until finished
	for i := 0; i < (cnt); i++ {
		go calculateDistance(races, i)
	}

	wg.Wait() // Wait for go routines to finish
	// Output execution time
	elapsed := time.Since(start)
	log.Printf("Execution time %s\n", elapsed)
}

func calculateDistance(races Races, i int) {
	defer wg.Done()
	winningCombinations := 0
	raceTime := races.Time[i]
	for timer := 1; timer < raceTime; timer++ {
		result := (raceTime - timer) * timer
		if result > races.Distance[i] {
			winningCombinations++
			//fmt.Printf("Wiining race distance: (RaceTime %d - Timer %d)*Timer %d = %d \n", raceTime, timer, timer, result)
		}
	}
	fmt.Println(winningCombinations)
}

func parseRaces(input readfile.InputFile) Races {
	races := Races{}
	rgx := regexp.MustCompile(`[0-9]+`)
	races.Time = convNum(rgx.FindAllString(input.InputRow[0], -1))
	races.Distance = convNum(rgx.FindAllString(input.InputRow[1], -1))
	return races
}

func convNum(in []string) []int {
	var out []int
	for _, v := range in {
		n, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		out = append(out, n)
	}
	return out
}
