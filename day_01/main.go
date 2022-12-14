package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/jaaanko/advent-of-code-2022/imath"
)

func solvePart1(elves [][]int) int {
	maxCalories := 0
	for _, calories := range elves {
		curr := 0
		for _, value := range calories {
			curr += value
		}
		maxCalories = imath.MaxInt(maxCalories, curr)
	}

	return maxCalories
}

func solvePart2(elves [][]int) int {
	totalCalories := []int{}
	for _, calories := range elves {
		curr := 0
		for _, value := range calories {
			curr += value
		}
		totalCalories = append(totalCalories, curr)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(totalCalories)))
	return totalCalories[0] + totalCalories[1] + totalCalories[2]
}

func main() {
	file, err := os.Open("input_01.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	elves := parseInput(file)
	fmt.Println("Part 1:", solvePart1(elves))
	fmt.Println("Part 2:", solvePart2(elves))
}

func parseInput(r io.Reader) [][]int {
	elves := [][]int{}
	calories := []int{}
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			elves = append(elves, calories)
			calories = []int{}
		} else {
			calories = append(calories, num)
		}
	}

	elves = append(elves, calories)
	return elves
}
