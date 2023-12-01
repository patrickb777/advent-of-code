package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

type Calibrations struct {
	Input          []string
	CalibrationNum []int
}

func main() {
	start := time.Now()
	fmt.Println("Trebuchet")
	f := flag.String("f", "none", "Input file")
	flag.Parse()
	Calibrations := readFile(*f)
	calibrations := getCalibrationNum(Calibrations)
	total := 0
	for _, val := range calibrations.CalibrationNum {
		total = total + val
	}
	fmt.Printf("Result: %d\n", total)
	elapsed := time.Since(start)
	log.Printf("Execution time %s\n", elapsed)
}

func readFile(file string) Calibrations {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	// Read each line of the file via a scanner and to the input slice of the Calibrations struct
	c := Calibrations{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		c.Input = append(c.Input, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return c
}

func getCalibrationNum(c Calibrations) Calibrations {
	re := regexp.MustCompile("[0-9]")
	for _, input := range c.Input {
		result := re.FindAllString(input, -1)
		val, err := strconv.Atoi(result[0] + result[len(result)-1])
		if err != nil {
			log.Fatal(err)
		}
		c.CalibrationNum = append(c.CalibrationNum, val)
	}
	return c
}
