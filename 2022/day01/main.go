package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	//"regexp"
	//"strconv"
	"github.com/appleshampooid/adventofcode/util"
)

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

	blockScanner := bufio.NewScanner(inputFile)
	blockScanner.Split(util.ScanBlocks)
	elfCalories := make([]int, 0, 5)
	for blockScanner.Scan() {
		block := blockScanner.Text()
		scanner := bufio.NewScanner(strings.NewReader(block))
		elfTotal := 0
		for scanner.Scan() {
			line := scanner.Text()
			calories, err := strconv.Atoi(line)
			if err != nil {
				panic("fock you, AOC")
			}
			elfTotal += calories
		}
		elfCalories = append(elfCalories, elfTotal)
	}
	sort.Ints(elfCalories)
	fmt.Printf("%v\n", elfCalories)
	fmt.Printf("Max elf calories are %d\n", elfCalories[len(elfCalories)-1])
	fmt.Printf("And top 3 sum... %v\n", elfCalories[len(elfCalories)-1]+elfCalories[len(elfCalories)-2]+elfCalories[len(elfCalories)-3])
}
