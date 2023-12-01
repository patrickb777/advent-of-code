package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Coordinates struct {
	Input []string
	Coord []int
}

func main() {
	fmt.Println("Advent of Code Day 1 - Coordinates pt 1")
	f := flag.String("f", "none", "Input file")
	flag.Parse()
	coordinates := readFile(*f)
	coord := getCoord(coordinates)
	total := 0
	for _, c := range coord.Coord {
		total = total + c
	}
	fmt.Println(total)
}

func readFile(file string) Coordinates {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	// Read each line of the file via a scanner and to the input slice of the Coordinates struct
	c := Coordinates{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		c.Input = append(c.Input, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return c
}

func getCoord(c Coordinates) Coordinates {
	re := regexp.MustCompile("[0-9]")
	for _, v := range c.Input {
		s := re.FindAllString(v, -1)
		coord, err := strconv.Atoi(s[0] + s[len(s)-1])
		if err != nil {
			log.Fatal(err)
		}
		c.Coord = append(c.Coord, coord)
	}
	return c
}
