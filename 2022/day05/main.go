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
	stackLabel := regexp.MustCompile(`^(?: *(\d+) *)+$`)
	move := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()
		groups := stackLabel.FindStringSubmatch(line)
		if groups != nil {
			numStacks, _ := strconv.Atoi(groups[1])
			fmt.Printf("We have reached the stack labels, and we have %d of them\n", numStacks)
			fmt.Printf("Raw stacks are:\n")
			rawStacks.Print()
			// for (int i = 0; i < 10; i++)
			for stackLine, err := rawStacks.Pop(); err == nil; stackLine, err = rawStacks.Pop() {
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
		} else if len(stacks) == 0 {
			rawStacks.Push(line)
		} else {
			if line == "" {
				continue
			}
			moves := move.FindStringSubmatch(line)
			if moves == nil || len(moves) != 4 {
				panic(fmt.Sprintf("malformed input: length %d", len(moves)))
			}
			quantity, _ := strconv.Atoi(moves[1])
			source, _ := strconv.Atoi(moves[2])
			dest, _ := strconv.Atoi(moves[3])
			fmt.Printf("%d %d %d\n", quantity, source, dest)
			// part 1 answer
			// executeMove9000(stacks, quantity, source, dest)
			// part2 answer
			executeMove9001(stacks, quantity, source, dest)
		}
	}
	var tops []byte
	for _, stack := range stacks {
		top, err := stack.Peek()
		if err != nil {
			panic("Malformed input when peeking at end condition")
		}
		tops = append(tops, top)
	}
	fmt.Printf("Top of the stacks are: %s\n", tops)
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

func executeMove9000(stacks []util.Stack[byte], quantity, source, dest int) {
	for i := 0; i < quantity; i++ {
		b, err := stacks[source-1].Pop()
		if err != nil {
			panic("Malformed stacks/instructions")
		}
		stacks[dest-1].Push(b)
		fmt.Printf("Moved 1 block from %d to %d\n", source, dest)
	}
}

func executeMove9001(stacks []util.Stack[byte], quantity, source, dest int) {
	b, err := stacks[source-1].PopN(quantity)
	if err != nil {
		panic("Malformed stacks/instructions")
	}
	stacks[dest-1].PushN(b)
	fmt.Printf("Moved %d block from %d to %d\n", quantity, source, dest)
}
