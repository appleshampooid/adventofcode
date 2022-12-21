package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
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
	}
	defer inputFile.Close()
	var input []int
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf("Bad input, could not convert to int: %s\n", scanner.Text())
			os.Exit(1)
		}
		input = append(input, line)
	}
	fmt.Printf("%v\n", input)
	for i, v := range input {
		for j, w := range input[i:] {
			if v+w == 2020 {
				fmt.Printf("Part 1: The two values are %d and %d, and the product is %d\n", v, w, v*w)
			}
			for _, x := range input[j:] {
				if v+w+x == 2020 {
					fmt.Printf("Part 2: The three values are %d %d %d, and the product is %d\n", v, w, x, v*w*x)
				}
			}
		}
	}
}
