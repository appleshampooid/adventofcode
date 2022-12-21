package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	//"regexp"
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

	scanner := bufio.NewScanner(inputFile)
	var highestID uint64 = 0
	// cheating a bit since I already know the max
	var seatMap [849]bool
	for scanner.Scan() {
		line := scanner.Text()
		var row_spec string = line[0:7]
		row_bin := strings.ReplaceAll(row_spec, "F", "0")
		row_bin = strings.ReplaceAll(row_bin, "B", "1")
		row, err := strconv.ParseUint(row_bin, 2, 8)
		if err != nil {
			fmt.Printf("Error decoding %s as binary\n", row_bin)
		}
		var col_spec string = line[7:]
		col_bin := strings.ReplaceAll(col_spec, "L", "0")
		col_bin = strings.ReplaceAll(col_bin, "R", "1")
		col, err := strconv.ParseUint(col_bin, 2, 8)
		if err != nil {
			fmt.Printf("Error decoding %s as binary\n", col_bin)
		}

		// fmt.Printf("row is %v, col is %v\n", row, col)
		// fmt.Printf("seat ID is %d\n", SeatID(row, col))
		seatID := SeatID(row, col)
		seatMap[seatID] = true
		if seatID > highestID {
			highestID = seatID
		}
	}
	fmt.Printf("Highest ID is %d\n", highestID)
	started := false
	for id, present := range seatMap {
		// fmt.Printf("%d: %v\n", id, present)
		if !started && present {
			started = true
		} else if started && !present {
			fmt.Printf("My seat is %d\n", id)
			break
		}
	}
}

func SeatID(row, col uint64) uint64 {
	return row*8 + col
}
