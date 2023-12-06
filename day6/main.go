package main

import (
	"advent-of-code/readfile"
	"flag"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Boat Races")
	f := flag.String("f", "none", "Input file")
	flag.Parse()
	inputFile := readfile.ReadFile(*f)
	fmt.Println(inputFile)

	// Output execution time
	elapsed := time.Since(start)
	log.Printf("Execution time %s\n", elapsed)
}
