package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
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
		os.Exit(1)
	}
	defer inputFile.Close()
	password_regex := regexp.MustCompile(`(\d+)-(\d+) (\w): (\w+)`)
	var valid int
	var valid_pt2 int
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		matches := password_regex.FindStringSubmatch(line)
		if matches == nil || len(matches) != 5 {
			fmt.Printf("This string did not matth the spec: %s\n", line)
			continue
		}
		min, err := strconv.Atoi(matches[1])
		max, err2 := strconv.Atoi(matches[2])
		if err != nil || err2 != nil {
			fmt.Printf("These should be ints :shrug:: '%s' '%s'\n", matches[1], matches[2])
			continue
		}
		char := matches[3]
		password := matches[4]
		var count int
		for _, c := range password {
			if string(c) == char {
				count++
			}
		}
		if count <= max && count >= min {
			valid++
		}
		one_match := string(password[min-1]) == char
		two_match := string(password[max-1]) == char
		if one_match != two_match {
			valid_pt2++
		} else {
			fmt.Printf("spec %s returned invalid, check it?\none_match: %v two_match: %v char: %v min: %v max:%v password: %v\n\n", line, one_match, two_match, char, min, max, password)
		}
	}
	fmt.Printf("There are %d valid passwords for part 1\n", valid)
	fmt.Printf("There are %d valid passwords for part 2\n", valid_pt2)
}
