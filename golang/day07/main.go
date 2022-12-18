package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	//"strconv"
)

type quantityOfBag struct {
	bag   bag
	count int
}

type bag struct {
	contents []quantityOfBag
	color    string
}

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

	scanner := bufio.NewScanner(inputFile)
	bagRegexp := regexp.MustCompile("(.+) bags contain (\\d+) (.+) bags?[\\.,]( (\\d+) (.+) bags\\.)?")
	for scanner.Scan() {
		line := scanner.Text()
		match := bagRegexp.MatchString(line)
		fmt.Println(fmt.Sprintf("Line \"%s\" had a match: %t\n", line, match))
	}
}
