package main

import (
	"fmt"
	"log"
	"os"

	set "github.com/deckarep/golang-set/v2"
)

func main() {
	bytes, err := os.ReadFile("input_01.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := string(bytes)
	fmt.Println("Part 1:", findMarker(input, 4))
	fmt.Println("Part 2:", findMarker(input, 14))
}

func findMarker(input string, k int) int {
	for i := k; i < len(input); i++ {
		if set.NewSet([]rune(input[i-k:i])...).Cardinality() == k {
			return i
		}
	}
	return -1
}
