package main

import (
	"advent-of-code/readfile"
	"flag"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Seeds struct {
	Seeds []int
}

type Maps struct {
	Mapdata []Mapdata
}
type Mapdata struct {
	Map    string
	Dest   int
	Source int
	Range  int
}

func main() {
	start := time.Now()
	fmt.Println("The Gardener & The Almanac")
	f := flag.String("f", "none", "Input file")
	flag.Parse()
	inputFile := readfile.ReadFile(*f)

	// Processing
	loc := 0
	lowestLoc := 0
	seeds, maps := parseAlmanac(inputFile)
	for i, s := range seeds.Seeds {
		loc = getLocation(s, maps)
		//fmt.Println("Returned location:", loc)
		if i == 0 {
			lowestLoc = loc
		} else if loc < lowestLoc {
			lowestLoc = loc
		}
	}
	fmt.Printf("Lowest location: %d \n", lowestLoc)
	// //fmt.Println(locations)
	// fmt.Println("Closet location: ", loc)

	// Output execution time
	elapsed := time.Since(start)
	log.Printf("Execution time %s\n", elapsed)
	fmt.Sprintln(seeds, maps)
}

func getLocation(seed int, maps Maps) int {
	loc := 0
	soil := mapLookup(seed, "seed-to-soil", maps)
	fert := mapLookup(soil, "soil-to-fertilizer", maps)
	water := mapLookup(fert, "fertilizer-to-water", maps)
	light := mapLookup(water, "water-to-light", maps)
	temp := mapLookup(light, "light-to-temperature", maps)
	humi := mapLookup(temp, "temperature-to-humidity", maps)
	loc = mapLookup(humi, "humidity-to-location", maps)
	//log.Println(seed, soil, fert, water, light, temp, humi, loc)
	return loc
}

func mapLookup(source int, m string, maps Maps) int {
	//lookup := make(map[int]int)
	dest := 404
	for _, md := range maps.Mapdata {
		switch md.Map {
		case m:
			sourceStart := md.Source
			sourceEnd := md.Source + (md.Range - 1)
			if source >= sourceStart && source <= sourceEnd {
				delta := source - md.Source
				dest = md.Dest + delta
			}
		}
	}
	if dest == 404 {
		dest = source
	}
	return dest
}

func parseAlmanac(almanac readfile.InputFile) (Seeds, Maps) {
	seeds := Seeds{}
	maps := Maps{}
	md := Mapdata{}
	rgx := regexp.MustCompile(`[0-9]+`)
	m := ""

	for i, v := range almanac.InputRow {
		// Extract seed listß
		if i == 0 {
			seedStr := strings.Split(v, "seeds: ")[1]
			seedList := rgx.FindAllString(seedStr, -1)
			seedListConv := convNumbers(seedList)
			seeds.Seeds = append(seeds.Seeds, seedListConv...)
		} else {
			switch v {
			case "seed-to-soil map:":
				m = "seed-to-soil"
			case "soil-to-fertilizer map:":
				m = "soil-to-fertilizer"
			case "fertilizer-to-water map:":
				m = "fertilizer-to-water"
			case "water-to-light map:":
				m = "water-to-light"
			case "light-to-temperature map:":
				m = "light-to-temperature"
			case "temperature-to-humidity map:":
				m = "temperature-to-humidity"
			case "humidity-to-location map:":
				m = "humidity-to-location"
			}
			row := convNumbers(rgx.FindAllString(v, -1))
			if row != nil {
				md = Mapdata{Map: m, Dest: row[0], Source: row[1], Range: row[2]}
				maps.Mapdata = append(maps.Mapdata, md)
			}
		}
	}
	return seeds, maps
}

func convNumbers(str []string) []int {
	var numList []int
	for _, v := range str {
		num, err := strconv.Atoi(v)
		if err != nil {
			log.Print(err)
			return numList
		}
		numList = append(numList, num)
	}
	return numList
}
