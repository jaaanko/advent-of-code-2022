package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	set "github.com/deckarep/golang-set/v2"
)

type instruction struct {
	direction string
	dist      int
}

type point struct {
	x, y int
}

func main() {
	file, err := os.Open("input_01.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	instructions := []instruction{}
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		dist, _ := strconv.Atoi(line[1])
		instructions = append(instructions, instruction{line[0], dist})
	}

	fmt.Println("Part 1:", solvePart1(instructions))
	fmt.Println("Part 2:", solvePart2(instructions))
}

func solvePart1(instructions []instruction) int {
	seen := set.NewSet([]point{point{0, 0}}...)
	head := point{0, 0}
	tail := point{0, 0}

	for _, in := range instructions {
		for i := 0; i < in.dist; i++ {
			head = moveByOne(in.direction, head)
			if !adjacent(head, tail) {
				tail = chaseAfter(head, tail)
			}
			seen.Add(tail)
		}
	}

	return seen.Cardinality()
}

func moveByOne(direction string, p point) point {
	if direction == "R" {
		p.x += 1
	} else if direction == "L" {
		p.x -= 1
	} else if direction == "U" {
		p.y += 1
	} else {
		p.y -= 1
	}

	return p
}

func adjacent(head point, tail point) bool {
	for dx := -1; dx < 2; dx++ {
		for dy := -1; dy < 2; dy++ {
			nei := point{head.x + dx, head.y + dy}
			if nei == tail {
				return true
			}
		}
	}

	return false
}

func chaseAfter(head point, tail point) point {
	if tail.x != head.x {
		if tail.x < head.x {
			tail.x += 1
		} else {
			tail.x -= 1
		}
	}

	if tail.y != head.y {
		if tail.y < head.y {
			tail.y += 1
		} else {
			tail.y -= 1
		}
	}

	return tail
}

func solvePart2(instructions []instruction) int {
	const n = 10
	seen := set.NewSet([]point{point{0, 0}}...)
	knots := [n]point{}

	for _, in := range instructions {
		for i := 0; i < in.dist; i++ {
			knots[0] = moveByOne(in.direction, knots[0])
			for j := 1; j < n; j++ {
				if !adjacent(knots[j-1], knots[j]) {
					knots[j] = chaseAfter(knots[j-1], knots[j])
				}
			}
			seen.Add(knots[n-1])
		}
	}

	return seen.Cardinality()
}
