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
	GameID   []int
	CubeSet  [][]Cubes
	Possible []bool
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
	for i, _ := range games.GameID {
		fmt.Printf("Game Number: %d, Cube Sets: %v, Possible: %v\n", games.GameID[i], games.CubeSet[i], games.Possible[i])
	}
	total := sumOfPossibles(games)
	fmt.Printf("Sum of Possibles = %d\n", total)

	// Output execution time
	elapsed := time.Since(start)
	log.Printf("Execution time %s\n", elapsed)
}

func sumOfPossibles(g Game) int {
	total := 0
	for i, _ := range g.GameID {
		if g.Possible[i] == true {
			total = total + g.GameID[i]
		}
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
		for _, reveal := range revealCubes {
			cubes := strings.Split(reveal, ",")
			blues, greens, reds := getColours(cubes)
			c = append(c, Cubes{Blue: blues, Green: greens, Red: reds})
			if blues > BlueCubes || greens > GreenCubes || reds > RedCubes {
				p = false
			}
		}
		game.GameID = append(game.GameID, id)
		game.CubeSet = append(game.CubeSet, c)
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

// Move this to an external module
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
