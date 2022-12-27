package main

// I must say it once more: ONLY KNUTH CAN JUDGE ME

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
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

	var heights [][]int = make([][]int, 0, 5)
	row := 0
	numVisible := 0
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		heights = append(heights, make([]int, 0, 5))
		for _, r := range line {
			height, err := strconv.Atoi(string(r))
			if err != nil {
				panic("malformed input")
			}
			fmt.Printf("%d ", height)
			heights[row] = append(heights[row], height)
		}
		row++
		fmt.Println()
	}
	fmt.Printf("%v\n", heights)
	numVisible += len(heights) * 2
	numVisible += (len(heights[0]) - 2) * 2
	maxScenic := 0
	for i, row := range heights[1 : len(heights)-1] { //
		for j, height := range row[1 : len(row)-1] {
			scenicScore := 1
			westOccluded := false
			eastOccluded := false
			northOccluded := false
			southOccluded := false
			x, y := i+1, j+1
			fmt.Printf("Checking vis on %d, %d\n", x, y)
			var k int
			var count int
			//for _, o := range heights[x][0:y] {
			for k = y - 1; k >= 0; k-- {
				count++
				if heights[x][k] >= height {
					fmt.Printf("%d, %d is occluded from the west\n", x, y)
					westOccluded = true
					break
				}
			}
			scenicScore *= count
			count = 0
			// fmt.Printf("y is %d, len(heights[x]) is %d\n", y, len(heights[x]))
			// for _, o := range heights[x][y+1 : len(heights[x])] {
			for k = y + 1; k < len(heights[x]); k++ {
				count++
				// fmt.Printf("Checking %d, %d for occlusion from the east at y coords %d\n", x, y, j2)
				if heights[x][k] >= height {
					fmt.Printf("%d, %d is occluded from the east\n", x, y)
					eastOccluded = true
					break
				}
			}
			scenicScore *= count
			count = 0
			for k = x - 1; k >= 0; k-- {
				count++
				if heights[k][y] >= height {
					fmt.Printf("%d, %d is occluded from the north\n", x, y)
					northOccluded = true
					break
				}
			}
			scenicScore *= count
			count = 0
			for k = x + 1; k < len(heights); k++ {
				count++
				if heights[k][y] >= height {
					fmt.Printf("%d, %d is occluded from the south\n", x, y)
					southOccluded = true
					break
				}
			}
			scenicScore *= count
			if !(westOccluded && eastOccluded && northOccluded && southOccluded) {
				fmt.Printf("%d, %d with height %d is visible!\n", x, y, height)
				numVisible++
			}
			if scenicScore > maxScenic {
				maxScenic = scenicScore
			}
		}
	}
	fmt.Printf("numVisible is %d\n", numVisible)
	fmt.Printf("maximum scenic score is %d\n", maxScenic)
	fmt.Printf("dimensions are: %d x %d\n", len(heights), len(heights[0]))
}
