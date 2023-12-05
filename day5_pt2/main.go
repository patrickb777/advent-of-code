package main

import (
	"advent-of-code/readfile"
	"flag"
	"fmt"
	"log"
	"regexp"
	"slices"
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
	fmt.Println("The Gardener & The Almanac - PT2")
	f := flag.String("f", "none", "Input file")
	flag.Parse()
	inputFile := readfile.ReadFile(*f)

	// Processing
	locations := Locations{}
	seeds, maps := parseAlmanac(inputFile)
	for _, s := range seeds.Seeds {
		locations.Loc = append(locations.Loc, getLocation(s, maps))
	}
	//fmt.Println(locations)
	loc := getNearestLoc(locations.Loc)
	fmt.Println("Closet location: ", loc)

	// Output execution time
	elapsed := time.Since(start)
	log.Printf("Execution time %s\n", elapsed)
}

func expSeedList(seeds []int) []int {
	var seedList []int
	for i := 0; i < len(seeds); i++ {
		//log.Printf("Seed: %d, Range: %d", seeds[i], seeds[i+1])
		start := seeds[i]
		end := start + (seeds[i+1] - 1)
		fmt.Println(start, end)
		for x := start; x <= end; x++ {
			//fmt.Printf("x= %d, ", x)
			seedList = append(seedList, x)
		}
		i++
	}
	log.Println("\nSeed List: ", seedList)
	return seedList
}

func getNearestLoc(locs []int) int {
	loc := slices.Min(locs)
	return loc
}

func getLocation(seed int, maps Maps) int {
	loc := 0
	//log.Printf("Seed: %d", seed)
	soil := mapLookup(seed, "seed-to-soil", maps)
	//log.Printf("Soil: %d", soil)
	fert := mapLookup(soil, "soil-to-fertilizer", maps)
	//log.Printf("Fertilizer: %d", fert)
	water := mapLookup(fert, "fertilizer-to-water", maps)
	light := mapLookup(water, "water-to-light", maps)
	temp := mapLookup(light, "light-to-temperature", maps)
	humi := mapLookup(temp, "temperature-to-humidity", maps)
	loc = mapLookup(humi, "humidity-to-location", maps)
	//log.Printf("Location: %d\n\n", loc)
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
