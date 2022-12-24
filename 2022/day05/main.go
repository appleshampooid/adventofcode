package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/appleshampooid/adventofcode/util"
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
	var stacks []util.Stack[byte] = make([]util.Stack[byte], 0, 5)
	var rawStacks util.Stack[string] = util.NewStack[string]()
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		stackLabel := regexp.MustCompile(`^(?: *(\d+) *)+$`)
		groups := stackLabel.FindStringSubmatch(line)
		if groups != nil {
			numStacks, _ := strconv.Atoi(groups[1])
			fmt.Printf("We have reached the stack labels, and we have %d of them\n", numStacks)
			fmt.Printf("Raw stacks are:\n")
			rawStacks.Print()
			for {
				stackLine, err := rawStacks.Pop()
				if err != nil {
					break
				}
				stack := 0
				for {
					remaining := len(stackLine)
					if remaining < 3 {
						break
					}
					nextLen := 4
					if remaining == 3 {
						nextLen = 3
					}
					curr := stackLine[0:nextLen]
					stackLine = stackLine[nextLen:]
					r := curr[1]
					if len(stacks) < stack+1 {
						stacks = append(stacks, util.Stack[byte]{})
					}
					if r != ' ' {
						stacks[stack].Push(r)
					}
					stack++
				}
			}
			fmt.Println("stacks are:")
			printByteStacks(stacks)
		} else {
			rawStacks.Push(line)
		}
	}
}

func printStacks[T interface{}](stacks []util.Stack[T]) {
	for _, s := range stacks {
		fmt.Print("[")
		s.Print()
		fmt.Print("]\n")
	}
}

func printByteStacks(stacks []util.Stack[byte]) {
	for _, s := range stacks {
		fmt.Print("[")
		s.PrintByteStack()
		fmt.Print("]\n")
	}
}
