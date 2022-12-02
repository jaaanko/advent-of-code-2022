package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var translation = map[string]string{
	"A": "R",
	"B": "P",
	"C": "S",
	"X": "R",
	"Y": "P",
	"Z": "S",
}

var score = map[string]int{
	"R": 1,
	"P": 2,
	"S": 3,
}

var weakness = map[string]string{
	"R": "P",
	"P": "S",
	"S": "R",
}

var strength = map[string]string{
	"R": "S",
	"P": "R",
	"S": "P",
}

func solvePart1(input [][]string) int {
	total := 0
	for _, moves := range input {
		opponent := translation[moves[0]]
		me := translation[moves[1]]

		if opponent == me {
			total += 3
		} else if weakness[opponent] == me {
			total += 6
		}
		total += score[me]
	}
	return total
}

func solvePart2(input [][]string) int {
	total := 0
	for _, moves := range input {
		opponent := translation[moves[0]]
		me := moves[1]

		if me == "X" {
			total += score[strength[opponent]]
		} else if me == "Y" {
			total += score[opponent] + 3
		} else {
			total += score[weakness[opponent]] + 6
		}
	}
	return total
}

func main() {
	file, err := os.Open("input_01.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	input := [][]string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, strings.Fields(scanner.Text()))
	}

	fmt.Println("Part 1:", solvePart1(input))
	fmt.Println("Part 2:", solvePart2(input))
}
