package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/jaaanko/advent-of-code-2022/imath"
)

type point struct {
	row, col int
}

func main() {
	inputBytes, err := os.ReadFile("input_01.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Solve Part 1:", solvePart1(parseInput(bytes.NewReader(inputBytes))))
	fmt.Println("Solve Part 2:", solvePart2(parseInput(bytes.NewReader(inputBytes))))
}

func parseInput(r io.Reader) map[point]bool {
	grid := map[point]bool{}
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		var prevPoint point

		for i, p := range strings.Split(scanner.Text(), " -> ") {
			fields := strings.Split(p, ",")
			row, _ := strconv.Atoi(fields[1])
			col, _ := strconv.Atoi(fields[0])
			currPoint := point{row, col}

			if i > 0 {
				for i := imath.MinInt(currPoint.row, prevPoint.row); i <= imath.MaxInt(currPoint.row, prevPoint.row); i++ {
					grid[point{i, col}] = true
				}
				for j := imath.MinInt(currPoint.col, prevPoint.col); j <= imath.MaxInt(currPoint.col, prevPoint.col); j++ {
					grid[point{row, j}] = true
				}
			} else {
				grid[currPoint] = true
			}
			prevPoint = currPoint
		}
	}

	return grid
}

func solvePart1(grid map[point]bool) int {
	result := 0
	maxRow := -1
	for k, _ := range grid {
		maxRow = imath.MaxInt(maxRow, k.row)
	}

	var sand point
	for sand.row <= maxRow {
		sand = point{0, 500}

		for !grid[sand] && sand.row <= maxRow {
			if !grid[point{sand.row + 1, sand.col}] {
				sand.row += 1
			} else if !grid[point{sand.row + 1, sand.col - 1}] {
				sand.row += 1
				sand.col -= 1
			} else if !grid[point{sand.row + 1, sand.col + 1}] {
				sand.row += 1
				sand.col += 1
			} else {
				grid[sand] = true
				result += 1
			}
		}
	}

	return result
}

func solvePart2(grid map[point]bool) int {
	result := 0
	maxRow := -1
	for k, _ := range grid {
		maxRow = imath.MaxInt(maxRow, k.row)
	}

	floor := maxRow + 2
	for !grid[point{0, 500}] {
		sand := point{0, 500}

		for !grid[sand] && sand.row < floor-1 {
			if !grid[point{sand.row + 1, sand.col}] {
				sand.row += 1
			} else if !grid[point{sand.row + 1, sand.col - 1}] {
				sand.row += 1
				sand.col -= 1
			} else if !grid[point{sand.row + 1, sand.col + 1}] {
				sand.row += 1
				sand.col += 1
			} else {
				grid[sand] = true
				result += 1
			}
		}

		if sand.row == floor-1 {
			grid[sand] = true
			result += 1
		}
	}

	return result
}
