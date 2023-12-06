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
	"sync"
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
	log.Printf("Commencing calculations at: %v", start)
	f := flag.String("f", "none", "Input file")
	flag.Parse()
	inputFile := readfile.ReadFile(*f)

	// Processing
	locations := Locations{}
	seeds, maps := parseAlmanac(inputFile)
	expSeeds := expSeedList(seeds.Seeds)
	seedsTotal := len(expSeeds)
	log.Printf("Total Number of Seeds: %d", seedsTotal)
	u := 1000000
	for i, s := range expSeeds {
		loc := getLocation(s, maps)
		if len(locations.Loc) == 0 {
			// Store first location
			locations.Loc = append(locations.Loc, loc)
		} else if loc < locations.Loc[0] {
			locations.Loc[0] = loc
			log.Printf("New lowest location found: seed: %d @ location:  %d!", s, locations.Loc[0])
		}
		if i == u {
			log.Printf("Processed %d %% of %d seeds", (i / seedsTotal * 100), seedsTotal)
			u = u + 1000000
		}
	}
	loc := getNearestLoc(locations.Loc)
	fmt.Println("Closet location: ", loc)

	// Output execution time
	elapsed := time.Since(start)
	log.Printf("Execution time %s\n", elapsed)
}

func expSeedList(seeds []int) []int {
	var wg sync.WaitGroup
	totalSeeds := len(seeds)
	wg.Add(16)
	var seedList []int
	t := 0
	for i := 0; i < totalSeeds; i++ {
		go func(i int) {
			defer wg.Done()
			t = seeds[i]
			start := seeds[t]               // start the loop at the seed number
			end := start + (seeds[t+1] - 1) // end the loop at the end of the range
			for x := start; x <= end; x++ {
				seedList = append(seedList, x)
			}
			fmt.Println(start, end, t)
			t++
		}(i)
	}

	wg.Wait()
	//log.Println(seedList)
	seedList = append(seedList, 1)
	return seedList
}

func getNearestLoc(locs []int) int {
	loc := slices.Min(locs)
	return loc
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

	return loc
}

func mapLookup(source int, m string, maps Maps) int {
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
