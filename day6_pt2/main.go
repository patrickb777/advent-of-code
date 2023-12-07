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
	for timer := 1; timer < races.Time; timer++ {
		result := (races.Time - timer) * timer
		if result > races.Distance {
			winningCombinations++
			//fmt.Printf("Wiining race distance: (RaceTime %d - Timer %d)*Timer %d = %d \n", raceTime, timer, timer, result)
		}
	}
	fmt.Println(winningCombinations)
}

func parseRaces(input readfile.InputFile) Races {
	races := Races{}
	races.Time = convNumber(input.InputRow[0])
	races.Distance = convNumber(input.InputRow[1])
	return races
}

func convNumber(str string) int {
	rgx := regexp.MustCompile(`[0-9]+`)
	numbers := rgx.FindAllString(str, -1)
	number := ""
	for _, n := range numbers {
		number = number + n
	}
	num, err := strconv.Atoi(number)
	if err != nil {
		log.Fatal(err)
	}
	return num
}
