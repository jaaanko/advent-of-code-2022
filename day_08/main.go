package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/jaaanko/advent-of-code-2022/imath"
)

var directions = [][]int{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1},
}

func main() {
	file, err := os.Open("input_01.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	grid := [][]int{}

	for scanner.Scan() {
		row := []int{}
		for _, char := range scanner.Text() {
			row = append(row, int(char-'0'))
		}
		grid = append(grid, row)
	}

	fmt.Println(solvePart1(grid))
	fmt.Println(solvePart2(grid))
}

func solvePart1(grid [][]int) int {
	res := 0
	m := len(grid)
	n := len(grid[0])

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if visible(grid, r, c) {
				res++
			}
		}
	}

	return res
}

func visible(grid [][]int, r int, c int) bool {
	m := len(grid)
	n := len(grid[0])

	for _, d := range directions {
		nr := r + d[0]
		nc := c + d[1]

		for nr >= 0 && nr < m && nc >= 0 && nc < n && grid[r][c] > grid[nr][nc] {
			nr += d[0]
			nc += d[1]
		}

		if nr < 0 || nr == m || nc < 0 || nc == n {
			return true
		}
	}

	return false
}

func solvePart2(grid [][]int) int {
	res := 0
	m := len(grid)
	n := len(grid[0])

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			res = imath.MaxInt(res, getScore(grid, r, c))
		}
	}

	return res
}

func getScore(grid [][]int, r int, c int) int {
	m := len(grid)
	n := len(grid[0])
	score := 1

	for _, d := range directions {
		nr := r + d[0]
		nc := c + d[1]
		curr := 0

		for nr >= 0 && nr < m && nc >= 0 && nc < n {
			curr++
			if grid[r][c] <= grid[nr][nc] {
				break
			}
			nr += d[0]
			nc += d[1]
		}
		score *= curr
	}

	return score
}
