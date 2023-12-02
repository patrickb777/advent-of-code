// Package will read a text based input file line by line and store those in a struct

package readfile

import (
	"bufio"
	"log"
	"os"
)

type InputFile struct {
	InputRow []string
}

func readFile(file string) InputFile {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	// Scan each line of input file with scanner and append to InputFile struct
	i := InputFile{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		i.InputRow = append(i.InputRow, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return i
}
