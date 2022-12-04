package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type interval struct {
	start, end int
}

func fullyIntersect(a int, b int, c int, d int) bool {
	return a <= c && d <= b || c <= a && b <= d
}

func intersect(a int, b int, c int, d int) bool {
	return a <= c && c <= b || c <= a && a <= d
}

func solvePart1(input [][]interval) int {
	total := 0
	for _, intervals := range input {
		firstInterval := intervals[0]
		secondInterval := intervals[1]
		if fullyIntersect(firstInterval.start, firstInterval.end, secondInterval.start, secondInterval.end) {
			total += 1
		}
	}
	return total
}

func solvePart2(input [][]interval) int {
	total := 0
	for _, intervals := range input {
		firstInterval := intervals[0]
		secondInterval := intervals[1]
		if intersect(firstInterval.start, firstInterval.end, secondInterval.start, secondInterval.end) {
			total += 1
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
	scanner := bufio.NewScanner(file)
	input := [][]interval{}

	for scanner.Scan() {
		intervals := []interval{}
		for _, group := range strings.Split(scanner.Text(), ",") {
			points := strings.Split(group, "-")
			start, _ := strconv.Atoi(points[0])
			end, _ := strconv.Atoi(points[1])
			intervals = append(intervals, interval{
				start: start,
				end:   end,
			})
		}
		input = append(input, intervals)
	}

	fmt.Println("Part 1:", solvePart1(input))
	fmt.Println("Part 2:", solvePart2(input))
}
