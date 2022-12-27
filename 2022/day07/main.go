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

	dirSizes := make(map[string]int)
	var cwd []string = make([]string, 0, 5)
	cdRe := regexp.MustCompile(`^\$ cd (.+)$`)
	// dirRe := regexp.MustCompile(`^dir .+$`)
	fileRe := regexp.MustCompile(`^(\d+) .+$`)
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		if groups := cdRe.FindStringSubmatch(line); groups != nil {
			if groups[1] == ".." {
				cwd = cwd[0 : len(cwd)-1]
			} else if groups[1] == "/" {
				cwd = []string{"/"}
			} else {
				cwd = append(cwd, groups[1])
			}
			fmt.Printf("dir stack is %v\n", cwd)
		} else if groups := fileRe.FindStringSubmatch(line); groups != nil {
			size, err := strconv.Atoi(groups[1])
			if err != nil {
				panic("Error converting size to int")
			}
			fullPath := ""
			for _, d := range cwd {
				fullPath += fmt.Sprintf("%s/", d)
				if _, present := dirSizes[fullPath]; !present {
					dirSizes[fullPath] = size
				} else {
					dirSizes[fullPath] += size
				}
			}
		}
	}
	fmt.Printf("dirSizes are %v\n", dirSizes)
	freeSpace := 70000000 - dirSizes["//"]
	spaceNeeded := 30000000 - freeSpace
	fmt.Printf("total space needed to clear is %d\n", spaceNeeded)
	totalSmolDirs := 0
	deletionCandidate := ""
	deletionCandidateSize := 0
	for dir, size := range dirSizes {
		if size <= 100000 {
			totalSmolDirs += size
		}
		if size >= spaceNeeded && (deletionCandidate == "" || size < deletionCandidateSize) {
			deletionCandidateSize = size
			deletionCandidate = dir
		}
	}
	fmt.Printf("total Smol Dirs is %d\n", totalSmolDirs)
	fmt.Printf("Deletion candidate is %s with size %d\n", deletionCandidate, deletionCandidateSize)
	// 1656621 is too lowv
}
