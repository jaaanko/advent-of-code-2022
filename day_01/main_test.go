package main

import (
	"strings"
	"testing"
)

var example = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

func TestPart1(t *testing.T) {
	expected := 24000
	actual := solvePart1(parseInput(strings.NewReader(example)))

	if actual != expected {
		t.Errorf("Error in part 1, expected: %v, actual: %v", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	expected := 45000
	actual := solvePart2(parseInput(strings.NewReader(example)))

	if actual != expected {
		t.Errorf("Error in part 2, expected: %v, actual: %v", expected, actual)
	}
}
