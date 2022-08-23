package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	//"regexp"
	//"strconv"
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
	trees := make([]bool, 0)
	var j int = 0
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		trees = append(trees, make([]bool, 0)
		for i, data := range line {
			fmt.Printf("%d, %d\n", i, j)
			trees[i] = 
			if data == '#' {
				//trees[i][j] = true
			}
		}
		j++
	}
}
