package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type operation struct {
	name string
	arg  int
}

type sprite struct {
	left, right int
}

type pixel struct {
	row, col int
}

func main() {
	file, err := os.Open("input_01.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	operations := parseInput(file)
	fmt.Println("Part 1:", solvePart1(operations))
	fmt.Println("Part 2:", "\n"+solvePart2(operations))
}

func parseInput(r io.Reader) []operation {
	scanner := bufio.NewScanner(r)
	operations := []operation{}

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		operations = append(operations, operation{name: "noop"})

		if fields[0] == "addx" {
			arg, _ := strconv.Atoi(fields[1])
			operations = append(operations, operation{name: "addx", arg: arg})
		}
	}

	return operations
}

func solvePart1(operations []operation) int {
	total := 0
	cycle := 1
	x := 1
	for _, op := range operations {
		if (cycle-20)%40 == 0 {
			total += cycle * x
		}
		if op.name == "addx" {
			x += op.arg
		}
		cycle += 1
	}

	return total
}

func solvePart2(operations []operation) string {
	board := []rune{}
	cycle := 1
	x := 1
	currPixel := pixel{0, 0}
	sprite := sprite{0, 2}

	for _, op := range operations {
		cycle += 1

		if currPixel.col >= sprite.left && currPixel.col <= sprite.right {
			board = append(board, '#')
		} else {
			board = append(board, '.')
		}

		if op.name == "addx" {
			x += op.arg
			sprite.left = x - 1
			sprite.right = x + 1
		}

		currPixel.col += 1
		if cycle%40 == 1 {
			currPixel.row += 1
			currPixel.col = 0
			board = append(board, '\n')
		}
	}

	return string(board[:len(board)-1])
}
