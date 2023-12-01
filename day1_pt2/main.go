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
	fmt.Println("Trebuchet pt2")
	f := flag.String("f", "none", "Input file")
	flag.Parse()
	Calibrations := readFile(*f)
	CalibrationNum := getCalibrationNum(Calibrations)
	total := 0
	for _, val := range CalibrationNum.CalibrationNum {
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
	// Overlapping strings are causing a problem with my regex approach, Golang doesn't support look aheads (?=()) that would identiy these
	// The following cleans the inputs of overlapping strings, it's pretty dirty and honestly a hack to solve the puzzle
	regex := regexp.MustCompile("(?:zerone|oneight|twone|threeight|fiveight|eightwo|eighthree|nineight)")
	overlaps := map[string]string{
		"zerone":    "zeroone",
		"oneight":   "oneeight",
		"twone":     "twoone",
		"threeight": "threeight",
		"fiveight":  "fiveeight",
		"eightwo":   "eighttwo",
		"eighthree": "eightthree",
		"nineight":  "nineeight",
	}
	for i, input := range c.Input {
		stringOverlap := regex.FindAllString(input, -1)
		if len(stringOverlap) > 0 {
			for _, key := range stringOverlap {
				overlapCorrection, exists := overlaps[key]
				if exists {
					c.Input[i] = regex.ReplaceAllString(c.Input[i], overlapCorrection)
				}
			}
		}
	}
	// iterate over the cleansed inputs and convert the spelt out numbers to digits
	regex = regexp.MustCompile("(?:zero|one|two|three|four|five|six|seven|eight|nine|[0-9])")
	numConv := map[string]string{
		"zero":  "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	for _, input := range c.Input {
		result := regex.FindAllString(input, -1)
		for i, key := range result {
			num, exists := numConv[key]
			if exists {
				result[i] = num
			}
		}
		CalibrationNum, err := strconv.Atoi(result[0] + result[len(result)-1])
		if err != nil {
			log.Fatal(err)
		}
		c.CalibrationNum = append(c.CalibrationNum, CalibrationNum)
	}
	return c
}
