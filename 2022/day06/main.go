package main

import (
	"flag"
	"fmt"
	"io/ioutil"
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

	input, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic("Error reading string")
	}
	accumulator := make(map[byte]int)
	// for part 1 answer:
	// markerSize := 4
	markerSize := 14
	for i, char := range input {
		fmt.Printf("map is %+v\n", accumulator)
		if _, present := accumulator[char]; present {
			accumulator[char]++
		} else {
			accumulator[char] = 1
		}
		if len(accumulator) == markerSize {
			fmt.Printf("Found the marker at %d\n", i+1)
			break
		}
		if i >= markerSize-1 {
			if count, present := accumulator[input[i-(markerSize-1)]]; !present {
				panic("Malformed input/logic lol")
			} else if count == 1 {
				delete(accumulator, input[i-(markerSize-1)])
			} else {
				accumulator[input[i-(markerSize-1)]]--
			}
		}
	}
	// fmt.Printf("Reached end of string without finding marker\nmap is %+v\n", last3)
}
