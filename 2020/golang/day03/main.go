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
	var trees [][]bool = make([][]bool, 0, 10)
	var y int = 0
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		trees = append(trees, make([]bool, 0, 10))
		for _, data := range line {
			// fmt.Printf("%d, %d\n", x, y)
			if data == '#' {

				trees[y] = append(trees[y], true)
			} else {
				trees[y] = append(trees[y], false)
			}
		}
		y++
	}
	treeCount := countTrees(trees, 3, 1)
	fmt.Printf("tree count for slope 3, 1 is: %d\n", treeCount)
	product := treeCount * countTrees(trees, 1, 1) * countTrees(trees, 5, 1) * countTrees(trees, 7, 1) * countTrees(trees, 1, 2)
	fmt.Printf("product of all tree counts is: %d\n", product)
}

func countTrees(trees [][]bool, slope_x, slope_y int) int {
	treeCount, pos_x, pos_y := 0, 0, 0
	width := len(trees[0])
	height := len(trees)
	for true {
		pos_x += slope_x
		pos_y += slope_y
		if pos_y >= height {
			break
		}
		if trees[pos_y][pos_x%width] == true {
			treeCount++
		}
	}
	return treeCount
}
