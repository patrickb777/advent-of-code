package main

import (
	"advent-of-code/readfile"
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type Game struct {
	Hand []Hands
}

type Hands struct {
	Hand     string
	Bid      int
	Strength int
}

type Rankings struct {
	Hand []string
}

func main() {
	start := time.Now()
	fmt.Println("[♥]]] [♦]]] [♣]]] [♠]]]")
	f := flag.String("f", "none", "Input file")
	flag.Parse()

	// Parse Input
	inputFile := readfile.ReadFile(*f)
	//fmt.Println(inputFile)
	game := parseCards(inputFile)

	// Processing

	// fmt.Printf("Pre-Sort order:\n")
	// for i := range game.Hand {
	// 	fmt.Println(game.Hand[i].Hand, game.Hand[i].Strength)
	// }

	for {
		sortFlag := 0
		for i := range game.Hand {
			swap := Hands{}
			if i != len(game.Hand)-1 {
				swap = game.Hand[i]
				if game.Hand[i].Strength > game.Hand[i+1].Strength {
					game.Hand[i] = game.Hand[i+1]
					game.Hand[i+1] = swap
					sortFlag = 1
				} else if game.Hand[i].Strength == game.Hand[i+1].Strength {
					//log.Println(game.Hand[i].Hand, game.Hand[i].Strength, "<<>>", game.Hand[i+1].Hand, game.Hand[i].Strength)
					//calculate highest card rule
					winner := highestCard(game.Hand[i].Hand, game.Hand[i+1].Hand)
					//log.Println(winner)
					switch game.Hand[i+1].Hand == winner {
					case true:
						game.Hand[i] = game.Hand[i+1]
						game.Hand[i+1] = swap
						sortFlag = 1
					}
				}
			}
		}

		if sortFlag == 0 {
			break
		}
	}

	// fmt.Printf("Post-Sort order:\n")
	// for i := range game.Hand {
	// 	fmt.Println(game.Hand[i].Hand, game.Hand[i].Strength)
	// }

	total := 0
	for i := range game.Hand {
		total = total + ((i + 1) * game.Hand[i].Bid)
	}
	fmt.Println("Total Winnings:", total)
	// Output execution time
	elapsed := time.Since(start)
	log.Printf("Execution time %s\n", elapsed)
}

func highestCard(hand1 string, hand2 string) string {
	cardValues := map[string]int{"2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "T": 10, "J": 11, "Q": 12, "K": 13, "A": 14}

	var loser string
	for i := 0; i < 5; i++ {
		//log.Println(cardValues[string(hand1[i])], "<<>>", cardValues[string(hand2[i])])
		if cardValues[string(hand1[i])] != cardValues[string(hand2[i])] {
			switch cardValues[string(hand1[i])] < cardValues[string(hand2[i])] {
			case true:
				loser = hand1
			case false:
				loser = hand2
			}
			break
		}
	}
	return loser
}

func parseCards(input readfile.InputFile) Game {
	game := Game{}
	hand := Hands{}
	for _, v := range input.InputRow {
		h := strings.Split(v, " ")
		hand.Hand = h[0]
		hand.Bid = convNum(h[1])
		hand.Strength = calcStrength(hand.Hand)
		game.Hand = append(game.Hand, hand)
	}
	return game
}

func calcStrength(cards string) int {
	/*
	   Five of a kind :: 7
	   Four of a kind :: 6
	   Full house (three of a kind and 1 pair) :: 5
	   Three of a kind :: 4
	   Two pair :: 3
	   One pair :: 2
	   high card :: 1
	*/

	cardMap := make(map[string]int)
	for i := 0; i < len(cards); i++ {
		c := string(cards[i])
		if _, exist := cardMap[c]; !exist {
			cardMap[c] = strings.Count(cards, c)
		}
	}

	// Count number of pairs
	p := 0
	for _, v := range cardMap {
		if v == 2 {
			p++
		}
	}

	// Get card strength
	strength := 1
	for _, v := range cardMap {
		switch v {
		case 5:
			strength = 7
		case 4:
			strength = 6
		case 3:
			//log.Println(k, v, "Checking for full house: pais:", p)
			if p == 1 && v == 3 { // Check for full house
				strength = 5
				//log.Println(k, v, "full house found! Strength =", strength)
			} else {
				strength = 4
			}
		case 2:
			//log.Println(k, v, "Checking # pairs:", p)
			if p == 2 {
				strength = 3
			} else {
				if strength <= 2 {
					strength = 2
					//log.Println(k, v, "1 pair assigning st:", strength)
				}
			}
		}
	}
	//log.Println(strength)
	return strength
}

func convNum(in string) int {
	out, err := strconv.Atoi(in)
	if err != nil {
		log.Fatal(err)
	}
	return out
}
