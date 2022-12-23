package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type domain struct {
	begin int
	end   int
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
	totalFullOverlap := 0
	totalOverlap := 0
	for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile(`(\d+)-(\d+),(\d+)-(\d+)`)
		groups := re.FindStringSubmatch(line)
		ints := make([]int, 4)
		if groups == nil {
			panic("Malformed input")
		}
		for i, delimeter := range groups[1:] {
			// I can ignore this error because the regex guarantees they can be converted :grin:
			ints[i], _ = strconv.Atoi(delimeter)
		}
		r1 := domain{begin: ints[0], end: ints[1]}
		r2 := domain{begin: ints[2], end: ints[3]}
		//fmt.Printf("ranges are %+v and %+v\n", r1, r2)
		if eitherContains(r1, r2) {
			totalFullOverlap++
		}
		if anyOverlap(r1, r2) {
			totalOverlap++
		}
	}
	fmt.Printf("There are %d total  overlapping pairs\n", totalFullOverlap)
	fmt.Printf("There are %d any overlapping pairs\n", totalOverlap)
}

func eitherContains(r1, r2 domain) bool {
	return r1.contains(r2) || r2.contains(r1)
}

func anyOverlap(r1, r2 domain) bool {
	return r1.end >= r2.begin && r2.end >= r1.begin
}

func (d domain) contains(o domain) bool {
	return d.begin <= o.begin && d.end >= o.end
}
