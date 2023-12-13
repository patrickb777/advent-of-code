package main

import (
	"advent-of-code/readfile"
	"flag"
	"fmt"
	"log"
	"math"
	"time"
)

type Galaxy struct {
	Id     int
	Xcoord int
	Ycoord int
}
type Distance struct {
	Source      int
	Destination int
	Distance    float64
}

func main() {
	// Init
	start := time.Now()
	fmt.Printf("\n.   *  .  .     .  * ..   [Cosmic Expansion] .        .        .   *    ..\n")
	f := flag.String("f", "none", "Input file")
	flag.Parse()

	// Parse Input
	inputFile := readfile.ReadFile(*f)
	fmt.Sprintln(inputFile) //Does not output to stdout
	universe := parseData(inputFile)
	// for i := range universe {
	// 	fmt.Println(universe[i])
	// }

	// Processing

	//Find Galaxies
	galaxies := []Galaxy{}
	id := 0
	yExpPoints := []int{}
	for y := range universe {
		yExpansion := true
		for x := range universe {
			if universe[y][x] == 35 {
				//fmt.Println("Found Galaxy @", x, y)
				galaxies = append(galaxies, Galaxy{Id: id, Xcoord: x, Ycoord: y})
				id++
				yExpansion = false
			}
		}
		if yExpansion == true {
			yExpPoints = append(yExpPoints, y)
		}
	}
	//fmt.Println(galaxies)

	xExpPoints := []int{}
	for x := range universe {
		xExpansion := true
		for y := range universe {
			if universe[y][x] == 35 {
				xExpansion = false
			}
		}
		if xExpansion == true {
			xExpPoints = append(xExpPoints, x)
		}
	}
	//fmt.Println(xExpPoints)
	//fmt.Println(yExpPoints)

	//Shift Galaxy Positions
	shift := 999999
	// X Axis
	for g, gal := range galaxies {
		cnt := 1
		adjXcoord := 0
		shiftFlag := 0
		for _, ep := range xExpPoints {
			if gal.Xcoord > ep {
				//fmt.Printf("Galaxy: %d, X: %d, Exp Point: %d \n", gal.Id, gal.Xcoord, ep)
				adjXcoord = gal.Xcoord + (shift * cnt)
				//fmt.Printf("Galaxy: %d, X: %d, adjusted X: %d \n", gal.Id, gal.Xcoord, adjXcoord)
				cnt++
				shiftFlag = 1
			}
		}
		if shiftFlag == 1 {
			galaxies[g].Xcoord = adjXcoord
		}
	}
	//fmt.Println(galaxies)
	// Y Axis
	for g, gal := range galaxies {
		cnt := 1
		adjYcoord := 0
		shiftFlag := 0
		for _, ep := range yExpPoints {
			if gal.Ycoord > ep {
				//fmt.Printf("Galaxy: %d, Y: %d, Exp Point: %d \n", gal.Id, gal.Ycoord, ep)
				adjYcoord = gal.Ycoord + (shift * cnt)
				//fmt.Printf("Galaxy: %d, Y: %d, adjusted Y: %d \n", gal.Id, gal.Ycoord, adjYcoord)
				cnt++
				shiftFlag = 1
			}
		}
		if shiftFlag == 1 {
			galaxies[g].Ycoord = adjYcoord
		}
	}
	//fmt.Println(galaxies)

	//Calculate distances
	distances := []Distance{}
	totalDistance := 0
	for _, srcGal := range galaxies {
		for destGal := srcGal.Id; destGal < len(galaxies); destGal++ {
			if srcGal.Id != destGal {
				stepsY := math.Abs(float64(galaxies[destGal].Ycoord - srcGal.Ycoord))
				stepsX := math.Abs(float64(galaxies[destGal].Xcoord - srcGal.Xcoord))
				distances = append(distances, Distance{Source: srcGal.Id, Destination: destGal, Distance: stepsY + stepsX})
				totalDistance = totalDistance + int(stepsX+stepsY)
			}

		}
	}
	fmt.Printf("\nTotal galaxy distances: %d\n", totalDistance)

	// Output execution time
	elapsed := time.Since(start)
	fmt.Printf("\n")
	log.Printf("Execution time %s\n", elapsed)
}

func parseData(input readfile.InputFile) [][]byte {
	var universe [][]byte

	for _, row := range input.InputRow {
		var dataRow []byte
		for i := 0; i < len(row); i++ {
			dataPoint := row[i]
			dataRow = append(dataRow, dataPoint)
		}
		universe = append(universe, dataRow)
	}
	return universe
}
