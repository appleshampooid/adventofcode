package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	//"regexp"
)

const (
	rock     int = 1
	paper        = 2
	scissors     = 3
)

var mapper map[string]int

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Println("Need an input file")
		os.Exit(1)
	}
	fileName := flag.Args()[0]
	inputFile, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error opening file %s\n", fileName)
		os.Exit(1)
	}
	defer inputFile.Close()

	mapper = map[string]int{"A": rock,
		"X": rock,
		"B": paper,
		"Y": paper,
		"C": scissors,
		"Z": scissors}
	scanner := bufio.NewScanner(inputFile)
	totalScorePart1 := 0
	totalScorePart2 := 0
	for scanner.Scan() {
		plays := strings.Split(scanner.Text(), " ")
		totalScorePart1 += scoreRound(mapper[plays[0]], mapper[plays[1]])
		totalScorePart2 += scoreRound(mapper[plays[0]], determineMyPlay(mapper[plays[0]], plays[1]))
	}
	fmt.Printf("Total score for part 1 is %d\n", totalScorePart1)
	fmt.Printf("Total score for part 2 is %d\n", totalScorePart2)
}

func scoreRound(theirs, mine int) int {
	score := 0
	if (mine == rock && theirs == scissors) || (mine == paper && theirs == rock) || (mine == scissors && theirs == paper) {
		score += 6
	} else if mine == theirs {
		score += 3
	}
	score += mine
	return score
}

func determineMyPlay(theirs int, result string) int {
	switch result {
	case "X":
		switch theirs {
		case rock:
			return scissors
		case paper:
			return rock
		case scissors:
			return paper
		}
	case "Y":
		return theirs
	case "Z":
		switch theirs {
		case rock:
			return paper
		case paper:
			return scissors
		case scissors:
			return rock
		}
	}
	panic("whoops")
}
