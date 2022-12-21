package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type bag struct {
	contents []quantityOfBag
	color    string
}

type quantityOfBag struct {
	bag   *bag
	count int
}

var bagIndex map[string]*bag

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

	bagIndex = make(map[string]*bag)
	scanner := bufio.NewScanner(inputFile)
	outerBagRegexp := regexp.MustCompile(`(.+) bags contain((?: (?:(?:\d+ [^\d]+)|no other) bags?[\.,])+)`)
	innerBagRegexp := regexp.MustCompile(`(\d+) ([^\d]+) bags?[\.,]`)
	for scanner.Scan() {
		line := scanner.Text()
		outerGroups := outerBagRegexp.FindStringSubmatch(line)
		if outerGroups != nil {
			fmt.Printf("Line \"%s\" had a match, slice of size %d, subgroups:\n", line, len(outerGroups))
			color := outerGroups[1]
			outerBag := getOrAddBag(color)
			for i, group := range outerGroups {
				fmt.Printf(`"%s", `, group)
				fmt.Println()
				if i < 2 {
					continue
				}
				innerGroups := innerBagRegexp.FindAllStringSubmatch(group, -1)
				if innerGroups != nil {
					fmt.Printf("Subline \"%s\" had matches\n", group)
					for _, ig := range innerGroups {
						fmt.Printf("%s bags of color %s\n", ig[1], ig[2])
						innerBag := getOrAddBag(ig[2])
						quantity, err := strconv.Atoi(ig[1])
						if err != nil {
							panic("Hahah, this MUST Be a number based on regex but it ain't")
						}
						outerBag.contents = append(outerBag.contents, quantityOfBag{bag: innerBag, count: quantity})
					}
					fmt.Println()
				}
			}
		}
	}
	fmt.Printf("bagIndex is %+v\n", bagIndex)
	count := 0
	for _, bag := range bagIndex {
		bag.print()
		if bag.contains("shiny gold") {
			count++
		}
	}
	fmt.Printf("%d bags can hold a shiny gold bag\n", count)
	fmt.Printf("a shiny gold bag holds %d total bags\n", bagIndex["shiny gold"].totalBags())
}

func getOrAddBag(color string) *bag {
	var candidateBag *bag
	var present bool
	if candidateBag, present = bagIndex[color]; !present {
		candidateBag = &bag{color: color, contents: make([]quantityOfBag, 0, 3)}
		bagIndex[color] = candidateBag
	}
	fmt.Printf("candidateBag is: %+v\n", candidateBag)
	return candidateBag
}

func (this bag) contains(color string) bool {
	for _, content := range this.contents {
		if content.bag.color == color {
			return true
		}
		if content.bag.contains(color) {
			return true
		}
	}
	return false
}

func (this *bag) print() {
	fmt.Printf("My color is %s and my memory location is %p.\n I contain:", this.color, this)
	for _, content := range this.contents {
		fmt.Printf("%d %s bags at memory location %p, ", content.count, content.bag.color, content.bag)
	}
	fmt.Print("\n\n")
}

func (this *bag) totalBags() int {
	total := 0
	for _, content := range this.contents {
		total += content.count
		total += content.count * content.bag.totalBags()
	}
	return total
}
