package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	//"regexp"
	//"strconv"
)

type void struct{}

var p void

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
	prioritySum := 0
	prioritySumPart2 := 0
	groupCounter := 0
	var elf1Contents map[rune]void
	var elf2Intersection map[rune]void
	for scanner.Scan() {
		line := scanner.Text()

		// part 1
		compartmentSize := len(line) / 2
		compartmentOne := line[:compartmentSize]
		compartmentTwo := line[compartmentSize:]
		fmt.Printf("1: '%s', 2: '%s'\n", compartmentOne, compartmentTwo)
		compartmentOneContents := make(map[rune]bool)
		for _, item := range compartmentOne {
			// fmt.Printf("Rune is %v\n", item)
			compartmentOneContents[item] = true
		}
		for _, item := range compartmentTwo {
			if _, present := compartmentOneContents[item]; present {
				prioritySum += runePriority(item)
				break
			}
		}

		// part 2
		if groupCounter == 0 {
			elf1Contents = make(map[rune]void)
			elf2Intersection = make(map[rune]void)
			for _, item := range line {
				elf1Contents[item] = p
			}
			groupCounter++
		} else if groupCounter == 1 {
			for _, item := range line {
				if _, present := elf1Contents[item]; present {
					elf2Intersection[item] = p
				}
			}
			groupCounter++
		} else if groupCounter == 2 {
			for _, item := range line {
				if _, present := elf2Intersection[item]; present {
					prioritySumPart2 += runePriority(item)
					break
				}
			}
			groupCounter = 0
		}
	}
	fmt.Printf("Priority sum is %d\n", prioritySum)
	fmt.Printf("Priority sum part2 is %d\n", prioritySumPart2)
	// fmt.Printf("a b c z A B Runes are %v %v %v %v %v %v\n", 'a', 'b', 'c', 'z', 'A', 'B')
}

func runePriority(r rune) int {
	if r >= 97 && r <= 122 {
		return int(r - 96)
	} else {
		return int(r - 38)
	}
}
