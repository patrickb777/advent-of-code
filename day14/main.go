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
	fmt.Printf("\n[The Parabolic Refelctor Dish]\n\n")
	f := flag.String("f", "none", "Input file")
	flag.Parse()

	// Parse Input
	inputFile := readfile.ReadFile(*f)
	fmt.Sprintln(inputFile) //Does not output to stdout
	for _, i := range inputFile.InputRow {
		fmt.Println(i)
	}
	//input := parseData(inputFile)

	// Output execution time
	elapsed := time.Since(start)
	fmt.Printf("\n")
	log.Printf("Execution time %s\n", elapsed)
}

func parseData(input readfile.InputFile) {

}
