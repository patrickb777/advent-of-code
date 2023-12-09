package main

import (
	"advent-of-code/readfile"
	"flag"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"time"
)

func main() {
	// Init
	start := time.Now()
	fmt.Println("mirage maintenance")
	f := flag.String("f", "none", "Input file")
	flag.Parse()

	// Parse Input
	inputFile := readfile.ReadFile(*f)
	fmt.Sprintln(inputFile) //Does not output to stdout

	// Processing
	inputs := parseData(inputFile)
	//fmt.Printf("%v\n", inputs)
	var forecast int64
	cnt := 0
	for _, row := range inputs {
		//log.Println(row)
		forecast = forecast + analyseData(row)
		cnt++
	}
	fmt.Printf("\nResult: %d.  Data Points Forecased: %d\n", forecast, cnt)
	// Output execution time
	elapsed := time.Since(start)
	log.Printf("Execution time %s\n", elapsed)
}

func analyseData(data []int64) int64 {
	var forecast [][]int64
	var t int
	forecast = append(forecast, data)
	for i := 0; i < len(forecast); i++ {
		fmt.Print("\n", forecast[i])
		tmp := []int64{}
		t = len(forecast[i]) - 1
		for x := 0; x < t; x++ {
			tmp = append(tmp, forecast[i][x+1]-forecast[i][x])
		}
		forecast = append(forecast, tmp)
		zeroCnt := 0
		target := len(forecast[i])
		for _, x := range forecast[i] {
			if x == 0 {
				zeroCnt++
			}
		}
		if target == zeroCnt {
			break
		}
	}
	y := len(forecast) - 1

	// Forewards Forecast
	forecast[y] = append(forecast[y], 0)
	for i := y - 1; i >= 0; i-- {
		dA := len(forecast[i]) - 1
		dB := len(forecast[i+1]) - 1
		dP := forecast[i][dA] + forecast[i+1][dB]
		forecast[i] = append(forecast[i], dP)
	}

	// Backwards Forecast
	forecast[y] = append([]int64{0}, forecast[y]...)
	for i := y - 1; i >= 0; i-- {
		dA := 0
		dB := 0
		dP := forecast[i][dA] - forecast[i+1][dB]
		forecast[i] = append([]int64{dP}, forecast[i]...)
	}

	//f := len(forecast[0]) - 1
	return forecast[0][0]
}

func parseData(input readfile.InputFile) [][]int64 {
	var inputFile [][]int64
	rgx := regexp.MustCompile(`-?[0-9]+`)
	for _, row := range input.InputRow {
		rNum := []int64{}
		rStr := rgx.FindAllString(row, -1)
		for i := 0; i < len(rStr); i++ {
			num := convNum(rStr[i])
			rNum = append(rNum, num)
		}
		inputFile = append(inputFile, rNum)
	}
	return inputFile
}

func convNum(in string) int64 {
	num, err := strconv.Atoi(in)
	out := int64(num)
	if err != nil {
		log.Fatal(err)
	}
	return out
}
