package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"

	set "github.com/deckarep/golang-set/v2"
)

func priority(char rune) int {
	if unicode.IsLower(char) {
		return int(char - 'a' + 1)
	}
	return int(char - 'A' + 27)
}

func solvePart1(sacks []string) int {
	total := 0
	for _, sack := range sacks {
		first := set.NewSet[rune]()
		second := set.NewSet[rune]()

		for i := 0; i < len(sack)/2; i++ {
			first.Add(rune(sack[i]))
			second.Add(rune(sack[len(sack)/2+i]))
		}

		common := first.Intersect(second)
		for char := range common.Iter() {
			total += priority(char)
		}
	}
	return total
}

func solvePart2(sacks []string) int {
	total := 0
	if len(sacks)%3 != 0 {
		log.Fatal("Invalid input")
	}

	for i := 0; i < len(sacks); i += 3 {
		first := set.NewSet[rune]()
		for _, char := range sacks[i] {
			first.Add(rune(char))
		}
		second := set.NewSet[rune]()
		for _, char := range sacks[i+1] {
			second.Add(rune(char))
		}
		third := set.NewSet[rune]()
		for _, char := range sacks[i+2] {
			third.Add(rune(char))
		}

		common := first.Intersect(second)
		common = common.Intersect(third)
		for char := range common.Iter() {
			total += priority(char)
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
	sacks := []string{}

	for scanner.Scan() {
		sacks = append(sacks, scanner.Text())
	}

	fmt.Println("Part 1:", solvePart1(sacks))
	fmt.Println("Part 2:", solvePart2(sacks))
}
