package main

import (
	"strings"
	"testing"
)

var example = `A Y
B X
C Z`

func TestPart1(t *testing.T) {
	expected := 15
	actual := solvePart1(parseInput(strings.NewReader(example)))
	if actual != expected {
		t.Errorf("Error in part 1, expected: %v, actual: %v", expected, actual)
	}
}

func TestPart2(t *testing.T) {
	expected := 12
	actual := solvePart2(parseInput(strings.NewReader(example)))
	if actual != expected {
		t.Errorf("Error in part 2, expected: %v, actual: %v", expected, actual)
	}
}
