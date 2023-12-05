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

type Locations struct {
	Loc []int
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
	inputFile := readfile.ReadFile(*f) // moved the reading of the input file to an external module

	// Processing
	locations := Locations{}
	seeds, maps := parseAlmanac(inputFile)
	for _, s := range seeds.Seeds {
		locations.Loc = append(locations.Loc, getLocation(s, maps))
	}
	fmt.Println(locations)

	// Output execution time
	elapsed := time.Since(start)
	log.Printf("Execution time %s\n", elapsed)
}

func getLocation(seed int, maps Maps) int {
	loc := 0
	// Get Soil
	soil := mapLookup(seed, "seed-to-soil", maps)
	fert := mapLookup(soil, "soil-to-fertilizer", maps)
	water := mapLookup(fert, "fertilizer-to-water", maps)
	light := mapLookup(water, "water-to-light", maps)
	temp := mapLookup(light, "light-to-temperature", maps)
	humi := mapLookup(temp, "temperature-to-humidity", maps)
	loc = mapLookup(humi, "humidity-to-location", maps)
	//fmt.Println(seed, soil, fert, water, light, temp, humi, loc)
	return loc
}

func mapLookup(source int, m string, maps Maps) int {
	lookup := make(map[int]int)
	dest := 404
	for _, md := range maps.Mapdata {
		switch md.Map {
		case m:
			// create temporary map for the source and detination range
			cnt := 0
			for s := md.Source; s <= md.Source+(md.Range-1); s++ {
				//fmt.Print(m, ":", s, "..", md.Dest+cnt, " || ")
				lookup[s] = md.Dest + cnt
				cnt++
				_, exists := lookup[source]
				if exists {
					dest = lookup[source]
				} else {
					dest = source
				}
			}
		}
	}
	// log.Println(lookup)
	// log.Println("Destination:", dest)
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
