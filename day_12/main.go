package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/jaaanko/advent-of-code-2022/imath"
)

type point struct {
	row, col int
}

var directions = [][]int{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1},
}

func main() {
	inputBytes, err := os.ReadFile("input_01.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1:", solvePart1(parseInput(bytes.NewReader(inputBytes))))
	fmt.Println("Part 2:", solvePart2(parseInput(bytes.NewReader(inputBytes))))
}

func parseInput(r io.Reader) [][]rune {
	scanner := bufio.NewScanner(r)
	grid := [][]rune{}

	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	return grid
}

func solvePart1(grid [][]rune) int {
	m := len(grid)
	n := len(grid[0])
	var start point
	var target point

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 'S' {
				start = point{i, j}
				grid[i][j] = 'a'
			} else if grid[i][j] == 'E' {
				target = point{i, j}
				grid[i][j] = 'z'
			}
		}
	}

	return bfs(grid, start, target)
}

func bfs(grid [][]rune, start point, target point) int {
	m := len(grid)
	n := len(grid[0])
	queue := []point{start}
	dist := make([][]int, m)

	for i := 0; i < m; i++ {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dist[i][j] = -1
		}
	}

	dist[start.row][start.col] = 0

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr == target {
			return dist[target.row][target.col]
		}

		for _, direction := range directions {
			next := point{direction[0] + curr.row, direction[1] + curr.col}

			if next.row < 0 || next.row == m || next.col < 0 ||
				next.col == n || dist[next.row][next.col] != -1 ||
				grid[next.row][next.col]-grid[curr.row][curr.col] > 1 {
				continue
			}

			dist[next.row][next.col] = dist[curr.row][curr.col] + 1
			queue = append(queue, next)
		}
	}

	return -1
}

func solvePart2(grid [][]rune) int {
	m := len(grid)
	n := len(grid[0])
	starts := []point{}
	var target point
	best := m * n

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 'S' || grid[i][j] == 'a' {
				starts = append(starts, point{i, j})
				grid[i][j] = 'a'
			} else if grid[i][j] == 'E' {
				target = point{i, j}
				grid[i][j] = 'z'
			}
		}
	}

	for _, start := range starts {
		curr := bfs(grid, start, target)
		if curr != -1 {
			best = imath.MinInt(best, curr)
		}
	}

	return best
}
