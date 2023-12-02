package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Cube Conundrum")

	// Output execution time
	elapsed := time.Since(start)
	log.Printf("Execution time %s\n", elapsed)
}
