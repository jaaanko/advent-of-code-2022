package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jaaanko/advent-of-code-2022/stack"
)

type instruction struct {
	count, source, dest int
}

func main() {
	bytes, err := os.ReadFile("input_01.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(string(bytes), "\r\n\r\n")
	moves := []instruction{}

	for _, line := range strings.Split(input[1], "\n") {
		var count, source, dest int
		fmt.Sscanf(line, "move %d from %d to %d", &count, &source, &dest)
		moves = append(moves, instruction{count, source, dest})
	}

	fmt.Println("Part 1:", solvePart1(createStacksFromInput(input[0]), moves))
	fmt.Println("Part 2:", solvePart2(createStacksFromInput(input[0]), moves))
}

func createStacksFromInput(input string) []stack.Stack[rune] {
	stacks := []stack.Stack[rune]{}
	lines := strings.Split(input, "\n")

	for i := 0; i < (len(lines[0])+1)/4; i++ {
		stacks = append(stacks, stack.New[rune]())
	}

	for i := len(lines)-1; i >= 0; i-- {
		line := []rune(lines[i])
		for j := 0; j < len(line); j += 4 {
			if line[j] == '[' {
				stacks[j/4].Push(line[j+1])
			}
		}
	}
	return stacks
}

func solvePart1(stacks []stack.Stack[rune], moves []instruction) string {
	for _,instruction := range moves {
		for i := 0; i < instruction.count; i++ {
			element,err := stacks[instruction.source-1].Pop()
			if err != nil {
				log.Fatal(err)
			}
			stacks[instruction.dest-1].Push(element)
		}
	}
	return readFromTopOfStacks(stacks)
}

func readFromTopOfStacks(stacks []stack.Stack[rune]) string {
	result := []rune{}
	for _,stack := range stacks {
		element,err := stack.Peek()
		if err != nil {
			log.Fatal(err)
		}
		result = append(result,element)
	}
	return string(result)
}

func solvePart2(stacks []stack.Stack[rune], moves []instruction) string {
	for _,instruction := range moves {
		tmp := []rune{}
		for i := 0; i < instruction.count; i++ {
			element,err := stacks[instruction.source-1].Pop()
			if err != nil {
				log.Fatal(err)
			}
			tmp = append(tmp,element)
		}

		for i := len(tmp)-1; i >= 0; i-- {
			stacks[instruction.dest-1].Push(tmp[i])
		}
	}
	return readFromTopOfStacks(stacks)
}
