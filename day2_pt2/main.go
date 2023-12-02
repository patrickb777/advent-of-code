package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type InputFile struct {
	InputRow []string
}

type Game struct {
	GameID     []int
	CubeSet    [][]Cubes
	MinCubeSet []Cubes
	MinPower   []int
	Possible   []bool
}

type Cubes struct {
	Blue  int
	Green int
	Red   int
}

const BlueCubes = 14
const GreenCubes = 13
const RedCubes = 12

func main() {
	start := time.Now()
	fmt.Println("Cube Conundrum")
	f := flag.String("f", "none", "Input file")
	flag.Parse()
	InputFile := readFile(*f)
	games := getGameStats(InputFile)
	total := sumOfPowers(games)
	fmt.Printf("Sum of Powers = %d\n", total)

	// Output execution time
	elapsed := time.Since(start)
	log.Printf("Execution time %s\n", elapsed)
}

func sumOfPowers(g Game) int {
	total := 0
	for i, _ := range g.GameID {
		total = total + g.MinPower[i]
	}
	return total
}

func getGameStats(inputs InputFile) Game {
	game := Game{}
	regex := regexp.MustCompile(`\d+`)
	for _, row := range inputs.InputRow {

		// Extract game number
		g := strings.Split(row, ":")
		result := regex.FindAllString(g[0], -1)
		id, err := strconv.Atoi(result[0])
		if err != nil {
			log.Fatal(err)
		}

		// Extract game stats
		revealCubes := strings.Split(g[1], ";")
		c := []Cubes{}
		p := true
		m := Cubes{Blue: 0, Green: 0, Red: 0}
		for _, reveal := range revealCubes {
			cubes := strings.Split(reveal, ",")
			blues, greens, reds := getColours(cubes)
			c = append(c, Cubes{Blue: blues, Green: greens, Red: reds})
			if blues > BlueCubes || greens > GreenCubes || reds > RedCubes {
				p = false
			}

			// Calculate minimum cube set
			if blues != 0 {
				if m.Blue == 0 || blues > m.Blue {
					m.Blue = blues
				}
			}
			if greens != 0 {
				if m.Green == 0 || greens > m.Green {
					m.Green = greens
				}
			}
			if reds != 0 {
				if m.Red == 0 || reds > m.Red {
					m.Red = reds
				}
			}
		}

		// Finalise game stats
		game.GameID = append(game.GameID, id)
		game.CubeSet = append(game.CubeSet, c)
		game.MinCubeSet = append(game.MinCubeSet, m)
		game.MinPower = append(game.MinPower, m.Blue*m.Green*m.Red)
		game.Possible = append(game.Possible, p)
	}
	return game
}

func getColours(cubes []string) (int, int, int) {
	colourRegex := regexp.MustCompile(`[a-z]+`)
	numberRegex := regexp.MustCompile(`[0-9]+`)
	blues := 0
	greens := 0
	reds := 0
	for _, colours := range cubes {
		colourName := colourRegex.FindAllString(colours, -1)[0]
		cubeCount, err := strconv.Atoi(numberRegex.FindAllString(colours, -1)[0])
		if err != nil {
			log.Fatal(err)
		}
		switch colourName {
		case "blue":
			blues = cubeCount
		case "green":
			greens = cubeCount
		case "red":
			reds = cubeCount
		}
	}
	return blues, greens, reds
}

// Need to move this to an external module for re-use in future puzzles
func readFile(file string) InputFile {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	// Scan each line of input file with scanner and append to InputFile struct
	input := InputFile{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input.InputRow = append(input.InputRow, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return input
}
