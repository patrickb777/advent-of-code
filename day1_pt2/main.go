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
	fmt.Println("Trebuchet")
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
	re := regexp.MustCompile("(?:zero|one|two|three|four|five|six|seven|eight|nine|[0-9])") // overlapping strings are causing a problem, Golang doesn't support look aheads (?=())
	nConv := map[string]string{
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
	for _, v := range c.Input {
		s := re.FindAllString(v, -1)
		for i, n := range s {
			mv, e := nConv[n]
			if e {
				s[i] = mv
			}
		}
		coord, err := strconv.Atoi(s[0] + s[len(s)-1])
		if err != nil {
			log.Fatal(err)
		}
		c.Coord = append(c.Coord, coord)
	}
	for i, _ := range c.Coord {
		log.Println(c.Input[i], " : ", c.Coord[i])
	}

	return c
}
