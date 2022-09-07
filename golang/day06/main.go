package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"strings"
	//"regexp"
	//"strconv"
)

func ScanBlocks(data []byte, atEOF bool) (int, []byte, error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.Index(data, []byte("\n\n")); i >= 0 {
		return i + 2, data[0:i], nil
	}
	// If we're at EOF, we have a final, non-terminated block. Return it. Strip the newline if it exists.
	if atEOF {
		originalLength := len(data)
		if originalLength > 0 && data[originalLength-1] == '\n' {
			data = data[0 : originalLength-1]
		}
		return originalLength, data, nil
	}
	// Request more data.
	return 0, nil, nil
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

	numAnyYes := 0
	numAllYes := 0
	scanner := bufio.NewScanner(inputFile)
	scanner.Split(ScanBlocks)
	for scanner.Scan() {
		yeses := make(map[rune]int)
		block := scanner.Text()
		responses := strings.Split(block, "\n")
		for _, response := range responses {
			for _, c := range response {
				yeses[c]++
			}
		}
		numResponses := len(responses)
		for _, numYes := range yeses {
			numAnyYes++
			if numYes == numResponses {
				numAllYes++
			}
		}
	}
	fmt.Printf("Total number of any yeses: %d\n", numAnyYes)
	fmt.Printf("Total number of all yeses: %d\n", numAllYes)
}
