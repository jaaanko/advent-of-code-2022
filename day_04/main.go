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

func fullOverlap(interval1 interval, interval2 interval) bool {
	return interval1.start <= interval2.start && interval2.end <= interval1.end ||
		interval2.start <= interval1.start && interval1.end <= interval2.end
}

func intersect(interval1 interval, interval2 interval) bool {
	return interval1.start <= interval2.end && interval2.start <= interval1.end
}

func solvePart1(input [][]interval) int {
	total := 0
	for _, intervals := range input {
		if fullOverlap(intervals[0], intervals[1]) {
			total += 1
		}
	}
	return total
}

func solvePart2(input [][]interval) int {
	total := 0
	for _, intervals := range input {
		if intersect(intervals[0], intervals[1]) {
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
